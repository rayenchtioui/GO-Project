package main

import (
	"go-project/data"
	"go-project/pkg/database"
)

func main() {
	db := database.GetDB()
	database.SetupDatabase(db)
	data.GenerateAndWriteData()
	database.InsertData()

	database.CloseDB()
}
