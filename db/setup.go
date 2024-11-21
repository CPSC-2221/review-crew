package db

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUsersTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE IF NOT EXISTS users ("+
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
		"CREATE TABLE IF NOT EXISTS restaurants ("+
			"restaurantID SERIAL PRIMARY KEY,"+
			"name TEXT NOT NULL UNIQUE,"+
			"location TEXT NOT NULL,"+
			"description TEXT NOT NULL"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func CreateHasPizzaImageTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE IF NOT EXISTS hasPizzaImage ("+
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

func CreateHasBurgerEmojiTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE IF NOT EXISTS hasBurgerEmoji ("+
			"username TEXT PRIMARY KEY,"+
			"hasBurgerEmoji BOOLEAN NOT NULL,"+
			"FOREIGN KEY(username) REFERENCES users(username)"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func CreateReviewTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE IF NOT EXISTS review ("+
			"reviewID SERIAL PRIMARY KEY,"+
			"email TEXT NOT NULL,"+
			"restaurantID INTEGER NOT NULL,"+
			"comment TEXT,"+
			"datetime TIMESTAMPTZ,"+
			"FOREIGN KEY(email) REFERENCES Users(email),"+
			"FOREIGN KEY(restaurantID) REFERENCES restaurants(restaurantID)"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func CreateRepliesToTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE IF NOT EXISTS repliesTo ("+
			"repliesToReviewID INTEGER,"+
			"isRepliedToReviewID INTEGER,"+
			"PRIMARY KEY (repliesToReviewID, isRepliedToReviewID),"+
			"FOREIGN KEY(repliesToReviewID) REFERENCES review(reviewID) ON DELETE CASCADE,"+
			"FOREIGN KEY(isRepliedToReviewID) REFERENCES review(reviewID) ON DELETE CASCADE"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func CreateManagesTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE IF NOT EXISTS manages ("+
			"email TEXT NOT NULL,"+
			"restaurantID INTEGER NOT NULL,"+
			"CanDeleteComments BOOLEAN NOT NULL,"+
			"CanUpdateListing BOOLEAN,"+
			"PRIMARY KEY (email, restaurantID),"+
			"FOREIGN KEY (email) REFERENCES users(email) ON DELETE CASCADE,"+
			"FOREIGN KEY (restaurantID) REFERENCES restaurants(restaurantID) ON DELETE CASCADE"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func CreateLikesTable(c *gin.Context) {
	_, err := dbpool.Exec(c,
		"CREATE TABLE IF NOT EXISTS likes ("+
			"reviewID INTEGER NOT NULL,"+
			"email TEXT NOT NULL,"+
			"PRIMARY KEY (reviewID, email),"+
			"FOREIGN KEY (reviewID) REFERENCES review(reviewID) ON DELETE CASCADE,"+
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
		"CREATE TABLE IF NOT EXISTS owns ("+
			"email TEXT NOT NULL,"+
			"restaurantID INTEGER,"+
			"PRIMARY KEY (email, restaurantID),"+
			"FOREIGN KEY (email) REFERENCES users(email),"+
			"FOREIGN KEY (restaurantID) REFERENCES restaurants(restaurantID)"+
			");")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
