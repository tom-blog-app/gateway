package main

import (
	"context"
	"fmt"
	logs "github.com/tom-blog-app/gataway/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"time"
)

func LogViaGrpc(name string, data string) {
	conn, err := grpc.Dial("log-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		fmt.Println("Error connecting to log service:", err)
		//c.JSON(http.StatusInternalServerError, err)
		return
	}

	defer conn.Close()

	logServiceClient := logs.NewLogServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = logServiceClient.WriteLog(ctx, &logs.LogRequest{
		LogEntry: &logs.Log{
			Name: name,
			Data: data,
		},
	})

	if err != nil {
		fmt.Println("Error connecting to log service:", err)
		//c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(http.StatusOK, "logged!")
	//c.JSON(http.StatusOK, "logged!")
}
