package db

import (
	"log"
	"net/http"
	"time"

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

func InsertDummyData(c *gin.Context) {

	dbpool.QueryRow(c, "INSERT INTO users(username, email) VALUES ($1, $2) RETURNING *;", "Michelin Inspector", "MI@gmail.com")

	parsedTime, err := time.Parse("2006-01-02 15:04:05", "2023-03-30 15:30:41")
	if err != nil {
		// Handle the error, e.g., log it or return it
		log.Fatalf("Error parsing time: %v", err)
	}

	dbpool.QueryRow(c, "INSERT INTO restaurants(restaurantID, name, location, description) VALUES ($1, $2, $3, $4) RETURNING *;", 1, "Cactus Club", "Downtown Vancouver", "Cactus Club Cafe is your go-to casual dining restaurant. We offer the best in global cuisine using local, fresh ingredients served in a vibrant, contemporary setting.")

	dbpool.QueryRow(c, "INSERT INTO users(username, email) VALUES ($1, $2) RETURNING *;", "CactusClub", "CactusClub@gmail.com")

	dbpool.QueryRow(c, "INSERT INTO owns(email, restaurantID) VALUES ($1, $2) RETURNING *;", "CactusClub@gmail.com", 1)

	dbpool.QueryRow(c, "INSERT INTO review(reviewID, email, restaurantID, comment, datetime) VALUES ($1, $2, $3, $4, $5) RETURNING *;", 100, "MI@gmail.com", 1, "AMAZING!", parsedTime)

	parsedTime2, err := time.Parse("2006-01-02 15:04:05", "2021-11-22 14:10:20")
	if err != nil {
		// Handle the error, e.g., log it or return it
		log.Fatalf("Error parsing time: %v", err)
	}

	dbpool.QueryRow(c, "INSERT INTO restaurants(restaurantID, name, location, description) VALUES ($1, $2, $3, $4) RETURNING *;", 2, "The Old Spaghetti Factory", "53 Water Street, Vancouver", "A customer favourite, known for good times and great food. This is our first location in Vancouver’s historic Gastown. Reservations recommended.")

	dbpool.QueryRow(c, "INSERT INTO users(username, email) VALUES ($1, $2) RETURNING *;", "OldSpaghettiFactory", "OldFactory@gmail.com")

	dbpool.QueryRow(c, "INSERT INTO owns(email, restaurantID) VALUES ($1, $2) RETURNING *;", "OldFactory@gmail.com", 2)

	dbpool.QueryRow(c, "INSERT INTO review(reviewID, email, restaurantID, comment, datetime) VALUES ($1, $2, $3, $4, $5) RETURNING *;", 200, "MI@gmail.com", 2, "Incredible food", parsedTime2)

	parsedTime3, err := time.Parse("2006-01-02 15:04:05", "2024-04-15 19:23:40")
	if err != nil {
		// Handle the error, e.g., log it or return it
		log.Fatalf("Error parsing time: %v", err)
	}

	dbpool.QueryRow(c, "INSERT INTO restaurants(restaurantID, name, location, description) VALUES ($1, $2, $3, $4) RETURNING *;", 3, "Zefferelli's Spaghetti Joint", "1136 Robson St 2nd Floor, Vancouver", "Over looking busy Robson street, a 2nd floor room that offers a cozy and charming atmosphere with friendly and efficient service that mostly prides itself on home-style, flavorful Italian dishes and pasta. Offering portion sizes that even Mama would be proud of.")

	dbpool.QueryRow(c, "INSERT INTO users(username, email) VALUES ($1, $2) RETURNING *;", "Zefferelis", "Zefferelis@gmail.com")

	dbpool.QueryRow(c, "INSERT INTO owns(email, restaurantID) VALUES ($1, $2) RETURNING *;", "Zefferelis@gmail.com", 3)

	dbpool.QueryRow(c, "INSERT INTO review(reviewID, email, restaurantID, comment, datetime) VALUES ($1, $2, $3, $4, $5) RETURNING *;", 300, "MI@gmail.com", 3, "good staff", parsedTime3)

	parsedTime4, err := time.Parse("2006-01-02 15:04:05", "2024-08-23 18:24:20")
	if err != nil {
		// Handle the error, e.g., log it or return it
		log.Fatalf("Error parsing time: %v", err)
	}

	dbpool.QueryRow(c, "INSERT INTO restaurants(restaurantID, name, location, description) VALUES ($1, $2, $3, $4) RETURNING *;", 4, "Oku Izakaya Bar", "2 Water St, Vancouver", "Your destination to experience the rich tapestry of Japanese delicacies, from sushi to small plates and creative dishes. At OKU IZAKAYA, immerse yourself in an array of flavors that highlight the best of Japanese cuisine.")

	dbpool.QueryRow(c, "INSERT INTO users(username, email) VALUES ($1, $2) RETURNING *;", "OkuBar", "OkuBar@gmail.com")

	dbpool.QueryRow(c, "INSERT INTO owns(email, restaurantID) VALUES ($1, $2) RETURNING *;", "OkuBar@gmail.com", 4)

	dbpool.QueryRow(c, "INSERT INTO review(reviewID, email, restaurantID, comment, datetime) VALUES ($1, $2, $3, $4, $5) RETURNING *;", 400, "MI@gmail.com", 4, "Cozy place", parsedTime4)

}
