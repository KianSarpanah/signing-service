package persistence

import (
	"fmt"
	"signaturesign/domain"
	"sync"
)

// InMemoryDeviceStore stores devices in memory
type InMemoryDeviceStore struct {
	mu      sync.Mutex
	devices map[string]*domain.Device
}

// NewInMemoryDeviceStore initializes a new InMemoryDeviceStore
func NewInMemoryDeviceStore() *InMemoryDeviceStore {
	return &InMemoryDeviceStore{
		devices: make(map[string]*domain.Device),
	}
}

// AddDevice adds a device to the store
func (store *InMemoryDeviceStore) AddDevice(device *domain.Device) {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.devices[device.ID] = device
}

// GetDevice retrieves a device by its ID
func (store *InMemoryDeviceStore) GetDevice(id string) (*domain.Device, bool) {
	store.mu.Lock()
	defer store.mu.Unlock()
	device, exists := store.devices[id]
	fmt.Println(device)
	return device, exists
}
