package xx

import (
	"fmt"
	"os"
	"time"

	"github.com/gogf/gf/os/gtime"

	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/gconv"
	"github.com/zut/x/xlog"
)

func T1(title ...interface{}) {
	if len(title) == 0 {
		title = append(title, "Time")
	}
	key := Str(title[0])
	gcache.Set(key, time.Now(), 0) // 改成先进后出, 剥洋葱的方式, 嵌套多个
}

func T2(title ...interface{}) {
	if len(title) == 0 {
		title = append(title, "Time")
	}
	key := Str(title[0])
	i, err := gcache.Get(key)
	if err != nil {
		xlog.Error(err)
	}
	elapsed := time.Since(gconv.Time(i))
	xlog.Debug(key+" elapsed = ", elapsed, "Skip1")
	//if len(gt) == 0 || time.Duration(gconv.Int(gt[0]*1e9)) <= elapsed {
	//	xlog.Debug("T2 elapsed = ", elapsed, "Skip1")
	//}
}

// Second
func Sleep(s float64, show ...int) {
	if FirstInt(show) == 1 {
		xlog.Debug("Sleep", time.Nanosecond*time.Duration(int64(s*1e9)), "Skip1")
	}
	time.Sleep(time.Nanosecond * time.Duration(s*1e9))
}
func Zzz(title ...string) {
	if !gfile.Exists("config/debug.Zzz") {
		return
	}
	if len(title) > 0 {
		xlog.Info(title, "Skip1")
	}
	xlog.Warning("Zzz ... Sleep ... Zzz", "Skip1")
	xlog.Info("Continue \n  1.Yes \n 2.Stop", "Skip1")
	i := 0
	_, _ = fmt.Scanln(&i)
	if i != 1 {
		os.Exit(0)
	}
}

// TodayDate ... Y-m-d
func TodayDate() string {
	return gtime.Now().Format("Y-m-d")
}
