package db

import (
	"fmt"
	"log"
	"os"
	"vaqua/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() *gorm.DB {

	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	//open DB
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to database successfully!")

	//migrate models to create db tables
	err = DB.AutoMigrate(&models.User{},
		&models.Transaction{},
		)
	if err != nil {
		log.Fatal("failed to migrate schema", err)
	}
	return DB
}

func Ping() error {
	sqlDB, err := DB.DB() 
	if err != nil {
		return err
	}
	//ping the database.
	return sqlDB.Ping()
}
