// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gsha1 provides useful API for SHA1 encryption algorithms.
package xx_test

import (
	"fmt"
	"github.com/zut/x/xx"
	"testing"
)

func Test_Sha3(t *testing.T) {
	r := xx.Sha3("123")
	fmt.Println(r)
	if r != "a03ab19b866fc585b5cb1812a2f63ca861e7e7643ee5d43fd7106b623725fd67" {
		t.Errorf(r)
	}
}
func Test_Sha3File(t *testing.T) {
	r := xx.Sha3File("sha3.go")
	fmt.Println("File:" + r)
	if r != "09127657119c70c0c43480996f702f878389c402ce59bebebdcd9efbe608cb0e" {
		t.Errorf(r)
	}
}
