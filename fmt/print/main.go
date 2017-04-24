package main

import "fmt"

func main() {

	// 两边都不是字符串时，才插入空格
	fmt.Print(1, 2, "\n")
	fmt.Print(1, "2", "\n")

	// 都插入空格
	fmt.Println(3, 4)
	fmt.Println(3, "4")
}
