package main

import (
	"example.com/go/dependency"
	"example.com/go/pkg/database"
	"example.com/go/tools"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Authorization,Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	app.Use(tools.Logger)
	tools.DistributionImages(app)

	redis, err := ConnectRedis()
	if err != nil {
		panic(err)
	}

	db, err := database.NewDataBaseConnection()
	if err != nil {
		panic("Connection failed")
	}

	defer func(db *gorm.DB) {
		err := database.CloseConnection(db)
		if err != nil {
			panic(err)
		}
	}(db)

	database.AutoMigrate(db)
	dependency.SettingDepInjection(app, db, redis)

	server := app.Listen("192.168.0.105:5000")
	if server != nil {
		panic(server)
	}
	fmt.Println("Server started")
	fmt.Println("DB connection")
	fmt.Println("Redis connection")
}
