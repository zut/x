package xdb

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/zut/x/xx"
)

var (
	ctx = context.Background()
	rdb *redis.Client

	host = "127.0.0.1"
	port = 9221
)

const Nil = redis.Nil

func Open(HostPort ...string) error {
	if len(HostPort) == 2 {
		host = HostPort[0]
		port = gconv.Int(HostPort[1])
	} else if xx.FileExists("config.toml") || xx.FileExists("config/config.toml") {
		host = g.Cfg().MustGet(gctx.New(), "db.host").String()
		port = g.Cfg().MustGet(gctx.New(), "db.port").Int()
	}

	glog.Info(gctx.New(), "rdb Open", host, port)
	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%v:%v", host, port),
		DialTimeout:  10 * time.Second,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
		//PoolSize:     100,
		//PoolTimeout:  20 * time.Second,
	})

	_, err := rdb.Ping(ctx).Result()
	return err
}

func Close() {
	glog.Info(gctx.New(), "rdb Close", host, port)
	if err := rdb.Close(); err != nil {
		glog.Error(gctx.New(), "rdb", err)
	}
}
func Conn() *redis.Conn {
	//func Conn() *redis.Client {
	conn := new(redis.Conn)
	for i := 0; i <= 1000; i++ {
		conn = rdb.Conn(ctx)
		err := conn.Ping(ctx).Err()
		if err == nil { // redis: Conn is in a bad state: EOF 会出错
			break
		}
		if err := conn.Close(); err != nil {
			glog.Error(gctx.New(), i, "conn.Close()", err)
		}
		glog.Debug(gctx.New(), i, err)
		time.Sleep(time.Millisecond * 10) // 10ms
	}
	//glog.Debug(rdb.PoolStats())
	return conn
	//return rdb
}
func ConnClose(c *redis.Conn) {
	//glog.Debug("ConnClose")
	if err := c.Close(); err != nil {
		glog.Error(gctx.New(), err, c.Info(ctx))
	}
}
func Info() string {
	//conn := rdb.Conn(ctx)
	return rdb.Info(ctx).String()
}

func FlushDB() {
	// 删除当前选定数据库中的所有key。这个命令的执行不会失败。
	glog.Info(gctx.New(), "rdb FlushDB", host, port)
	rdb.FlushDB(ctx)
}

func toIntErr(i int64, err error) (int, error) {
	return gconv.Int(i), err
}

func AddKK(h string) string {
	return fmt.Sprintf("%v_keys", h)
}

func cEmpty(s ...string) error {
	for _, i := range s {
		if i == "" {
			glog.Errorf(gctx.New(), "Key.Empty")
			return fmt.Errorf("Key.Empty")
		}
	}
	return nil
}

func VersionHashName(hashName string) string {
	return fmt.Sprintf("%v_%v", hashName, "Version")
}
