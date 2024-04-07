#!/bin/bash
set -ex
# Leave build/ directory
cd ..

# Specify the name of your Go source file (without the .go extension)
SOURCE_FILE="main"

# Specify the name of the output binary
OUTPUT_BINARY="RouteRaydar"

# Specify the name of the Docker image
IMAGE_NAME="route-raydar:1.0"

# TODO: Add protobuf compilation if changes occured

if [ $? -eq 0 ]; then
  echo "Compiling Proto buffers."
  protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  proto/routeRaydar.proto
  if [ $? -eq 0 ]; then

    echo "Compiling Go application. This may take a moment..."

    # Compile the Go source file
    go build -o $OUTPUT_BINARY $SOURCE_FILE.go
    # Check if compilation was successful
    if [ $? -eq 0 ]; then
        echo "Docker image build successful. Image name: $IMAGE_NAME"
    else
        echo "Docker image build failed."
        exit 1
    fi
  else
    echo "Go compilation failed."
    exit 1
  fi
else
  echo "Protobuf compilation failed."
  exit 1
fi

