package routes

import (
	"amelia/jackbox-server/src/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (cm *ClientManager) HandlePostUser(c *gin.Context) {
	body := models.User{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		// TODO: add logger
		fmt.Println(err)
		return
	}

	// TODO: add input validation
	// TODO: hash/encrypt password before storing

	body.UserID = uuid.NewString()

	_, err := cm.db.Exec("INSERT INTO Users VALUES(?,?,?,?,?,?)",
		body.UserID, body.FirstName, body.LastName, body.Email, body.Username, body.Password)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		// TODO: add logger
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"UserID": body.UserID})
}

func (cm *ClientManager) HandleGetUser(c *gin.Context) {
	userID := c.Param("UserID")

	user := &models.User{
		UserID: userID,
	}
	row := cm.db.QueryRow("SELECT FirstName, LastName, Email, Username FROM Users WHERE Id=?", userID)
	if err := row.Scan(&user.FirstName, &user.LastName, &user.Email, &user.Username); err != nil {
		c.Status(http.StatusInternalServerError)
		// TODO: add logger
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (cm *ClientManager) HandlePutUser(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}
