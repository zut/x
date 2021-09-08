package xdb_test

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/test/gtest"
	"github.com/zut/x/xdb"
	"github.com/zut/x/xx"
	"testing"
)

func init() {
	_ = xdb.HSet(name, key, value)
	_ = xdb.HSet(name, tk.Int64, tv.Int64)
	_ = xdb.HSet(name, tk.Float64, tv.Float64)
	_ = xdb.HSet(name, tk.String, tv.String)
	_ = xdb.HSet(name, tk.Bool, tv.Bool)
	_ = xdb.HSet(name, tk.Array, tv.Array)
	_ = xdb.HSet(name, tk.Slice, tv.Slice)
	_ = xdb.HSet(name, tk.Map, tv.Map)
	_ = xdb.HSet(nameUser, user1.Name, user1)
	_ = xdb.HSet(nameUser, user2.Name, user2)
	_ = xdb.HSet(nameUser, user3.Name, user3)
	// _ =xdb.HSet(namePJ, g.Slice{1, 1}, g.Map{"Id": 1, "Mode": "a1"})
	// _ =xdb.HSet(namePJ, g.Slice{11, 2}, g.Map{"Id": 11, "Mode": "a11"})
	// _ =xdb.HSet(namePJ, g.Slice{2, 11}, g.Map{"Id": 2, "Mode": "a2"})
	_ = xdb.HSet(namePJ, xx.Join(101, 1), g.Map{"Id": 101, "Mode": "a101"})
	_ = xdb.HSet(namePJ, xx.Join(111, 2), g.Map{"Id": 111, "Mode": "a111"})
	_ = xdb.HSet(namePJ, xx.Join(102, 1), g.Map{"Id": 102, "Mode": "a102"})
}

func TestHClear(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v, e := xdb.HLen(nameUser)
		t.Assert(e, nil)
		t.Assert(v, 3)
		e = xdb.HClear(nameUser)
		t.Assert(e, nil)
		v, e = xdb.HLen(nameUser)
		t.Assert(e, nil)
		t.Assert(v, 0)
	})
	// errors
	gtest.C(t, func(t *gtest.T) {
		err := xdb.HClear(nameNotExist)
		t.Assert(err, nil)
	})
	_ = xdb.HSet(nameUser, user1.Name, user1)
	_ = xdb.HSet(nameUser, user2.Name, user2)
	_ = xdb.HSet(nameUser, user3.Name, user3)
}

func TestHDel(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.HExists(name, key)
		t.Assert(err, nil)
		t.Assert(v, true)
		err = xdb.HDel(name, key)
		t.Assert(err, nil)
		v, err = xdb.HExists(name, key)
		t.Assert(err, nil)
		t.Assert(v, false)
	})
	// error
	gtest.C(t, func(t *gtest.T) {
		err := xdb.HDel(nameNotExist, tk.NotExist)
		t.Assert(err, nil)
		err = xdb.HDel(name, tk.NotExist)
		t.Assert(err, nil)
	})
	_ = xdb.HSet(name, key, value)
}
func TestHExists(t *testing.T) {
	// same as TestHDel
}

func TestHGet(t *testing.T) {
	// vString
	gtest.C(t, func(t *gtest.T) {
		v, e := xdb.HGet(name, tk.String)
		t.Assert(e, nil)
		t.AssertEQ(v, tv.String)
	})
	// vBool
	gtest.C(t, func(t *gtest.T) {
		v, e := xdb.HGet(name, tk.Bool)
		t.Assert(e, nil)
		t.AssertEQ(v, tv.Bool)
	})
	// errors
	gtest.C(t, func(t *gtest.T) {
		v, e := xdb.HGet(name, tk.NotExist)
		t.Assert(e, xdb.Nil)
		t.Assert(v, nil)
		v, e = xdb.HGet(nameNotExist, tk.NotExist)
		t.Assert(e, xdb.Nil)
		t.Assert(v, nil)
	})
}

func TestHGetALL(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		m, e := xdb.HGetAll(nameUser)
		t.Assert(e, nil)
		t.Assert(len(m), 3)
	})
	// errors
	gtest.C(t, func(t *gtest.T) {
		m, e := xdb.HGetAll(nameNotExist)
		t.Assert(e, nil)
		t.Assert(len(m), 0)
	})
}

func TestHGetTo(t *testing.T) {
	// int64
	gtest.C(t, func(t *gtest.T) {
		var v int64
		e := xdb.HGetTo(name, tk.Int64, &v)
		t.Assert(e, nil)
		t.AssertEQ(v, tv.Int64)
	})
	// vFloat64
	gtest.C(t, func(t *gtest.T) {
		var v float64
		e := xdb.HGetTo(name, tk.Float64, &v)
		t.Assert(e, nil)
		t.AssertEQ(v, tv.Float64)
	})
	// vArray
	gtest.C(t, func(t *gtest.T) {
		var v [2]float64
		e := xdb.HGetTo(name, tk.Array, &v)
		t.Assert(e, nil)
		t.Assert(v[0], tv.Array[0])
		t.Assert(v[1], tv.Array[1])
		t.Assert(len(v), len(tv.Array))
	})
	// vSlice
	gtest.C(t, func(t *gtest.T) {
		var v []int
		e := xdb.HGetTo(name, tk.Slice, &v)
		t.Assert(e, nil)
		t.AssertEQ(v, tv.Slice)
	})
	// vMap
	gtest.C(t, func(t *gtest.T) {
		var v map[string]int
		e := xdb.HGetTo(name, tk.Map, &v)
		t.Assert(e, nil)
		t.AssertEQ(v, tv.Map)
	})
	// struct
	gtest.C(t, func(t *gtest.T) {
		var aa User
		e := xdb.HGetTo(nameUser, user1.Name, &aa)
		t.Assert(e, nil)
		t.Assert(user1.Name, aa.Name)
		t.Assert(user1, aa)
	})
	// error
	gtest.C(t, func(t *gtest.T) {
		var aa User
		e := xdb.HGetTo(name, tk.NotExist, &aa)
		t.Assert(e, xdb.Nil)
		e = xdb.HGetTo(nameNotExist, tk.NotExist, &aa)
		t.Assert(e, xdb.Nil)
	})
}

func TestHIncr(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.HDel(name, tk.Incr), nil)
		v, e := xdb.HIncr(name, tk.Incr)
		t.Assert(e, nil)
		t.Assert(v, 1)
		v, e = xdb.HIncr(name, tk.Incr)
		t.Assert(e, nil)
		t.Assert(v, 2)
	})
}
func TestHIncrBy(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.HDel(name, tk.Incr), nil)
		v, e := xdb.HIncrBy(name, tk.Incr, 1)
		t.Assert(e, nil)
		t.Assert(v, 1)
		v, e = xdb.HIncrBy(name, tk.Incr, 5)
		t.Assert(e, nil)
		t.Assert(v, 6)
		v, e = xdb.HIncrBy(name, tk.Incr, -3)
		t.Assert(e, nil)
		t.Assert(v, 3)
		v, e = xdb.HIncrBy(name, tk.Incr, -3)
		t.Assert(e, nil)
		t.Assert(v, 0)
		v, e = xdb.HIncrBy(name, tk.Incr, -5)
		t.Assert(e, nil)
		t.Assert(v, -5)
	})
	gtest.C(t, func(t *gtest.T) {
		_ = xdb.HSet(name, tk.Incr, "aaa")
		v, e := xdb.HIncrBy(name, tk.Incr, 1)
		t.Assert(e, "ERR value is not an integer or out of range")
		t.Assert(v, 0)

	})
}

func TestHIncrByFloat(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.HDel(name, tk.Incr), nil)
		v, e := xdb.HIncrByFloat(name, tk.Incr, 1.5)
		t.Assert(e, nil)
		t.Assert(v, 1.5)
		v, e = xdb.HIncrByFloat(name, tk.Incr, 2.1)
		t.Assert(e, nil)
		t.Assert(v, 3.6)
	})
}
func TestHIncrId(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.HDel(name, tk.Incr), nil)
		v, e := xdb.HIncrBy(name, tk.Incr, 2)
		t.Assert(e, nil)
		t.Assert(v, 2)
		v, e = xdb.HIncrId(name, tk.Incr, 6)
		t.Assert(e, nil)
		t.Assert(v, 100001)
		v, e = xdb.HIncrId(name, tk.Incr, 6)
		t.Assert(e, nil)
		t.Assert(v, 100002)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.HDel(name, tk.Incr), nil)
		v, e := xdb.HIncrId(name, tk.Incr, 6)
		t.Assert(e, nil)
		t.Assert(v, 100001)
		v, e = xdb.HIncrId(name, tk.Incr, 6)
		t.Assert(e, nil)
		t.Assert(v, 100002)
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.HDel(name, tk.Incr), nil)
		v, e := xdb.HIncrId(name, tk.Incr, 3)
		t.Assert(e, nil)
		t.Assert(v, 101)
		v, e = xdb.HIncrId(name, tk.Incr, 3)
		t.Assert(e, nil)
		t.Assert(v, 102)
		v, e = xdb.HIncrId(name, tk.Incr, 3)
		t.Assert(e, nil)
		t.Assert(v, 103)
	})
	// error
}
func TestHKeys(t *testing.T) {
	//  users
	gtest.C(t, func(t *gtest.T) {
		keys, err := xdb.HKeys(nameUser)
		t.Assert(err, nil)
		t.Assert(keys, []string{"user1", "user2", "user3"})
	})
	// errors
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.HKeys(nameNotExist)
		t.Assert(err, nil)
		t.Assert(v, g.Slice{})
	})
}

func TestHLen(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.HLen(nameUser)
		t.Assert(err, nil)
		t.Assert(v, 3)
	})
	// errors
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.HLen(nameNotExist)
		t.Assert(err, nil)
		t.Assert(v, 0)
	})
}

func TestHMDel(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		keys, err := xdb.HKeys(nameUser)
		t.Assert(err, nil)
		t.Assert(keys, keysUser)
		e := xdb.HmDel(nameUser, keysUser)
		t.Assert(e, nil)
		keys, err = xdb.HKeys(nameUser)
		t.Assert(err, nil)
		t.Assert(keys, []string{})
	})
	_ = xdb.HmSet(nameUser, kvm)
}

func TestHMDelByPrefix(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		keys, err := xdb.HKeys(nameUser)
		t.Assert(err, nil)
		t.Assert(keys, keysUser)
		e := xdb.HmDelByPrefix(nameUser, "user")
		t.Assert(e, nil)
		keys, err = xdb.HKeys(nameUser)
		t.Assert(err, nil)
		t.Assert(keys, []string{})
	})
	_ = xdb.HmSet(nameUser, kvm)
}

func TestHMGet(t *testing.T) {
	_ = xdb.HClear(nameUser)
	gtest.C(t, func(t *gtest.T) {
		err := xdb.HmSet(nameUser, kvm)
		t.Assert(err, nil)
		vs, err := xdb.HmGet(nameUser, keysUser)
		t.Assert(err, nil)
		t.Assert(vs, []g.MapStrAny{
			{"Age": 1, "Id": 1, "Name": "user1"}, {"Age": 2, "Id": 2, "Name": "user2"}, {"Age": 3, "Id": 3, "Name": "user3"},
		})
	})
	gtest.C(t, func(t *gtest.T) {
		vs, err := xdb.HmGet(tk.NotExist, keysUser)
		t.Assert(err, nil)
		t.Assert(len(vs), 0)
	})
}

func TestHMGetTo(t *testing.T) {
	_ = xdb.HClear(name)
	gtest.C(t, func(t *gtest.T) {
		e := xdb.HmSet(name, kvm)
		t.Assert(e, nil)
		vs := make([]User, 0)
		err := xdb.HmGetTo(name, keysUser, &vs)
		t.Assert(err, nil)
		t.Assert(vs, []User{user1, user2, user3})
	})

}
func TestHMSet(t *testing.T) {
}

func TestHScan(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		_ = xdb.HSet(name, "S1", 111)
		_ = xdb.HSet(name, "S2", 222)
		_ = xdb.HSet(name, "S3", 333)
		kvm, _, err := xdb.HScan(name, 0, "S*", 0)
		t.Assert(err, nil)
		t.Assert(kvm["S1"], 111)
		t.Assert(kvm["S2"], 222)
	})

}

func TestHScanMatch(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		_ = xdb.HSet(name, "S1", 111)
		_ = xdb.HSet(name, "S2", 222)
		_ = xdb.HSet(name, "S3", 333)
		kvm, err := xdb.HScanMatch(name, "S*")
		t.Assert(err, nil)
		t.Assert(kvm["S1"], 111)
		t.Assert(kvm["S2"], 222)
	})

}

func TestHSet(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.HSet(name, key, value), nil)
		v, e := xdb.HGet(name, key)
		t.Assert(e, nil)
		t.Assert(v, value)
	})
}

func TestHValues(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		vs, err := xdb.HValues(nameUser)
		t.Assert(err, nil)
		t.Assert(vs, []g.MapStrAny{
			{"Age": 1, "Id": 1, "Name": "user1"}, {"Age": 2, "Id": 2, "Name": "user2"}, {"Age": 3, "Id": 3, "Name": "user3"},
		})
	})
}

func TestHValuesTo(t *testing.T) {
}
