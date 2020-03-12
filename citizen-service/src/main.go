package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/rizalhamdana/citizen-service/app"
	"github.com/rizalhamdana/citizen-service/config"
)

func main() {
	godotenv.Load(".env")
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)

	fmt.Println("Application is running on port 8080")
	app.Run(":8080")

}
