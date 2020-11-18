package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andanhm/anglebroking/pkg/platform/config"

	"github.com/andanhm/anglebroking/pkg/platform/db/redis"
	"github.com/andanhm/anglebroking/service"

	"github.com/andanhm/anglebroking/api"
	"github.com/andanhm/anglebroking/api/http"
	"github.com/bnkamalesh/webgo/v4"
)

func main() {
	webConfig := &webgo.Config{
		Port: "8081",
	}
	config, err := config.New()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	db, err := redis.New(&redis.Config{
		Host: "localhost",
		Port: 6379,
	})
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	service := service.New(config, db)
	apis := api.New(service)

	httpHandlerService := http.New(apis)
	routes := http.Routes(httpHandlerService)
	wgoRouter := webgo.NewRouter(webConfig, routes)

	wgoRouter.NotFound = httpHandlerService.NotFound
	log.Println(
		"Info",
		"Application",
		"STARTED.SUCCESSFULLY",
		fmt.Sprintf("Application started listening on %s %s port", os.Getenv("HOST"), webConfig.Port),
		"",
		"",
		nil,
	)
	wgoRouter.Start()
}
