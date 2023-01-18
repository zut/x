package xdb_test

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/zut/x/xdb"
	"github.com/zut/x/xx"
	"testing"
	"time"
)

func init() {
	_ = xdb.Set(key, value)
	_ = xdb.Set(tk.Int64, tv.Int64)
	_ = xdb.Set(tk.Float64, tv.Float64)
	_ = xdb.Set(tk.String, tv.String)
	_ = xdb.Set(tk.Bool, tv.Bool)
	_ = xdb.Set(tk.Array, tv.Array)
	_ = xdb.Set(tk.Slice, tv.Slice)
	_ = xdb.Set(tk.Map, tv.Map)
	_ = xdb.Set(tk.User1, tv.User1)
	_ = xdb.Set(tk.User2, tv.User2)
	_ = xdb.Set(tk.User3, tv.User3)
}

func TestDel(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.Del(tk.Int64), nil)
	})
	//error
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.Del(tk.NotExist), nil)
	})
	_ = xdb.Set(tk.Int64, tv.Int64)
}

func TestExists(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v, e := xdb.Exists(key)
		t.Assert(e, nil)
		t.Assert(v, true)
	})
	gtest.C(t, func(t *gtest.T) {
		v, e := xdb.Exists(tk.NotExist)
		t.Assert(e, nil)
		t.Assert(v, false)
	})
}
func TestExpire(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.Expire(key, 1)
		t.Assert(err, nil)
		t.Assert(v, true)
		xx.Sleep(2)
		v, err = xdb.Exists(key)
		t.Assert(err, nil)
		t.Assert(v, false)
	})
	_ = xdb.Set(key, value)
}
func TestExpireAt(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.ExpireAt(key, time.Now().Add(time.Second*10))
		t.Assert(err, nil)
		t.Assert(v, true)
		tm, err := xdb.TTL(key)
		t.Assert(err, nil)
		t.AssertGT(tm.Seconds(), 10-1)
		_ = xdb.Del(key)
	})
	_ = xdb.Set(key, value)
}

func TestGet(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v, e := xdb.Get(key)
		t.Assert(e, nil)
		t.Assert(v, value)
	})
	// string
	gtest.C(t, func(t *gtest.T) {
		v, e := xdb.Get(tk.String)
		t.Assert(e, nil)
		t.Assert(v, tv.String)
	})
	// bool
	gtest.C(t, func(t *gtest.T) {
		v, e := xdb.Get(tk.Bool)
		t.Assert(e, nil)
		t.Assert(v, tv.Bool)
	})
	//error
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.Get(tk.NotExist)
		t.Assert(err, xdb.Nil)
		t.Assert(v, nil)
	})
}

func TestGetStr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v0 := "123"
		t.Assert(xdb.Set(key, v0), nil)
		v, err := xdb.GetStr(key)
		t.Assert(err, nil)
		t.Assert(v, v0)
	})
}

func TestGetTo(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		_ = xdb.Set(tk.Int64, tv.Int64)
		var v int64
		e := xdb.GetTo(tk.Int64, &v)
		t.Assert(e, nil)
		t.Assert(v, tv.Int64)
	})
	// Float64
	gtest.C(t, func(t *gtest.T) {
		var v float64
		e := xdb.GetTo(tk.Float64, &v)
		t.Assert(e, nil)
		t.Assert(v, tv.Float64)
	})
	// Array
	gtest.C(t, func(t *gtest.T) {
		var v [2]float64
		e := xdb.GetTo(tk.Array, &v)
		t.Assert(e, nil)
		t.Assert(len(v), len(tv.Array))
		t.Assert(v == tv.Array, true)
		t.Assert(gconv.SliceAny(v), gconv.SliceAny(tv.Array))
	})
	// Slice
	gtest.C(t, func(t *gtest.T) {
		var v []int
		e := xdb.GetTo(tk.Slice, &v)
		t.Assert(e, nil)
		t.Assert(v, tv.Slice)
	})
	// Map
	gtest.C(t, func(t *gtest.T) {
		var v map[string]int
		e := xdb.GetTo(tk.Map, &v)
		t.Assert(e, nil)
		t.Assert(v, tv.Map)
	})
	// struct
	gtest.C(t, func(t *gtest.T) {
		var v User
		e := xdb.GetTo(tk.User1, &v)
		t.Assert(e, nil)
		t.Assert(v, tv.User1)
	})
	//error
	gtest.C(t, func(t *gtest.T) {
		var v string
		e := xdb.GetTo(tk.NotExist, &v)
		t.Assert(e, xdb.Nil)
		t.Assert(v, nil)
	})
}

func TestIncr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		_ = xdb.Del(tk.Incr)
		v, e := xdb.Incr(tk.Incr)
		t.Assert(e, nil)
		t.Assert(v, 1)
		v, e = xdb.Incr(tk.Incr)
		t.Assert(e, nil)
		t.Assert(v, 2)
	})
	// error
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.Set(tk.Incr, "aaa"), nil)
		v, e := xdb.Incr(tk.Incr)
		t.Assert(e, "ERR value is not an integer or out of range")
		t.Assert(v, 0)
	})
}
func TestIncrBy(t *testing.T) {
	_ = xdb.Del(key)
	gtest.C(t, func(t *gtest.T) {
		v, e := xdb.IncrBy(key, 1)
		t.Assert(e, nil)
		t.Assert(v, 1)
		v, e = xdb.IncrBy(key, 5)
		t.Assert(e, nil)
		t.Assert(v, 6)
		v, e = xdb.IncrBy(key, -9)
		t.Assert(e, nil)
		t.Assert(v, -3)
		v, e = xdb.IncrBy(key, 3)
		t.Assert(e, nil)
		t.Assert(v, 0)
	})
}

func TestKeys(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v, e := xdb.Keys("user*")
		t.Assert(e, nil)
		t.Assert(v, g.Slice{"user1", "user2", "user3"})
	})
}

func TestKeysByPrefix(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		_ = xdb.Set("K1", 111)
		_ = xdb.Set("K2", 222)
		_ = xdb.Set("K3", 333)
		keyList, err := xdb.KeysByPrefix("K")
		t.Assert(err, nil)
		t.Assert(keyList, []string{"K1", "K2", "K3"})
	})
}

func TestMDel(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		//same to MSet
	})
	// error
	gtest.C(t, func(t *gtest.T) {
		e := xdb.MDel(g.SliceStr{tk.NotExist})
		t.Assert(e, nil)
	})
}
func TestMGet(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		//same to MSet
	})
	// error
	gtest.C(t, func(t *gtest.T) {
		vs, e := xdb.MGet(g.SliceStr{tk.NotExist})
		t.Assert(e, nil)
		t.Assert(len(vs), 0)
	})
}
func TestMSet(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		kvs := g.Map{"user1": user1, "user2": user2, "user3": user3}
		t.Assert(xdb.MSet(kvs), nil)
		vs, e := xdb.MGet(keysUser)
		t.Assert(e, nil)
		t.Assert(vs, []g.Map{
			{"Age": 1, "Id": 1, "Name": "user1"},
			{"Age": 2, "Id": 2, "Name": "user2"},
			{"Age": 3, "Id": 3, "Name": "user3"},
		})
		e = xdb.MDel(keysUser)
		t.Assert(e, nil)
		vs, e = xdb.MGet(keysUser)
		t.Assert(e, nil)
		t.Assert(len(vs), 0)
	})
}

func TestScan(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		_ = xdb.Set("K1", 111)
		_ = xdb.Set("K2", 222)
		_ = xdb.Set("K3", 333)
		keyList, err := xdb.Scan("K*")
		t.Assert(err, nil)
		t.Assert(keyList, []string{"K1", "K2", "K3"})
	})

}

func TestSet(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.Set(key, 2), nil)
	})
}

func TestSetX(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		err := xdb.SetX(key, value, 1)
		t.Assert(err, nil)
		v, e := xdb.Get(key)
		t.Assert(e, nil)
		t.Assert(v, value)
		xx.Sleep(2)
		v, e = xdb.Get(key)
		t.Assert(e, xdb.Nil)
		t.Assert(v, nil)
	})
}

func TestTTL(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		tm, err := xdb.TTL(key)
		t.Assert(err, nil)
		t.Assert(tm, -1*time.Nanosecond)
	})
}
