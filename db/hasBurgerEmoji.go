package db

import (
	"github.com/gin-gonic/gin"
)

type HasBurgerEmoji struct {
	Username       string `json:"username"`
	HasBurgerEmoji bool   `json:"hasBurgerEmoji"`
}

func CreateHasBurgerEmoji(hasBurgerEmoji HasBurgerEmoji, c *gin.Context) (*HasBurgerEmoji, error) {
	var newHasBurgerEmoji HasBurgerEmoji
	row := dbpool.QueryRow(
		c,
		"INSERT INTO hasBurgerEmoji(username, hasBurgerEmoji) VALUES ($1,$2) RETURNING *;",
		hasBurgerEmoji.Username,
		hasBurgerEmoji.HasBurgerEmoji,
	)
	err := row.Scan(
		&newHasBurgerEmoji.Username,
		&newHasBurgerEmoji.HasBurgerEmoji,
	)
	if err != nil {
		return nil, err
	}
	return &newHasBurgerEmoji, nil
}

func GetHasBurgerEmoji(username string, c *gin.Context) (*HasBurgerEmoji, error) {
	var hasBurgerEmoji HasBurgerEmoji
	row := dbpool.QueryRow(c, "SELECT * FROM hasBurgerEmoji WHERE name = $1;", username)
	err := row.Scan(
		&hasBurgerEmoji.Username,
		&hasBurgerEmoji.HasBurgerEmoji,
	)
	if err != nil {
		return nil, err
	}
	return &hasBurgerEmoji, nil
}

func DeleteHasBurgerEmoji(name string, c *gin.Context) (*HasBurgerEmoji, error) {
	var deletedHasBurgerEmoji HasBurgerEmoji
	row := dbpool.QueryRow(
		c,
		"DELETE FROM hasBurgerEmoji WHERE name = $1 RETURNING *;",
		name,
	)
	err := row.Scan(
		&deletedHasBurgerEmoji.Username,
		&deletedHasBurgerEmoji.HasBurgerEmoji,
	)
	if err != nil {
		return nil, err
	}
	return &deletedHasBurgerEmoji, nil
}

func UpdateHasBurgerEmoji(replaceWith HasBurgerEmoji, username string, c *gin.Context) (*HasBurgerEmoji, error) {
	var hasBurgerEmoji HasBurgerEmoji
	row := dbpool.QueryRow(
		c,
		"UPDATE hasBurgerEmoji SET hasBurgerEmoji=$1 where name=$2 RETURNING *;",
		replaceWith.HasBurgerEmoji,
		username,
	)
	err := row.Scan(
		&hasBurgerEmoji.Username,
		&hasBurgerEmoji.HasBurgerEmoji,
	)
	if err != nil {
		return nil, err
	}
	return &hasBurgerEmoji, nil
}
