package config

import (
	"fmt"
	"go-keycloak/src/domains"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func DbConnect() {
	dns := fmt.Sprintf(
		"root:root@tcp(%s:%s)/go_keycloak?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_PROVIDER_HOST"),
		os.Getenv("MYSQL_PROVIDER_PORT"),
	)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	Database = db

	if err != nil {
		panic("failed to connect database")
	}

	runMigrations()
}

func runMigrations() {
	Database.AutoMigrate(&domains.Event{})
}
