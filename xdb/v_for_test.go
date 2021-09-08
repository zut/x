package xdb_test

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/grand"
	"math"
)

type Base struct {
	Age        int
	privateAge int // 私有属性不会进行转换
}
type User struct {
	Id   int
	Name string
	Base
}

var (
	user1 = User{Id: 1, Name: "user1", Base: Base{Age: 1, privateAge: 1}}
	user2 = User{Id: 2, Name: "user2", Base: Base{Age: 2, privateAge: 2}}
	user3 = User{Id: 3, Name: "user3", Base: Base{Age: 3, privateAge: 3}}
)

type testKey struct {
	String   string
	Bool     string
	Int64    string
	Float64  string
	Array    string
	Slice    string
	Map      string
	User1    string
	User2    string
	User3    string
	NotExist string
	Incr     string
	Lists    string
}

var tk = testKey{
	String: "String", // not to
	Bool:   "Bool",   // not to

	Int64:    "Int64",
	Float64:  "Float64",
	Array:    "Array",
	Slice:    "Slice",
	Map:      "Map",
	User1:    "user1",
	User2:    "user2",
	User3:    "user3",
	NotExist: "NotExist",
	Incr:     "Incr",
	Lists:    "Lists",
}

type testValue struct {
	Int64   int64
	Float64 float64
	String  string
	Bool    bool
	Array   [2]float64
	Slice   []int
	Map     map[string]int
	User1   User
	User2   User
	User3   User
}

var tv = testValue{
	String:  grand.S(10),
	Bool:    true,
	Int64:   64,
	Float64: math.Pi,
	Array:   [2]float64{.03, .02},
	Slice:   []int{1, 2},
	Map:     map[string]int{"a": 1, "b": 222},
	User1:   user1,
	User2:   user2,
	User3:   user3,
}

var name = "name"
var namePJ = "namePJ"
var nameZ = "nameZ"
var nameQ = "nameQ"
var nameUser = "nameUser"
var key = "kkk"
var value = "0123456789"
var nameNotExist = "nameNotExist"
var keysUser = g.SliceStr{tk.User1, tk.User2, tk.User3}
var kvm = g.Map{
	tk.User1: user1,
	tk.User2: user2,
	tk.User3: user3,
}

func init() {
	//_ = xdb.Clear()

	//ns, _ := xdb.ZList("", "", 1e6)
	//for _, n := range ns {
	//	_ = xdb.ZClear(n)
	//}
	//_ = xdb.ZSet(name, key, 1)
	//_ = xdb.ZSet(name, g.Slice{"a", 1, "z"}, 111)
	//_ = xdb.ZSet(name, g.Slice{"a", 2, "x"}, 333)
	//_ = xdb.ZSet(name, g.Slice{"a", 3, "y"}, 222)
	//_ = xdb.ZSet(name, g.Slice{"a", 4, "s"}, 222)
	//_ = xdb.ZSet(name, g.Slice{"a", 5, "b"}, 555)
	//_ = xdb.ZSet(name, g.Slice{"b", 11, "b"}, 1000)
	//_ = xdb.ZSet(name, g.Slice{"c", 22, "b"}, 2000)
	//_ = xdb.ZSet(name, g.Slice{"x", 1, "x"}, 1)
	//_ = xdb.ZSet(name, g.Slice{"z", 0, "z"}, 0)
	//_ = xdb.ZSet(nameZ, 1, 1)
	//_ = xdb.ZSet(nameZ, 2, 2)
	//_ = xdb.ZSet(nameZ, 3, 3)
	//_ = xdb.ZSet(nameZ, 4, 4)
	//_ = xdb.ZSet(nameZ, 5, 5)

}
