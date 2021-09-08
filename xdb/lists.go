package xdb

func LIndex(key string, index int) (string, error) {
	// 返回列表里的元素的索引 index 存储在 key 里面。 下标是从0开始索引的，所以
	//  0 是表示第一个元素，1 表示第二个元素
	//  -1 表示最后一个元素，-2 表示倒数第二个元素
	// Error:
	// index 超过范围的时候
	// key 位置的值不是一个列表的时候
	c := Conn()
	defer ConnClose(c)
	return c.LIndex(ctx, key, int64(index)).Result()
}

func LInsertAfter(key string, pivot, value interface{}) (int, error) {
	c := Conn()
	defer ConnClose(c)
	return toIntErr(c.LInsertAfter(ctx, key, pivot, value).Result())
}

func LInsertBefore(key string, pivot, value interface{}) (int, error) {
	// 1 把 value 插入存于 key 的列表中在基准值 pivot 的前面或后面。
	//  1.1 当 key 不存在时，这个list会被看作是空list，任何操作都不会发生。
	//  1.2 当 key 存在，但保存的不是一个list的时候，会返回error。
	// 2 返回值
	//  2.1 integer-reply: 经过插入操作后的list长度，或者当 pivot 值找不到的时候返回 -1。
	c := Conn()
	defer ConnClose(c)
	return toIntErr(c.LInsertBefore(ctx, key, pivot, value).Result())
}

func LLen(key string) (int, error) {
	c := Conn()
	defer ConnClose(c)
	return toIntErr(c.LLen(ctx, key).Result())
}

func LPop(key string) (string, error) {
	// 移除并且返回 key 对应的 list 的第一个元素。
	c := Conn()
	defer ConnClose(c)
	return c.LPop(ctx, key).Result()
}

func LPush(key string, value interface{}) error {
	// 将所有指定的值插入到存于 key 的列表的头部。
	//  如果 key 不存在，那么在进行 push 操作前会创建一个空列表。
	//  如果 key 对应的值不是一个 list 的话，那么会返回一个错误。

	// 可以使用一个命令把多个元素 push 进入列表，只需在命令末尾加上多个指定的参数。
	//  元素是从最左端的到最右端的、一个接一个被插入到 list 的头部。
	//  所以对于这个命令例子 LPUSH mylist a b c，返回的列表是
	//  c 为第一个元素， b 为第二个元素， a 为第三个元素。
	c := Conn()
	defer ConnClose(c)
	return c.LPush(ctx, key, value).Err()
}

func LPushX(key string, value interface{}) error {
	// 只有当 key 已经存在并且存着一个 list 的时候，在这个 key 下面的 list 的头部插入 value。
	//  与 LPUSH 相反，当 key 不存在的时候不会进行任何操作。
	c := Conn()
	defer ConnClose(c)
	return c.LPushX(ctx, key, value).Err()
}

func LRange(key string, start, stop int) ([]string, error) {
	// 0/1 -1/-2
	// 如果你有一个List，里面的元素是从0到100，那么 LRange List 0 10 这个命令会返回11个元素，即最右边的那个元素也会被包含在内。
	c := Conn()
	defer ConnClose(c)
	return c.LRange(ctx, key, int64(start), int64(stop)).Result()
}

func LRem(key string, count int, value interface{}) (int, error) {
	// remove
	// 从存于 key 的列表里移除前 count 次出现的值为 value 的元素。 这个 count 参数通过下面几种方式影响这个操作：
	//  count > 0: 从头往尾移除值为 value 的元素。
	//  count < 0: 从尾往头移除值为 value 的元素。
	//  count = 0: 移除所有值为 value 的元素。
	// 比如， LRem list -2 “hello” 会从存于 list 的列表里移除最后两个出现的 “hello”。
	// 需要注意的是，如果list里没有存在key就会被当作空list处理，所以当 key 不存在的时候，这个命令会返回 0。
	// 返回值:
	//  integer-reply: 被移除的元素个数。
	c := Conn()
	defer ConnClose(c)
	return toIntErr(c.LRem(ctx, key, int64(count), value).Result())
}

func LSet(key string, index int, value interface{}) error {
	// 设置 index 位置的list元素的值为 value。 更多关于 Index 参数的信息，详见 LIndex。
	// 当index超出范围时会返回一个error。 (ERR index out of range)
	c := Conn()
	defer ConnClose(c)
	return c.LSet(ctx, key, int64(index), value).Err()
}

func LTrim(key string, start, stop int) error {
	// 修剪(trim)一个已存在的 list
	//  start 和 stop 都是由0开始计数的，  0 是 第一个元素（表头），1 是第二个元素， -1 表示列表里的最后一个元素
	//  例如： LTrim foobar 0 2 将会对存储在 foobar 的列表进行修剪，只保留列表里的前3个元素。
	//   超过范围的下标并不会产生错误：如果 start 超过列表尾部，或者 start > end，结果会是列表变成空表（即该 key 会被移除）。
	//   如果 end 超过列表尾部，Redis 会将其当作列表的最后一个元素。
	// LTrim 的一个常见用法是和 LPush / RPush 一起使用。 例如：
	//  LPush MyList Some Element
	//  LTrim MyList 0 99
	// 这一对命令会将一个新的元素 push 进列表里，并保证该列表不会增长到超过100个元素。
	// 这个是很有用的，比如当用 Redis 来存储日志。 因为平均情况下，每次只有一个元素会被移除。
	c := Conn()
	defer ConnClose(c)
	return c.LTrim(ctx, key, int64(start), int64(stop)).Err()
}

func RPop(key string) (string, error) {
	// 移除并返回存于 key 的 list 的 最后 一个元素。
	// bulk-string-reply: 最后一个元素的值，或者当 key 不存在的时候返回 nil。
	c := Conn()
	defer ConnClose(c)
	return c.RPop(ctx, key).Result()
}

func RPopLPush(src, dst string) (string, error) {
	// 原子性地返回并移除存储在 src 的列表的最后一个元素（列表尾部元素）
	// 并把该元素放入存储在 dst 的列表的第一个元素位置（列表头部）。
	// 例如：
	//      src  a,b,c
	//      dst  x,y,z
	// 执行 RPopLPush 得到的结果是
	//      src  a,b
	//      dst  c,x,y,z
	c := Conn()
	defer ConnClose(c)
	return c.RPopLPush(ctx, src, dst).Result()
}

func RPush(key string, value interface{}) error {
	// 向存于 key 的列表的尾部插入所有指定的值。
	c := Conn()
	defer ConnClose(c)
	return c.RPush(ctx, key, value).Err()
}

func RPushX(key string, value interface{}) error {
	// 将值 value 插入到列表 key 的表尾, 当且仅当 key 存在并且是一个列表。
	//  和 RPush 命令相反, 当 key 不存在时，RPushX 命令什么也不做。
	c := Conn()
	defer ConnClose(c)
	return c.RPushX(ctx, key, value).Err()
}
