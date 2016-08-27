package sorting

// Merge sort - http://en.wikipedia.org/wiki/Merge_sort
func Merge(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	i := len(arr) / 2

	fromLeft := Merge(arr[:i])
	fromRight := Merge(arr[i:])

	return mergeArr(fromLeft, fromRight)
}

/// PRIVATE

func mergeArr(partLeft, partRight []int) []int {
	result := make([]int, 0, len(partLeft)+len(partRight))

	for len(partLeft) > 0 || len(partRight) > 0 {
		if len(partLeft) == 0 {
			return append(result, partRight...)
		}
		if len(partRight) == 0 {
			return append(result, partLeft...)
		}
		if partLeft[0] <= partRight[0] {
			result = append(result, partLeft[0])
			partLeft = partLeft[1:]
		} else {
			result = append(result, partRight[0])
			partRight = partRight[1:]
		}
	}

	return result
}
