package main

import (
	"context"
	pb "github.com/mpedrozoduran/gogrpctest/timeproto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:9091", opts...)
	if err != nil {
		log.Fatalf("Fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewTimeManagerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	res, err := client.GetTime(ctx, &pb.TimeRequest{ClientTime:"2020-04-23 14:47:00"})
	if err != nil {
		log.Printf("Error when calling GetTime: %v", err)
	}
	log.Println("GetTime response: ", res)
}
