package v1

import (
	"context"
	"database/sql"
	v1 "github.com/amsokol/go-grpc-http-rest-microservice-tutorial/pkg/api/v1"
	"log"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type maltaBEServer struct {
	db *sql.DB
}

func NewMaltaServiceServer(db *sql.DB) v1.MaltaBEServer {
	return &maltaBEServer{db: db}
}

func (s *maltaBEServer) Calculate(ctx context.Context, req *v1.CalculateRequest) (*v1.CalculateResponse, error) {
	log.Print("Recieved Calculate. Sending Response...")
	return &v1.CalculateResponse{
		Api:   apiVersion,
		Success: true,
	}, nil
}