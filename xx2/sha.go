package xx2

import (
	"encoding/hex"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/util/gconv"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
	"io"
	"os"
)

// Sha1 returns the SHA1 digest of the data. (Result Length: 40 bits)
func Sha1(v interface{}) string {
	return gsha1.Encrypt(v)
}

// Sha3 returns the SHA3-256 digest of the data. (Result Length : 64 bits)
func Sha3(v interface{}) string {
	return Sha3Bytes(gconv.Bytes(gconv.String(v))) // 一定要 gconv.String(v) 否者每次都不同sha3
}

// Sha3Bytes returns the SHA3-256 digest of the data. (Result Length : 64 bits)
func Sha3Bytes(v []byte) string {
	r := sha3.Sum256(v) //
	return hex.EncodeToString(r[:])
}

// Sha3File returns the SHA3-256 digest of the data. (Result Length : 64 bits)
func Sha3File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", errors.Wrap(err, "Sha3File.os.Open(path)")
	}
	defer f.Close()
	h := sha3.New256()
	if _, err = io.Copy(h, f); err != nil {
		return "", errors.Wrap(err, "Sha3File.io.Copy(h, f)")
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
