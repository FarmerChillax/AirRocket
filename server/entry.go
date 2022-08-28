package server

import (
	"errors"
	"sync"

	"github.com/FarmerChillax/AirRocket/proto/rocket_server"
)

type EntryCache struct {
	entries map[string]*rocket_server.EntryMessage
	mutex   sync.RWMutex
}

var (
	entryCache *EntryCache
	once       sync.Once
)

var (
	ErrAccessCodeRepeated error = errors.New("access code repeated")
	ErrNotFound           error = errors.New("access code not found")
)

func NewEntryCache() *EntryCache {
	if entryCache == nil {
		once.Do(func() {
			entryCache = &EntryCache{
				entries: make(map[string]*rocket_server.EntryMessage),
				mutex:   sync.RWMutex{},
			}
		})
	}
	return entryCache
}

func (ec *EntryCache) Set(accessCode string, entry *rocket_server.EntryMessage) error {
	ec.mutex.Lock()
	defer ec.mutex.Unlock()
	if _, ok := ec.entries[accessCode]; ok {
		return ErrAccessCodeRepeated
	}
	ec.entries[accessCode] = entry
	return nil
}

func (ec *EntryCache) Get(accessCode string) (*rocket_server.EntryMessage, error) {
	ec.mutex.RLock()
	defer ec.mutex.RUnlock()
	entry, ok := ec.entries[accessCode]
	if !ok {
		return nil, nil
	}
	return entry, nil
}

func (ec *EntryCache) All() (map[string]*rocket_server.EntryMessage, error) {
	ec.mutex.RLock()
	defer ec.mutex.RUnlock()
	result := make(map[string]*rocket_server.EntryMessage)
	for key, Value := range ec.entries {
		result[key] = Value
	}
	return result, nil
}
