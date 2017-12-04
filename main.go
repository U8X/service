package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/U8X/service/shorten"
	"github.com/go-redis/redis"
	"gopkg.in/mgo.v2"
)

var (
	collection *mgo.Collection
	rc         *redis.Client

	addr          = flag.String("addr", ":5995", "服务监听地址")
	debug         = flag.Bool("debug", false, "是否开打Debug模式")
	prefix        = flag.String("prefix", "http://u0x.cn/", "短链接前缀")
	idStart       = flag.Int64("id", 600000, "自增ID起始值")
	mongoAddr     = flag.String("mongo.addr", "localhost:27017", "MongoDB服务地址")
	mongoUser     = flag.String("mongo.user", "", "MongoDB用户名")
	mongoPassword = flag.String("mongo.password", "", "MongoDB密码")
	mongoDatabase = flag.String("mongo.database", "test", "MongoDB数据库")
	redisAddr     = flag.String("redis.addr", "localhost:6379", "Redis服务地址")
	redisPassword = flag.String("redis.password", "", "Redis密码")
	redisDatabase = flag.Int("redis.database", 1, "Redis数据库")
)

func init() {
	flag.Parse()

	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logrus.SetLevel(logrus.WarnLevel)
	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.Debugf("正在连接Redis: addr=%s, dbname=%d", *redisAddr, *redisDatabase)
	rc = redis.NewClient(&redis.Options{
		Addr:     *redisAddr,
		DB:       *redisDatabase,
		Password: *redisPassword,
	})

	if err := rc.Ping().Err(); err != nil {
		logrus.Panic(err)
	}

	logrus.Debugf("正在连接MongoDB: addr=%s, dbname=%s", *mongoAddr, *mongoDatabase)
	mongo, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{*mongoAddr},
		Username: *mongoUser,
		Password: *mongoPassword,
		Timeout:  30 * time.Second,
	})
	if err != nil {
		logrus.Panic(err)
	}
	collection = mongo.DB(*mongoDatabase).C("short_url_records")
}

func main() {
	defer collection.Database.Session.Close()
	defer rc.Close()

	ch := make(chan error)
	id := shorten.NewRedisSerialNoGenerator(rc, *debug, uint64(*idStart))
	bs := shorten.NewB62EncShorten(id)
	st := &shorten.MongoDBStore{C: collection}
	rt := shorten.NewRedisStore(st, rc)
	sh := shorten.NewPersistentShorten(bs, rt)
	sr := &server{Service: sh, Prefix: *prefix}

	go func() {
		logrus.Debugf("启动短链接服务: addr=%s, prefix=%s", *addr, *prefix)
		ch <- http.ListenAndServe(*addr, sr)
	}()

	logrus.Error(<-ch)
}
