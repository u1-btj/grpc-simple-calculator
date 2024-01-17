package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	pb "github.com/u1-btj/push_notifications/topic_pb2"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8080, "Server Port")
)

type server struct {
	pb.UnimplementedTopicSelectionServer
}

type FactsData struct {
	Data []string `json:"data"`
}

func getFactsData(n int32) FactsData {
	count := strconv.Itoa(int(n))
	url := "https://meowfacts.herokuapp.com/?count=" + count

	// Create an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data FactsData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (s *server) StreamMeowFacts(req *pb.FactRequest, stream pb.TopicSelection_StreamMeowFactsServer) error {
	log.Printf("Received request to send %d cat facts every %d second for %d times", req.Count, req.Second, req.Limit)
	for i := 0; i < int(req.Limit); i++ {
		allData := getFactsData(req.Count)
		response := &pb.FactResponse{Facts: allData.Data}

		if err := stream.Send(response); err != nil {
			log.Printf("Error sending data to stream: %v", err)
			return err
		}

		time.Sleep(time.Duration(req.Second) * time.Second) // Simulating delay for each day's forecast.
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTopicSelectionServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
