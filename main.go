package main

import (
	"context"
	"log"
	"net"

	pb "github.com/BalliAsghar/micro/product"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type productServiceServer struct {
	pb.UnimplementedProductServiceServer
}

// Colour variables
var (
	Success = color.New(color.FgHiGreen)
	Error   = color.New(color.FgHiRed, color.Bold)
	Info    = color.New(color.FgHiMagenta)
	Warn    = color.New(color.FgYellow)
)

// list of products
var products = []*pb.Product{
	{Id: "1", Name: "Product 1", Description: "Product 1 Description", Price: "100"},
	{Id: "2", Name: "Product 2", Description: "Product 2 Description", Price: "200"},
	{Id: "3", Name: "Product 3", Description: "Product 3 Description", Price: "300"},
}

func (s *productServiceServer) GetProduct(ctx context.Context, req *pb.ProductRequest) (*pb.Product, error) {
	Info.Println("GetProduct called with id: ", req.Id)
	// TODO: Implement GetProduct method
	for _, product := range products {
		if product.Id == req.Id {
			return product, nil
		}
	}
	return nil, nil
}

func (s *productServiceServer) GetProductList(ctx context.Context, req *pb.ProductListRequest) (*pb.ProductList, error) {
	Info.Println("GetProductList called")
	return &pb.ProductList{Products: products}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &productServiceServer{})
	reflection.Register(s)

	log.Printf("Starting gRPC server on port 8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
