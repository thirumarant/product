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
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}
