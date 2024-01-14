package main

import (
	"context"
	"fmt"
	logProto "github.com/tom-blog-app/blog-proto/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"time"
)

var logServiceClient logProto.LogServiceClient

func init() {
	conn, err := grpc.Dial("log-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		fmt.Println("Error connecting to log service:", err)
		return
	}
	logServiceClient = logProto.NewLogServiceClient(conn)
}
func LogViaGrpc(name string, data string) error {
	log.Println("Logging via gRPC...")
	conn, err := grpc.Dial("log-service:50001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println("Error connecting to log service:", err)
		return err
	}
	logServiceClient = logProto.NewLogServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	log.Println("Sending log to log service...")
	_, err = logServiceClient.CreateLog(ctx, &logProto.LogRequest{
		Log: &logProto.Log{
			Name:    name,
			Content: data,
		},
	})

	if err != nil {
		log.Println("Log sent to log service!" + err.Error())
		return fmt.Errorf("Error connecting to log service: %v", err)
	}

	fmt.Println(http.StatusOK, "logged!")
	return nil
}
