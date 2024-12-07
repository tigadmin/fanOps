package main

import (
	"github.com/your-username/fanOps/internal/database"
	"github.com/your-username/fanOps/internal/grpc"
	"log"
)

func main() {
	// Initialize PostgreSQL database connection
	_, err := database.InitDB()
	if err != nil {
		log.Fatalf("could not initialize database: %v", err)
	}

	// Start the gRPC server
	grpc.StartGRPCServer()
}
