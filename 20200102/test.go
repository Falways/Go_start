package main

import "fmt"

func main() {
	var s string = "str"
	fmt.Println(s)
	var a *string = &s
	var b string = *a
	fmt.Println(&s==&b)
}
