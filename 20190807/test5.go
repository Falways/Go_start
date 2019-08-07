package main

import "fmt"

func main() {
	var a int = 10
	var b string = "xuhang"

	fmt.Printf("变量的地址: %x, %x\n", &a, &b  )
}