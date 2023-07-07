package main

import (
	"amelia/jackbox-server/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.New()

	cm, err := routes.NewClientManager("users.db")
	if err != nil {
		panic(err)
	}

	e.POST("/users", cm.HandlePostUser)
	e.GET("/users/:UserID", cm.HandleGetUser)
	e.PUT("/users/:UserID", cm.HandlePutUser)

	e.Run(":8080")
}
