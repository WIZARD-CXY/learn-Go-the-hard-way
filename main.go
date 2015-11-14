package main

import (
	_ "fmt"
	"reflect"
)

func MakeMap(fpt interface{}) {
	fnV := reflect.ValueOf(fpt).Elem()
	//fmt.Println(fnV.Type().NumOut())
	fnI := reflect.MakeFunc(fnV.Type(), implMap)
	//fmt.Println("haha1", fnI.Type().NumOut())
	fnV.Set(fnI)
}

//TODO:completes implMap function.
// note the declarition
// func MakeFunc(typ Type, fn func(args []Value) (results []Value)) Value
var implMap func(in []reflect.Value) []reflect.Value = func(in []reflect.Value) (res []reflect.Value) {
	//in[0] is func, in[1] is slice or map
	//fmt.Println(in[0].Type(), in[1].Type())

	if in[1].Kind() == reflect.Slice {
		for i := 0; i < in[1].Len(); i++ {
			inslice := make([]reflect.Value, 1)
			inslice[0] = in[1].Index(i)
			in[1].Index(i).Set(in[0].Call(inslice)[0])
		}

	} else {
		//in[1] is map
		for _, k := range in[1].MapKeys() {
			inslice := make([]reflect.Value, 1)
			inslice[0] = in[1].MapIndex(k)
			in[1].SetMapIndex(k, in[0].Call(inslice)[0])
		}

	}
	// The result Value slice returned by fn (res here)must have the number and type of results given by typ(fnV.Type() here).
	// In this case f is func(func(e int) int, []int) []int, so len(res) must be 1
	// if f2 is func(func(e int) int, []int) ([]int, []int) then len(res) should be 2
	res = make([]reflect.Value, 1)
	res[0] = in[1]

	return
}

func main() {
	println("It is said that Go has no generics.\nHowever we have many other ways to implement generics like using a library if it seems less smoothly, one is reflect.MakeFunc.\nUnderscore is a very useful js library,and now let's implement part of it-map,it will help you to understand how reflect works.\nPlease finish the 'implMap' function and pass the test.")
}
