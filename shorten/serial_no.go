package shorten

import (
	"sync"
)

// SerialNoGenerator 自增ID生成器
type SerialNoGenerator interface {
	ID() uint64
}

// BasicSerialNoGenerator 通过锁实现的自增ID
type BasicSerialNoGenerator struct {
	i     uint64
	mutex sync.Mutex
}

// NewBasicSerialNoGenerator 生成基于内存的自增ID生成器
func NewBasicSerialNoGenerator(i uint64) *BasicSerialNoGenerator {
	return &BasicSerialNoGenerator{i: i}
}

// ID 返回全局自增ID
func (g *BasicSerialNoGenerator) ID() uint64 {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	v := g.i
	g.i++
	return v
}
