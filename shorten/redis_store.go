package shorten

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/go-redis/redis"
)

func init() {
	gob.Register(&Record{})
}

// RedisStore redis cached store
type RedisStore struct {
	Store
	R *redis.Client
}

func NewRedisStore(s Store, rc *redis.Client) *RedisStore {
	return &RedisStore{Store: s, R: rc}
}

// Put doc
func (rs *RedisStore) Put(r *Record) error {
	if err := rs.Store.Put(r); err != nil {
		return err
	}
	return rs.remeber(rs.key(r.ShortUrl), r)
}

// Get doc
func (rs *RedisStore) Get(s string) (*Record, error) {
	k := rs.key(s)
	v, err := rs.R.Get(k).Result()
	if err != nil {
		if err != redis.Nil {
			return nil, err
		}
		r, err := rs.Store.Get(s)
		if err != nil {
			return nil, err
		}
		rs.remeber(k, r)
		return r, nil
	}
	var r Record
	b := bytes.NewBufferString(v)
	if err := gob.NewDecoder(b).Decode(&r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (rs *RedisStore) key(s string) string {
	return fmt.Sprintf("_shortUrl:key:%s", s)
}

func (rs *RedisStore) remeber(k string, r *Record) error {
	b := new(bytes.Buffer)
	if err := gob.NewEncoder(b).Encode(r); err != nil {
		return err
	}
	return rs.R.Set(k, b.String(), 0).Err()
}
