# Welcome to Route Raydar

Route Raydar is an Demo applications aimed to mimic similar behavior to ride sharing apps. Currently the progress of the application is focused on the backend services side, with a focus on searching and infrastructure necessary to support it.
The application is currently in Version 1, which involves the general gRPC services that supports creating a search grid, discovering a route (Shortest path) from 2 coordinates, and streaming the route as a "ride".

Version 2 development is under way and aims to add persistent storage to the application and add a messaging service to support subscribing to active rides.

## Setup

Will be added after version 2...

## Run the application

# Build the go server binary

```
docker build -t route-raydar:1.0 .
```

# Run App within container Docker

```
docker run -d -p 50051:50051 -p 8080:8080 --name grpc_server route-raydar
```

# Run App separately

```
go run main.go
```

# Test the application

Will be added after version 2...
