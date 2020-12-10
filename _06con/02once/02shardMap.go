package main

import "sync"

var SHARD_COUNT = 32

// 分成SHARD_COUNT个分片的map
type ConcurrentMap []*ConcurrentMapShared

// 通过RWMutex保护的线程安全的分片，包含一个map
type ConcurrentMapShared struct {
	items        map[string]interface{}
	sync.RWMutex // Read Write mutex, guards access to internal map.
}

// 创建并发map
func New() ConcurrentMap {
	m := make(ConcurrentMap, SHARD_COUNT)
	for i := 0; i < SHARD_COUNT; i++ {
		m[i] = &ConcurrentMapShared{items: make(map[string]interface{})}
	}
	return m
}


// 根据key计算分片索引
func (m ConcurrentMap) GetShard(key string) *ConcurrentMapShared {
	return m[uint(fnv32(key))%uint(SHARD_COUNT)]
}

func (m ConcurrentMap) Set(key string, value interface{}) {
	// 根据key计算出对应的分片
	shard := m.GetShard(key)
	shard.Lock() //对这个分片加锁，执行业务操作
	shard.items[key] = value
	shard.Unlock()
}

func (m ConcurrentMap) Get(key string) (interface{}, bool) {
	// 根据key计算出对应的分片
	shard := m.GetShard(key)
	shard.RLock()
	// 从这个分片读取key的值
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}