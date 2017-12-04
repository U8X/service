package shorten

import (
	"bytes"
)

const (
	divisor = 62
	tables  = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// DecTo62 将10进制数字转换成62进制字符串
func DecTo62(n uint64) string {
	c := make([]int, 0)
	r := n
	b := new(bytes.Buffer)
	for r >= divisor {
		i := r % divisor
		r = r / divisor
		c = append(c, int(i))
	}
	c = append(c, int(r))
	for i := len(c) - 1; i >= 0; i-- {
		b.WriteByte(tables[c[i]])
	}
	return b.String()
}
