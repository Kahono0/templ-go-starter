package main

import (
	"starter/db"
)

func main() {
	db.Connect("")

	db.DB.AutoMigrate(
		// add models to migrate here
	)
}
