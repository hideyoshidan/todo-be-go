package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
	"todo.com/config"
	"todo.com/infrastructure/database"
	pb "todo.com/proto/statusmaster"
	"todo.com/services/statusmaster/rpc"
)

func main() {
	log.Print("StatusMaster Server Starting")
	listenProt, err := net.Listen("tcp", fmt.Sprintf(":%v", os.Getenv("RPC_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		os.Exit(1)
	}

	log.Print("DB set up is started")
	db, err := connDB()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	grpc := grpc.NewServer()
	rpc := rpc.NewRPC(db)
	pb.RegisterStatusMasterServer(grpc, rpc)
	reflection.Register(grpc)

	if err := grpc.Serve(listenProt); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func connDB() (db *gorm.DB, err error) {
	return database.NewDBConn(
		database.NewWithNoOpt(
			config.DB_USER,
			config.DB_PASSWORD,
			config.DB_ADDR,
			config.DB_NAME,
			config.LOCATION,
		),
	)
}
