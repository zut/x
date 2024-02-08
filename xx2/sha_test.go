package xx2

import (
	"github.com/gogf/gf/test/gtest"
	"testing"
)

func TestSha1(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(Sha1("123"), "40bd001563085fc35165329ea1ff5c5ecbdbbeef")
		t.Assert(Sha1("abc"), "a9993e364706816aba3e25717850c26c9cd0d89d")
	})
}

func TestSha3(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(Sha3("123"), "a03ab19b866fc585b5cb1812a2f63ca861e7e7643ee5d43fd7106b623725fd67")
		t.Assert(Sha3("abc"), "3a985da74fe225b2045c172d6bd390bd855f086e3e9d525b46bfe24511431532")
	})
}

func TestSha3Bytes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(Sha3Bytes([]byte("123")), "a03ab19b866fc585b5cb1812a2f63ca861e7e7643ee5d43fd7106b623725fd67")
		t.Assert(Sha3Bytes([]byte("abc")), "3a985da74fe225b2045c172d6bd390bd855f086e3e9d525b46bfe24511431532")
	})
}

func TestSha3File(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		v, err := Sha3File("z.test.file")
		t.Assert(err, nil)
		t.Assert(v, "a03ab19b866fc585b5cb1812a2f63ca861e7e7643ee5d43fd7106b623725fd67")
	})
}
