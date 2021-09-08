module github.com/zut/x

go 1.16

require (
	github.com/go-redis/redis/v8 v8.11.3
	github.com/gogf/gf v1.16.5
	github.com/labstack/gommon v0.3.0
	github.com/montanaflynn/stats v0.6.6
	github.com/pkg/errors v0.9.1
	github.com/vmihailenco/msgpack/v5 v5.3.4
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
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
