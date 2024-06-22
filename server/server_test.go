package server

import (
	"context"
	pb "goRide/proto"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

// go test -run TestRPCServerConnection -v
func TestRPCServerConnection(t *testing.T) {
	client := StartServer()

	t.Run("Test client", func(t *testing.T) {
		require.NotEmpty(t, client)
	})
}

// go test -run TestSubmitGrid -v
func TestSubmitGrid(t *testing.T) {
	client := StartServer()
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
		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	t.Run("Negative input Grid", func(t *testing.T) {
		res, err := client.SubmitGrid(ctx, &pb.SubmitGridRequest{
			Width:  1,
			Height: -5,
		})
		require.Error(t, err)
		require.Empty(t, res)
	})
}

// go test -run TestStreamRide -v
func TestStreamRide(t *testing.T) {
	client := StartServer()
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
	var routeId string
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
		require.NotEmpty(t, res.GetRouteId())
		routeId = res.GetRouteId()
	})

	t.Run("Stream Ride", func(t *testing.T) {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			t.Fatalf("failed to dial server: %v", err)
		}
		defer conn.Close()
		a := pb.NewRouteServiceClient(conn)
		req := &pb.StreamRideRequest{
			RouteId: routeId,
		}
		stream, err := a.StreamRide(ctx, req)
		// a := pb.routeServiceStreamRideClient{}
		// var stream pb.RouteService_StreamRideServer
		// err := client.StreamRide(req, stream)
		require.NoError(t, err)
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				// End of stream
				t.Log("Sever has closed the stream, end of stream")
				break
			}
			if err != nil {
				t.Errorf("Error receiving stream response: %v", err)
				break
			}
			t.Logf("Received stream response: %v", res)
		}
	})
}

func TestSendCoordinates(t *testing.T) {
	// Need create mock server instance
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
	var routeId string
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
		require.NotEmpty(t, res.GetRouteId())
		routeId = res.GetRouteId()
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
		// Todo: Need to add mock kafka instance, shouldn't actually connect.
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

	t.Run("Get Route", func(t *testing.T) {
		req := &pb.GetRouteRequest{
			RouteId: routeId,
		}
		res, err := client.GetRoute(ctx, req)
		require.NoError(t, err)
		require.NotEmpty(t, res)
		t.Log(res)
	})
}

// go test -run TestValidateKafkaClient -v
func TestValidateKafkaClient(t *testing.T) {
	// var client pb.RouteServiceClient
	rs := StartServer()

	// Set up a gRPC client to test the server
	// conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	// if err != nil {
	// 	t.Fatalf("failed to dial server: %v", err)
	// }
	// defer conn.Close()

	// client := pb.NewRouteServiceClient(conn)

	t.Run("init client", func(t *testing.T) {
		require.NotEmpty(t, rs.Kc)
	})

	t.Run("kafka health check", func(t *testing.T) {

	})
}
