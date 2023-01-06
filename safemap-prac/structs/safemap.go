package structs

import "sync"

type SafeMap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

// NewSafeMap returns a new SafeMap.
func New[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V),
	}
}

// Insert inserts a key-value pair into the SafeMap.
func (sm *SafeMap[K, V]) Insert(key K, value V) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.data[key] = value
}

// Get returns a value from the SafeMap.
func (sm *SafeMap[K, V]) Get(key K) (V, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	value, ok := sm.data[key]

	return value, ok
}

// Update an existing key-value pair into the SafeMap.
func (sm *SafeMap[K, V]) Update(key K, newValue V) bool {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if _, ok := sm.data[key]; !ok {
		return false
	}

	sm.data[key] = newValue
	return true
}

// Remove an existing key-value pair from the SafeMap.
func (sm *SafeMap[K, V]) Delete(key K) bool {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if _, ok := sm.data[key]; !ok {
		return false
	}

	delete(sm.data, key)
	return true
}

// Checks if key exists on SafeMap
func (sm *SafeMap[K, V]) HasKey(key K) bool {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	_, ok := sm.data[key]

	return ok
}
