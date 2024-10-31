package db

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUsersTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE users ("+
			"email TEXT PRIMARY KEY,"+
			"username TEXT UNIQUE NOT NULL"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func CreateRestaurantTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE restaurants ("+
			"restaurantID SERIAL PRIMARY KEY,"+
			"name TEXT NOT NULL,"+
			"location TEXT NOT NULL,"+
			"description TEXT NOT NULL,"+
			"isPremium BOOLEAN NOT NULL"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func CreateHasPizzaImageTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE hasPizzaImage ("+
			"name TEXT PRIMARY KEY,"+
			"hasPizzaImage BOOLEAN NOT NULL,"+
			"FOREIGN KEY(name) REFERENCES restaurants(name)"+
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
			"reviewID SERIAL PRIMARY KEY,"+
			"email TEXT NOT NULL,"+
			"restaurantID INTEGER NOT NULL,"+
			"repliesToReviewID INTEGER,"+
			"comment TEXT,"+
			"datetime TIMESTAMPTZ,"+
			"FOREIGN KEY(email) REFERENCES Users(email),"+
			"FOREIGN KEY(restaurantID) REFERENCES restaurants(restaurantID),"+
			"FOREIGN KEY(repliesToReviewID) REFERENCES review(reviewID)"+
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
			"email TEXT NOT NULL,"+
			"restaurantID INTEGER NOT NULL,"+
			"CanDeleteComments BOOLEAN NOT NULL,"+
			"CanUpdateListing BOOLEAN,"+
			"PRIMARY KEY (email, restaurantID)"+
			"FOREIGN KEY (email) REFERENCES users(email),"+
			"FOREIGN KEY (restaurantID) REFERENCES restaurant(restaurantID)"+
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
			"reviewID INTEGER NOT NULL,"+
			"email TEXT NOT NULL,"+
			"PRIMARY KEY (reviewID, email),"+
			"FOREIGN KEY (reviewID) REFERENCES review(reviewID)"+
			"FOREIGN KEY (email) REFERENCES users(email)"+
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
			"email TEXT NOT NULL,"+
			"restaurantID INTEGER,"+
			"PRIMARY KEY (email, restaurantID),"+
			"FOREIGN KEY (email) REFERENCES users(email),"+
			"FOREIGN KEY (restaurantID) REFERENCES restaurant(restaurantID)"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
