package v1

import (
	"context"
	"database/sql"
	v1 "github.com/amsokol/go-grpc-http-rest-microservice-tutorial/pkg/api/v1"
	"log"
	"strings"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type maltaBEServer struct {
	db *sql.DB
	p Poker
}

func NewMaltaServiceServer(db *sql.DB) v1.MaltaBEServer {
	return &maltaBEServer{
		db: db,
		p: Poker{},
	}
}

func (s *maltaBEServer) Calculate(ctx context.Context, req *v1.CalculateRequest) (*v1.CalculateResponse, error) {
	log.Print("Calculating...")

	res1 := strings.Split(req.HackData, "\n")
	for _, game := range res1 {
		s.calculateGame(game)
	}

	return &v1.CalculateResponse{
		Api:   apiVersion,
		Data: &v1.Data{
			Data: "Hello world calculated",
		},
	}, nil
}

func (s *maltaBEServer) calculateGame(game string) (*v1.CalculateResponse, error) {
	res1 := strings.Split(game, "\n")
	for _, i2 := range res1 {
		s.p.compareHands(string(i2[:14]), string(i2[15:]))
	}

	return &v1.CalculateResponse{
		Api:   apiVersion,
		Data: &v1.Data{
			Data: "Hello world calculated",
		},
	}, nil
}
