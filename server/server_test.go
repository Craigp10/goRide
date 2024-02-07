package server

import (
	pb "RouteRaydar/routeRaydar"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestRPC(t *testing.T) {
	go StartServer()

	// Set up a gRPC client to test the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewRouteServiceClient(conn)
	ctx := context.Background()

	t.Run("Test Request", func(t *testing.T) {
		res, err := client.SubmitGrid(ctx, &pb.SubmitGridRequest{
			Width:  10,
			Height: 10,
		})

		require.NoError(t, err)
		require.NotEmpty(t, res)
	})

	// t.Run("Test Send Coordinates", func(t *testing.T) {
	// 	res, err := client.SendCoordinates(ctx, &pb.SendCoordinatesRequest{
	// 		Start: &pb.Coordinates{
	// 			X: 4,
	// 			Y: 8,
	// 		},
	// 		End: &pb.Coordinates{
	// 			X: 2,
	// 			Y: 7,
	// 		},
	// 	})
	// 	require.NoError(t, err)
	// 	require.NotEmpty(t, res)
	// })

	t.Run("Test Send Coordinates", func(t *testing.T) {
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

	t.Run("Test Send Coordinates", func(t *testing.T) {
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
}
