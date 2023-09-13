package xdb

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/zut/x/xx"
	"math"
)

func HClear(h string) error {
	if err := cEmpty(h); err != nil {
		return err
	}
	if err := Del(AddKK(h)); err != nil {
		return err
	}
	return Del(h)
}

func HDel(h, k string) error {
	if err := cEmpty(h, k); err != nil {
		return err
	}
	// 从 key 指定的哈希集中移除指定的域。在哈希集中不存在的域将被忽略。
	// 如果 key 指定的哈希集不存在，它将被认为是一个空的哈希集，该命令将返回0。
	c := Conn()
	defer ConnClose(c)
	if err := c.HDel(ctx, AddKK(h), k).Err(); err != nil {
		return err
	}
	return c.HDel(ctx, h, k).Err()
}

func HExists(h, k string) (bool, error) {
	if err := cEmpty(h, k); err != nil {
		return false, err
	}
	// 返回hash里面field是否存在
	// true  hash里面包含该field。
	// false hash里面不包含该field或key不存在。
	c := Conn()
	defer ConnClose(c)
	return c.HExists(ctx, h, k).Result()
}
func HExistsNotErr(h, k string) bool {
	b, _ := HExists(h, k)
	return b
}

func HGet(h, k string) (v interface{}, err error) {
	err = HGetTo(h, k, &v)
	return
}
func HGetStr(h, k string) (v string, err error) {
	err = HGetTo(h, k, &v)
	return
}
func HGetF64(h, k string) (v float64, err error) {
	err = HGetTo(h, k, &v)
	return
}
func HGetInt(h, k string) (v int, err error) {
	err = HGetTo(h, k, &v)
	return
}
func HGetInt64(h, k string) (v int64, err error) {
	err = HGetTo(h, k, &v)
	return
}

func HGetAllOriginal(h string) (map[string]string, error) {
	if err := cEmpty(h); err != nil {
		return nil, err
	}
	c := Conn()
	defer ConnClose(c)
	return c.HGetAll(ctx, h).Result()
}

func HGetAll(h string) (map[string]interface{}, error) {
	m, err := HGetAllOriginal(h)
	if err != nil {
		return nil, err
	}
	m2 := make(map[string]interface{}, len(m))
	for k, v := range m {
		v2, err := xx.Unpack(gconv.Bytes(v))
		if err != nil {
			return nil, err
		}
		m2[k] = v2
	}
	return m2, err
}

func HGetAllMapStrInt(h string) (map[string]int, error) {
	m, err := HGetAllOriginal(h)
	if err != nil {
		return nil, err
	}
	m2 := make(map[string]int, len(m))
	for k, v := range m {
		m2[k] = xx.Int(v)
	}
	return m2, err
}

func HGetNotUnpack(h, k string) (string, error) {
	if err := cEmpty(h, k); err != nil {
		return "", err
	}
	c := Conn()
	defer ConnClose(c)
	str, err := c.HGet(ctx, h, k).Result()

	return str, err
}
func HGetTo(h, k string, v interface{}) error {
	str, err := HGetNotUnpack(h, k)
	if err != nil {
		return err
	}
	return xx.UnpackTo(gconv.Bytes(str), &v)
}
func HIncrSet(h, k string, v int) (int, error) {
	c := Conn()
	defer ConnClose(c)
	r, err := c.HSet(ctx, h, k, v).Result()
	return gconv.Int(r), err
}

func HIncrGet(h, k string) (int, error) {
	c := Conn()
	defer ConnClose(c)
	r, err := c.HGet(ctx, h, k).Result()
	//return gconv.Int(r), errors.Wrap(err, "HIncrGet")
	return gconv.Int(r), err
}
func HIncr(h, k string) (int, error) {
	return HIncrBy(h, k, 1)
}
func HIncrBy(h, k string, v int) (int, error) {
	if err := cEmpty(h, k); err != nil {
		return 0, err
	}
	// 1. 如果key不存在，操作之前，key就会被置为0。
	// 2. 如果key的value类型错误或是个不能表示成数字 返回错误: ERR value is not an integer or out of range
	c := Conn()
	defer ConnClose(c)
	r, err := c.HIncrBy(ctx, h, k, int64(v)).Result()
	return int(r), err
}

func HIncrByFloat(key, field string, incr float64) (float64, error) {
	// 为指定key的hash的field字段值执行float类型的increment加。
	//  如果field不存在，则在执行该操作前设置为0.如果出现下列情况之一，则返回错误：field的值包含的类型错误(不是字符串)。
	//  当前field或者increment不能解析为一个float类型。
	//  此命令的确切行为与 IncrByFloat 命令相同，请参 IncrByFloat 命令获取更多信息。
	// 返回值
	//  bulk-string-reply： field执行increment加后的值
	c := Conn()
	defer ConnClose(c)
	return c.HIncrByFloat(ctx, key, field, incr).Result()
}

func HIncrId(h, k string, digit int) (int, error) {
	if err := cEmpty(h, k); err != nil {
		return 0, err
	}
	v, err := HIncrBy(h, k, 1)
	if err != nil {
		return v, err
	}
	//if decimals and v < 10 ** (decimals - 1):
	//v = self.execute_command('zincr', c(name, 1), c(key, 1), 10 ** (decimals - 1))
	if v < int(math.Pow10(digit-1)) {
		v, err = HIncrBy(h, k, int(math.Pow10(digit-1))-v+1)
	}
	return v, err
}

func HKeys(h string) ([]string, error) {
	if err := cEmpty(h); err != nil {
		return nil, err
	}
	// 返回 key 指定的哈希集中所有字段的名字。
	// Keys 要注意, 二 HKeys是可以用的
	c := Conn()
	defer ConnClose(c)
	s, err := c.HKeys(ctx, AddKK(h)).Result()
	if err != nil {
		return nil, err
	}
	size, err := HLen(h)
	if err != nil {
		return nil, err
	}
	if len(s) != size {
		return nil, fmt.Errorf("HKeys.Error,Hash=%v.HLen=%v,HKeys=%v", h, size, len(s))
	}
	return s, nil
}

func HKeysPrefix(h, prefix string) ([]string, error) {
	if err := cEmpty(h); err != nil {
		return nil, err
	}
	// 返回 key 指定的哈希集中所有字段的名字。
	kvm, _, err := HScan(AddKK(h), 0, fmt.Sprintf("%v*", prefix), 1e6, true)
	s := xx.MapKeys(kvm)
	return s, err
}

func HkeysPrefixIteratorOriginal(h, prefix string) ([]string, error) {
	// 用 Iterator 能保证取出来, 但是数量越大,时间越多
	c := Conn()
	defer ConnClose(c)
	iter := c.HScan(ctx, h, 0, fmt.Sprintf("%v*", prefix), 0).Iterator()
	s := make([]string, 0)
	n := 0
	for iter.Next(ctx) {
		n++
		if n%1000 == 0 {
			fmt.Sprintln(n)
		}
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

func HLen(h string) (int, error) {
	if err := cEmpty(h); err != nil {
		return 0, err
	}
	c := Conn()
	defer ConnClose(c)
	v, err := c.HLen(ctx, h).Result()
	return gconv.Int(v), err
}

func HmDel(h string, ks []string) error {
	if err := cEmpty(h); err != nil {
		return err
	}
	if len(ks) == 0 {
		return nil
	}
	c := Conn()
	defer ConnClose(c)
	if err := c.HDel(ctx, AddKK(h), ks...).Err(); err != nil {
		return err
	}
	return c.HDel(ctx, h, ks...).Err()
}

func HmDelByPrefix(h string, prefix string) error {
	ks, err := HKeysPrefix(h, prefix)
	if err != nil {
		return err
	}
	if len(ks) == 0 {
		return nil
	}
	c := Conn()
	defer ConnClose(c)
	if err := c.HDel(ctx, AddKK(h), ks...).Err(); err != nil {
		return err
	}
	return c.HDel(ctx, h, ks...).Err()
}

func HmGet(h string, ks []string) ([]interface{}, error) {
	if err := cEmpty(h); err != nil {
		return nil, err
	}
	c := Conn()
	defer ConnClose(c)
	vsUnpack, err := c.HMGet(ctx, h, ks...).Result()
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
func HmGetTo(h string, ks []string, p interface{}) error {
	if err := cEmpty(h); err != nil {
		return err
	}
	xx.IsPointer(p)
	if len(ks) == 0 {
		return nil
	}
	vs, err := HmGet(h, ks)
	if err != nil {
		return err
	}
	return gconv.Structs(vs, p)
}

func hmSet(h string, kvm map[string]interface{}, key string) error {
	err := cEmpty(h)
	if err != nil {
		return err
	}
	if len(kvm) == 0 {
		return nil
	}
	kvmPack := make(map[string]interface{}, len(kvm))
	kvmPackForKeys := make(map[string]interface{}, len(kvm))
	for k, v := range kvm {
		var b []byte
		if key == "" {
			b, err = xx.Pack(v)
			if err != nil {
				return err
			}
		} else {
			b, err = xx.PackCompressEncrypt(v, key)
			if err != nil {
				return err
			}
		}
		kvmPack[k] = b
		kvmPackForKeys[k] = 1
	}
	c := Conn()
	defer ConnClose(c)
	if err := c.HMSet(ctx, AddKK(h), kvmPackForKeys).Err(); err != nil {
		return err
	}
	return c.HMSet(ctx, h, kvmPack).Err()
}

func HmSet(h string, kvm map[string]interface{}) error {
	return hmSet(h, kvm, "")
}
func HmSetEncrypt(h string, kvm map[string]interface{}, key string) error {
	return hmSet(h, kvm, key)
}

func HmSetOriginal(h string, kvm map[string]interface{}) error {
	// tst pass  AddKK
	if len(kvm) == 0 {
		return nil
	}
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
	return c.HMSet(ctx, h, kvmPack).Err()
}

func HScanNotUnpack(h string, cursor uint64, prefix string, count int) ([]string, uint64, error) {
	c := Conn()
	defer ConnClose(c)
	kvList, cursorResp, err := c.HScan(ctx, h, cursor, fmt.Sprintf("%v*", prefix), int64(count)).Result()
	return kvList, cursorResp, err
}
func HScan(h string, cursor uint64, match string, count int, onlyKeys ...bool) (map[string]interface{}, uint64, error) {
	if err := cEmpty(h); err != nil {
		return nil, 0, err
	}
	// kvm  cursor err
	c := Conn()
	defer ConnClose(c)
	kvList, cursorResp, err := c.HScan(ctx, h, cursor, match, int64(count)).Result()
	if err != nil {
		return nil, 0, err
	}
	kvm := make(map[string]interface{}, 0)
	onlyKey := xx.FirstBool(onlyKeys)
	for i := 0; i < len(kvList); i += 2 {
		if onlyKey {
			kvm[kvList[i]] = 0
		} else {
			v, err := xx.Unpack(gconv.Bytes(kvList[i+1]))
			if err != nil {
				return nil, 0, err
			}
			kvm[kvList[i]] = v
		}

	}
	return kvm, cursorResp, err
}
func HScanMatch(h string, match string) (map[string]interface{}, error) {
	kvm, _, err := HScan(h, 0, match, 1e6)
	return kvm, err
}
func HScanPrefix(h string, prefix string) (map[string]interface{}, error) {
	kvm, _, err := HScan(h, 0, fmt.Sprintf("%v*", prefix), 1e6)
	return kvm, err
}

func hSet(h, k string, v interface{}, key string) error {
	err := cEmpty(h, k)
	if err != nil {
		return err
	}
	var b []byte
	if key == "" {
		b, err = xx.Pack(v)
		if err != nil {
			return err
		}
	} else {
		b, err = xx.PackCompressEncrypt(v, key)
		if err != nil {
			return err
		}
	}
	c := Conn()
	defer ConnClose(c)
	if err := c.HSet(ctx, AddKK(h), k, 1).Err(); err != nil {
		return err
	}
	return c.HSet(ctx, h, k, b).Err()
}

func HSet(h, k string, v interface{}) error {
	return hSet(h, k, v, "")
}
func HSetEncrypt(h, k string, v interface{}, key string) error {
	return hSet(h, k, v, key)
}

func HValuesNotUnpack(h string) ([]string, error) {
	if err := cEmpty(h); err != nil {
		return nil, err
	}
	c := Conn()
	defer ConnClose(c)
	return c.HVals(ctx, h).Result()
}

func HValues(h string) ([]interface{}, error) {
	vStrList, err := HValuesNotUnpack(h)
	if err != nil {
		return nil, err
	}
	vs := make([]interface{}, len(vStrList))
	for n, vStr := range vStrList {
		v, err := xx.Unpack(gconv.Bytes(vStr))
		if err != nil {
			return nil, err
		}
		vs[n] = v
	}
	return vs, err
}

func HValuesTo(h string, p interface{}) error {
	if err := cEmpty(h); err != nil {
		return err
	}
	xx.IsPointer(p)
	vs, err := HValues(h)
	if err != nil {
		return err
	}
	return gconv.Structs(vs, p)
}

func HValuesToByPrefix(h, prefix string, p interface{}) error {
	if err := cEmpty(h); err != nil {
		return err
	}
	xx.IsPointer(p)
	kvm, err := HScanPrefix(h, prefix)
	if err != nil {
		return err
	}
	vs := make([]interface{}, len(kvm))
	n := 0
	for k := range kvm {
		vs[n] = kvm[k]
		n++
	}
	return gconv.Structs(vs, p)
}

func HValuesByPrefix(h, prefix string) ([]interface{}, error) {
	if err := cEmpty(h); err != nil {
		return nil, err
	}
	kvm, err := HScanPrefix(h, prefix)
	if err != nil {
		return nil, err
	}
	vs := make([]interface{}, len(kvm))
	n := 0
	for k := range kvm {
		vs[n] = kvm[k]
		n++
	}
	return vs, err
}

func HValuesByPrefixStr(h, prefix string) ([]string, error) {
	s1, err := HValuesByPrefix(h, prefix)
	if err != nil {
		return nil, err
	}
	s2 := make([]string, len(s1))
	for n := range s1 {
		s2[n] = xx.Str(s1[n])
	}
	return s2, nil
}
