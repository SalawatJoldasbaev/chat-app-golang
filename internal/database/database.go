package database

import (
	"fmt"
	"github.com/SalawatJoldasbaev/chat-app-golang/configs"
	"github.com/SalawatJoldasbaev/chat-app-golang/pkg/utility"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DB *gorm.DB

func ConnectDatabase(config *configs.Config) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.DbName,
		config.Database.Password,
	)
	logMode := logger.Silent
	if config.Database.Debug {
		logMode = logger.Info
	}
	dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		utility.Logger.Error("GORM Error : " + err.Error())
		log.Fatal("Open database connection failed")
	}

	sqlDB, err := dbInstance.DB()
	if err != nil {
		log.Fatalf("Error connection database! %s", err.Error())
	}

	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.Database.MaxConnLifeTime) * time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Duration(config.Database.MaxConnIdleTime) * time.Hour)

	DB = dbInstance
}
