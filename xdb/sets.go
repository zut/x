package xdb

func SAdd(k string, v interface{}) error {
	c := Conn()
	defer ConnClose(c)
	return c.SAdd(ctx, k, v).Err()
}
func SCard(k string) (int, error) {
	// Returns the set cardinality (number of elements) of the set stored at key.
	// Return value
	// Integer reply: the cardinality (number of elements) of the set, or 0 if key does not exist.
	c := Conn()
	defer ConnClose(c)
	return toIntErr(c.SCard(ctx, k).Result())
}

func SDiff(k string) ([]string, error) {
	// Returns the members of the set resulting from the difference between the first set and all the successive sets.
	// For example:
	//  key1 = {a,b,c,d}
	//  key2 = {c}
	//  key3 = {a,c,e}
	//  SDiff key1 key2 key3 = {b,d}
	//  Keys that do not exist are considered to be empty sets.
	c := Conn()
	defer ConnClose(c)
	return c.SDiff(ctx, k).Result()
}
