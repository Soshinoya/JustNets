package handlers

import (
	"fmt"
	"net/http"
	"net/mail"

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
	h.DataBase.Init("emails", datastructs.SubscriberEmail{})
	h.log = logger.MyLog{}
	h.log.Init("router_logs", "Handlers")
}

// ###### Middlewares ######

func (h *Handlers) HaveAcess() gin.HandlerFunc {
	auth := "Fuck You Stupid Hacker!"
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA)+1:]
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

}

func (h *Handlers) AddEmail(c *gin.Context) {
	email, ok := c.GetPostForm("email")
	if !ok {
		h.log.LogWarn("Empty Form Body")
		c.JSON(400, gin.H{
			"status": "EROOR Empty Form Body",
		})
		return
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		c.JSON(400, gin.H{
			"status": "This is not a email",
		})
		return
	}
	var eemail = datastructs.SubscriberEmail{Email: email}
	r := h.DataBase.DB.Where("email = ?", email).First(&eemail)
	if r.RowsAffected > 0 {
		c.JSON(400, gin.H{
			"status": "Already have this Email",
		})
		return
	}

	result := h.DataBase.DB.Create(&datastructs.SubscriberEmail{Email: email})
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

}

func (h *Handlers) GetEmail(c *gin.Context) {
	res := []datastructs.SubscriberEmail{}
	r := h.DataBase.DB.Find(&res)
	if r.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(200, gin.H{
		"emails": res,
	})

}
