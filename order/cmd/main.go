package main

import (
	"log"

	"github.com/vteja-code/microservices/order/config"
	"github.com/vteja-code/microservices/order/internal/adapters/db"
	"github.com/vteja-code/microservices/order/internal/adapters/grpc"
	"github.com/vteja-code/microservices/order/internal/application/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
