package main

import (
	"./app"
	"./database"
)

func main() {
	database.ConnectDB()
	app.StartApp()
}
