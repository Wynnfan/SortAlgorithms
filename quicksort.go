package main

import (
	"fmt"
)

 var a = make([]int, 11)
// var a [101] int


func quicksort(left, right int) {
	var i, j, temp int
	if left > right {
		return
	}

	temp = a[left]
	i, j = left, right
	for i != j {
		for a[j] >= temp && i < j {
			j--
		}
		for a[i] <= temp && i < j {
			i++
		}

		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}

	a[left], a[i] = a[i], a[left]

	quicksort(left, i-1)
	quicksort(i+1, right)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(n)
	// for i, _ := range a {
	// 	fmt.Scanf()
	// }
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	quicksort(0, n-1)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
	}
}
