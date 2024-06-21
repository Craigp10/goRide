# Welcome to Go Rides

Route Raydar is an Demo applications aimed to mimic similar behavior to ride sharing apps. Currently the progress of the application is focused on the backend services side, with a focus on searching and infrastructure necessary to support it.
The application is currently in Version 1, which involves the general gRPC services that supports creating a search grid, discovering a route (Shortest path) from 2 coordinates, and streaming the route as a "ride".

Version 2 development is under way and aims to add persistent storage to the application and add a messaging service to support subscribing to active rides.

## Setup

Will be added after version 2...

## Run the application

### To build the application a few services need to compile

#### Proto buffers

proto/

#### Go application

main.go

#### Build the Docker image

Dockerfile

### Build the application

Run the build script to compile the executables

```
cd build/
./build.sh
```

Then run the containerized application run:

```
docker-compose up -d
```

To bring down the containerized application run:

```
docker-compose down
```
