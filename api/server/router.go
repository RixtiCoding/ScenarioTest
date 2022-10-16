package server

import (
	"MessageFilter/api/helpers"
	"MessageFilter/db"
	"MessageFilter/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	dtb, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}

	router.GET("/badwords", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, helpers.FetchBannedWords().Words)
	})

	router.POST("/images", func(c *gin.Context) {
		var ImageMessage utility.ImageMessage

		err = c.BindJSON(&ImageMessage)
		if err != nil {
			fmt.Println(err)
		}

		dtb.SetToRedis(db.Ctx, ImageMessage.Body, "IMAGE")

	})

	return router
}
