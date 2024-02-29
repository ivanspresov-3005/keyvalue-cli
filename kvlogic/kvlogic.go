package kvlogic

import (
	"sync"
	"time"
)

// KeyValue represents the key-value storage system.
type KeyValue struct {
	data sync.Map
}

// entry represents a key-value pair with its TTL (Time to Live).
type entry struct {
	value string
	ttl   time.Time
}

// NewKeyValue initializes a new KeyValue instance.
func NewKeyValue() *KeyValue {
	return &KeyValue{}
}

// Set stores a key-value pair with an optional TTL.
func (kv *KeyValue) Set(key, value string, ttl int) {
	// Calculate the expiration time based on the provided TTL
	expiration := time.Now().Add(time.Duration(ttl) * time.Second)

	// Store the key-value pair in the map
	kv.data.Store(key, entry{
		value: value,
		ttl:   expiration,
	})
}

// Get retrieves the value associated with the given key.
func (kv *KeyValue) Get(key string) string {
	// Load the entry from the map
	entryInterface, found := kv.data.Load(key)
	if !found {
		return "Key not found"
	}

	entry := entryInterface.(entry)

	// Check if the entry has expired
	if entry.ttl.IsZero() || time.Now().After(entry.ttl) {
		// Delete the entry if it has expired
		kv.data.Delete(key)
		return "Key not found (expired)"
	}

	return entry.value
}

// Delete removes a key-value pair from the storage.
func (kv *KeyValue) Delete(key string) {
	// Delete the entry from the map
	kv.data.Delete(key)
}
