package database

import (
	"fmt"

	"github.com/iki-rumondor/init-golang-service/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() (*gorm.DB, error) {

	env, err := utils.GetDatabaseEnv()
	if err != nil{
		return nil, err
	}
	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname%s sslmode=%s", env["host"], env["port"], env["user"], env["password"], env["name"], env["sslmode"])

	gormDB, err := gorm.Open(postgres.Open(strConn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
