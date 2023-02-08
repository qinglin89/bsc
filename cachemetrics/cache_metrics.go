package cachemetrics

import (
	"time"

	"github.com/ethereum/go-ethereum/metrics"
)

type cacheLayerName string

const (
	CacheL1ACCOUNT cacheLayerName = "CACHE_L1_ACCOUNT"
	CacheL2ACCOUNT cacheLayerName = "CACHE_L2_ACCOUNT"
	CacheL3ACCOUNT cacheLayerName = "CACHE_L3_ACCOUNT"
	DiskL4ACCOUNT  cacheLayerName = "DISK_L4_ACCOUNT"
	CacheL1STORAGE cacheLayerName = "CACHE_L1_STORAGE"
	CacheL2STORAGE cacheLayerName = "CACHE_L2_STORAGE"
	CacheL3STORAGE cacheLayerName = "CACHE_L3_STORAGE"
	DiskL4STORAGE  cacheLayerName = "DISK_L4_STORAGE"
)

var (
	cacheL1AccountTimer = metrics.NewRegisteredTimer("cache/cost/account/layer1", nil)
	cacheL2AccountTimer = metrics.NewRegisteredTimer("cache/cost/account/layer2", nil)
	cacheL3AccountTimer = metrics.NewRegisteredTimer("cache/cost/account/layer3", nil)
	diskL4AccountTimer  = metrics.NewRegisteredTimer("cache/cost/account/layer4", nil)
	cacheL1StorageTimer = metrics.NewRegisteredTimer("cache/cost/storage/layer1", nil)
	cacheL2StorageTimer = metrics.NewRegisteredTimer("cache/cost/storage/layer2", nil)
	cacheL3StorageTimer = metrics.NewRegisteredTimer("cache/cost/storage/layer3", nil)
	diskL4StorageTimer  = metrics.NewRegisteredTimer("cache/cost/storage/layer4", nil)

	cacheL1AccountCounter = metrics.NewRegisteredCounter("cache/count/account/layer1", nil)
	cacheL2AccountCounter = metrics.NewRegisteredCounter("cache/count/account/layer2", nil)
	cacheL3AccountCounter = metrics.NewRegisteredCounter("cache/count/account/layer3", nil)
	diskL4AccountCounter  = metrics.NewRegisteredCounter("cache/count/account/layer4", nil)
	cacheL1StorageCounter = metrics.NewRegisteredCounter("cache/count/storage/layer1", nil)
	cacheL2StorageCounter = metrics.NewRegisteredCounter("cache/count/storage/layer2", nil)
	cacheL3StorageCounter = metrics.NewRegisteredCounter("cache/count/storage/layer3", nil)
	diskL4StorageCounter  = metrics.NewRegisteredCounter("cache/count/storage/layer4", nil)

	cacheL1AccountCostCounter = metrics.NewRegisteredCounter("cache/totalcost/account/layer1", nil)
	cacheL2AccountCostCounter = metrics.NewRegisteredCounter("cache/totalcost/account/layer2", nil)
	cacheL3AccountCostCounter = metrics.NewRegisteredCounter("cache/totalcost/account/layer3", nil)
	diskL4AccountCostCounter  = metrics.NewRegisteredCounter("cache/totalcost/account/layer4", nil)
	cacheL1StorageCostCounter = metrics.NewRegisteredCounter("cache/totalcost/storage/layer1", nil)
	cacheL2StorageCostCounter = metrics.NewRegisteredCounter("cache/totalcost/storage/layer2", nil)
	cacheL3StorageCostCounter = metrics.NewRegisteredCounter("cache/totalcost/storage/layer3", nil)
	diskL4StorageCostCounter  = metrics.NewRegisteredCounter("cache/totalcost/storage/layer4", nil)

	SyncL1AccountCounterL    = metrics.NewRegisteredCounter("/sync/account/l1/delay/l", nil)
	MinerL1AccountCounterL   = metrics.NewRegisteredCounter("/miner/account/l1/delay/l", nil)
	SyncL1AccountCounterIDL  = metrics.NewRegisteredCounter("/sync/account/l1/delay/idisk/l", nil)
	MinerL1AccountCounterIDL = metrics.NewRegisteredCounter("/miner/account/l1/delay/idisk/l", nil)

	SyncL1StorageCounterL    = metrics.NewRegisteredCounter("/sync/storage/l1/delay/l", nil)
	MinerL1StorageCounterL   = metrics.NewRegisteredCounter("/miner/storage/l1/delay/l", nil)
	SyncL1StorageCounterIDL  = metrics.NewRegisteredCounter("/sync/storage/l1/delay/idisk/l", nil)
	MinerL1StorageCounterIDL = metrics.NewRegisteredCounter("/miner/storage/l1/delay/idisk/l", nil)

	SyncL2AccountCounterL    = metrics.NewRegisteredCounter("/sync/account/l2/delay/l", nil)
	MinerL2AccountCounterL   = metrics.NewRegisteredCounter("/miner/account/l2/delay/l", nil)
	SyncL2AccountCounterIDL  = metrics.NewRegisteredCounter("/sync/account/l2/delay/idisk/l", nil)
	MinerL2AccountCounterIDL = metrics.NewRegisteredCounter("/miner/account/l2/delay/idisk/l", nil)

	SyncL2StorageCounterL    = metrics.NewRegisteredCounter("/sync/storage/l2/delay/l", nil)
	MinerL2StorageCounterL   = metrics.NewRegisteredCounter("/miner/storage/l2/delay/l", nil)
	SyncL2StorageCounterIDL  = metrics.NewRegisteredCounter("/sync/storage/l2/delay/idisk/l", nil)
	MinerL2StorageCounterIDL = metrics.NewRegisteredCounter("/miner/storage/l2/delay/idisk/l", nil)

	SyncL3AccountCounterL    = metrics.NewRegisteredCounter("/sync/account/l3/delay/l", nil)
	MinerL3AccountCounterL   = metrics.NewRegisteredCounter("/miner/account/l3/delay/l", nil)
	SyncL3AccountCounterIDL  = metrics.NewRegisteredCounter("/sync/account/l3/delay/idisk/l", nil)
	MinerL3AccountCounterIDL = metrics.NewRegisteredCounter("/miner/account/l3/delay/idisk/l", nil)

	SyncL3StorageCounterL    = metrics.NewRegisteredCounter("/sync/storage/l3/delay/l", nil)
	MinerL3StorageCounterL   = metrics.NewRegisteredCounter("/miner/storage/l3/delay/l", nil)
	SyncL3StorageCounterIDL  = metrics.NewRegisteredCounter("/sync/storage/l3/delay/idisk/l", nil)
	MinerL3StorageCounterIDL = metrics.NewRegisteredCounter("/miner/storage/l3/delay/idisk/l", nil)

	SyncL4AccountCounterL    = metrics.NewRegisteredCounter("/sync/account/l4/delay/l", nil)
	MinerL4AccountCounterL   = metrics.NewRegisteredCounter("/miner/account/l4/delay/l", nil)
	SyncL4AccountCounterIDL  = metrics.NewRegisteredCounter("/sync/account/l4/delay/idisk/l", nil)
	MinerL4AccountCounterIDL = metrics.NewRegisteredCounter("/miner/account/l4/delay/idisk/l", nil)

	SyncL4StorageCounterL    = metrics.NewRegisteredCounter("/sync/storage/l4/delay/l", nil)
	MinerL4StorageCounterL   = metrics.NewRegisteredCounter("/miner/storage/l4/delay/l", nil)
	SyncL4StorageCounterIDL  = metrics.NewRegisteredCounter("/sync/storage/l4/delay/idisk/l", nil)
	MinerL4StorageCounterIDL = metrics.NewRegisteredCounter("/miner/storage/l4/delay/idisk/l", nil)

	SyncL1AccountCounter  = metrics.NewRegisteredCounter("/sync/account/l1/hit/l", nil)
	MinerL1AccountCounter = metrics.NewRegisteredCounter("/miner/account/l1/hit/l", nil)

	SyncL1StorageCounter  = metrics.NewRegisteredCounter("/sync/storage/l1/hit/l", nil)
	MinerL1StorageCounter = metrics.NewRegisteredCounter("/miner/storage/l1/hit/l", nil)

	SyncL2AccountCounter  = metrics.NewRegisteredCounter("/sync/account/l2/hit/l", nil)
	MinerL2AccountCounter = metrics.NewRegisteredCounter("/miner/account/l2/hit/l", nil)

	SyncL2StorageCounter  = metrics.NewRegisteredCounter("/sync/storage/l2/hit/l", nil)
	MinerL2StorageCounter = metrics.NewRegisteredCounter("/miner/storage/l2/hit/l", nil)

	SyncL3AccountCounter  = metrics.NewRegisteredCounter("/sync/account/l3/hit/l", nil)
	MinerL3AccountCounter = metrics.NewRegisteredCounter("/miner/account/l3/hit/l", nil)

	SyncL3StorageCounter  = metrics.NewRegisteredCounter("/sync/storage/l3/hit/l", nil)
	MinerL3StorageCounter = metrics.NewRegisteredCounter("/miner/storage/l3/hit/l", nil)

	SyncL4AccountCounter  = metrics.NewRegisteredCounter("/sync/account/l4/hit/l", nil)
	MinerL4AccountCounter = metrics.NewRegisteredCounter("/miner/account/l4/hit/l", nil)

	SyncL4StorageCounter  = metrics.NewRegisteredCounter("/sync/storage/l4/hit/l", nil)
	MinerL4StorageCounter = metrics.NewRegisteredCounter("/miner/storage/l4/hit/l", nil)
)

// mark the info of total hit counts of each layers
func RecordCacheDepth(metricsName cacheLayerName) {
	switch metricsName {
	case CacheL1ACCOUNT:
		cacheL1AccountCounter.Inc(1)
	case CacheL2ACCOUNT:
		cacheL2AccountCounter.Inc(1)
	case CacheL3ACCOUNT:
		cacheL3AccountCounter.Inc(1)
	case DiskL4ACCOUNT:
		diskL4AccountCounter.Inc(1)
	case CacheL1STORAGE:
		cacheL1StorageCounter.Inc(1)
	case CacheL2STORAGE:
		cacheL2StorageCounter.Inc(1)
	case CacheL3STORAGE:
		cacheL3StorageCounter.Inc(1)
	case DiskL4STORAGE:
		diskL4StorageCounter.Inc(1)
	}
}

// mark the dalays of each layers
func RecordCacheMetrics(metricsName cacheLayerName, start time.Time) {
	switch metricsName {
	case CacheL1ACCOUNT:
		recordCost(cacheL1AccountTimer, start)
	case CacheL2ACCOUNT:
		recordCost(cacheL2AccountTimer, start)
	case CacheL3ACCOUNT:
		recordCost(cacheL3AccountTimer, start)
	case DiskL4ACCOUNT:
		recordCost(diskL4AccountTimer, start)
	case CacheL1STORAGE:
		recordCost(cacheL1StorageTimer, start)
	case CacheL2STORAGE:
		recordCost(cacheL2StorageTimer, start)
	case CacheL3STORAGE:
		recordCost(cacheL3StorageTimer, start)
	case DiskL4STORAGE:
		recordCost(diskL4StorageTimer, start)

	}
}

// accumulate the total dalays of each layers
func RecordTotalCosts(metricsName cacheLayerName, start time.Time) {
	switch metricsName {
	case CacheL1ACCOUNT:
		accumulateCost(cacheL1AccountCostCounter, start)
	case CacheL2ACCOUNT:
		accumulateCost(cacheL2AccountCostCounter, start)
	case CacheL3ACCOUNT:
		accumulateCost(cacheL3AccountCostCounter, start)
	case DiskL4ACCOUNT:
		accumulateCost(diskL4AccountCostCounter, start)
	case CacheL1STORAGE:
		accumulateCost(cacheL1StorageCostCounter, start)
	case CacheL2STORAGE:
		accumulateCost(cacheL2StorageCostCounter, start)
	case CacheL3STORAGE:
		accumulateCost(cacheL3StorageCostCounter, start)
	case DiskL4STORAGE:
		accumulateCost(diskL4StorageCostCounter, start)

	}
}

func recordCost(timer metrics.Timer, start time.Time) {
	timer.Update(time.Since(start))
}

func accumulateCost(totalcost metrics.Counter, start time.Time) {
	totalcost.Inc(time.Since(start).Nanoseconds())
}
