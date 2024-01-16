package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/u1-btj/grpc-simple-calculator/calculator_protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 8080, "Server Port")
)

// Implement calculator_protoc.CalculationServer
type server struct {
	pb.UnimplementedCalculationServer
}

// Printing number info on server
func print_num_info(in *pb.NumRequest) {
	log.Printf("Num 1: %v", in.GetNumA())
	log.Printf("Num 2: %v", in.GetNumB())
}

// Implements rpc CalculationServer.Addition
func (s *server) Addition(ctx context.Context, in *pb.NumRequest) (*pb.NumReply, error) {
	num_a := in.GetNumA()
	num_b := in.GetNumB()
	result := num_a + num_b
	print_num_info(in)
	log.Printf("Addition: %v", result)
	return &pb.NumReply{Result: result}, nil
}

// Implements rpc CalculationServer.Substraction
func (s *server) Substraction(ctx context.Context, in *pb.NumRequest) (*pb.NumReply, error) {
	num_a := in.GetNumA()
	num_b := in.GetNumB()
	result := num_a - num_b
	log.Printf("Substraction: %v", result)
	return &pb.NumReply{Result: result}, nil
}

// Implements rpc CalculationServer.Multiplication
func (s *server) Multiplication(ctx context.Context, in *pb.NumRequest) (*pb.NumReply, error) {
	num_a := in.GetNumA()
	num_b := in.GetNumB()
	result := num_a * num_b
	log.Printf("Multiplication: %v", result)
	return &pb.NumReply{Result: result}, nil
}

// Implements rpc CalculationServer.Division
func (s *server) Division(ctx context.Context, in *pb.NumRequest) (*pb.NumReply, error) {
	num_a := in.GetNumA()
	num_b := in.GetNumB()

	// Handle zero division
	if num_b == 0 {
		err := status.Error(codes.InvalidArgument, "Cannot divided by zero")
		log.Printf("Division Error: %v", err)
		return &pb.NumReply{Result: 0}, err
	}

	result := num_a / num_b
	log.Printf("Division: %v", result)
	return &pb.NumReply{Result: result}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculationServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
