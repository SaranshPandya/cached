package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/saranshpandya/cached/internal/storage"
)

type Store struct {
	mu   sync.RWMutex
	data map[string][]byte
}

func New() *Store {
	return &Store{
		data: make(map[string][]byte),
	}
}

func (store *Store) Set(ctx context.Context, key string, value []byte) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	store.mu.Lock()
	defer store.mu.Unlock()

	store.data[key] = value

	return nil
}

func (store *Store) Get(ctx context.Context, key string) ([]byte, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	store.mu.RLock()
	defer store.mu.RUnlock()

	value, exists := store.data[key]

	if !exists {
		fmt.Printf("\nKey - %s does not exists\n", key)
		return nil, storage.ErrKeyNotFound
	}

	return value, nil
}

func (store *Store) Delete(ctx context.Context, key string) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	store.mu.Lock()
	defer store.mu.Unlock()

	delete(store.data, key)

	return nil
}
