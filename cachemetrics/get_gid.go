package cachemetrics

import (
	"sync/atomic"

	"github.com/petermattis/goid"
)

var (
	MiningRoutineId  int64 // mining main process routine id
	SyncingRoutineId int64 // syncing main process routine id
// Mining           int64
// Importing        int64
)

func Goid() int64 {
	return goid.Get()
}

func UpdateMiningRoutineID(id int64) {
	atomic.StoreInt64(&MiningRoutineId, id)
}

//func StartMiningProcess() {
//	atomic.StoreInt64(&Mining, 1)
//}
//func StopMiningProcess() {
//	atomic.StoreInt64(&Mining, 0)
//}
//func StartImportingProcess() {
//	atomic.StoreInt64(&Importing, 1)
//}
//func StopImportingProcess() {
//	atomic.StoreInt64(&Importing, 0)
//}
//func DuringMining() bool {
//	return atomic.LoadInt64(&Mining) == 1
//}
//func DuringImporting() bool {
//	return atomic.LoadInt64(&Importing) == 1
//}

// judge if it is main process of mining
func IsMinerMainRoutineID(id int64) bool {
	if id == atomic.LoadInt64(&MiningRoutineId) {
		return true
	}
	return false
}

func UpdateSyncingRoutineID(id int64) {
	atomic.StoreInt64(&SyncingRoutineId, id)
}

// judge if it is main process of syncing
func IsSyncMainRoutineID(id int64) bool {
	if id == atomic.LoadInt64(&SyncingRoutineId) {
		return true
	}
	return false
}
