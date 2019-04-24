package main

import "fmt"

func main() {
	a := []int{3, 6, 7, 4, 2, 1, 5}

	InsertSort(7, a)

	for i := 0; i < 7; i++ {
		fmt.Printf("%d ", a[i])
	}

}

// 插入排序
// 核心  边界情况  temp值为最小时
func InsertSort(n int, a []int)  {

	var j int

	// 需要从索引为1开始往前插  如果从0开始 数组会产生越界
	for i := 1; i < n; i++ {

		temp := a[i]

		// j要大于0 这样temp最小时 不会数组越界
		for j = i; j > 0 && temp < a[j-1]; j-- {

			a[j] = a[j-1]

		}

		a[j] = temp

	}
}
