package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/rizalhamdana/family-service/app"
)

func main() {
	godotenv.Load(".env")
	app := &app.App{}
	app.Initialize()
	fmt.Println("Application is running on port 8082")
	app.Run(":8082")
}
