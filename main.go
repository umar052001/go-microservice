package main

import (
	"context"
	"fmt"

	"github.com/umar052001/go-microservice.git/application"
)

func main() {
	app := application.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start the server %w", err)
	}
}
