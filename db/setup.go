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

func CreateManagesTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE manages ("+
			"Uemail VARCHAR (255),"+
			"restaurantID VARCHAR (50),"+
			"CanDeleteComments BOOL,"+
			"CanUpdateListing BOOL,"+
			"PRIMARY KEY(email,restaurantID)"+
			"FOREIGN KEY(Uemail) REFERENCES users(email),"+
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
			"userEmail VARCHAR (255),"+
			"PRIMARY KEY(ReviewID,userEmail),"+
			"FOREIGN KEY(ReviewID) REFERENCES review(ReviewID)"+
			"FOREIGN KEY(userEmail) REFERENCES users(email)"+
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
			"email VARCHAR (255),"+
			"restaurantID INTEGER,"+
			"PRIMARY KEY (email, restaurantID),"+
			"FOREIGN KEY (email) REFERENCES Users(email),"+
			"FOREIGN KEY (restaurantID) REFERENCES Restaurant(restaurantID)"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
