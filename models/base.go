package models

import (
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	e := godotenv.Load() // Load .env file
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(e)

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	// Build connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Println(err)
	}

	db = conn

	// AutoMigrate will automatically migrate our scheme.
	// It will automatically create the table based on our provided model
	// We dont need to create the table manually.

	db.Debug().AutoMigrate(&Account{}, &Contact{})
	fmt.Println("base")
	// db.Debug().AutoMigrate(&Account{}, &Contact{})

}

// GetDB returns a handle to DB object
func GetDB() *gorm.DB{
	return db
}