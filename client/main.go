package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc_todo/genproto/store"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Store struct {
	Id          int
	Name        string
	Description string
	Addresses   []string
	IsOpen      bool
}

const (
	ServerAddress = "localhost:8013"
)

func main() {
	conn, err := grpc.Dial(ServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	defer conn.Close()

	c := pb.NewStoreServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// store, err := c.CreateStore(ctx, &pb.Store{
	// 	Name:        "Name",
	// 	Description: "Description",
	// 	Addresses: []string{
	// 		"address-1",
	// 		"address-2",
	// 	},
	// 	IsOpen: true,
	// })
	// if err != nil {
	// 	log.Fatalf("Failed to createStrore: %v", err)
	// }
	// fmt.Println(store)

	// _, err = c.UpdateStore(ctx, &pb.Store{})
	// if err != nil {
	// 	log.Fatalf("Failed to updateStore: %v", err)
	// }

	// _, err = c.DeleteStore(ctx, &pb.GetStoreRequest{Id: 6})
	// if err != nil {
	// 	log.Fatalf("Failed to delete from client: %v", err)
	// }

	store, err := c.GetStore(ctx, &pb.GetStoreRequest{Id: 5})
	if err != nil {
		log.Fatalf("Failed to get store from client: %v", err)
	}
	fmt.Println(store)
	
}
