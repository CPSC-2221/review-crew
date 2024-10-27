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

func CreateReviewTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE review ("+
			"review_ID INT PRIMARY KEY"+
			"email VARCHAR(255) NOT NULL,"+
			"restaurant_ID INT NOT NULL,"+
			"comment TEXT,"+
			"date DATE,"+
			"FOREIGN KEY(email) REFERENCES Users(email)"+
			"FOREIGN KEY(restaurant_ID) REFRENCES Restaurant(restaurant_ID)"+");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}


func CreateReviewCharacterLimitTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE reviewCharacterLimit ("+
			"restaurant_ID INT PRIMARY KEY"+
			"character_limit INT NOT NULL,"+
			"FOREIGN KEY(restaurant_ID) REFRENCE Review(restaurant_ID)"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}


