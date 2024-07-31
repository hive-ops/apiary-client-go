package main

import (
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	c = apiary.NewClient("localhost:8080", insecure.NewCredentials())
}
