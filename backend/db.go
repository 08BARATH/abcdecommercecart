package main

import (
	"log"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// Use this even with modernc — same code
	DB, err = gorm.Open(sqlite.Open("ecom.db?_pragma=foreign_keys(1)"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&User{}, &Item{}, &Cart{}, &CartItem{}, &Order{})
	if err != nil {
		log.Fatal("AutoMigrate error:", err)
	}
}
