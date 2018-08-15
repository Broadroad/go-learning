/*
This package tell us that point and struct will call different function
if they have the same function sign except pass by value or pointer
*/
package main

import (
	"fmt"
)

type Orange struct {
	a int
}

func (o *Orange) Incream(num int) {
	o.a += num
}

func (o *Orange) Decream(num int) {
	o.a -= num
}

func (o *Orange) String() string {
	fmt.Println("1")
	return fmt.Sprintf("int(%d)", o.a)
}

func main() {
	// in go, does not exist uninitialized variable. This is nil in struct for each struct element
	var o Orange
	o.Incream(10)
	o.Decream(5)
	fmt.Println(o)
	// here call func(o Orange)String() string
	fmt.Printf("%v", o)

	// {5 <nil>}
}
