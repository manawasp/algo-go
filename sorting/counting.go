package sort

func countArr(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	k := arr[0]

	for _, v := range arr {
		if v > k {
			k = v
		}
	}

	return k
}

// Counting - https://en.wikipedia.org/wiki/Countingi_sort
func Counting(arr []int) {
	k := countArr(arr)
	count := make([]int, k+1)

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	for i, j := 0, 0; i < k; i++ {
		for {
			if count[i] > 0 {
				arr[j] = i
				j++
				count[i]--
				continue
			}
			break
		}
	}
}
