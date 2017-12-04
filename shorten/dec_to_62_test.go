package shorten

import (
	"testing"
)

func TestDecTo62(t *testing.T) {
	{
		n := 62
		s := DecTo62(uint64(n))
		if s != "10" {
			t.Errorf("进制转换错误，62转成62进制应该是 [10], 得到 [%s]", s)
		}
	}

	{
		n := 15
		s := DecTo62(uint64(n))
		if s != "F" {
			t.Errorf("进制转换错误，15转成62进制应该是 [F], 得到 [%s]", s)
		}
	}

	{
		n := 38902341
		s := DecTo62(uint64(n))
		if s != "2dEH7" {
			t.Errorf("进制转换错误，38902341转成62进制应该是 [2dEH7], 得到 [%s]", s)
		}
	}
}

func BenchmarkDecTo62(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DecTo62(38902341)
	}
}
