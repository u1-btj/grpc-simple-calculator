package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/u1-btj/grpc-simple-calculator/calculator_protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultNum = 1
)

var (
	addr = flag.String("addr", "localhost:8080", "Address To Connect")
	num1 = flag.Float64("num1", defaultNum, "Num 1")
	num2 = flag.Float64("num2", defaultNum, "Num 2")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did Not Connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewCalculationClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	add, _ := c.Addition(ctx, &pb.NumRequest{NumA: *num1, NumB: *num2})
	substract, _ := c.Substraction(ctx, &pb.NumRequest{NumA: *num1, NumB: *num2})
	multiply, _ := c.Multiplication(ctx, &pb.NumRequest{NumA: *num1, NumB: *num2})
	div, err := c.Division(ctx, &pb.NumRequest{NumA: *num1, NumB: *num2})

	log.Printf("Addition: %v", add.GetResult())
	log.Printf("Substraction: %v", substract.GetResult())
	log.Printf("Multiplication: %v", multiply.GetResult())
	if err != nil {
		log.Fatalf("Division Error: %v", err)
	}
	log.Printf("Division: %v", div.GetResult())
}
