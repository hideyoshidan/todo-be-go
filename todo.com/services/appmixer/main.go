package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"todo.com/lib/connection"
	pb "todo.com/proto/appmixer"
	"todo.com/proto/statusmaster"
	"todo.com/services/appmixer/rpc"
)

func main() {
	log.Print("Appmixer gRPC Server Starting")
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", os.Getenv("RPC_PORT")))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	sConn := initConnectToStatusMaster()
	defer sConn.Close()

	// Create a gRPC server object
	rpc := rpc.NewRPC(statusmaster.NewStatusMasterClient(sConn))
	s := grpc.NewServer()
	pb.RegisterAppmixerServer(s, rpc)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("0.0.0.0:%v", os.Getenv("RPC_PORT")),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	log.Print("Appmixer grpc server waiting proxy...")

	mux := runtime.NewServeMux()
	err = pb.RegisterAppmixerHandler(context.Background(), mux, conn)
	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", os.Getenv("GATEWAY_PORT")),
		Handler: mux,
	}

	log.Print("Appmixer grpc-gateway waiting request...")
	log.Fatalln(gwServer.ListenAndServe())
}

func initConnectToStatusMaster() *grpc.ClientConn {
	conn := connection.NewConnectRPC("app-statusmaster", os.Getenv("STATUSMASTER_PORT"))
	sConn, err := conn.InitConnection()
	if err != nil {
		log.Fatalln("Cannot connect with StatusMaster RPC Server:", err)
	}
	return sConn
}
