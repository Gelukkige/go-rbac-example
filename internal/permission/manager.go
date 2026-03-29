package permission

import (
	"context"
	"encoding/json"
	"fmt"
	"go-rbac-example/internal/global"
	"go-rbac-example/internal/model"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// 权限初始化
func Init() {
	ctx := context.Background()
	redisClient := global.RedisClient
	if redisClient == nil {
		log.Println("Redis client is not initialized")
		return
	}

	var batchSize = 500

	// 用户-角色关系 (Set 存储)
	err := global.DB.Model(&model.User{}).Preload("Roles").FindInBatches(&[]model.User{}, batchSize, func(tx *gorm.DB, batch int) error {
		usersPtr, ok := tx.Statement.Dest.(*[]model.User)
		if !ok {
			return fmt.Errorf("failed to assert destination to *[]model.User")
		}

		pipe := redisClient.Pipeline()
		for _, user := range *usersPtr {
			key := fmt.Sprintf("user:%d", user.ID)
			pipe.Del(ctx, key) // 清理旧缓存

			if len(user.Roles) > 0 {
				roleIds := make([]interface{}, len(user.Roles))
				for i, r := range user.Roles {
					roleIds[i] = r.ID
				}
				pipe.SAdd(ctx, key, roleIds...)
				expireTime := rand.Intn(30) + 30 // 随机30-60分钟过期
				pipe.Expire(ctx, key, time.Duration(expireTime)*time.Minute)
			}
		}
		_, err := pipe.Exec(ctx)
		return err
	}).Error

	if err != nil {
		log.Printf("Error processing user-role batches: %v", err)
	}

	// 角色-权限关系 (Hash 存储)
	var roles []model.Role
	if err := global.DB.Preload("Permissions").Find(&roles).Error; err != nil {
		log.Printf("Error loading roles: %v", err)
		return
	}

	pipe := redisClient.Pipeline()
	for _, role := range roles {
		key := fmt.Sprintf("role:%d", role.ID)
		pipe.Del(ctx, key) // 清理旧缓存

		if len(role.Permissions) > 0 {
			for _, perm := range role.Permissions {
				field := fmt.Sprintf("%s:%s", perm.Page, perm.Action)
				columnVal, _ := json.Marshal(perm.Columns)
				pipe.HSet(ctx, key, field, columnVal)
			}
			expireTime := rand.Intn(30) + 30 // 随机30-60分钟过期
			pipe.Expire(ctx, key, time.Duration(expireTime)*time.Minute)
		}
	}

	if _, err := pipe.Exec(ctx); err != nil {
		log.Printf("Error executing role-permission pipeline: %v", err)
	}
}

// 辅助函数：切片去重合并
func mergeSlices(s1, s2 []string) []string {
	m := make(map[string]bool)
	for _, v := range s1 {
		m[v] = true
	}
	for _, v := range s2 {
		m[v] = true
	}
	res := make([]string, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}

// 从数据库加载用户角色并写入 Redis
func loadUserRolesToCache(ctx context.Context, userID uint64) ([]string, error) {
	var user model.User
	err := global.DB.Preload("Roles").First(&user, userID).Error
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("user:%d", userID)
	if len(user.Roles) == 0 {
		return []string{}, nil // 无角色直接返回
	}

	roleIds := make([]interface{}, len(user.Roles))
	roleIdStrings := make([]string, len(user.Roles))
	for i, r := range user.Roles {
		roleIds[i] = r.ID
		roleIdStrings[i] = fmt.Sprintf("%d", r.ID)
	}

	pipe := global.RedisClient.Pipeline()
	pipe.SAdd(ctx, key, roleIds...)
	expireTime := rand.Intn(30) + 30
	pipe.Expire(ctx, key, time.Duration(expireTime)*time.Minute)
	_, err = pipe.Exec(ctx)

	return roleIdStrings, err
}

// 从数据库加载角色权限并写入 Redis
func loadRolePermissionsToCache(ctx context.Context, roleIdStr string) error {
	var role model.Role
	// 注意：这里需要根据你的 roleId 类型进行转换
	err := global.DB.Preload("Permissions").First(&role, "id = ?", roleIdStr).Error
	if err != nil {
		return err
	}

	key := fmt.Sprintf("role:%s", roleIdStr)
	pipe := global.RedisClient.Pipeline()
	pipe.Del(ctx, key)

	for _, perm := range role.Permissions {
		field := fmt.Sprintf("%s:%s", perm.Page, perm.Action)
		columnVal, _ := json.Marshal(perm.Columns)
		pipe.HSet(ctx, key, field, columnVal)
	}
	expireTime := rand.Intn(30) + 30
	pipe.Expire(ctx, key, time.Duration(expireTime)*time.Minute)
	_, err = pipe.Exec(ctx)
	return err
}

// 获取用户的权限列表
func GetUserAllPages(ctx context.Context, userID uint64) ([]string, error) {
	redisClient := global.RedisClient
	if redisClient == nil {
		return nil, fmt.Errorf("Redis client is not initialized")
	}

	userKey := fmt.Sprintf("user:%d", userID)

	// 获取用户所有角色 ID
	roleIds, err := redisClient.SMembers(ctx, userKey).Result()

	// 如果 Redis 中没有用户角色数据，触发回源加载
	if err != nil || len(roleIds) == 0 {
		roleIds, err = loadUserRolesToCache(ctx, userID)
		if err != nil {
			return nil, fmt.Errorf("failed to load user roles from DB: %v", err)
		}
		if len(roleIds) == 0 {
			return []string{}, nil
		}
	}

	// 使用 map 去重页面名称
	pageMap := make(map[string]bool)

	// 批量获取所有角色的权限字段 (HKeys)
	pipe := redisClient.Pipeline()
	for _, roleId := range roleIds {
		pipe.HKeys(ctx, fmt.Sprintf("role:%s", roleId))
	}

	cmds, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return nil, err
	}

	for i, cmd := range cmds {
		fields, cmdErr := cmd.(*redis.StringSliceCmd).Result()

		if cmdErr == redis.Nil || len(fields) == 0 {
			currentRoleId := roleIds[i]
			// 触发该角色的回源加载
			err := loadRolePermissionsToCache(ctx, currentRoleId)
			if err != nil {
				log.Printf("Failed to reload role %s to cache: %v", currentRoleId, err)
				continue
			}
			// 重新获取一次该角色的权限字段
			fields, _ = redisClient.HKeys(ctx, fmt.Sprintf("role:%s", currentRoleId)).Result()
		}

		// 解析 page:action 并提取 page
		for _, field := range fields {
			parts := strings.Split(field, ":")
			if len(parts) > 0 {
				pageMap[parts[0]] = true
			}
		}
	}

	pages := make([]string, 0, len(pageMap))
	for page := range pageMap {
		pages = append(pages, page)
	}

	return pages, nil
}

// 获取用户在特定页面特定方法的列权限
func GetUserPagePermissions(ctx context.Context, userID uint64, page string, action string) ([]string, error) {
	redisClient := global.RedisClient
	if redisClient == nil {
		return nil, fmt.Errorf("Redis client is not initialized")
	}
	userKey := fmt.Sprintf("user:%d", userID)
	roleIds, err := redisClient.SMembers(ctx, userKey).Result()
	if err != nil || len(roleIds) == 0 {
		// 从数据库加载并写回redis
		roleIds, err = loadUserRolesToCache(ctx, userID)
		if err != nil || len(roleIds) == 0 {
			return nil, err
		}
	}

	field := fmt.Sprintf("%s:%s", page, action)
	var finalColumns []string

	pipe := redisClient.Pipeline()
	for _, roleId := range roleIds {
		pipe.HGet(ctx, fmt.Sprintf("role:%s", roleId), field)
	}
	cmds, _ := pipe.Exec(ctx)

	for i, cmd := range cmds {
		val, err := cmd.(*redis.StringCmd).Result()

		if err == redis.Nil {
			// 缓存未命中，尝试从数据库加载并写回 Redis
			roleId := roleIds[i]
			_ = loadRolePermissionsToCache(ctx, roleId)
			// 再次尝试从 Redis
			val, _ = redisClient.HGet(ctx, fmt.Sprintf("role:%s", roleId), field).Result()
		}

		if val != "" {
			var columns []string
			json.Unmarshal([]byte(val), &columns)
			finalColumns = mergeSlices(finalColumns, columns)
		}
	}

	return finalColumns, nil
}

// 删除指定用户的角色缓存
func DeleteUserCache(ctx context.Context, userID uint64) error {
	key := fmt.Sprintf("user:%d", userID)
	return global.RedisClient.Del(ctx, key).Err()
}

// 删除指定角色的权限缓存
func DeleteRoleCache(ctx context.Context, roleID uint64) error {
	key := fmt.Sprintf("role:%d", roleID)
	return global.RedisClient.Del(ctx, key).Err()
}
