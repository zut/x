package xx

import (
	"context"
	"encoding/hex"
	"github.com/gogf/gf/v2/crypto/gsha1"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/crypto/sha3"
	"io"
	"os"
)

func Sha1(v interface{}) string {
	return gsha1.Encrypt(v)
}
func Sha3(v interface{}) string {
	r := sha3.Sum256(gconv.Bytes(Str(v))) // 一定要 xx.Str(v) 否者每次都不同sha3
	return hex.EncodeToString(r[:])
}
func Sha3Bytes(v []byte) string {
	r := sha3.Sum256(v) // 一定要 xx.Str(v) 否者每次都不同sha3
	return hex.EncodeToString(r[:])
}

func Sha3File(path string) string {
	ctx := context.TODO()
	f, err := os.Open(path)
	if err != nil {
		glog.Panic(ctx, err)
	}
	defer f.Close()
	h := sha3.New256()
	_, err = io.Copy(h, f)
	if err != nil {
		glog.Panic(ctx, err)
	}
	return hex.EncodeToString(h.Sum(nil))
}
