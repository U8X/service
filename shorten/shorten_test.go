package shorten

import (
	"testing"
)

func TestB62EncShorten(t *testing.T) {
	g := NewBasicSerialNoGenerator(10000)
	s := NewB62EncShorten(g)
	v := s.Create("http://www.baidu.com")
	t.Log(v)
	if v == "" {
		t.Error("短链接生成失败")
	}
}
