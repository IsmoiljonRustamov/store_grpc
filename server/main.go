package main

import (
	"context"
	pb "grpc_todo/genproto/store"
	"grpc_todo/server/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StoreServer struct {
	pb.UnimplementedStoreServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":8013")

	if err != nil {
		log.Fatalf("Failed listening: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterStoreServiceServer(s, &StoreServer{})

	log.Printf("server listening: %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

}

func (s *StoreServer) CreateStore(ctx context.Context, in *pb.Store) (*pb.Store, error) {
	log.Printf("Received: %v", in.GetName())
	store, err := postgres.CreateStore(&pb.Store{
		Id:          1,
		Name:        in.Name,
		Description: in.Description,
		Addresses:   in.Addresses,
		IsOpen:      in.IsOpen,
	})
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (s *StoreServer) UpdateStore(ctx context.Context, in *pb.Store) (*emptypb.Empty, error) {
	err := postgres.UpdateStore(in)
	if err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
	return &emptypb.Empty{}, err
}

func (s *StoreServer) DeleteStore(ctx context.Context, in *pb.GetStoreRequest) (*emptypb.Empty, error) {
	err := postgres.DeleteStore(in.Id)
	if err != nil {
		log.Fatalf("Failed to delete from server: %v", err)
	}

	return &emptypb.Empty{}, err
}

func (s *StoreServer) GetStore(ctx context.Context, in *pb.GetStoreRequest) (*pb.Store, error) {
	store, err := postgres.GetStore(in.Id)
	if err != nil {
		log.Fatalf("Failed to get store from server: %v", err)
	}
	return store, nil
}
