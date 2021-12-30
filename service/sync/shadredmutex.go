package sync

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"sync"
)

var (
	NumberOfShard = 256
)

type ShardedRWMutex struct {
	mutexes []*sync.RWMutex
}

func NewShardedRWMutex() *ShardedRWMutex {
	s := &ShardedRWMutex{
		mutexes: make([]*sync.RWMutex, NumberOfShard),
	}

	for i := 0; i < NumberOfShard; i++ {
		s.mutexes[i] = new(sync.RWMutex)
	}

	return s
}

func (c *ShardedRWMutex) Lock(key string) {
	c.getShardedMutex(key).Lock()
}

func (c *ShardedRWMutex) Unlock(key string) {
	c.getShardedMutex(key).Unlock()
}

func (c *ShardedRWMutex) RLock(key string) {
	c.getShardedMutex(key).RLock()
}

func (c *ShardedRWMutex) RUnlock(key string) {
	c.getShardedMutex(key).RUnlock()
}

func (c *ShardedRWMutex) getShardedMutex(key string) *sync.RWMutex {
	shardId := getShard(key)
	return c.mutexes[shardId]
}

func getShard(key string) int {
	// Hash
	hasher := sha1.New()
	hasher.Write([]byte(key))
	hash := hasher.Sum(nil)

	// Get the first 8 bytes and convert it to uint32
	buffer := bytes.NewReader(hash[0:4])
	var result uint32
	binary.Read(buffer, binary.BigEndian, &result)

	// Mod it to the number of shard
	return int(result) % NumberOfShard
}
