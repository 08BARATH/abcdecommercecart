package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/users", func(c *gin.Context) {
		var user User
		c.BindJSON(&user)
		DB.Create(&user)
		c.JSON(http.StatusOK, user)
	})

	r.GET("/users", func(c *gin.Context) {
		var users []User
		DB.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	r.POST("/users/login", func(c *gin.Context) {
		var input User
		c.BindJSON(&input)
		var user User
		if err := DB.Where("username = ? AND password = ?", input.Username, input.Password).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		token := "token_" + strconv.Itoa(rand.Intn(999999))
		user.Token = token
		DB.Save(&user)
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	r.POST("/items", func(c *gin.Context) {
		var item Item
		c.BindJSON(&item)
		DB.Create(&item)
		c.JSON(http.StatusOK, item)
	})

	r.GET("/items", func(c *gin.Context) {
		var items []Item
		DB.Find(&items)
		c.JSON(http.StatusOK, items)
	})

	authorized := r.Group("/")
	authorized.Use(AuthMiddleware())

	authorized.POST("/carts", func(c *gin.Context) {
		user := c.MustGet("user").(User)
		var body struct {
			ItemID uint
		}
		c.BindJSON(&body)

		var cart Cart
		if err := DB.Where("user_id = ?", user.ID).First(&cart).Error; err != nil {
			cart = Cart{UserID: user.ID}
			DB.Create(&cart)
		}
		cartItem := CartItem{CartID: cart.ID, ItemID: body.ItemID}
		DB.Create(&cartItem)
		c.JSON(http.StatusOK, cartItem)
	})

	authorized.GET("/carts", func(c *gin.Context) {
		var carts []Cart
		DB.Preload("Items").Find(&carts)
		c.JSON(http.StatusOK, carts)
	})

	authorized.POST("/orders", func(c *gin.Context) {
		user := c.MustGet("user").(User)
		var cart Cart
		if err := DB.Where("user_id = ?", user.ID).First(&cart).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cart not found"})
			return
		}
		order := Order{UserID: user.ID, CartID: cart.ID}
		DB.Create(&order)
		DB.Delete(&cart)
		c.JSON(http.StatusOK, gin.H{"message": "Order Successful", "order_id": order.ID})
	})

	authorized.GET("/orders", func(c *gin.Context) {
		user := c.MustGet("user").(User)
		var orders []Order
		DB.Where("user_id = ?", user.ID).Find(&orders)
		c.JSON(http.StatusOK, orders)
	})
}
