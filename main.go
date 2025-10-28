package main

import (
	"notification_service/cmd"
)

func main() {
	// run consumer in another goroutine
	cmd.RunConsumer()

	// run grpc server in main goroutine
	cmd.RunGrpcServer()
}
