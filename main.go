package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/umar052001/go-microservice.git/application"
)

func main() {
	app := application.New()

	ctx, cancle := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancle()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("Failed to start the server %w", err)
	}

}
