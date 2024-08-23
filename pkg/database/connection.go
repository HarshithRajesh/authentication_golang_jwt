package database

import (
	"log"
	"os"
	"github.com/HarshithRajesh/idea1/pkg/models"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectionDB() (*gorm.DB,error){
	err := godotenv.Load(".env")
	 if err !=nil{
		log.Fatal("error loading .env files")
	 }
	dsn := os.Getenv("POSTGRES_URL")

	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		return nil,err
	}

	DB = db
	db.AutoMigrate(&models.User{})

	return db,nil
}