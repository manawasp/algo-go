package sort

// Bubble sort - http://en.wikipedia.org/wiki/Bubble_sort
func Bubble(arr []int) {
	tmp := 0

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
}
