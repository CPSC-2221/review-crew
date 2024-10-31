package db

import (
	"github.com/gin-gonic/gin"
)

type HasPizzaImage struct {
	Name          string `json:"name"`
	HasPizzaImage bool   `json:"hasPizzaImage"`
}

func CreateHasPizzaImage(hasPizzaImage HasPizzaImage, c *gin.Context) (*HasPizzaImage, error) {
	var newHasPizzaImage HasPizzaImage
	row := dbpool.QueryRow(
		c,
		"INSERT INTO hasPizzaImage(name, hasPizzaImage) VALUES ($1,$2) RETURNING *;",
		hasPizzaImage.Name,
		hasPizzaImage.HasPizzaImage,
	)
	err := row.Scan(
		&newHasPizzaImage.Name,
		&newHasPizzaImage.HasPizzaImage,
	)
	if err != nil {
		return nil, err
	}
	return &newHasPizzaImage, nil
}

func GetHasPizzaImage(name string, c *gin.Context) (*HasPizzaImage, error) {
	var hasPizzaImage HasPizzaImage
	row := dbpool.QueryRow(c, "SELECT * FROM hasPizzaImage WHERE name = $1;", name)
	err := row.Scan(
		&hasPizzaImage.Name,
		&hasPizzaImage.HasPizzaImage,
	)
	if err != nil {
		return nil, err
	}
	return &hasPizzaImage, nil
}

func DeleteHasPizzaImage(name string, c *gin.Context) (*HasPizzaImage, error) {
	var deletedHasPizzaImage HasPizzaImage
	row := dbpool.QueryRow(
		c,
		"DELETE FROM hasPizzaImage WHERE name = $1 RETURNING *;",
		name,
	)
	err := row.Scan(
		&deletedHasPizzaImage.Name,
		&deletedHasPizzaImage.HasPizzaImage,
	)
	if err != nil {
		return nil, err
	}
	return &deletedHasPizzaImage, nil
}

func UpdateHasPizzaImage(replaceWith HasPizzaImage, name string, c *gin.Context) (*HasPizzaImage, error) {
	var hasPizzaImage HasPizzaImage
	row := dbpool.QueryRow(
		c,
		"UPDATE hasPizzaImage SET hasPizzaImage=$1 where name=$2 RETURNING *;",
		replaceWith.HasPizzaImage,
		name,
	)
	err := row.Scan(
		&hasPizzaImage.Name,
		&hasPizzaImage.HasPizzaImage,
	)
	if err != nil {
		return nil, err
	}
	return &hasPizzaImage, nil
}
