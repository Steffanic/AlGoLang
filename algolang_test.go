package AlGoLang

import (
	"fmt"
	"sort"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{8, 7, 6, 4, 3, 2, 1}
	sortarr := InsertionSort(arr)
	fmt.Println(sortarr)
	if !sort.IntsAreSorted(sortarr) {
		t.Error("The slice was not sorted!")
	}
}

func TestReturnMissingNumber(t *testing.T) {
	arr := make([]int, 99)
	for k := 1; k < 99; k++ {
		offset := 0
		for i := 0; i < 100; i++ {
			if i == k {
				offset = 1
				continue
			}
			arr[i-offset] = i
		}
		fmt.Println(arr)
		m := ReturnMissingNumber(arr)
		fmt.Println(m)
		if m != k {
			t.Error("Failure at k=", k)
		}
	}
}
