package main

import (
	"fmt"
	"reflect"
)

//Reverse reverses a slice.
var Reverse func(slice interface{}) = func(slice interface{}) {
	fmt.Println(reflect.ValueOf(slice).Kind())
	switch reflect.ValueOf(slice).Kind() {
	case reflect.Ptr:
		v := reflect.ValueOf(slice)
		fmt.Println(v.Elem().Len())
		n := v.Elem().Len()

		for i := 0; i < n/2; i++ {
			t := v.Elem().Index(i).Interface()
			// need Set
			v.Elem().Index(i).Set(reflect.ValueOf(v.Elem().Index(n - 1 - i).Interface()))
			v.Elem().Index(n - 1 - i).Set(reflect.ValueOf(t))
		}
	}
}

func main() {
	println("Please edit main.go,and complete the 'Reverse' function to pass the test.\nYou should use reflect package to reflect the slice type and make it apply to any type.\nTo run test,please run 'go test'\nIf you pass the test,please run 'git checkout l2' ")
}
