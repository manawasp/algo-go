package sorting

// Selection sort - http://en.wikipedia.org/wiki/Selection_sort
func Selection(arr []int) {
	min := 0
	tmp := 0

	for i := 0; i < len(arr); i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		tmp = arr[i]
		arr[i] = arr[min]
		arr[min] = tmp
	}
}
