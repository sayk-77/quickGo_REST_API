package main

import (
	"example.com/go/dependency"
	"example.com/go/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db, err := database.NewDataBaseConnection()
	if err != nil {
		panic("Connection failed")
	}

	defer database.CloseConnection(db)

	database.AutoMigrate(db)
	dependency.SettingDepInjection(app, db)

	app.Listen(":4000")
}
