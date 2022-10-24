package database

import (
	"fmt"
	"log"
	"os"

	"github.com/imajean/go-todo-postresql/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func InitDbConnection() {
	// load env variables
	if envErr := godotenv.Load(); envErr != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("MYHOST")
	dbPort := os.Getenv("MYDBPORT")
	dbName := os.Getenv("MYDBNAME")
	user := os.Getenv("MYUSER")
	password := os.Getenv("MYPASSWORD")

	// database connection string
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, dbPort, dbName, user, password)
	fmt.Println(dsn)

	// open connection to database
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database!")
	}

	// close connection to database when the main function finishes
	// postgresDB, _ := db.DB()
	// defer postgresDB.Close()

	// make migrations to the database if they have not been created
	db.AutoMigrate(&models.Todo{})
}

func GetDB() *gorm.DB {
	return db
}
