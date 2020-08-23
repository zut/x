package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	foo int
}

func main() {
	t := Test{5}
	fmt.Println(reflect.TypeOf(t) == reflect.TypeOf(Test{}))
	fmt.Println(reflect.TypeOf(t) == reflect.TypeOf((*Test)(nil)).Elem())
}
