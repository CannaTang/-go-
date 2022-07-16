package main

import "fmt"

var flag [160010]bool

func main() {
	var a [10][10]byte //地图
	var (
		x1 int //人坐标
		y1 int
		x2 int //牛坐标
		y2 int
	)
	var ( //
		dir  = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} //转向
		sub1 = 0                                           //人
		sub2 = 0                                           //牛
	)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Scanf("%c", &a[i][j])
			if a[i][j] == 'F' {
				y1 = i
				x1 = j
			}
			if a[i][j] == 'C' {
				y2 = i
				x2 = j
			}
		}
		fmt.Scanf("\n")
	}

	time := 0
	for {
		//println(y1, x1, ' ', y2, x2, time)

		if time >= 20000 {
			fmt.Printf("0")
			break
		}

		if x1 == x2 && y1 == y2 {
			fmt.Printf("%d", time)
			break
		}

		time++

		ny1 := y1 + dir[sub1][0]
		nx1 := x1 + dir[sub1][1]
		ny2 := y2 + dir[sub2][0]
		nx2 := x2 + dir[sub2][1]

		if ny1 < 0 || ny1 > 9 || nx1 < 0 || nx1 > 9 || a[ny1][nx1] == '*' {
			nx1 = x1
			ny1 = y1
			sub1 = (sub1 + 1) % 4
		}

		if ny2 < 0 || ny2 > 9 || nx2 < 0 || nx2 > 9 || a[ny2][nx2] == '*' {
			nx2 = x2
			ny2 = y2
			sub2 = (sub2 + 1) % 4
		}

		x1 = nx1
		y1 = ny1
		x2 = nx2
		y2 = ny2
	}
}
