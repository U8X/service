package shorten

import (
	"time"
)

// PersistentShorten 生成短链接，然后将记录永久保存
type PersistentShorten struct {
	Interface
	Store
}

// NewPersistentShorten 生成持久化短链接服务实例
func NewPersistentShorten(i Interface, s Store) *PersistentShorten {
	return &PersistentShorten{Interface: i, Store: s}
}

// Create 生成短链接
func (sh *PersistentShorten) Create(long string) string {
	i := 0
	for {
		v := sh.Interface.Create(long)
		if err := sh.Store.Put(&Record{LongUrl: long, ShortUrl: v}); err != nil {
			i++
			time.Sleep(time.Duration(i) * 10 * time.Millisecond)
			continue
		}
		return v
	}
}

// Expand Convert short string to long url
func (sh *PersistentShorten) Expand(s string) (string, error) {
	r, err := sh.Store.Get(s)
	if err != nil {
		return "", err
	}
	return r.LongUrl, nil
}
