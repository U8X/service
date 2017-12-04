package shorten

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	// ErrDuplicatedShortURL 数据库中已存在该短链接记录
	ErrDuplicatedShortURL = errors.New("短链接已经存在")
)

// Store 保存短链接记录的仓库接口
type Store interface {
	// Put 保存短链接
	Put(r *Record) error
	// Get 通过给定的短链接，获取对应的长链接记录
	Get(s string) (*Record, error)
}

// MongoDBStore store of mongodb
type MongoDBStore struct {
	C *mgo.Collection
}

// Put 将短链接保存到 mongodb 中
func (s *MongoDBStore) Put(r *Record) error {
	n, err := s.C.Find(bson.M{"short_url": r.ShortUrl}).Count()
	if err != nil {
		return err
	}
	if n > 0 {
		return ErrDuplicatedShortURL
	}
	return s.C.Insert(r)
}

// Get 从 mongodb 中获取短链接
func (st *MongoDBStore) Get(s string) (*Record, error) {
	var r Record
	err := st.C.Find(bson.M{"short_url": s}).One(&r)
	return &r, err
}
