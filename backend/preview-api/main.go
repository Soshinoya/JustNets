package main

import (
	"github.com/INebotov/JustNets/backend/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	h := handlers.Handlers{}
	h.Init()

	r.GET("/ping", h.Ping)

	notSpam := r.Group("/")
	notSpam.Use(h.NotSpam())
	{
		r.POST("/addemail", h.AddEmail)

		authorized := r.Group("/private")
		authorized.Use(h.HaveAcess())
		{
			authorized.GET("/getmails", h.GetEmail)
		}

	}

	r.Run()
}
