package main

import (
	"fmt"
	"reflect"
	_ "unsafe"
)

//Reverse reverses a slice.
var Reverse func(slice interface{}) = func(slice interface{}) {
	// type assertion
	sli := slice.(*[]myType)

	for i := 0; i < len(*sli)/2; i++ {
		(*sli)[i], (*sli)[len(*sli)-1-i] = (*sli)[len(*sli)-1-i], (*sli)[i]
	}
}

func main() {
	println("Please edit main.go,and complete the 'Reverse' function to pass the test.\nYou should use reflect package to reflect the slice type and make it apply to any type.\nTo run test,please run 'go test'\nIf you pass the test,please run 'git checkout l2' ")
}
