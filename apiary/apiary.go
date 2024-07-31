package client

import (
	apiaryv1 "github.com/hive-ops/go-apiary/pb/apiary/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type ApiaryClient struct {
	apiaryv1.ApiaryServiceClient
	conn *grpc.ClientConn
}

func NewClient(address string) *ApiaryClient {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &ApiaryClient{conn: conn, ApiaryServiceClient: apiaryv1.NewApiaryServiceClient(conn)}
}

func (c *ApiaryClient) close() {
	err := c.conn.Close()
	if err != nil {
		return
	}
}
