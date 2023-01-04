package state

import (
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
)

type syncMap struct {
	*sync.Map
	*StoragePool
}

func (s *syncMap) Count() {
	atomic.AddInt32(&s.StoragePool.length, 1)
}

// sharedPool is used to store maps of originStorage of stateObjects
type StoragePool struct {
	sync.RWMutex
	//	sharedMap map[common.Address]*sync.Map
	sharedMap map[common.Address]*syncMap
	//	length    int
	length int32
}

func NewStoragePool() *StoragePool {
	//	sharedMap := make(map[common.Address]*sync.Map)
	sharedMap := make(map[common.Address]*syncMap)
	return &StoragePool{
		sync.RWMutex{},
		sharedMap,
		0,
	}
}

// getStorage Check whether the storage exist in pool,
// new one if not exist, the content of storage will be fetched in stateObjects.GetCommittedState()
func (s *StoragePool) getStorage(address common.Address) *syncMap {
	s.RLock()
	storageMap, ok := s.sharedMap[address]
	s.RUnlock()
	if !ok {
		s.Lock()
		defer s.Unlock()
		if storageMap, ok = s.sharedMap[address]; !ok {
			//			m := new(sync.Map)
			//			m := new(syncMap)
			m := &syncMap{
				new(sync.Map),
				s,
			}
			s.sharedMap[address] = m
			return m
		}
	}
	return storageMap
}

func (s *StoragePool) Length() int {
	return int(s.length)
}
