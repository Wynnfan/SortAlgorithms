package main

import (
	"fmt"
)

// var a = make([]int, 101)

func main() {
	a := []int{1, 5, 3, 7, 2}

	BubbleSort(5, a)

	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", a[i])
	}
}

// 冒泡过程
func BubbleSort(n int, a []int) {

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] < a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

}
