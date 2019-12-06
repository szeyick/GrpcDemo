package main

import (
	"context"
	"fmt"
	"github.com/szeyick/GrpcDemo/src/clients"
	"github.com/szeyick/GrpcDemo/src/models"
	"go.uber.org/fx"
	"log"
	"time"
)

func Sum(x int, y int) int {
	return x + y
}

func main() {
	app := fx.New(
		clients.In,
		fx.Invoke(Register),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}


	m := &models.Impl{}
	fmt.Println(m.Bar(5))

	stopCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}

func Register() {
	fmt.Println(Sum(5, 5))
}
