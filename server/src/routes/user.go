package routes

import (
	"amelia/jackbox-server/src/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (cm *ClientManager) HandleLogin(c *gin.Context) {
	type HandleLoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	body := HandleLoginInput{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		// TODO: add logger
		fmt.Println(err)
		return
	}

	row := cm.db.QueryRow("SELECT Id FROM Users WHERE Username=? AND Password=?", body.Username, body.Password)
	userID := ""
	if err := row.Scan(&userID); err != nil {
		c.Status(http.StatusInternalServerError)
		// TODO: add logger
		fmt.Println(err)
		return
	}

	// TODO: Normally, I'd create a jwt key and set a cookie for the client making the request
	// this is an incredibly insecure workaround to show a proof-of-concept, must change after

	c.JSON(http.StatusOK, gin.H{"UserID": userID})
}

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
