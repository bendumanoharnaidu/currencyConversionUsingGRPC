package main

import (
	currency_conversion "currencyConversion/currencyconversion"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func main() {
	// Start gRPC server
	go func() {
		if err := createGrpcServer(); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	log.Println("Received termination signal. Shutting down servers...")
}

func createGrpcServer() error {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	converterServer := currency_conversion.Server{}
	currency_conversion.RegisterConverterServer(grpcServer, &converterServer)

	log.Println("gRPC server started on port 9090")
	return grpcServer.Serve(lis)
}
