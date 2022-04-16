package main

import "github.com/INebotov/JustNets/backend/db"

func main() {

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run()
	db := db.DataBase{}
	db.Init()
}
