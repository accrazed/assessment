package main

import (
	"amelia/jackbox-server/src/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.New()

	cm, err := routes.NewClientManager("users.db")
	if err != nil {
		panic(err)
	}

	e.Use(gin.Logger())
	e.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT"},
		AllowHeaders: []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Accept, Origin, Cache-Control, X-Requested-With"},
	}))

	e.POST("/session", cm.HandleLogin)

	e.POST("/users", cm.HandlePostUser)
	e.GET("/users/:UserID", cm.HandleGetUser)
	e.PUT("/users/:UserID", cm.HandlePutUser)

	e.Run(":8080")
}
