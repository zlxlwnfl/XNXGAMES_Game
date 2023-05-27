package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	dbConfig, err := LoadDBConfig()
	if err != nil {
		panic(fmt.Sprintf("DB 설정 로드 실패: %v", err))
	}

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbConfig.DBUsername,
		dbConfig.DBPassword,
		dbConfig.DBHost,
		dbConfig.DBPort,
		dbConfig.DBName,
	)

	db, err = gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("DB 연결 실패: %v", err))
	}
}

func DB() *gorm.DB {
	if db == nil {
		InitDB()
	}

	return db
}
