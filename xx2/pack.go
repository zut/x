package xx2

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"

	"github.com/gogf/gf/crypto/gaes"
	"github.com/vmihailenco/msgpack/v5"
)

func Pack(src interface{}) ([]byte, error) {
	return msgpack.Marshal(src)
}
func Unpack(src []byte) (dst interface{}, err error) {
	err = UnpackTo(src, &dst)
	return
}
func UnpackTo(src []byte, dst interface{}) (err error) {
	//if err = IsPointer(dst); err != nil { // 为了速度, 不检查指针
	//	return err
	//}
	return msgpack.Unmarshal(src, dst)
}

func PackCompress(src interface{}) ([]byte, error) {
	b, err := msgpack.Marshal(src)
	if err != nil {
		return nil, err
	}
	return CompressZLib(b)
}

func UnCompressUnpack(src []byte) (dst interface{}, err error) {
	src, err = UnCompressZLib(src)
	if err != nil {
		return nil, err
	}
	err = msgpack.Unmarshal(src, &dst)
	return
}

func UnCompressUnpackTo(src []byte, dst interface{}) (err error) {
	src, err = UnCompressZLib(src)
	if err != nil {
		return err
	}
	err = UnpackTo(src, &dst)
	return
}

func CompressZLib(src []byte) ([]byte, error) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	if _, err := w.Write(src); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func UnCompressZLib(src []byte) ([]byte, error) {
	b := bytes.NewReader(src)
	var out bytes.Buffer
	r, err := zlib.NewReader(b)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(&out, r); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

// PackCompressEncrypt
// Note that the key must be 16/24/32 bit length.
func PackCompressEncrypt(src interface{}, key, iv string) (dst []byte, err error) {
	dst, err = PackCompress(src)
	if err != nil {
		return
	}
	return gaes.Encrypt(dst, []byte(key), []byte(iv))
}
func DecryptUnpack(src []byte, key, iv string) (dst interface{}, err error) {
	if len(src) == 0 {
		return nil, fmt.Errorf("data.empty")
	}
	err = DecryptUnCompressUnpackTo(src, &dst, key, iv)
	return
}
func DecryptUnCompressUnpackTo(src []byte, dst interface{}, key, iv string) (err error) {
	if len(src) == 0 {
		return fmt.Errorf("data.empty")
	}
	if err = IsPointer(dst); err != nil {
		return err
	}
	var data []byte
	if data, err = gaes.Decrypt(src, []byte(key), []byte(iv)); err != nil {
		fmt.Println(err)
		return
	}
	return UnCompressUnpackTo(data, dst)
}

func Encrypt(plainText, key, iv string) (dst []byte, err error) {
	return gaes.Encrypt([]byte(plainText), []byte(key), []byte(iv))
}
