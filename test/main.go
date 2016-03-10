package main

import (
	"fmt"
	"reflect"
)

type Testfn func()

type test struct {
	Fn Testfn
	ty reflect.Value
}

func (t *test) a() {
	fmt.Println("none")
}

func (t *test) callA() {
	met := t.ty.Addr().MethodByName("a")
	p := make([]reflect.Value, 0)
	met.Call(p)
}

func main() {
	obj := &test{
		Fn: func() {
			fmt.Println("testing")
		},
	}
	obj.ty = reflect.ValueOf(obj).Elem()

	// p[0] = reflect.ValueOf("test param")

	obj.callA()
}
