package xdb_test

import (
	"github.com/gogf/gf/test/gtest"
	"github.com/zut/x/xdb"
	"testing"
)

func init() {
	_ = xdb.Del(tk.Lists)
	_ = xdb.RPush(tk.Lists, "a")
	_ = xdb.RPush(tk.Lists, "b")
	_ = xdb.RPush(tk.Lists, "c")
	// ["a","b","c"]
}

func TestLIndex(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.LIndex(tk.Lists, 0)
		t.Assert(err, nil)
		t.Assert(v, "a")
	})
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.LIndex(tk.Lists, 1)
		t.Assert(err, nil)
		t.Assert(v, "b")
	})
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.LIndex(tk.Lists, -1)
		t.Assert(err, nil)
		t.Assert(v, "c")
	})
	// error 当 index 超过范围的时候，会返回一个error
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.LIndex(key, 4)
		t.Assert(err, xdb.Nil)
		t.Assert(v, "")
	})
	// error 当 key 位置的值不是一个列表的时候
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.Set(key, value), nil)
		v, err := xdb.LIndex(key, 0)
		t.Assert(err, xdb.Nil)
		t.Assert(v, "")
	})
}

func TestLInsertBefore(t *testing.T) {
	// 1.1 当 key 不存在时，这个list会被看作是空list，任何操作都不会发生。
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.LInsertBefore(tk.Lists, "a", "vBefore")
		t.Assert(err, nil)
		t.Assert(v, 3+1)
		vs, err := xdb.LRange(tk.Lists, 0, -1)
		t.Assert(err, nil)
		t.Assert(vs, []string{"vBefore", "a", "b", "c"})
	})
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.LInsertBefore(tk.NotExist, "a", "vBefore")
		t.Assert(err, nil)
		t.Assert(v, 0)
	})
	// error
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.LInsertBefore(tk.Lists, "pivot", "vBefore")
		t.Assert(err, nil)
		t.Assert(v, -1) // 当 pivot 值找不到的时候返回 -1。
	})
	_, _ = xdb.LRem(tk.Lists, 0, "vBefore")
}

func TestLInsertAfter(t *testing.T) {
	// 1.1 当 key 不存在时，这个list会被看作是空list，任何操作都不会发生。
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.LInsertAfter(tk.Lists, "a", "vAfter")
		t.Assert(err, nil)
		t.Assert(v, 3+1)
		vs, err := xdb.LRange(tk.Lists, 0, -1)
		t.Assert(err, nil)
		t.Assert(vs, []string{"a", "vAfter", "b", "c"})
	})
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.LInsertAfter(tk.NotExist, "a", "vAfter")
		t.Assert(err, nil)
		t.Assert(v, 0)
	})
	// error
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.LInsertAfter(tk.Lists, "pivot", "vAfter")
		t.Assert(err, nil)
		t.Assert(v, -1) // 当 pivot 值找不到的时候返回 -1。
	})
	_, _ = xdb.LRem(tk.Lists, 0, "vAfter")
}

func TestLLen(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.LLen(tk.Lists)
		t.Assert(err, nil)
		t.Assert(v, 3)
	})
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.LLen(tk.NotExist)
		t.Assert(err, nil)
		t.Assert(v, 0)
	})
}

func TestLPop(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.LPop(tk.Lists)
		t.Assert(err, nil)
		t.Assert(v, "a")
	})
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.LPop(tk.NotExist)
		t.Assert(err, xdb.Nil)
		t.Assert(v, "")
	})
	_ = xdb.LPush(tk.Lists, "a")
}

func TestLPush(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.LPush(tk.Lists, "d"), nil)
		vs, err := xdb.LRange(tk.Lists, 0, -1)
		t.Assert(err, nil)
		t.Assert(vs, []string{"d", "a", "b", "c"})
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.LPush(tk.NotExist, "d"), nil)
		vs, err := xdb.LRange(tk.NotExist, 0, -1)
		t.Assert(err, nil)
		t.Assert(vs, []string{"d"})
	})
	_ = xdb.Del(tk.NotExist)
	_, _ = xdb.LRem(tk.Lists, 0, "d")
}

func TestLPushX(t *testing.T) {
	_ = xdb.Del(tk.NotExist)
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.LPushX(tk.NotExist, "d"), nil)
		vs, err := xdb.LRange(tk.NotExist, 0, -1)
		t.Assert(err, nil)
		t.Assert(vs, []string{})
	})
}

func TestLRange(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		vs, err := xdb.LRange(tk.Lists, 0, -1)
		t.Assert(err, nil)
		t.Assert(vs, []string{"a", "b", "c"})
	})
}

func TestLRem(t *testing.T) {
	_ = xdb.LPush(tk.Lists, "x")
	_ = xdb.RPush(tk.Lists, "y")
	_ = xdb.RPush(tk.Lists, "y")
	_ = xdb.RPush(tk.Lists, "z")
	// "x", "a", "b", "c", "y", "y", "z"

	// count > 0: 从头往尾移除值为 value 的元素。
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.LRem(tk.Lists, 2, "x")
		t.Assert(err, nil)
		t.Assert(v, 1)

		vs, err := xdb.LRange(tk.Lists, 0, -1)
		t.Assert(vs, []string{"a", "b", "c", "y", "y", "z"})
	})
	// count < 0: 从尾往头移除值为 value 的元素。
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.LRem(tk.Lists, -2, "y")
		t.Assert(err, nil)
		t.Assert(v, 2)

		vs, err := xdb.LRange(tk.Lists, 0, -1)
		t.Assert(vs, []string{"a", "b", "c", "z"})
	})
	//  count = 0: 移除所有值为 value 的元素。
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.LRem(tk.Lists, 0, "z")
		t.Assert(err, nil)
		t.Assert(v, 1)

		vs, err := xdb.LRange(tk.Lists, 0, -1)
		t.Assert(vs, []string{"a", "b", "c"})
	})
}

func TestLSet(t *testing.T) {
	// count > 0: 从头往尾移除值为 value 的元素。
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.LSet(tk.Lists, 1, "bbb"), nil)
		vs, err := xdb.LRange(tk.Lists, 0, -1)
		t.Assert(err, nil)
		t.Assert(vs, []string{"a", "bbb", "c"})
		t.Assert(xdb.LSet(tk.Lists, 1, "b"), nil)
	})
	// errors
	gtest.C(t, func(t *gtest.T) {
		err := xdb.LSet(tk.Lists, 999, "222")
		t.Assert(err, "ERR index out of range")
	})
}

func TestLTrim(t *testing.T) {
	_ = xdb.LPush(tk.Lists, "x")
	_ = xdb.RPush(tk.Lists, "y")
	_ = xdb.RPush(tk.Lists, "y")
	_ = xdb.RPush(tk.Lists, "z")
	// "x", "a", "b", "c", "y", "y", "z"

	// count > 0: 从头往尾移除值为 value 的元素。
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.LTrim(tk.Lists, 1, -1), nil)
		vs, err := xdb.LRange(tk.Lists, 0, -1)
		t.Assert(err, nil)
		t.Assert(vs, []string{"a", "b", "c", "y", "y", "z"})
	})
	// count < 0: 从尾往头移除值为 value 的元素。
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.LTrim(tk.Lists, 0, -4), nil)

		vs, err := xdb.LRange(tk.Lists, 0, -1)
		t.Assert(err, nil)
		t.Assert(vs, []string{"a", "b", "c"})
	})
}

func TestRPop(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.RPop(tk.Lists)
		t.Assert(err, nil)
		t.Assert(v, "c")
	})
	// count < 0: 从尾往头移除值为 value 的元素。
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.RPop(tk.NotExist)
		t.Assert(err, xdb.Nil)
		t.Assert(v, "")
	})
	_ = xdb.RPush(tk.Lists, "c")
}

func TestRPopLPush(t *testing.T) {
	dst := "dst"
	_ = xdb.Del(dst)
	_ = xdb.LPush(dst, "x")
	_ = xdb.RPush(dst, "y")
	_ = xdb.RPush(dst, "z")
	gtest.C(t, func(t *gtest.T) {
		v, err := xdb.RPopLPush(tk.Lists, dst)
		t.Assert(err, nil)
		t.Assert(v, "c")

		vs, err := xdb.LRange(tk.Lists, 0, -1)
		t.Assert(err, nil)
		t.Assert(vs, []string{"a", "b"})

		vs, err = xdb.LRange(dst, 0, -1)
		t.Assert(err, nil)
		t.Assert(vs, []string{"c", "x", "y", "z"})
	})

	_ = xdb.RPush(tk.Lists, "c")
	_ = xdb.Del(dst)
}

func TestRPush(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.RPush(tk.Lists, "d"), nil)
		vs, err := xdb.LRange(tk.Lists, 0, -1)
		t.Assert(err, nil)
		t.Assert(vs, []string{"a", "b", "c", "d"})
	})
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.RPush(tk.NotExist, "d"), nil)
		vs, err := xdb.LRange(tk.NotExist, 0, -1)
		t.Assert(err, nil)
		t.Assert(vs, []string{"d"})
	})
	_ = xdb.Del(tk.NotExist)
	_, _ = xdb.LRem(tk.Lists, 0, "d")
}

func TestRPushX(t *testing.T) {
	_ = xdb.Del(tk.NotExist)
	gtest.C(t, func(t *gtest.T) {
		t.Assert(xdb.RPushX(tk.NotExist, "d"), nil)
		vs, err := xdb.LRange(tk.NotExist, 0, -1)
		t.Assert(err, nil)
		t.Assert(vs, []string{})
	})
}
