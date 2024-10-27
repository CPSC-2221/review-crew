package db

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUsersTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE users ("+
			"email VARCHAR (255) PRIMARY KEY"+
			"username VARCHAR (50) UNIQUE NOT NULL,"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
