package main

import "fmt"

func main() {
	x := [2]int{0, 0} //x[0]是常数系数,x[1]是变量系数

	var c byte
	var word byte
	num := 0
	turn := 1 //在方程右边
	flag := 0 //0为常数, 1为变量
	mark := 1
	for {
		fmt.Scanf("%c", &c)
		if c >= '0' && c <= '9' {
			num = num*10 + int(c) - 48
		} else if c >= 'a' && c <= 'z' {
			flag = 1
			word = c
		} else if c == '+' {
			if num == 0 && flag == 1 {
				num = 1
			}
			x[flag] += mark * num * turn
			num = 0
			flag = 0
			mark = 1
		} else if c == '-' {
			if num == 0 && flag == 1 {
				num = 1
			}
			x[flag] += mark * num * turn
			num = 0
			flag = 0
			mark = -1
		} else if c == '=' {
			if num == 0 && flag == 1 {
				num = 1
			}
			x[flag] += mark * num * turn
			turn = -1
			num = 0
			flag = 0
			mark = 1
		} else if c == '\n' {
			x[flag] += mark * num * turn
			break
		}
	}
	var ans float32
	ans = -float32(x[0]) / float32(x[1])
	fmt.Printf("%c=%.3f\n", word, ans)
}
