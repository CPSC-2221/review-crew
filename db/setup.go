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

func CreateOwnsTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE Owns ("+
			"email VARCHAR (255)"+
			"restaurantID INTEGER"+
			"PRIMARY KEY (email, restaurantID)"+
			"FOREIGN KEY (email) REFERENCES Users(email)"+
			"FOREIGN KEY (restaurantID) REFERENCES Restaurant(restaurantID))"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
