package db

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
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

func InsertDummyData(c *gin.Context) (err error) {
	var row pgx.Row
	var restaurant Restaurant

	dbpool.QueryRow(c, "INSERT INTO users(username, email) VALUES ($1, $2);", "Michelin Inspector", "MI@gmail.com")

	parsedTime, err := time.Parse("2006-01-02 15:04:05", "2023-03-30 15:30:41")
	if err != nil {
		return err
	}

	row = dbpool.QueryRow(c, "INSERT INTO restaurants(restaurantID, name, location, description) VALUES (DEFAULT, $1, $2, $3) RETURNING *;", "Cactus Club", "Downtown Vancouver", "Cactus Club Cafe is your go-to casual dining restaurant. We offer the best in global cuisine using local, fresh ingredients served in a vibrant, contemporary setting.")

	err = row.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Location, &restaurant.Description)
	if err != nil {
		return err
	}

	dbpool.QueryRow(c, "INSERT INTO users(username, email) VALUES ($1, $2);", "CactusClub", "CactusClub@gmail.com")

	dbpool.QueryRow(c, "INSERT INTO owns(email, restaurantID) VALUES ($1, $2);", "CactusClub@gmail.com", restaurant.ID)

	dbpool.QueryRow(c, "INSERT INTO review(reviewID, email, restaurantID, comment, datetime) VALUES (DEFAULT, $1, $2, $3, $4);", "MI@gmail.com", restaurant.ID, "AMAZING!", parsedTime)

	parsedTime2, err := time.Parse("2006-01-02 15:04:05", "2021-11-22 14:10:20")
	if err != nil {
		return err
	}

	row = dbpool.QueryRow(c, "INSERT INTO restaurants(restaurantID, name, location, description) VALUES (DEFAULT, $1, $2, $3) RETURNING *;", "The Old Spaghetti Factory", "53 Water Street, Vancouver", "A customer favourite, known for good times and great food. This is our first location in Vancouver’s historic Gastown. Reservations recommended.")

	err = row.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Location, &restaurant.Description)
	if err != nil {
		return err
	}

	dbpool.QueryRow(c, "INSERT INTO users(username, email) VALUES ($1, $2);", "OldSpaghettiFactory", "OldFactory@gmail.com")

	dbpool.QueryRow(c, "INSERT INTO owns(email, restaurantID) VALUES ($1, $2);", "OldFactory@gmail.com", restaurant.ID)

	dbpool.QueryRow(c, "INSERT INTO review(reviewID, email, restaurantID, comment, datetime) VALUES (DEFAULT, $1, $2, $3, $4);", "MI@gmail.com", restaurant.ID, "Incredible food", parsedTime2)

	parsedTime3, err := time.Parse("2006-01-02 15:04:05", "2024-04-15 19:23:40")
	if err != nil {
		return err
	}

	row = dbpool.QueryRow(c, "INSERT INTO restaurants(restaurantID, name, location, description) VALUES (DEFAULT, $1, $2, $3) RETURNING *;", "Zefferelli's Spaghetti Joint", "1136 Robson St 2nd Floor, Vancouver", "Over looking busy Robson street, a 2nd floor room that offers a cozy and charming atmosphere with friendly and efficient service that mostly prides itself on home-style, flavorful Italian dishes and pasta. Offering portion sizes that even Mama would be proud of.")

	err = row.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Location, &restaurant.Description)
	if err != nil {
		return err
	}

	userRow := dbpool.QueryRow(c, "INSERT INTO users(username, email) VALUES ($1, $2) RETURNING *;", "Zefferelis", "Zefferelis@gmail.com")
	var user User
	err = userRow.Scan(&user.Username, &user.Email)
	if err != nil {
		return err
	}

	ownsRow := dbpool.QueryRow(c, "INSERT INTO owns(email, restaurantID) VALUES ($1, $2) RETURNING *;", "Zefferelis@gmail.com", restaurant.ID)
	var owns Own
	err = ownsRow.Scan(&owns.Email, &owns.RestaurantID)
	if err != nil {
		return err
	}

	reviewRow := dbpool.QueryRow(c, "INSERT INTO review(reviewID, email, restaurantID, comment, datetime) VALUES (DEFAULT, $1, $2, $3, $4) RETURNING *;", "MI@gmail.com", restaurant.ID, "good staff", parsedTime3)
	var review Review
	err = reviewRow.Scan(&review.ID, &review.Email, &review.RestaurantID, &review.Comment, &review.Datetime)
	if err != nil {
		return err
	}

	parsedTime4, err := time.Parse("2006-01-02 15:04:05", "2024-08-23 18:24:20")
	if err != nil {
		return err
	}

	row = dbpool.QueryRow(c, "INSERT INTO restaurants(restaurantID, name, location, description) VALUES (DEFAULT, $1, $2, $3) RETURNING *;", "Oku Izakaya Bar", "2 Water St, Vancouver", "Your destination to experience the rich tapestry of Japanese delicacies, from sushi to small plates and creative dishes. At OKU IZAKAYA, immerse yourself in an array of flavors that highlight the best of Japanese cuisine.")

	err = row.Scan(&restaurant.ID, &restaurant.Name, &restaurant.Location, &restaurant.Description)
	if err != nil {
		return err
	}

	userRow = dbpool.QueryRow(c, "INSERT INTO users(username, email) VALUES ($1, $2) RETURNING *;", "OkuBar", "OkuBar@gmail.com")
	err = userRow.Scan(&user.Username, &user.Email)
	if err != nil {
		return err
	}

	ownsRow = dbpool.QueryRow(c, "INSERT INTO owns(email, restaurantID) VALUES ($1, $2) RETURNING *;", "OkuBar@gmail.com", restaurant.ID)
	err = ownsRow.Scan(&owns.Email, &owns.RestaurantID)
	if err != nil {
		return err
	}

	reviewRow = dbpool.QueryRow(c, "INSERT INTO review(reviewID, email, restaurantID, comment, datetime) VALUES (DEFAULT, $1, $2, $3, $4) RETURNING *;", "MI@gmail.com", restaurant.ID, "Cozy place", parsedTime4)
	err = reviewRow.Scan(&review.ID, &review.Email, &review.RestaurantID, &review.Comment, &review.Datetime)
	if err != nil {
		return err
	}

	return nil
}
