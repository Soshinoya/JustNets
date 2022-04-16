package handlers

import (
	"fmt"
	"net/http"

	"github.com/INebotov/JustNets/backend/datastructs"
	"github.com/INebotov/JustNets/backend/db"
	"github.com/INebotov/JustNets/backend/logger"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	log      logger.MyLog
	DataBase db.DataBase
}

func (h *Handlers) Init() {
	h.DataBase = db.DataBase{DEBUG: true}
	h.DataBase.Init()
	h.log = logger.MyLog{}
	h.log.Init("db_logs", "Handlers")
}

// ###### Middlewares ######

func (h *Handlers) HaveAcess() gin.HandlerFunc {
	auth := "Fuck You Stupid Hacker!"
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		if tokenString != auth {
			h.log.LogWarn(fmt.Sprintf("! Status Unauthorized IP: %s !", c.ClientIP()))
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}

func (h *Handlers) NotSpam() gin.HandlerFunc { // TODO !!!
	return func(c *gin.Context) {
		c.Next()
	}
}

// ###### Handlers ######

func (h *Handlers) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
	return
}

func (h *Handlers) AddEmail(c *gin.Context) {
	if email, ok := c.GetPostForm("email"); ok {
		result := h.DataBase.DB.Create(datastructs.SubscriberEmail{Email: email})
		if result.Error != nil {
			h.log.LogWarn(fmt.Sprintf("Error While Adding Email: %s", result.Error)) // Not Secure
			c.JSON(500, gin.H{
				"status": fmt.Sprintf("EROOR %s", result.Error),
			})
			return
		}
		c.JSON(200, gin.H{
			"status": "Sucsess!!",
		})
		return
	}
	h.log.LogWarn(fmt.Sprintf("Empty Form Body"))
	c.JSON(400, gin.H{
		"status": "EROOR Empty Form Body",
	})
	return
}

func (h *Handlers) GetEmail(c *gin.Context) {
	res := datastructs.SubscriberEmail{}
	r := h.DataBase.DB.Find(&res)
	if r.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(200, gin.H{
		"emails": res,
	})
	return
}
