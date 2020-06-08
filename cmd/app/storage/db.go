package storage

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// New constructor for the DB
func New() *gorm.DB {
	// Start a new connection with data source
	db, err := gorm.Open("sqlite3", "./data/products.db")

	// Check for errors
	if err != nil {

		// Output issues
		fmt.Println("Storage Error: ", err)
	}

	// Set number of connection
	db.DB().SetMaxIdleConns(2)

	// DB setup for debugging
	db.LogMode(false)
	db.Debug()

	return db
}

// Run a DB connection test
func TestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./data/test/products_test.db")
	if err != nil {
		fmt.Println("Storage Error: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	db.Debug()
	return db
}
