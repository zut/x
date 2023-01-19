package main

import (
	"context"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
)

type gCache struct {
	Title    string
	Duration int
}

func tGCache1() error {
	a := gCache{
		Title:    "test",
		Duration: 10,
	}
	if err := gcache.Set(context.TODO(), "a", a, 0); err != nil {
		return err
	}
	if v, err := gcache.Get(context.TODO(), "a"); err != nil {
		return err
	} else {
		b := v.Val().(gCache)
		glog.Debug(context.TODO(), b)
		glog.Debug(context.TODO(), b.Title)
	}
	return nil
}

func tGCache2() error {
	a := &gCache{
		Title:    "test",
		Duration: 10,
	}
	if err := gcache.Set(context.TODO(), "a", a, 0); err != nil {
		return err
	}
	if v, err := gcache.Get(context.TODO(), "a"); err != nil {
		return err
	} else {
		b := v.Val().(*gCache)
		glog.Debug(context.TODO(), b)
		glog.Debug(context.TODO(), b.Title)
	}
	return nil
}

func main() {
	if err := tGCache1(); err != nil {
		glog.Error(context.TODO(), err)
	}

	if err := tGCache2(); err != nil {
		glog.Error(context.TODO(), err)
	}

}
