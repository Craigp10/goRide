package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	pb "RouteRaydar/proto"

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
	if height < 0 || width < 0 {
		return nil, fmt.Errorf("invalid grid input - negative plane")
	}
	m, err := s.storeGrid(height, width)
	if err != nil {
		return nil, err
	}
	grid := m.ToProto()
	return &pb.SubmitGridResponse{Grid: grid}, nil
}

func search(grid Matrix, st, ed *pb.Coordinates) *pb.SendCoordinatesResponse {
	// If start and end coordinates are same, return 0 distance and empty route -- May remove
	if st.GetX() == ed.GetX() && st.GetY() == ed.GetY() {
		return &pb.SendCoordinatesResponse{
			Distance: 0,
			Route:    []*pb.Coordinates{},
		}
	}

	// Initialize structures to manage the search

	// Visited - maintain a map of 'seen' coordinates to avoid duplicate work
	visited := make(map[string]bool)

	// Counted distance of movements made in the path
	dist := make(map[*pb.Coordinates]int)

	// Adjancy List to track paths during the search -- used to reverse the shortest path and return it
	path := make(map[string]*pb.Coordinates)

	// Queue to execute the BFS search.
	queue := []*pb.Coordinates{st}

	// Potential directions allow to search in.
	dirs := [][]int64{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// Set initial values to begin search with
	dist[st] = 0
	path[st.String()] = st
	queue = append(queue, st)

	var shortestPathDistance int
	var shortestRoute []*pb.Coordinates

	for len(queue) > 0 {

		popped := queue[0]
		queue = queue[1:]

		// Base case -- search has reached the target coordinate, return.
		if popped.GetX() == ed.X && popped.GetY() == ed.Y {
			shortestPathDistance = dist[popped]
			shortestRoute = reconstructPath(st, ed, path)
			return &pb.SendCoordinatesResponse{
				Route:    shortestRoute,
				Distance: int64(shortestPathDistance),
			}
		}

		// Current coordinate has been seen before, skip.
		if visited[popped.String()] == true {
			continue
		}
		visited[popped.String()] = true

		// iterate for each direction from current coordinate.
		for _, dir := range dirs {
			dX, dY := popped.GetX()+dir[0], popped.GetY()+dir[1]
			newCoord := &pb.Coordinates{
				X: dX,
				Y: dY,
			}

			if validCoordinates(grid, newCoord) != true || visited[newCoord.String()] == true {
				continue
			}

			// Manage structures with current coordinate directions
			path[newCoord.String()] = popped
			dist[newCoord] = dist[popped] + 1
			queue = append(queue, newCoord)
		}
	}
	return nil
}

// Implement the SendNewPoints method of the RouteRaydarServer interface.
func (s *RouteRaydarServer) SendCoordinates(ctx context.Context, req *pb.SendCoordinatesRequest) (*pb.SendCoordinatesResponse, error) {
	st := req.GetStart()
	ed := req.GetEnd()

	if !validCoordinates(s.grid, st) || !validCoordinates(s.grid, ed) {
		return nil, fmt.Errorf("the provided start or end coordinate is not on the grid plane.")
	}

	// Begin searching the grid
	res := search(s.grid, st, ed)
	if res == nil {
		return nil, fmt.Errorf("Path not found")
	}

	return res, nil
}

// Re construct the shortest path utilizing the path adjancey matrix.
func reconstructPath(start, end *pb.Coordinates, path map[string]*pb.Coordinates) []*pb.Coordinates {
	var route []*pb.Coordinates
	// Start from end (work backwards through the route).
	current := end

	// While current is not equal to start, iterate.
	for current != start {
		// append the current coordinate at the beginning of the route.
		route = append([]*pb.Coordinates{current}, route...)
		current = path[current.String()]
		// If we've reached the start then append the start and end iteration.
		if current == start {
			route = append([]*pb.Coordinates{start}, route...)
		}
	}

	return route
}

// TODO: Implement proirity Queue to improve performance on search.
// func dequeue(queue []*pb.Coordinates) (*pb.Coordinates, []*pb.Coordinates) {
// 	if len(queue) == 0 {
// 		return nil, queue // Return 0 or handle empty queue case
// 	}
// 	first := queue[0]
// 	queue = queue[1:] // Remove the first element
// 	return first, queue
// }

// validCoordinates validates that the provided coordinates are valid on the provided matrix.
func validCoordinates(grid Matrix, coord *pb.Coordinates) bool {
	if coord.X < 0 || coord.Y < 0 || coord.X >= grid.Rows || coord.Y >= grid.Rows {
		return false
	}
	return true
}
