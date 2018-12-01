package main

import (
	"fmt"
)

var a = make([]int, 101)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	BubbleSort(n, a)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
	}
}

// 冒泡过程
func BubbleSort(n int, a []int) {
	var i int
	for i = 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if a[j] < a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
