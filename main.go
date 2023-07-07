package jackboxassessment

import (
	"amelia/jackbox-server/src"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.New()

	e.POST("/users", src.HandlePostUser)
	e.PUT("/users/:UserID", src.HandlePutUser)

	e.Run(":8080")
}
