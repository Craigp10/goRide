package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "RouteRaydar/routeRaydar"

	"google.golang.org/grpc"
)

// server is used to implement routeRaydar.RouteServiceServer
type server struct{}

type RouteRaydarServer struct {
	pb.RouteServiceServer

	mu sync.Mutex // protects routeNotes
}

// GetRoute implements routeRaydar.RouteServiceServer
func (s *server) GetRoute(ctx context.Context, req *pb.GetRouteRequest) (*pb.GetRouteResponse, error) {
	// Implement logic to retrieve route based on request parameters
	routeID := req.GetRouteId()
	// Your logic here...

	// Return the route in the response
	return &pb.GetRouteResponse{
		RouteId: routeID,
		// Populate other fields as needed
	}, nil
}

// StartServer starts the gRPC server
func StartServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRouteServiceServer(s, &RouteRaydarServer{})
	log.Println("Starting routeRaydar gRPC server on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
