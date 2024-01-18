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
	"strings"
	"time"

	pb "github.com/u1-btj/push_notifications/topic_pb2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

type Color struct {
	Status string `json:"status"`
	Base   struct {
		Hex struct {
			Value string
		} `json:"hex"`
		Rgb struct {
			Value string
		} `json:"rgb"`
		Hsl struct {
			Value string
		} `json:"hsl"`
	} `json:"base"`
	Error struct {
		Message string
	} `json:"error"`
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

func getColorInfo(name string) Color {
	url := "https://color.serialif.com/" + strings.ToLower(name)

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

	var color Color
	err = json.Unmarshal(body, &color)
	if err != nil {
		log.Fatal(err)
	}
	return color
}

func (s *server) StreamMeowFacts(req *pb.FactRequest, stream pb.TopicSelection_StreamMeowFactsServer) error {
	log.Printf("Received request to send %d cat facts every %d second for %d times", req.Count, req.Second, req.Limit)
	for i := 0; i < int(req.Limit); i++ {
		allData := getFactsData(req.Count)
		response := &pb.FactResponse{Facts: allData.Data}

		if err := stream.Send(response); err != nil {
			log.Printf("Error sending data to fact stream: %v", err)
			return err
		}

		time.Sleep(time.Duration(req.Second) * time.Second) // Delay per second for each request
	}
	return nil
}

func (s *server) StreamColorInfo(req *pb.ColorRequest, stream pb.TopicSelection_StreamColorInfoServer) error {
	log.Printf("Received request to send color information every %d second for total of %d colors. Colors : %v.", req.Second, len(req.Name), req.Name)
	for i := 0; i < len(req.Name); i++ {
		color := getColorInfo(req.Name[i])

		// Handle error status from API (when color name not found)
		if color.Status == "error" {
			msg := req.Name[i] + " got this error " + color.Error.Message
			return status.Error(codes.InvalidArgument, msg)
		}

		response := &pb.ColorResponse{Hex: color.Base.Hex.Value, Rgb: color.Base.Rgb.Value, Hsl: color.Base.Hsl.Value}

		if err := stream.Send(response); err != nil {
			log.Printf("Error sending data to color stream: %v", err)
			return err
		}

		time.Sleep(time.Duration(req.Second) * time.Second) // Delay per second for each request
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
