package xx

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"os"
	"time"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
)

func T1(title ...interface{}) {
	if len(title) == 0 {
		title = append(title, "Time")
	}
	key := Str(title[0])
	_ = gcache.Set(gctx.New(), key, time.Now(), 0) // 改成先进后出, 剥洋葱的方式, 嵌套多个
}

func T2(title ...interface{}) {
	ctx := gctx.New()
	if len(title) == 0 {
		title = append(title, "Time")
	}
	key := Str(title[0])
	i, err := gcache.Get(ctx, key)
	if err != nil {
		glog.Error(ctx, err)
	}
	elapsed := time.Since(gconv.Time(i))
	glog.Debug(ctx, key+" elapsed = ", elapsed, "Skip1")
	//if len(gt) == 0 || time.Duration(gconv.Int(gt[0]*1e9)) <= elapsed {
	//	glog.Debug("T2 elapsed = ", elapsed, "Skip1")
	//}
}

// Sleep  Second
func Sleep(s float64, show ...int) {
	if FirstInt(show) == 1 {
		glog.Debug(gctx.New(), "Sleep", time.Nanosecond*time.Duration(int64(s*1e9)), "Skip1")
	}
	time.Sleep(time.Nanosecond * time.Duration(s*1e9))
}
func Zzz(title ...string) {
	if !gfile.Exists("config/debug.Zzz") {
		return
	}
	if len(title) > 0 {
		glog.Info(gctx.New(), title, "Skip1")
	}
	glog.Warning(gctx.New(), "Zzz ... Sleep ... Zzz", "Skip1")
	glog.Info(gctx.New(), "Continue \n  1.Yes \n 2.Stop", "Skip1")
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
