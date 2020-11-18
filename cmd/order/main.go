package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/andanhm/anglebroking/api"
	internal "github.com/andanhm/anglebroking/api/grpc"
	v1 "github.com/andanhm/anglebroking/grpc/anglebroking/stock/order/v1"
	"github.com/andanhm/anglebroking/pkg/platform/config"
	"github.com/andanhm/anglebroking/pkg/platform/db/mysql"
	"github.com/andanhm/anglebroking/service"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	db, err := mysql.New(&mysql.Config{
		Host:     "localhost",
		Username: "root",
		Password: "1234",
		DBName:   "anglebrokerdb",
		Port:     3306,
	})
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	service := service.NewOrder(config, db.Client)
	apis := api.New(service)

	grpcServer := grpc.NewServer()
	gRPCService := internal.NewOrder(apis)
	v1.RegisterOrderServer(grpcServer, gRPCService)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", "5002"))
	if err != nil {
		log.Println(
			"Fatal",
			"Application",
			"TCP.CONNECTION.ERROR",
			err.Error(),
			"ConnectionError",
			"",
			nil,
			true,
		)
		return
	}
	if err := grpcServer.Serve(listener); err != nil {
		log.Println(
			"Fatal",
			"Application",
			"GRPC.SERVER.ERROR",
			err.Error(),
			"ConnectionError",
			"",
			nil,
			true,
		)
	}
}
