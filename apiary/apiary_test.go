package apiary

import (
	"context"
	apiaryv1 "github.com/hive-ops/go-apiary/pb/apiary/v1"
	"github.com/hive-ops/go-apiary/utils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
	"time"
)

var keyspace = "benchmark"

func TestClient(t *testing.T) {

	client := NewClient("127.0.0.1:2468", insecure.NewCredentials())

	keyspace := utils.RandomString(10, true, false, false)
	entries := []*apiaryv1.Entry{
		{
			Key:   "test1",
			Value: []byte("test1"),
		},
		{
			Key:   "test2",
			Value: []byte("test2"),
		},
	}
	keys := make([]string, 0)
	for _, entry := range entries {
		keys = append(keys, entry.Key)
	}

	values := make(map[string][]byte)
	for _, entry := range entries {
		values[entry.Key] = entry.Value
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Run("GetEntrySuccessfully - Empty", func(t *testing.T) {
		res, err := client.GetEntries(ctx, apiaryv1.NewGetEntriesRequest(keyspace, keys))

		assert.NoError(t, err)
		assert.Empty(t, res.Entries)
		assert.NotEmpty(t, res.NotFound)
	})

	t.Run("SetEntrySuccessfully", func(t *testing.T) {
		res, err := client.SetEntries(ctx, apiaryv1.NewSetEntriesRequest(keyspace, entries))

		assert.NoError(t, err)
		assert.Equal(t, keys, res.Successful)
	})

	t.Run("GetEntrySuccessfully - Non-empty", func(t *testing.T) {
		res, err := client.GetEntries(ctx, apiaryv1.NewGetEntriesRequest(keyspace, []string{"test1", "invalid-key"}))

		assert.NoError(t, err)
		assert.NotEmpty(t, res.Entries)
		assert.NotEmpty(t, res.NotFound)
	})

	t.Run("DeleteEntrySuccessfully", func(t *testing.T) {
		keysToBeDeleted := []string{"test1"}
		res, err := client.DeleteEntries(ctx, apiaryv1.NewDeleteEntriesRequest(keyspace, keysToBeDeleted))

		assert.NoError(t, err)
		assert.Equal(t, keysToBeDeleted, res.Successful)

		getRes, getErr := client.GetEntries(ctx, apiaryv1.NewGetEntriesRequest(keyspace, keys))

		assert.NoError(t, getErr)
		assert.NotEmpty(t, getRes.Entries)
		assert.NotEmpty(t, getRes.NotFound)

	})

	t.Run("ClearEntriesSuccessfully", func(t *testing.T) {
		res, err := client.ClearEntries(ctx, apiaryv1.NewClearEntriesRequest(keyspace))

		assert.NoError(t, err)
		assert.True(t, res.Successful)

		getRes, getErr := client.GetEntries(ctx, apiaryv1.NewGetEntriesRequest(keyspace, keys))

		assert.NoError(t, getErr)
		assert.Empty(t, getRes.Entries)
		assert.NotEmpty(t, getRes.NotFound)
	})

	client.close()
}
