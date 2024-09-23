package main

import (
	"server-api/db"
	"server-api/router"
)

func main() {
	db.InitPostgresDB()
	r := router.InitRouter()
	r.Run(":8080")
	// r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}
