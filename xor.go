package main

import (
	"fmt"
)

func main() {
	maxInt := int(^uint(0) >> 1) // ^x    bitwise complement    is m ^ x  with m = "all bits set to 1" for unsigned x
                                         // and  m = -1 for signed x
		
	fmt.Println(maxInt)
	// 2147483647
}
