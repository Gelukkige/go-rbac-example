package permission

import (
	"context"
	"encoding/json"
	"fmt"
	"go-rbac-example/internal/global"
	"go-rbac-example/internal/model"
	"log"
)

func Init() {
	var users []model.User
	err := global.DB.Preload("Roles").Find(&users).Error
	if err != nil {
		log.Printf("Failed to load users: %v", err)
		return
	}

	var roles []model.Role
	err = global.DB.Find(&roles).Error
	if err != nil {
		log.Printf("Failed to load roles: %v", err)
		return
	}

	redisClient := global.RedisClient
	pipe := redisClient.Pipeline()

	ctx := context.Background()

	for _, user := range users {
		var userRoles []uint64
		for _, role := range user.Roles {
			userRoles = append(userRoles, role.ID)
		}
		key := fmt.Sprintf("user:%d", user.ID)
		roleVal, err := json.Marshal(userRoles)
		if err != nil {
			log.Printf("Failed to marshal roles for user %d: %v", user.ID, err)
			continue
		}
		err = pipe.Set(ctx, key, roleVal, 0).Err()
		if err != nil {
			log.Printf("Failed to set permissions for user %d: %v", user.ID, err)
		}
	}

	for _, role := range roles {
		for _, permission := range role.Permissions {
			key := fmt.Sprintf("role:%d", role.ID)
			field := fmt.Sprintf("%s:%s", permission.Page, permission.Action)
			columnVal, err := json.Marshal(permission.Columns)
			if err != nil {
				log.Printf("Failed to marshal columns for role %d: %v", role.ID, err)
				continue
			}
			err = pipe.HSet(ctx, key, field, columnVal).Err()
			if err != nil {
				log.Printf("Failed to set permissions for role %d: %v", role.ID, err)
			}
		}
	}
}
