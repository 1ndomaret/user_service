package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	env, err := godotenv.Read("secret.env")
	if err != nil {
		log.Fatal("Error loading secret.env file")
	}

	host := env["DB_HOST"]
	port := env["DB_PORT"]
	user := env["DB_USER"]
	password := env["DB_PASSWORD"]
	dbname := env["DB_NAME"]

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("cant connect db:", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("cant get sql.DB instance:", err)
		return nil, err
	}

	if err = sqlDB.Ping(); err != nil {
		fmt.Println("cant ping db:", err)
		return nil, err
	}

	return db, nil
}
