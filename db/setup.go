package db

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUsersTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE users ("+
			"email VARCHAR (255) PRIMARY KEY,"+
			"username VARCHAR (50) UNIQUE NOT NULL"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func CreateRestaurantTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE restaurant ("+
			"restaurantID INTEGER PRIMARY KEY,"+
			"name VARCHAR (255) NOT NULL,"+
			"location VARCHAR (255) NOT NULL,"+
			"description VARCHAR (255) NOT NULL,"+
			"isPremium BOOLEAN NOT NULL"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func CreatePizzaImageTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE hasPizzaImage ("+
			"name VARCHAR (255) PRIMARY KEY,"+
			"haPizzaImage BOOLEAN NOT NULL"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func CreateRepliesToTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE RepliesTo ("+
			"reviewIDRepliesTo INTEGER PRIMARY KEY REFERENCES review(reviewID),"+
			"reviewIDIsRepliedTo INTEGER PRIMARY KEY REFERENCES review(reviewID)"+
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
			"reviewID INT PRIMARY KEY,"+
			"email VARCHAR(255) NOT NULL,"+
			"restaurantID INT NOT NULL,"+
			"comment TEXT,"+
			"date DATE,"+
			"FOREIGN KEY(email) REFERENCES Users(email),"+
			"FOREIGN KEY(restaurantID) REFERENCES restaurant(restaurantID)"+");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func CreateReviewCharacterLimitTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE reviewCharacterLimit ("+
			"restaurantID INT PRIMARY KEY"+
			"characterLimit INT NOT NULL,"+
			"FOREIGN KEY(restaurantID) REFERENCES review(restaurantID)"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func CreateManagesTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE manages ("+
			"email VARCHAR (255),"+
			"restaurantID VARCHAR (50),"+
			"CanDeleteComments BOOLEAN,"+
			"CanUpdateListing BOOLEAN,"+
			"PRIMARY KEY(email, restaurantID)"+
			"FOREIGN KEY(email) REFERENCES users(email),"+
			"FOREIGN KEY(restaurantID) REFERENCES restaurant(restaurantID)"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func CreateLikesTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE likes ("+
			"ReviewID INTEGER,"+
			"email VARCHAR (255),"+
			"PRIMARY KEY(reviewID,email),"+
			"FOREIGN KEY(reviewID) REFERENCES review(reviewID)"+
			"FOREIGN KEY(email) REFERENCES users(email)"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func CreateOwnsTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE owns ("+
			"email VARCHAR (255)"+
			"restaurantID INTEGER"+
			"PRIMARY KEY (email, restaurantID)"+
			"FOREIGN KEY (email) REFERENCES Users(email)"+
			"FOREIGN KEY (restaurantID) REFERENCES Restaurant(restaurantID)"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
