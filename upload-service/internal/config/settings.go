package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	EnvDbHost        = "DB_HOST"
	EnvDbPort        = "DB_PORT"
	EnvDbName        = "DB_NAME"
	EnvDbUser        = "DB_USER"
	EnvDbPassword    = "DB_PASSWORD"
	EnvMinioEndpoint = "MINIO_ENDPOINT"
	EnvMinioAccess   = "MINIO_ACCESSKEY"
	EnvMinioSecret   = "MINIO_SECRETKEY"
	EnvMinioBucket   = "MINIO_BUCKET"
)

const (
	DefaultDbHost        = "localhost"
	DefaultDbPort        = "5432"
	DefaultDbName        = "upload"
	DefaultDbUser        = "postgres"
	DefaultDbPassword    = "postgres"
	DefaultMinioEndpoint = "localhost:9000"
	DefaultMinioBucket   = "videos"
)

func init() {
	viper.SetDefault(EnvDbHost, DefaultDbHost)
	viper.SetDefault(EnvDbPort, DefaultDbPort)
	viper.SetDefault(EnvDbName, DefaultDbName)
	viper.SetDefault(EnvDbUser, DefaultDbUser)
	viper.SetDefault(EnvDbPassword, DefaultDbPassword)
	viper.SetDefault(EnvMinioSecret, DefaultMinioEndpoint)
	viper.SetDefault(EnvMinioBucket, DefaultMinioBucket)
	viper.AutomaticEnv()
}

func DbHost() string {
	return viper.GetString(EnvDbHost)
}

func DbPort() string {
	return viper.GetString(EnvDbPort)
}

func DbName() string {
	return viper.GetString(EnvDbName)
}

func DbUser() string {
	return viper.GetString(EnvDbUser)
}

func DbPassword() string {
	return viper.GetString(EnvDbPassword)
}

func PostgresDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		DbHost(),
		DbUser(),
		DbPassword(),
		DbName(),
		DbPort(),
	)
}

func MinioEndpoint() string {
	return viper.GetString(EnvMinioEndpoint)
}

func MinioAccessKey() string {
	return viper.GetString(EnvMinioAccess)
}

func MinioSecretKey() string {
	return viper.GetString(EnvMinioSecret)
}

func MinioBucket() string {
	return viper.GetString(EnvMinioBucket)
}