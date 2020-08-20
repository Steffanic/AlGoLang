package AlGoLang

/*Problem: Given a set of 99 unique numbers from 1-100, find the missing number.*/

func allIntegersPresent(arr []int) bool {
	return arr[len(arr)-1]-arr[0]+1 == len(arr)
}

func splitAndTest(arr []int) int {
	arr1 := arr[:int(len(arr)/2)]
	arr2 := arr[int(len(arr)/2):]
	if arr1[len(arr1)-1] == arr2[0]-2 {
		return arr2[0] - 1
	}
	if allIntegersPresent(arr1) {
		return splitAndTest(arr2)
	} else {
		return splitAndTest(arr1)
	}
}

func ReturnMissingNumber(arr []int) int {
	sortedarr := InsertionSort(arr)
	return splitAndTest(sortedarr)
}
