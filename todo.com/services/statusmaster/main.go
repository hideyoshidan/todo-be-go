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
		log.Fatalln(fmt.Sprintf("failed to listen: %v", err))
	}

	log.Print("DB set up is started")
	db, err := connDB()
	if err != nil {
		log.Fatalln(err)
	}

	grpc := grpc.NewServer()
	rpc := rpc.NewRPC(db)
	pb.RegisterStatusMasterServer(grpc, rpc)
	reflection.Register(grpc)

	if err := grpc.Serve(listenProt); err != nil {
		log.Fatalln(err)
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
