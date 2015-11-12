package main

import (
	"fmt"
	"reflect"
)

func MakeMap(fpt interface{}) {
	fnV := reflect.ValueOf(fpt).Elem()
	fmt.Println(fnV)
	fnI := reflect.MakeFunc(fnV.Type(), implMap)
	fmt.Println(fnI)
	fnV.Set(fnI)
}

//TODO:completes implMap function.
var implMap func(in []reflect.Value) []reflect.Value = func(in []reflect.Value) []reflect.Value {
	fmt.Println(in[0].Type(), "haha", in[1])

	for _, ele := range in[1] {
		fmt.Println(ele)
	}
	return nil
}

func main() {
	println("It is said that Go has no generics.\nHowever we have many other ways to implement generics like using a library if it seems less smoothly, one is reflect.MakeFunc.\nUnderscore is a very useful js library,and now let's implement part of it-map,it will help you to understand how reflect works.\nPlease finish the 'implMap' function and pass the test.")
}
