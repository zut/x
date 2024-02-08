package xx2

import (
	"fmt"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/util/gconv"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/vmihailenco/msgpack/v5"
	"gonum.org/v1/gonum/floats/scalar"
	"reflect"
	"runtime"

	"github.com/ip2location/ip2location-go"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type PsInfoItem struct {
	Total       float64
	Used        float64
	Free        float64
	UsedPercent float64
}
type PsInfoLoad struct {
	Avg1  float64
	Avg5  float64
	Avg15 float64
}
type PsInfo struct {
	Memory PsInfoItem
	Disk   PsInfoItem
	Load   PsInfoLoad
	Uptime uint64
}

// GetPsInfo ...
func GetPsInfo() (a PsInfo) {
	if v, err := mem.VirtualMemory(); err == nil {
		a.Memory.Total = scalar.Round(float64(v.Total)/GB, 2)
		a.Memory.Free = scalar.Round(float64(v.Free)/GB, 2)
		a.Memory.Used = scalar.Round(float64(v.Used)/GB, 2)
		a.Memory.UsedPercent = scalar.Round(v.UsedPercent, 2)
	}
	if v, err := disk.Usage("/"); err == nil {
		a.Disk.Total = scalar.Round(float64(v.Total)/GB, 2)
		a.Disk.Free = scalar.Round(float64(v.Free)/GB, 2)
		a.Disk.Used = scalar.Round(float64(v.Used)/GB, 2)
		a.Disk.UsedPercent = scalar.Round(v.UsedPercent, 2)
	}
	if v, err := load.Avg(); err == nil {
		a.Load.Avg1 = scalar.Round(v.Load1, 2)
		a.Load.Avg5 = scalar.Round(v.Load5, 2)
		a.Load.Avg15 = scalar.Round(v.Load15, 2)
	}
	if uptime, err := host.Uptime(); err == nil {
		a.Uptime = uptime
	}
	return
}

type Ip2LocationRecord struct {
	CountryShort string
	Country      string
	Province     string // Region Province
	City         string
}

// Ip2Location ... Country / CountryShort / Province / City
func Ip2Location(ipList []string) ([]*Ip2LocationRecord, error) {
	folder := lo.Ternary(runtime.GOOS == "darwin", "/Users/d/z", "/d/z")
	dbPath := fmt.Sprintf("%v/%v", folder, "IP2LOCATION-LITE-DB3.BIN")
	if !gfile.Exists(dbPath) {
		return nil, fmt.Errorf("db file not found: %v", dbPath)
	}
	db, err := ip2location.OpenDB(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	s := make([]*Ip2LocationRecord, len(ipList))
	for n := range ipList {
		r, err := db.Get_all(ipList[n])
		if err != nil {
			return nil, err
		}
		s[n] = &Ip2LocationRecord{
			CountryShort: r.Country_short,
			Country:      r.Country_long,
			Province:     r.Region,
			City:         r.City,
		}
	}
	return s, nil
}

func IsPointer(value interface{}) error {
	k := reflect.ValueOf(value).Kind()
	if k != reflect.Ptr {
		return errors.Errorf("v is not Pointer: %v", gconv.String(k))
	}
	return nil
}

func IsUuid(i string) bool {
	// 标准的UUID格式为：123e4567-e89b-12d3-a456-426655440000 (8-4-4-4-12)
	// 简单校验
	//return gregex.IsMatchString(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`, i)
	return gregex.IsMatchString(`^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$`, i)
}

// Copy msgpack.Marshal to msgpack.Unmarshal
func Copy(src interface{}, dst interface{}) error {
	//IsPointer(dst) // dst must be a pointer
	data, err := msgpack.Marshal(src)
	if err != nil {
		return err
	}
	return msgpack.Unmarshal(data, dst)
}
