package config

import (
	"log"

	"github.com/spf13/viper"
)

type DBConfig struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

// LoadDBConfig 함수는 DB 설정을 로드하여 DBConfig 구조체를 반환합니다.
func LoadDBConfig() (*DBConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	dbConfig := &DBConfig{
		DBUsername: viper.GetString("db.username"),
		DBPassword: viper.GetString("db.password"),
		DBHost:     viper.GetString("db.host"),
		DBPort:     viper.GetString("db.port"),
		DBName:     viper.GetString("db.name"),
	}
	log.Printf("DB 설정 불러오기 완료")

	return dbConfig, nil
}
