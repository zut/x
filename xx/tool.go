package xx

import (
	"bytes"
	"errors"
	"fmt"
	"image/png"
	"math/rand"
	"os"
	"reflect"
	"runtime"
	"sort"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/test/gtest"
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/labstack/gommon/color"
	"github.com/vmihailenco/msgpack/v5"
	"github.com/zut/x/vv"
	"github.com/zut/x/xlog"
)

func TryExpect() {
	if err := recover(); err != nil {
		//xlog.Errorf("%s: %s", e, debug.Stack()) // line 20
		xlog.Error(err)
		for i := 0; ; i++ {
			pc, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			if gregex.IsMatchString(`/(runtime|go/pkg)/`, file) {
				continue
			}
			fmt.Printf("%v %v:%v\n", pc, file, line)
		}
	}
}

func Reverse(s S) S {
	d := make([]interface{}, len(s))
	copy(d, s)
	last := len(s) - 1
	for i := 0; i < len(d)/2.0; i++ {
		d[i], d[last-i] = d[last-i], d[i]
	}
	return d
}

func ReverseST(s []string) []string {
	d := make([]string, len(s))
	copy(d, s)
	last := len(s) - 1
	for i := 0; i < len(d)/2.0; i++ {
		d[i], d[last-i] = d[last-i], d[i]
	}
	return d
}
func ReverseF64(s []float64) []float64 {
	d := make([]float64, len(s))
	copy(d, s)
	last := len(s) - 1
	for i := 0; i < len(d)/2.0; i++ {
		d[i], d[last-i] = d[last-i], d[i]
	}
	return d
}
func IsPointer(value interface{}) {
	if reflect.ValueOf(value).Kind() != reflect.Ptr {
		xlog.Panic("v is not Pointer: " + Str(reflect.ValueOf(value).Kind()))
	}
}

func Show(i interface{}) {
	xlog.Info(i, "Skip1")
	fmt.Println(color.Magenta(fmt.Sprintf("Type:(%T) +Value=(%+v)", i, i)))
}
func ShowBytes(i []byte) {
	xlog.Info(ST{Str(i)}, vv.Skip1)
}
func ShowDetail(i interface{}) {
	fmt.Printf("Type:(%T) +Value=(%+v)\n", i, i)
	if i == nil {
	} else if GetType(i) == GetType("") {
	} else if GetType(i) == GetType(errors.New("")) {
	} else if GetType(i) == GetType(1) {
	} else if GetType(i) == GetType(1.0) {
	} else {
		fmt.Printf("Value=%v\n", i)
		fmt.Printf("+Value=%s\n\n", Str(i))
	}
}
func ShowType(i interface{}) {
	fmt.Printf("Type=%T\n", i)
	//fmt.Println(reflect.TypeOf(i))
}
func GetType(i interface{}) reflect.Type {
	//fmt.Printf("%T\n", i)
	return reflect.TypeOf(i)
}

func RandomLetter(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
func RandomNumber(length int) string {
	const charset = "0123456789"
	var seededRand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
func RandomF64(min, max float64, decimals ...int) float64 {
	rand.Seed(time.Now().UnixNano())
	v := min + rand.Float64()*(max-min)
	return IfF64(len(decimals) > 0, Round(v, decimals...), v)
}

func RandomF64s(min, max float64, n int, decimals ...int) []float64 {
	s := make([]float64, n)
	for i := range s {
		s[i] = RandomF64(min, max, decimals...)
	}
	return s
}

// RandomInt [min to max] (include min max)
func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
func RandomInSF64(s ...float64) float64 {
	return s[RandomInt(0, len(s)-1)]
}
func RandomIntInSI(s []int) int {
	return s[RandomInt(0, len(s)-1)]
}

func RandomIntS(min, max int, n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = RandomInt(min, max)
	}
	return s
}
func RandomStrInST(s []string) string {
	if len(s) == 0 {
		return ""
	}
	return s[rand.Intn(len(s))]
}

func Join(s ...interface{}) string {
	return joinSep("_", 0, s...)
}

func JoinRmEmpty(sep string, s ...interface{}) string {
	return joinSep(sep, 1, s...)
}
func joinSep(sep string, removeEmptyStr int, s ...interface{}) string {

	s2 := make(g.Slice, 0)
	for _, i := range s {
		if removeEmptyStr == 1 && InST(Str(i), []string{"", "/"}) {
			continue
		}
		switch i.(type) {
		case int, int64:
		case float64:
		case string:
		default:
			xlog.Panic("joinSep Type Error", GetType(i), i)
		}
		s2 = append(s2, gstr.JoinAny(i, sep))
	}
	return gstr.JoinAny(s2, sep)
}

func MustIn(i string, s ...string) error { //contains
	for _, a := range s {
		if a == i {
			return nil
		}
	}
	return fmt.Errorf("ArgsError (%v) %v", i, s)
}
func MustInInt(i int, s ...int) error { //contains
	for _, a := range s {
		if a == i {
			return nil
		}
	}
	return fmt.Errorf("ArgsError (%v) %v", i, s)
}
func MustInF64(i float64, s ...float64) error { //contains
	for _, a := range s {
		if a == i {
			return nil
		}
	}
	return fmt.Errorf("ArgsError (%v) %v", i, s)
}
func IdxST(i string, s []string) int {
	for n, a := range s {
		if a == i {
			return n
		}
	}
	return 0
}

func InST(i string, s []string) bool { //contains
	for _, a := range s {
		if a == i {
			return true
		}
	}
	return false
}
func InSTIgnore(i string, s ...string) bool { //contains
	for _, a := range s {
		if gstr.ToUpper(a) == gstr.ToUpper(i) {
			return true
		}
	}
	return false
}
func SliceEqual(a S, b S) bool { //contains
	if len(a) != len(b) {
		return false
	}
	for n := range a {
		if a[n] != b[n] {
			return false
		}
	}
	return true
}
func SliceEqualST(a []string, b []string) bool { //contains
	if len(a) != len(b) {
		return false
	}
	for n := range a {
		if a[n] != b[n] {
			return false
		}
	}
	return true
}
func EqualST(a []string, b []string) bool {
	return SliceEqual(gconv.SliceAny(a), gconv.SliceAny(b))
}
func SliceFloat64Equal(a []float64, b []float64) bool { //contains
	return SliceEqual(gconv.SliceAny(a), gconv.SliceAny(b))
}

func SameMapStrStr(a map[string]string, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}
func InSSI(si []int, ssi [][]int) bool { //contains
	for _, a := range ssi {
		if EqualSI(si, a) {
			return true
		}
	}
	return false
}
func InSI(i int, s []int) bool { //contains
	for _, a := range s {
		if a == i {
			return true
		}
	}
	return false
}
func InSF(i float64, s []float64) bool { //contains
	for _, a := range s {
		if a == i {
			return true
		}
	}
	return false
}
func InSSF(i []float64, ss [][]float64) bool { //contains
	for _, a := range ss {
		if EqualSF(a, i) {
			return true
		}
	}
	return false
}
func EqualSF(a, b []float64) bool {
	if len(a) == len(b) {
		for n, i := range a {
			if i != b[n] {
				return false
			}
		}
		return true
	}
	return false
}

func EqualSI(a, b []int) bool {
	if len(a) == len(b) {
		for n, i := range a {
			if i != b[n] {
				return false
			}
		}
		return true
	}
	return false
}

// If evaluates a condition, if true returns the first parameter otherwise the second
func IfPanic(cdt bool, s ...interface{}) {
	if cdt {
		xlog.Panic(Str(s), vv.Skip1)
	}
}
func If(cdt bool, a interface{}, b interface{}) interface{} {
	if cdt {
		return a
	}
	return b
}

func IfStr(cdt bool, a string, b string) string {
	if cdt {
		return a
	}
	return b
}
func IfST(cdt bool, a []string, b []string) []string {
	if cdt {
		return a
	}
	return b
}
func IfS(cdt bool, a S, b S) S {
	if cdt {
		return a
	}
	return b
}
func IfSST(cdt bool, a [][]string, b [][]string) [][]string {
	if cdt {
		return a
	}
	return b
}
func IfInt(cdt bool, a int, b int) int {
	if cdt {
		return a
	}
	return b
}
func IfInt64(cdt bool, a int64, b int64) int64 {
	if cdt {
		return a
	}
	return b
}
func IfF64(cdt bool, a float64, b float64) float64 {
	if cdt {
		return a
	}
	return b
}
func IfSF(cdt bool, a []float64, b []float64) []float64 {
	if cdt {
		return a
	}
	return b
}
func IfSI(cdt bool, a []int, b []int) []int {
	if cdt {
		return a
	}
	return b
}
func IfErr(cdt bool, a error) error {
	if cdt {
		return a
	}
	return nil
}

func IfErr2(cdt bool, a, b error) error {
	if cdt {
		return a
	}
	return b
}

func OrInt(a int, b ...int) int {
	if len(b) > 0 {
		return b[0]
	}
	return a
}

func OrFloat(a float64, b ...float64) float64 {
	if len(b) > 0 {
		return b[0]
	}
	return a
}

func OrString(a string, b ...string) string {
	if len(b) > 0 {
		return b[0]
	}
	return a
}

func MapKvReverse(kv g.MapStrStr) g.MapStrStr {
	kvReverse := g.MapStrStr{}
	for k, v := range kv {
		kvReverse[v] = k
	}
	return kvReverse
}

func TestAssert12(t *gtest.T, e1, e2 error, v0, v1, v2 interface{}) {
	t.Assert(e1, nil)
	t.Assert(e2, nil)
	t.Assert(v1, v0)
	t.Assert(v2, v0)
}

// InInt value in Min Max
func InInt(i, min, max int) int {
	if i < min {
		i = min
	}
	if i > max {
		i = max
	}
	return i
}

// InF64 value in Min Max
func InF64(i, min, max float64) float64 {
	if i < min {
		i = min
	}
	if i > max {
		i = max
	}
	return i
}

func CompareStruct(aa interface{}, oo interface{}, ExceptKeys ...string) (keys []string) {
	IsPointer(aa)
	IsPointer(oo)
	val := reflect.ValueOf(aa).Elem()
	ooFields := reflect.Indirect(reflect.ValueOf(oo))
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		if InST(typeField.Name, ExceptKeys) {
			continue
		}
		v1 := val.Field(i)
		if oo != nil {
			v2 := ooFields.FieldByName(typeField.Name)
			if v1.Interface() == v2.Interface() {
				continue
			}
		}
		keys = append(keys, typeField.Name)
	}
	return keys
}

func ByteToPng(imgByte []byte, path string) error {
	img, err := png.Decode(bytes.NewReader(imgByte))
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	encoder := png.Encoder{CompressionLevel: png.BestCompression}
	//if err := png.Encode(f, img); err != nil {
	if err := encoder.Encode(f, img); err != nil {
		_ = f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func First(s S, defaultValue ...interface{}) interface{} {
	if len(s) > 0 {
		return s[0]
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return nil
}

func FirstF64(s []float64, defaultV ...float64) float64 {
	if len(s) > 0 {
		return s[0]
	} else if len(defaultV) > 0 {
		return defaultV[0]
	}
	return 0
}
func LastF64(s []float64, defaultV ...float64) float64 {
	if len(s) > 0 {
		return s[0]
	} else if len(defaultV) > 0 {
		return defaultV[0]
	}
	return 0
}
func FirstStr(s []string, defaultV ...string) string {
	if len(s) > 0 {
		return s[0]
	} else if len(defaultV) > 0 {
		return defaultV[0]
	}
	return ""
}
func LastStr(s []string, defaultV ...string) string {
	if len(s) > 0 {
		return s[len(s)-1]
	} else if len(defaultV) > 0 {
		return defaultV[0]
	}
	return ""
}

func FirstInt(s []int, defaultValue ...int) int {
	if len(s) > 0 {
		return s[0]
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}
func PartSI(s []int, maxLength int) []int {
	if len(s) > maxLength {
		return s[:maxLength-1]
	}
	return s
}
func PartSF(s []float64, maxLength int) []float64 {
	if len(s) > maxLength {
		return s[:maxLength-1]
	}
	return s
}
func FirstBool(s []bool, defaultValue ...bool) bool {
	if len(s) > 0 {
		return s[0]
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return false
}

func GetST(s []string, p int) string {
	if p >= len(s) {
		return "/"
	}
	if p < 0 {
		p = len(s) + p
	}
	return s[p]
}
func GetS(s []interface{}, p int) interface{} {
	if p >= len(s) {
		return ""
	}
	if p < 0 {
		p = len(s) + p
	}
	return s[p]
}
func GetSF(s []float64, p int) float64 {
	if p >= len(s) {
		return 999999
	}
	if p < 0 {
		p = len(s) + p
	}
	return s[p]
}
func GetSInt(s []int, p int) int {
	if p >= len(s) {
		return 999999
	}
	if p < 0 {
		p = len(s) + p
	}
	return s[p]
}

func F64Str(v float64, decimals ...int) string {
	p := FirstInt(decimals, 2)
	return fmt.Sprintf(fmt.Sprintf("%%0.%df", p), v)
}

func IsErr(e1 error, e2 error) bool {
	if e1 == nil && e2 == nil {
		return true
	} else if e1 != nil && e2 != nil && e1.Error() == e2.Error() {
		return true
	}
	return false
}

// Cross2Line line1: x1,x2 line2: y1,y2
func Cross2Line(x1, x2, y1, y2 float64) bool {
	if (y1 <= x2 && x2 <= y2) || (y1 <= x1 && x1 <= y2) || (x1 <= y1 && y2 <= x2) {
		return true
	}
	return false
}
func CrossPoint(s []float64, limit float64) (int, []int) {
	last := 0.0
	crossPoint := 0
	crossPosList := SI{}

	for n := range s {
		i := s[n]
		if n == 0 {
			last = i
			continue
		}
		if last > limit && i < limit {
			crossPoint++
			crossPosList = append(crossPosList, n)
			//fmt.Println(crossPoint,last, limit, i)
		} else if last < limit && i > limit {
			crossPoint++
			crossPosList = append(crossPosList, n)
			//fmt.Println(crossPoint,last, limit, i)
		}
		if i != limit {
			last = i
		}
	}
	return crossPoint, crossPosList
}
func Copy(src interface{}, dst interface{}) error {
	IsPointer(dst)
	data, err := msgpack.Marshal(src)
	if err != nil {
		return err
	}
	err = msgpack.Unmarshal(data, dst)
	return err
}

func IdxSI(i int, s SI) int {
	for n, aa := range s {
		if aa == i {
			return n
		}
	}
	xlog.Error("IdxSI Error", i, s)
	return -1
}
func IdxS(i interface{}, s S) int {
	for n, aa := range s {
		if aa == i {
			return n
		}
	}
	xlog.Error("SIndex Error", i, s)
	return -1
}
func SStrIndex(i string, s []string) int {
	for n, aa := range s {
		if aa == i {
			return n
		}
	}
	xlog.Error("SStrIndex Error", ST{i}, s, vv.Skip2)
	return -1
}
func SF64Index(i float64, s []float64) int {
	for n, aa := range s {
		if aa == i {
			return n
		}
	}
	xlog.Error("SStrIndex Error", i, s)
	return -1
}

func MapKeys(m map[string]interface{}) []string {
	s := make([]string, len(m))
	n := 0
	for k := range m {
		s[n] = k
		n++
	}
	sort.Strings(s)
	return s
}

func MapValueF64s(m map[string]float64) []float64 {
	s := make([]float64, len(m))
	n := 0
	for k := range m {
		s[n] = m[k]
		n++
	}
	return s
}

func FileExists(f string) bool {
	info, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	xlog.Debugf("Alloc = %vMB TotalAlloc = %vMB Sys = %vMB NumGC = %v",
		bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
func MergeStr(a, b []string) []string {
	s := CopySStr(a)
	s = append(s, b...)
	s = RemoveDuplicateStr(s)
	sort.Strings(s)
	return s
}
func RemoveDuplicateStr(s []string) []string {
	allKeys := make(map[string]bool)
	var s2 []string
	for _, item := range s {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			s2 = append(s2, item)
		}
	}
	return s2
}
func RemoveItemInST(s []string, removeItem string) []string {
	var s2 []string
	for _, item := range s {
		if item == removeItem {
			continue
		}
		s2 = append(s2, item)
	}
	return s2
}
func RemoveDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	var s2 []int
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			s2 = append(s2, item)
		}
	}
	return s2
}

func Default(OriginalVaule, v string) string {
	return IfStr(OriginalVaule == "", v, OriginalVaule)
}

func DefaultF64(OriginalVaule, v float64) float64 {
	return IfF64(OriginalVaule == 0, v, OriginalVaule)
}

func DefaultInt(OriginalVaule, v int) int {
	return IfInt(OriginalVaule == 0, v, OriginalVaule)
}
