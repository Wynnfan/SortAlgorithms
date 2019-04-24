package main

import "fmt"

func main() {
	a := []int{1, 5, 3, 7, 2}

	SelectSort(5, a)

	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", a[i])
	}
}

// 选择排序过程
// 选择排序就是每次选择数组中待排序部分的最大值或最小值  将最大值或最小值放到数组最后
// 下面排序是从数组最后开始选择  依次和数组前面的值进行比较
func SelectSort(n int, a []int) {

	for i := n - 1; i > 0; i-- {

		temp := i

		for j := 0; j < i; j++ {

			if a[temp] < a[j] {
				temp = j
			}

		}

		if temp != i {
			a[temp], a[i] = a[i], a[temp]
		}

	}

}

// func SelectSort(n int, a []int) {
//
// 	c := 0
// 	for i := 0; i < n; i++ {
//
// 		temp := i
//
// 		for j := 0; j < n; j++ {
// 			c++
// 			if a[temp] < a[j] {
// 				a[temp], a[j] = a[j], a[temp]
// 			}
// 		}
//
// 	}
// 	fmt.Println(c)
// }
