package storage

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// New constructor for the DB
func New() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./data/products.db")
	if err != nil {
		fmt.Println("Storage Error: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	db.Debug()
	return db
}

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
