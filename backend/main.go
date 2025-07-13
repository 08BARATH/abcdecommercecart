package main

import "github.com/gin-gonic/gin"

func main() {
	InitDB()
	r := gin.Default()
	SetupRoutes(r)
	r.Run(":8080")
}
