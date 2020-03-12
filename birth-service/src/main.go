package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/rizalhamdana/birth-service/app"
)

func main() {
	godotenv.Load(".env")
	app := &app.App{}
	app.Initialize()
	fmt.Println("Application is running on port 8081")
	app.Run(":8081")
}
