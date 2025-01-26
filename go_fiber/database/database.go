package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	// Database connection

	db, err := gorm.Open(sqlite.Open("myDatabase.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = db
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		panic("Could not close the database")
	}
	sqlDB.Close()
}

func Migrate(val interface{}) {
	DB.AutoMigrate(val)
}
