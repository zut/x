package main

import (
	"github.com/zut/x/xlog"
	"net"
	"time"
)

func main() {
	conn, err := net.DialTimeout("tcp", "10.0.0.22:5025", time.Second*3) // 无效
	//conn, err := net.Dial("tcp", "10.0.0.22:5025") // 无效
	xlog.Info(conn)
	xlog.Info(err.Error())
}
