package xx

import (
	"bytes"
	"compress/zlib"
	"github.com/vmihailenco/msgpack/v5"
	"io"
)

func Pack(src interface{}) ([]byte, error) {
	return msgpack.Marshal(src)
}
func Unpack(src []byte) (dst interface{}, err error) {
	//if len(src) == 0 {
	//	return nil, errors.New("NotFound")
	//}
	err = UnpackTo(src, &dst)
	return
}
func UnpackTo(src []byte, dst interface{}) (err error) {
	//if len(src) == 0 {
	//	return errors.New("NotFound")
	//}
	IsPointer(dst)
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

//func PackEncrypt(src interface{}) (dst []byte, err error) {
//	dst, err = msgpack.Marshal(src)
//	if err != nil {
//		return
//	}
//	return gaes.Encrypt(dst, []byte("8byNbjJWdTlExWeWrFd0arLdSY8I1NrU"))
//}
//func UnpackDecrypt(src []byte) (dst interface{}, err error) {
//	if len(src) == 0 {
//		return nil, errors.New("NotFound")
//	}
//	err = UnpackDecryptTo(src, &dst)
//	return
//}
//func UnpackDecryptTo(src []byte, dst interface{}) (err error) {
//	if len(src) == 0 {
//		return errors.New("NotFound")
//	}
//	IsPointer(dst)
//	var data []byte
//	if data, err = gaes.Decrypt(src, []byte("8byNbjJWdTlExWeWrFd0arLdSY8I1NrU")); err != nil {
//		return
//	}
//	return msgpack.Unmarshal(data, dst)
//}

//func PackEncrypt(src interface{}) (data []byte, err error) {
//	data, err = msgpack.Marshal(src)
//	if err != nil {
//		return
//	}
//	//compressLevel := zstd.BestCompression
//	//if len(data) < 1024*100 {
//	//	compressLevel = zstd.DefaultCompression
//	//}
//	//data, err = zstd.CompressLevel(nil, data, compressLevel)
//	//if err != nil {
//	//	return
//	//}
//	//xlog.Debugf("Save: %.2f%%", 100-100*float64(len(data))/float64(len(gconv.Bytes(src))))
//	return gaes.Encrypt(data, []byte("8byNbjJWdTlExWeWrFd0arLdSY8I1NrU"))
//}
//
//func UnpackToDecrypt(src []byte, item interface{}) (err error) {
//	if len(src) == 0 {
//		return errors.New("NotFound")
//	}
//	IsPointer(item)
//	var data []byte
//	if data, err = gaes.Decrypt(src, []byte("8byNbjJWdTlExWeWrFd0arLdSY8I1NrU")); err != nil {
//		return
//	}
//	//data, err = gcompress.UnZlib(data)
//	//data, err = zstd.Decompress(nil, data)
//	//if err != nil {
//	//	return
//	//}
//	return msgpack.Unmarshal(data, item)
//}
