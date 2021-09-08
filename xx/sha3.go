package xx

import (
	"encoding/hex"
	"github.com/gogf/gf/util/gconv"
	"github.com/zut/x/xlog"
	"golang.org/x/crypto/sha3"
	"io"
	"os"
)

func Sha3(v interface{}) string {
	r := sha3.Sum256(gconv.Bytes(Str(v))) // 一定要 xx.Str(v) 否者每次都不同sha3
	return hex.EncodeToString(r[:])
}

func Sha3File(path string) string {
	f, err := os.Open(path)
	if err != nil {
		xlog.Panic(err)
	}
	defer f.Close()
	h := sha3.New256()
	_, err = io.Copy(h, f)
	if err != nil {
		xlog.Panic(err)
	}
	return hex.EncodeToString(h.Sum(nil))
}
