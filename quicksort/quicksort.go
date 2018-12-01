package main

import (
	"fmt"
)

 var a = make([]int, 11)
// var a [101] int

func quickSort(left, right int) {
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

	quickSort(left, i-1)
	quickSort(i+1, right)
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	quickSort(0, n-1)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
	}
}
