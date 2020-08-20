package AlGoLang

func insert(array []int, element int, i int) []int {
	array = append(array, 0)
	copy(array[i+1:], array[i:len(array)-1])
	array[i] = element
	return array
}

func remove(arr []int, i int) []int {
	copy(arr[i:], arr[i+1:])
	arr[len(arr)-1] = 0
	return arr[:len(arr)-1]
}

func InsertionSort(arr []int) []int {
	sortedarr := []int{arr[0]}

	remove(arr, 0)
	for arr[0] != 0 {
		c := arr[0]
		ins := false
		for i, b := range sortedarr {
			if c < b {
				ins = true
				sortedarr = insert(sortedarr, c, i)
				remove(arr, 0)
				break
			}
		}
		if ins == false {
			sortedarr = append(sortedarr, c)
			remove(arr, 0)
		}
	}
	return sortedarr
}
