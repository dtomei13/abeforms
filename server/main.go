package main

import (
	"github.com/austinlhx/server/app"
	"github.com/austinlhx/server/database"
)

func main() {
	database.ConnectDB()
	app.StartApp()
}
