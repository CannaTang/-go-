package main

import "fmt"

func main() {
	var n int
	var str []byte
	fmt.Scanln(&n)
	fmt.Scanln(&str)
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", byte((int(str[i])-'a'+n)%26+'a'))
	}
}
