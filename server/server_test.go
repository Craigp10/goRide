package server

import (
	pb "RouteRaydar/routeRaydar"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

// go test -run TestRPCServerConnection -v
func TestRPCServerConnection(t *testing.T) {
	go StartServer()

	// Set up a gRPC client to test the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewRouteServiceClient(conn)

	t.Run("Test client", func(t *testing.T) {
		require.NotEmpty(t, client)
	})
}

// go test -run TestSubmitGrid -v
func TestSubmitGrid(t *testing.T) {
	go StartServer()

	// Set up a gRPC client to test the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewRouteServiceClient(conn)
	ctx := context.Background()

	t.Run("Standard Grid", func(t *testing.T) {
		res, err := client.SubmitGrid(ctx, &pb.SubmitGridRequest{
			Width:  10,
			Height: 10,
		})

		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	t.Run("Standard Grid", func(t *testing.T) {
		res, err := client.SubmitGrid(ctx, &pb.SubmitGridRequest{
			Width:  1,
			Height: 5,
		})
		t.Log(res)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	t.Run("Negative input Grid", func(t *testing.T) {
		res, err := client.SubmitGrid(ctx, &pb.SubmitGridRequest{
			Width:  1,
			Height: -5,
		})
		t.Log(res)
		require.Error(t, err)
	})
}

// go test -run TestSendCoordinates -v
func TestSendCoordinates(t *testing.T) {
	go StartServer()

	// Set up a gRPC client to test the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewRouteServiceClient(conn)
	ctx := context.Background()
	const GRID_WIDTH = 10
	const GRID_HEIGHT = 10

	t.Run("Test Request", func(t *testing.T) {
		res, err := client.SubmitGrid(ctx, &pb.SubmitGridRequest{
			Width:  GRID_WIDTH,
			Height: GRID_HEIGHT,
		})

		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	t.Run("Send Standard Coordinates", func(t *testing.T) {
		res, err := client.SendCoordinates(ctx, &pb.SendCoordinatesRequest{
			Start: &pb.Coordinates{
				X: 2,
				Y: 4,
			},
			End: &pb.Coordinates{
				X: 8,
				Y: 9,
			},
		})
		require.NoError(t, err)
		require.NotEmpty(t, res)
		require.Equal(t, int64(11), res.GetDistance())
	})

	t.Run("Send Standard Coordinates", func(t *testing.T) {
		res, err := client.SendCoordinates(ctx, &pb.SendCoordinatesRequest{
			Start: &pb.Coordinates{
				X: 2,
				Y: 4,
			},
			End: &pb.Coordinates{
				X: 0,
				Y: 0,
			},
		})
		require.NoError(t, err)
		require.NotEmpty(t, res)
		require.Equal(t, int64(6), res.GetDistance())
	})

	t.Run("Send Same Coordinates", func(t *testing.T) {
		res, err := client.SendCoordinates(ctx, &pb.SendCoordinatesRequest{
			Start: &pb.Coordinates{
				X: 2,
				Y: 4,
			},
			End: &pb.Coordinates{
				X: 2,
				Y: 4,
			},
		})
		require.NoError(t, err)
		require.NotEmpty(t, res)
		t.Log(res)
		require.Equal(t, int64(0), res.GetDistance())
		require.Equal(t, 0, len(res.GetRoute()))
	})

	t.Run("Send max distance Coordinates", func(t *testing.T) {
		const GRID_HEIGHT_coord = GRID_HEIGHT - 1
		const GRID_WIDTH_coord = GRID_WIDTH - 1

		res, err := client.SendCoordinates(ctx, &pb.SendCoordinatesRequest{
			Start: &pb.Coordinates{
				X: 0,
				Y: 0,
			},
			End: &pb.Coordinates{
				X: GRID_HEIGHT_coord,
				Y: GRID_WIDTH_coord,
			},
		})
		require.NoError(t, err)
		require.NotEmpty(t, res)
		require.Equal(t, int64(GRID_HEIGHT_coord+GRID_WIDTH_coord), res.GetDistance())
		require.Equal(t, GRID_HEIGHT_coord+GRID_WIDTH_coord+1, len(res.GetRoute()))
	})

	t.Run("Send off grid Coordinates", func(t *testing.T) {
		res, err := client.SendCoordinates(ctx, &pb.SendCoordinatesRequest{
			Start: &pb.Coordinates{
				X: 0,
				Y: 0,
			},
			End: &pb.Coordinates{
				X: 10,
				Y: -4,
			},
		})
		require.Error(t, err)
		require.Empty(t, res)
	})
}