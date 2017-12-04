package shorten

// B62EncShorten 采用通过全局自增ID，然后对其进行62进制编码获取短网址
type B62EncShorten struct {
	g SerialNoGenerator
}

// NewB62EncShorten Get instance of b62EncShorten
func NewB62EncShorten(g SerialNoGenerator) *B62EncShorten {
	return &B62EncShorten{g: g}
}

// Create 将长字符串转换成短字符串
func (sh *B62EncShorten) Create(long string) string {
	id := sh.g.ID()
	return DecTo62(id)
}

// Expand Deprecated
func (sh *B62EncShorten) Expand(s string) (string, error) {
	return "", nil
}
