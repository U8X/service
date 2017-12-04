package shorten

import (
	"testing"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func TestMongoDBStore(t *testing.T) {
	sess, err := mgo.DialWithTimeout("localhost", time.Second*30)
	defer sess.Close()
	if err != nil {
		panic(err)
	}
	db := sess.DB("test")
	collection := db.C("short_url_records")
	store := MongoDBStore{C: collection}

	collection.Remove(bson.M{"short_url": "f8"})

	r1 := &Record{ShortUrl: "f8", LongUrl: "http://yuez.me"}
	if err = store.Put(r1); err != nil {
		t.Error("Put error in MongoDBStore: ", err.Error())
	}

	if err = store.Put(r1); err != ErrDuplicatedShortURL {
		t.Error("Put error in MongoDBStore: Should raise duplicated error, but none.")
	}

	r2, err := store.Get("f8")
	if err != nil {
		t.Error("Get error in MongoDBStore: ", err.Error())
	}
	if r2.ShortUrl != "f8" {
		t.Errorf("Get error in MongoDBStore: expected %s, got %s", "f8", r2.ShortUrl)
	}
}
