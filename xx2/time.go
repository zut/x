package xx2

import (
	"fmt"
	"github.com/gogf/gf/os/glog"
	"os"
	"time"

	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/gfile"
)

func T1(title ...string) {
	key := "Time"
	if len(title) > 0 {
		key = title[0]
	}
	_ = gcache.Set(key, time.Now(), 0) // 改成先进后出, 剥洋葱的方式, 嵌套多个
}

func T2(title ...string) float64 {
	key := "Time"
	if len(title) > 0 {
		key = title[0]
	}
	i, err := gcache.Get(key)
	if err != nil {
		glog.Error(key, err)
		return 0
	}
	elapsed := time.Since(i.(time.Time))
	glog.Debug(key+" elapsed = ", elapsed)
	return elapsed.Seconds()
}

// Sleep  Second
func Sleep(s float64, show ...int) {
	if First(show) == 1 {
		glog.Debug("Sleep", s)
	}
	time.Sleep(time.Nanosecond * time.Duration(s*1e9))
}

func Zzz(title ...string) {
	if !gfile.Exists("config/debug.Zzz") {
		glog.Debug("Skip Zzz", "(not config/debug.Zzz)")
		return
	}
	if len(title) > 0 {
		glog.Info(title)
	}
	glog.Warning("Zzz ... Sleep ... Zzz")
	glog.Info("Continue \n  1.Yes \n 2.Stop")
	i := 0
	_, _ = fmt.Scanln(&i)
	if i != 1 {
		os.Exit(0)
	}
}
