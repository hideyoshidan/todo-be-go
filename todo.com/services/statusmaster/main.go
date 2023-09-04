package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "todo.com/proto/statusmaster"
	"todo.com/services/statusmaster/rpc"
)

func main() {
	log.Print("StatusMaster Server Starting")

	listenProt, err := net.Listen("tcp", fmt.Sprintf(":%v", os.Getenv("RPC_PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpc := grpc.NewServer()
	rpc := rpc.NewRPC()
	pb.RegisterStatusMasterServer(grpc, rpc)

	reflection.Register(grpc)

	if err := grpc.Serve(listenProt); err != nil {
		log.Fatal(err)
	}
}
