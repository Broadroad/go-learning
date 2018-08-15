package main

import (
	"fmt"
)

type Orange struct {
	a int
	c chan int
}

func(o *Orange)Incream(num int){
	o.a += num
}

func(o *Orange)Decream(num int){
	o.a -= num
}

func(o *Orange)String() string {
	fmt.Println("1")
	return string(o.a)
}

func main() {
	var o Orange
	o.Incream(10)
	o.Decream(5)
	fmt.Printf("%v", o)
  
  // {5 <nil>}
}
