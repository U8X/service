package shorten

import (
	"testing"
)

func TestBasicSerialNoGenerator(t *testing.T) {
	g := NewBasicSerialNoGenerator(10000)
	for i := 10000; i < 10010; i++ {
		v := g.ID()
		if v != uint64(i) {
			t.Error("BasicSerialNoGenerator generate ID error")
		}
	}
}
