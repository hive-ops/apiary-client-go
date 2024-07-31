package apiaryv1

func NewEntry(key string, value []byte) *Entry {
	return &Entry{
		Key:   key,
		Value: value,
	}
}

func NewGetEntriesRequest(keyspace string, keys []string) *GetEntriesRequest {
	return &GetEntriesRequest{
		Keyspace: keyspace,
		Keys:     keys,
	}
}

func NewSetEntriesRequest(keyspace string, entries []*Entry) *SetEntriesRequest {
	return &SetEntriesRequest{
		Keyspace: keyspace,
		Entries:  entries,
	}
}

func NewDeleteEntriesRequest(keyspace string, keys []string) *DeleteEntriesRequest {
	return &DeleteEntriesRequest{
		Keyspace: keyspace,
		Keys:     keys,
	}
}

func NewClearEntriesRequest(keyspace string) *ClearEntriesRequest {
	return &ClearEntriesRequest{
		Keyspace: keyspace,
	}
}
