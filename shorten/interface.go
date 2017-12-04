package shorten

// Interface 短网址生成接口
type Interface interface {
	// Create 将长链接转换成短链接
	Create(string) string

	// Expand 将短链接转成长链接
	Expand(string) (string, error)
}
