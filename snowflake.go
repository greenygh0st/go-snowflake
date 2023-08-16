package go-snowflake

import (
	"fmt"
	"sync"
	"time"
)

const (
	epoch            = int64(1577836800000) // set to any timestamp you want, this one is 01-01-2020
	workerIDBits     = uint64(5)            // 5 bits allows for 32 worker nodes
	dataCenterIDBits = uint64(5)            // 5 bits allows for 32 datacenters
	sequenceBits     = uint64(12)           // 12 bits allows for 4096 unique IDs per millisecond

	workerIDShift      = sequenceBits
	dataCenterIDShift  = sequenceBits + workerIDBits
	timestampLeftShift = sequenceBits + workerIDBits + dataCenterIDBits
	sequenceMask       = int64(-1) ^ (int64(-1) << sequenceBits)
)

type Snowflake struct {
	mu            sync.Mutex
	lastTimestamp int64
	workerID      int64
	dataCenterID  int64
	sequence      int64
}

func NewSnowflake(workerID, dataCenterID int64) *Snowflake {
	return &Snowflake{
		lastTimestamp: 0,
		workerID:      workerID,
		dataCenterID:  dataCenterID,
		sequence:      0,
	}
}

func (s *Snowflake) Generate() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	ts := time.Now().UnixNano() / int64(time.Millisecond)

	if ts < s.lastTimestamp {
		fmt.Println("Invalid system clock")
		return 0
	}

	if s.lastTimestamp == ts {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			for ts <= s.lastTimestamp {
				ts = time.Now().UnixNano() / int64(time.Millisecond)
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastTimestamp = ts

	return ((ts - epoch) << timestampLeftShift) | (s.dataCenterID << dataCenterIDShift) | (s.workerID << workerIDShift) | s.sequence
}

func main() {
	node := NewSnowflake(1, 1)
	for i := 0; i < 10; i++ {
		fmt.Println(node.Generate())
	}
}
