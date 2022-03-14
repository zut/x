module github.com/zut/x

go 1.16

require (
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/go-redis/redis/v8 v8.11.4
	github.com/gogf/gf v1.16.6
	github.com/ip2location/ip2location-go v8.3.0+incompatible // indirect
	github.com/labstack/gommon v0.3.1
	github.com/pkg/errors v0.9.1
	github.com/shirou/gopsutil/v3 v3.21.12
	github.com/vmihailenco/msgpack/v5 v5.3.5
	golang.org/x/crypto v0.0.0-20211117183948-ae814b36b871
	golang.org/x/sys v0.0.0-20211117180635-dee7805ff2e1 // indirect
	gonum.org/v1/gonum v0.9.3
)

// latest
// go get github.com/zut/x@master
// go get -u -v github.com/gogf/gf

// Go update all modules
// go get -u
// go mod tidy
// >>> change version to latest
// go mod download
