package sort

import "fmt"

//插入排序
func InsertTionSort(a []int) {
	for j := 1; j < len(a); j++ {
		key := a[j]
		for i := j - 1; i >= 0; i-- {
			if a[i] > key {
				a[i+1] = a[i]
				a[i] = key
				fmt.Println(a, i, j, key)
			}
		}

	}
	fmt.Println(a)
}
