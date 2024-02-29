package kvlogic

import (
	"testing"
	"time"
)

// TestKeyValue_SetAndGet tests the Set and Get methods of the KeyValue type.
func TestKeyValue_SetAndGet(t *testing.T) {
	// Create a new KeyValue instance
	kv := NewKeyValue()

	// Test Set method
	kv.Set("key1", "value1", 5)
	kv.Set("key2", "value2", 10)

	// Test Get method
	value1 := kv.Get("key1")
	if value1 != "value1" {
		t.Errorf("Expected value1 to be 'value1', but got '%s'", value1)
	}

	value2 := kv.Get("key2")
	if value2 != "value2" {
		t.Errorf("Expected value2 to be 'value2', but got '%s'", value2)
	}

	// Test Get method for an expired key
	time.Sleep(6 * time.Second)
	value1Expired := kv.Get("key1")
	if value1Expired != "Key not found (expired)" {
		t.Errorf("Expected value1Expired to be 'Key not found (expired)', but got '%s'", value1Expired)
	}
}

// TestKeyValue_Delete tests the Delete method of the KeyValue type.
func TestKeyValue_Delete(t *testing.T) {
	// Create a new KeyValue instance
	kv := NewKeyValue()

	// Set key-value pair
	kv.Set("key", "value", 5)

	// Test Delete method
	kv.Delete("key")

	// Attempt to get the value after deletion
	value := kv.Get("key")
	if value != "Key not found" {
		t.Errorf("Expected value to be 'Key not found', but got '%s'", value)
	}
}

// TestKeyValue_GetNonExistentKey tests the Get method for a non-existent key.
func TestKeyValue_GetNonExistentKey(t *testing.T) {
	// Create a new KeyValue instance
	kv := NewKeyValue()

	// Test Get method for a non-existent key
	value := kv.Get("nonexistent")
	if value != "Key not found" {
		t.Errorf("Expected value to be 'Key not found', but got '%s'", value)
	}
}

// TestKeyValue_SetWithNegativeTTL tests the Set method with a negative TTL.
func TestKeyValue_SetWithNegativeTTL(t *testing.T) {
	// Create a new KeyValue instance
	kv := NewKeyValue()

	// Attempt to set a key-value pair with negative TTL
	kv.Set("key", "value", -1)

	// The key should not be stored due to the negative TTL
	value := kv.Get("key")
	if value != "Key not found (expired)" {
		t.Errorf("Expected value to be 'Key not found (expired)', but got '%s'", value)
	}
}
