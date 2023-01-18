package xx

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
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
		a.Memory.Total = R2(float64(v.Total) / GB)
		a.Memory.Free = R2(float64(v.Free) / GB)
		a.Memory.Used = R2(float64(v.Used) / GB)
		a.Memory.UsedPercent = R2(v.UsedPercent)
	}
	if v, err := disk.Usage("/"); err == nil {
		a.Disk.Total = R2(float64(v.Total) / GB)
		a.Disk.Free = R2(float64(v.Free) / GB)
		a.Disk.Used = R2(float64(v.Used) / GB)
		a.Disk.UsedPercent = R2(v.UsedPercent)
	}
	if v, err := load.Avg(); err == nil {
		a.Load.Avg1 = R2(v.Load1)
		a.Load.Avg5 = R2(v.Load5)
		a.Load.Avg15 = R2(v.Load15)
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
	folder := IfStr(runtime.GOOS == "darwin", "/Users/d/z", "/d/z")
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
