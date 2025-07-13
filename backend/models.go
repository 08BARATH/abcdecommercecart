
package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Token    string
}

type Item struct {
	gorm.Model
	Name  string
	Price float64
}

type Cart struct {
	gorm.Model
	UserID uint
	Items  []CartItem
}

type CartItem struct {
	gorm.Model
	CartID uint
	ItemID uint
}

type Order struct {
	gorm.Model
	UserID uint
	CartID uint
}
