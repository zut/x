module github.com/zut/x

go 1.18

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/gogf/gf v1.16.6
	github.com/ip2location/ip2location-go v8.3.0+incompatible
	github.com/labstack/gommon v0.3.1
	github.com/pkg/errors v0.9.1
	github.com/shirou/gopsutil/v3 v3.21.12
	github.com/vmihailenco/msgpack/v5 v5.3.5
	golang.org/x/crypto v0.0.0-20211117183948-ae814b36b871
	gonum.org/v1/gonum v0.11.0
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/clbanning/mxj v1.8.5-0.20200714211355-ff02cfb8ea28 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fatih/color v1.12.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gomodule/redigo v1.8.5 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/grokify/html-strip-tags-go v0.0.0-20190921062105-daaa06bf1aaf // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	github.com/tklauser/numcpus v0.3.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go.opentelemetry.io/otel v1.0.0-RC2 // indirect
	go.opentelemetry.io/otel/trace v1.0.0-RC2 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

// latest
// go get github.com/zut/x@master
// go get -u -v github.com/gogf/gf

// Go update all modules
// go get -u
// go mod tidy
// >>> change version to latest
// go mod download
