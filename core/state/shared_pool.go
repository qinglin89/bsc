package state

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// sharedPool is used to store maps of originStorage of stateObjects
type StoragePool struct {
	sync.RWMutex
	sharedMap map[common.Address]*sync.Map
}

func NewStoragePool() *StoragePool {
	sharedMap := make(map[common.Address]*sync.Map)
	return &StoragePool{
		sync.RWMutex{},
		sharedMap,
	}
}

// getStorage Check whether the storage exist in pool,
// new one if not exist, the content of storage will be fetched in stateObjects.GetCommittedState()
func (s *StoragePool) getStorage(address common.Address) *sync.Map {
	s.RLock()
	storageMap, ok := s.sharedMap[address]
	s.RUnlock()
	if !ok {
		s.Lock()
		defer s.Unlock()
		if storageMap, ok = s.sharedMap[address]; !ok {
			m := new(sync.Map)
			s.sharedMap[address] = m
			return m
		}
	}
	return storageMap
}

type sharedStateObjects struct {
	stateObjects map[common.Address]*StateObject
	mu           sync.RWMutex
}

func (s *sharedStateObjects) set(so *StateObject) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.stateObjects[so.address] == nil {
		s.stateObjects[so.address] = &StateObject{
			db:                  so.db,
			address:             so.address,
			addrHash:            so.addrHash,
			data:                so.data,
			sharedOriginStorage: so.sharedOriginStorage,
		}
	}
}
func (s *sharedStateObjects) get(addr common.Address) *StateObject {
	s.mu.RLock()
	so := s.stateObjects[addr]
	s.mu.RUnlock()
	if so == nil {
		return nil
	}
	return &StateObject{
		db:                  so.db,
		address:             so.address,
		addrHash:            so.addrHash,
		data:                so.data,
		sharedOriginStorage: so.sharedOriginStorage,
		originStorage:       make(Storage),
		pendingStorage:      make(Storage),
		dirtyStorage:        make(Storage),
	}
}

func copyStateObject(so *StateObject) *StateObject {
	return &StateObject{
		db:                  so.db,
		address:             so.address,
		addrHash:            so.addrHash,
		data:                so.data,
		sharedOriginStorage: so.sharedOriginStorage,
	}
}
