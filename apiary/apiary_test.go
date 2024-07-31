package client

import (
	"context"
	apiaryv1 "github.com/hive-ops/apiary-client-go/pb/apiary/v1"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var keyspace = "benchmark"

func TestClient(t *testing.T) {

	client := NewClient("127.0.0.1:24681")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	key := "testKey"

	req := &apiaryv1.GetEntriesRequest{Keyspace: keyspace, Keys: []string{key}}

	res, err := client.GetEntries(ctx, req)

	assert.NoError(t, err)
	assert.Contains(t, res.NotFound, key)

	client.close()
}
