package server

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/FarmerChillax/AirRocket/proto/rocket_server"
)

type Entry struct {
	payload    *rocket_server.EntryMessage
	createTime int64
	expired    int64
}

type EntryCache struct {
	entries map[string]*Entry
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
				entries: make(map[string]*Entry),
				mutex:   sync.RWMutex{},
			}
			go func() {
				timer := time.NewTimer(time.Second)
				for ; true; <-timer.C {
					log.Println("Staring delete function...")
					entryCache.run()
					timer.Reset(time.Second)
				}
			}()
		})
	}
	return entryCache
}

func (ec *EntryCache) Set(accessCode string, entry *rocket_server.EntryMessage, expired int64) error {
	ec.mutex.Lock()
	defer ec.mutex.Unlock()
	if _, ok := ec.entries[accessCode]; ok {
		return ErrAccessCodeRepeated
	}
	ec.entries[accessCode] = &Entry{
		payload:    entry,
		createTime: time.Now().Unix(),
		expired:    expired,
	}
	return nil
}

func (ec *EntryCache) Get(accessCode string) (*rocket_server.EntryMessage, error) {
	ec.mutex.RLock()
	defer ec.mutex.RUnlock()
	entry, ok := ec.entries[accessCode]
	if !ok {
		return nil, nil
	}
	if entry.createTime+entry.expired < time.Now().Unix() {
		delete(ec.entries, accessCode)
		return nil, nil
	}
	return entry.payload, nil
}

func (ec *EntryCache) Delete(accessCode string) {
	ec.mutex.Lock()
	defer ec.mutex.Unlock()
	delete(ec.entries, accessCode)
	log.Printf("delete accessCode: %v", accessCode)
}

func (ec *EntryCache) All() (map[string]*rocket_server.EntryMessage, error) {
	ec.mutex.RLock()
	defer ec.mutex.RUnlock()
	result := make(map[string]*rocket_server.EntryMessage)
	for key, Value := range ec.entries {
		result[key] = Value.payload
	}
	return result, nil
}

func (ec *EntryCache) run() {
	for accessCode, entry := range ec.entries {
		if time.Now().Unix()-entry.createTime > entry.expired {
			// 过期删除字段
			ec.mutex.Lock()
			delete(ec.entries, accessCode)
			ec.mutex.Unlock()
		}
	}
}
