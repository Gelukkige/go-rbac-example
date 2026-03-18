package initialize

import (
	"fmt"
	"go-rbac-example/internal/global"
	"go-rbac-example/internal/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBInit() {
	dbConfig := global.Config.Database

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
		dbConfig.SSLMode,
		dbConfig.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	err = db.AutoMigrate(
		model.User{},
		model.Role{},
		model.Permission{},
	)

	global.DB = db
	log.Println("数据库连接成功并完成迁移!")
}
