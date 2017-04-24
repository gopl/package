package main

import "fmt"

func main() {
	type T struct {
		a int
		b float64
		c string
	}
	// exactly what Print and Println would produce
	t := T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
	// %+v	annotates the fields of the structure with their names
	fmt.Printf("%+v\n", t)
	// %#v	prints the value in full Go syntax
	fmt.Printf("%#v\n", t)

	// %q	quoted string format
	fmt.Printf("%q\n", "hello")
	// single-quoted rune constant.
	fmt.Printf("%q\n", []int{1234, 1235})
	// %#q	backquoted string format
	fmt.Printf("%#q\n", "hello")

	// %x	hexadecimal string
	fmt.Printf("%x\n", "abc")
	// % x	puts spaces between the bytes
	fmt.Printf("% x\n", "abc")
}
