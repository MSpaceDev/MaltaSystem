package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/amsokol/go-grpc-http-rest-microservice-tutorial/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

func main() {
	// get configuration
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewMaltaBEClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Ping
	req1 := v1.CalculateRequest{
		Api: apiVersion,
		HackData: "8C TS KC 9H 4S 7D 2S 5D 3S AC\n5C AD 5D AC 9C 7C 5H 8D TD KS\n3H 7H 6S KC JS QH TD JC 2D 8S\nTH 8H 5C QS TC 9H 4D JC KS JS",
		}

	res1, _ := c.Calculate(ctx, &req1)

	log.Printf("Calculated: %v", res1.Data)
}