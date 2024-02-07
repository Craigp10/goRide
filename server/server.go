package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	pb "RouteRaydar/routeRaydar"

	"google.golang.org/grpc"
)

// server is used to implement routeRaydar.RouteServiceServer
type RouteRaydarServer struct {
	pb.RouteServiceServer
	grid Matrix
	mu   sync.Mutex // protects routeNotes
}

// GetRoute implements routeRaydar.RouteServiceServer
func (s *RouteRaydarServer) GetRoute(ctx context.Context, req *pb.GetRouteRequest) (*pb.GetRouteResponse, error) {
	// Implement logic to retrieve route based on request parameters
	routeID := req.GetRouteId()

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

	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
	// Start gRPC server via separate go Routine (so that a http server can run in the current routine)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, RouteRaydar!\n")
	})

	log.Println("Starting routeRaydar HTTP server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}

}

func (s *RouteRaydarServer) storeGrid(rows, cols int64) (Matrix, error) {
	grid := *NewMatrix(rows, cols)
	s.grid = grid
	return grid, nil
}

// Implement the SubmitGrid method of the RouteRaydarServer interface
func (s *RouteRaydarServer) SubmitGrid(ctx context.Context, req *pb.SubmitGridRequest) (*pb.SubmitGridResponse, error) {
	// Your implementation for SubmitGrid
	height, width := req.GetHeight(), req.GetWidth()

	m, err := s.storeGrid(height, width)
	if err != nil {
		return nil, err
	}
	grid := m.ToProto()
	return &pb.SubmitGridResponse{Grid: grid}, nil
}

// Implement the SendNewPoints method of the RouteRaydarServer interface
func (s *RouteRaydarServer) SendCoordinates(ctx context.Context, req *pb.SendCoordinatesRequest) (*pb.SendCoordinatesResponse, error) {
	// Your implementation for SendNewPoints
	st := req.GetStart()
	ed := req.GetEnd()

	// Validate the coordinates are within the grid is within the
	if coordinatesOutOfPlane(s.grid, st) || coordinatesOutOfPlane(s.grid, ed) {
		return nil, fmt.Errorf("the provided start or end coordinate is not on the grid plane.")
	}

	// Else calculate route
	visited := make(map[string]bool)
	dist := make(map[*pb.Coordinates]int)
	dist[st] = 0
	queue := []*pb.Coordinates{st}
	queue = append(queue, st)
	dirs := [][]int64{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var res int
	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]
		// popped, _ := dequeue(queue)
		if popped.GetX() == ed.X && popped.GetY() == ed.Y {
			res = dist[popped]
			break
		}
		if popped == nil {
			// No path found.
			break
		}
		if visited[popped.String()] {
			continue
		}
		for _, dir := range dirs {
			dX, dY := popped.GetX()+dir[0], popped.GetY()+dir[1]
			// next := Coordinates{curr.X + dir.X, curr.Y + dir.Y}
			newCoord := &pb.Coordinates{
				X: dX,
				Y: dY,
			}
			if coordinatesOutOfPlane(s.grid, newCoord) || visited[newCoord.String()] == true {
				continue
			}
			visited[popped.String()] = true
			dist[newCoord] = dist[popped] + 1
			queue = append(queue, newCoord)
			// Need to make a break condition and record the route....
		}
	}

	// path := []*pb.Coordinates{ed}
	// current := ed
	// for current != st {
	// 	// current = dist[current.String()]
	// 	path = append(path, current)
	// }
	// fmt.Println(res, dist)
	return &pb.SendCoordinatesResponse{
		Distance: int64(res),
	}, nil
}

func dequeue(queue []*pb.Coordinates) (*pb.Coordinates, []*pb.Coordinates) {
	if len(queue) == 0 {
		return nil, queue // Return 0 or handle empty queue case
	}
	first := queue[0]
	queue = queue[1:] // Remove the first element
	return first, queue
}

func coordinatesOutOfPlane(grid Matrix, coord *pb.Coordinates) bool {
	if coord.X < 0 || coord.Y < 0 || coord.X >= grid.Rows || coord.Y >= grid.Rows {
		return true
	}
	return false
}
