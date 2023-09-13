package xdb

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/zut/x/xx"
	"time"
)

// 删除指定的一批keys，
// 如果删除中的某些key不存在，则直接忽略。
// 你无法通过返回值来判断被删除的 key 是否存在.
func Del(k string) error {
	c := Conn()
	defer ConnClose(c)
	return c.Del(ctx, k).Err()
}

func Exists(k string) (bool, error) {
	c := Conn()
	defer ConnClose(c)
	r, err := c.Exists(ctx, k).Result()
	return r == 1, err
}

// 1 如果成功设置过期时间。
// 0 如果key不存在或者不能设置过期时间。
func Expire(k string, ttl int) (bool, error) {
	c := Conn()
	defer ConnClose(c)
	return c.Expire(ctx, k, time.Duration(ttl)*time.Second).Result()
}

// ExpireAt 的作用和 Expire，都用于为 key 设置生存时间。
// 不同在于 ExpireAt 命令接受的时间参数是 Unix 时间戳 Unix timestamp 。
// time.Now().Add(time.Hour)
// time.Now().Add(900 * time.Millisecond)
func ExpireAt(k string, tm time.Time) (bool, error) {
	c := Conn()
	defer ConnClose(c)
	return c.ExpireAt(ctx, k, tm).Result()
}

// 返回key的value。
//
//	如果key不存在，返回特殊值nil。
//	如果key的value不是string，就返回错误，因为GET只处理string类型的values。
//
// 返回值
//
//	simple-string-reply:key对应的value，或者nil（key不存在时）
func Get(k string) (value interface{}, err error) {
	err = GetTo(k, &value)
	return
}
func GetStr(k string) (string, error) {
	v, err := Get(k)
	return xx.Str(v), err
}
func GetTo(k string, v interface{}) error {
	c := Conn()
	defer ConnClose(c)
	str, err := c.Get(ctx, k).Result()
	if err != nil {
		//xlog.Info("GetGo", k, err)
		return err
	}
	return xx.UnpackTo(gconv.Bytes(str), &v)
}

// 对存储在指定key的数值执行原子的加1操作。
func Incr(k string) (int, error) {
	return IncrBy(k, 1)
}

// 将key对应的数字加decrement。
// 1. 如果key不存在，操作之前，key就会被置为0。
// 2. 如果key的value类型错误或是个不能表示成数字 返回错误: ERR value is not an integer or out of range
func IncrBy(k string, v int) (int, error) {
	c := Conn()
	defer ConnClose(c)
	r, err := c.IncrBy(ctx, k, int64(v)).Result()
	return gconv.Int(r), err
}

// 查找所有符合给定模式pattern（正则表达式）的 key 。
// Warning: 生产环境使用 KEYS 命令需要非常小心。
// 在大的数据库上执行命令会影响性能, 不要在你的代码中使用 KEYS
// 如果你需要一个寻找键空间中的key子集，考虑使用SCAN 或 sets
//
// 支持的匹配模式 patterns:
//
//	h?llo 匹配 hello, hallo 和 hxllo
//	h*llo 匹配 hllo 和 heeeello
//	h[ae]llo 匹配 hello 和 hallo, 不匹配 hillo
//	h[^e]llo 匹配 hallo, hbllo, … 不匹配 hello
//	h[a-b]llo 匹配 hallo 和 hbllo
//	使用 \ 转义你想匹配的特殊字符。
func Keys(pattern string) ([]string, error) {
	c := Conn()
	defer ConnClose(c)
	return c.Keys(ctx, pattern).Result()
}

func KeysByPrefix(prefix string) ([]string, error) {
	return Scan(fmt.Sprintf("%v*", prefix))
}

func MDel(s []string) error {
	if len(s) == 0 {
		return nil
	}
	c := Conn()
	defer ConnClose(c)
	return c.Del(ctx, s...).Err()
}
func MDelByPrefix(prefix string) error {
	s, err := KeysByPrefix(prefix)
	if err != nil {
		return err
	}
	if len(s) == 0 {
		return nil
	}
	c := Conn()
	defer ConnClose(c)
	return c.Del(ctx, s...).Err()
}
func MGet(ks []string) ([]interface{}, error) {
	c := Conn()
	defer ConnClose(c)
	vsUnpack, err := c.MGet(ctx, ks...).Result()
	if err != nil {
		return nil, err
	}
	vs := make([]interface{}, 0)
	for _, v := range vsUnpack {
		if v == nil {
			continue
		}
		v2, err := xx.Unpack(gconv.Bytes(v))
		if err != nil {
			return nil, err
		}
		vs = append(vs, v2)
	}
	return vs, err
}

func MSet(kvm map[string]interface{}) error {
	kvmPack := make(map[string]interface{}, len(kvm))
	for k, v := range kvm {
		b, err := xx.Pack(v)
		if err != nil {
			return err
		}
		kvmPack[k] = b
	}
	c := Conn()
	defer ConnClose(c)
	return c.MSet(ctx, kvmPack).Err()
}
func Scan(match string) ([]string, error) {
	c := Conn()
	defer ConnClose(c)
	keyList, _, err := c.Scan(ctx, 0, match, 0).Result()
	return keyList, err
}

func Set(k string, v interface{}) error {
	return SetX(k, v, 0)
}

func SetX(k string, v interface{}, ttl int) error {
	b, err := xx.Pack(v)
	if err != nil {
		return err
	}
	c := Conn()
	defer ConnClose(c)
	return c.Set(ctx, k, b, time.Duration(ttl)*time.Second).Err()
}

// 如果key不存在或者已过期，返回 -2*time.Nanosecond
// 如果key存在并且没有设置过期时间（永久有效），返回 -1*time.Nanosecond
func TTL(k string) (time.Duration, error) {
	c := Conn()
	defer ConnClose(c)
	return c.TTL(ctx, k).Result()
}

// Type > string hash list set
// 如果key不存在或者已过期，返回 -2*time.Nanosecond
// 如果key存在并且没有设置过期时间（永久有效），返回 -1*time.Nanosecond
func Type(k string) (string, error) {
	c := Conn()
	defer ConnClose(c)
	return c.Type(ctx, k).Result()
}

// 包含 hash
func KeysIterator(prefix string) ([]string, error) {
	c := Conn()
	defer ConnClose(c)
	iter := c.Scan(ctx, 0, fmt.Sprintf("%v*", prefix), 1e6).Iterator()
	s := make([]string, 0)
	n := 0
	for iter.Next(ctx) {
		n++
		if n%2 == 0 {
			continue
		}
		s = append(s, iter.Val())
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return s, nil
}
