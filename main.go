package main

import (
	"fmt"
	"reflect"
)

func MakeMap(fpt interface{}) {
	fnV := reflect.ValueOf(fpt).Elem()
	fmt.Println(fnV.Type().NumOut())
	fnI := reflect.MakeFunc(fnV.Type(), implMap)
	fmt.Println("haha1", fnI.Type().NumOut())
	fnV.Set(fnI)
}

//TODO:completes implMap function.
var implMap func(in []reflect.Value) []reflect.Value = func(in []reflect.Value) (res []reflect.Value) {
	fmt.Println(in)
	fmt.Println(in[0].Type(), in[1].Len())
	//in[0] is func, in[1] is slice or map

	for i := 0; i < in[1].Len(); i++ {
		inslice := make([]reflect.Value, 1)
		inslice[0] = in[1].Index(i)
		res = append(res, in[0].Call(inslice)[0])
	}

	fmt.Println("len", len(res))
	return
}

func main() {
	println("It is said that Go has no generics.\nHowever we have many other ways to implement generics like using a library if it seems less smoothly, one is reflect.MakeFunc.\nUnderscore is a very useful js library,and now let's implement part of it-map,it will help you to understand how reflect works.\nPlease finish the 'implMap' function and pass the test.")
}
