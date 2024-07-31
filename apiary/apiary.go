package apiary

import (
	apiaryv1 "github.com/hive-ops/go-apiary/pb/apiary/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

type Client struct {
	apiaryv1.ApiaryServiceClient
	conn *grpc.ClientConn
}

func NewClient(address string, creds credentials.TransportCredentials) *Client {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &Client{conn: conn, ApiaryServiceClient: apiaryv1.NewApiaryServiceClient(conn)}
}

func (c *Client) close() {
	err := c.conn.Close()
	if err != nil {
		return
	}
}
