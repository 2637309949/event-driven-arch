package main

import (
	"errors"
	"sync"
	"time"
)

// Snowflake 实现（41+10+12 = 63 bits）
type Snowflake struct {
	mu           sync.Mutex
	epoch        int64 // 自定义 epoch（毫秒）
	nodeID       int64 // 节点 ID (0..maxNode)
	sequence     int64 // 当前序列
	lastTS       int64 // 上次生成 ID 的时间戳（毫秒）
	sequenceMask int64 // 序列掩码（最大序列）
	nodeBits     uint
	seqBits      uint
	timeShift    uint
	nodeShift    uint
	maxNode      int64
}

// 默认位分配
const (
	defaultTimeBits = 41
	defaultNodeBits = 10
	defaultSeqBits  = 12
)

// NewSnowflake 创建生成器
// nodeID 范围: 0 .. 2^nodeBits - 1
// epoch 为自定义起始时间（毫秒），通常用固定时间：例如 1577836800000 (2020-01-01T00:00:00Z)
func NewSnowflake(nodeID int64, epochMillis int64) (*Snowflake, error) {
	nodeBits := defaultNodeBits
	seqBits := defaultSeqBits

	maxNode := int64(-1) ^ (int64(-1) << nodeBits) // (1<<nodeBits)-1
	if nodeID < 0 || nodeID > maxNode {
		return nil, errors.New("nodeID out of range")
	}
	sequenceMask := int64(-1) ^ (int64(-1) << seqBits) // (1<<seqBits)-1

	sf := &Snowflake{
		epoch:        epochMillis,
		nodeID:       nodeID,
		sequence:     0,
		lastTS:       -1,
		sequenceMask: sequenceMask,
		nodeBits:     uint(nodeBits),
		seqBits:      uint(seqBits),
	}
	// 计算位移
	sf.nodeShift = uint(seqBits)
	sf.timeShift = uint(nodeBits) + sf.nodeShift
	sf.maxNode = maxNode

	return sf, nil
}

// current milliseconds
func nowMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// tilNextMillis 等到下一个毫秒
func tilNextMillis(last int64) int64 {
	ts := nowMillis()
	for ts <= last {
		ts = nowMillis()
	}
	return ts
}

// NextID 生成下一个 ID（线程安全）
func (s *Snowflake) NextID() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	ts := nowMillis()
	if ts < s.lastTS {
		// 时钟回拨：简单的处理是等待直到恢复（也可以返回错误或使用备用序列策略）
		ts = tilNextMillis(s.lastTS)
	}

	if ts == s.lastTS {
		s.sequence = (s.sequence + 1) & s.sequenceMask
		if s.sequence == 0 {
			// 序列溢出，等待下一个毫秒
			ts = tilNextMillis(s.lastTS)
			s.lastTS = ts
			s.sequence = 0
		}
	} else {
		// 新时间戳，重置序列
		s.sequence = 0
		s.lastTS = ts
	}

	id := ((ts - s.epoch) << s.timeShift) |
		(s.nodeID << s.nodeShift) |
		s.sequence

	return id
}

// ParseID 解析 ID，返回生成时间戳（ms）、nodeID、sequence
func (s *Snowflake) ParseID(id int64) (tsMillis int64, nodeID int64, seq int64) {
	seq = id & ((1 << s.seqBits) - 1)
	nodeID = (id >> s.nodeShift) & ((1 << s.nodeBits) - 1)
	tsPart := (id >> s.timeShift) & ((1 << defaultTimeBits) - 1)
	tsMillis = tsPart + s.epoch
	return
}
