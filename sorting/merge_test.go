package sort

import (
	"fmt"
	"testing"

	"github.com/manawasp/algo-go/utils"
)

func TestMergeSort(t *testing.T) {
	list := utils.RandArray(100)

	list = Merge(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

func BenchmarkMergeSort100(b *testing.B)    { benchmarkMergeSort(100, b) }
func BenchmarkMergeSort1000(b *testing.B)   { benchmarkMergeSort(1000, b) }
func BenchmarkMergeSort10000(b *testing.B)  { benchmarkMergeSort(10000, b) }
func BenchmarkMergeSort100000(b *testing.B) { benchmarkMergeSort(100000, b) }

/// PRIVATE

func benchmarkMergeSort(n int, b *testing.B) {
	list := utils.RandArray(n)
	for i := 0; i < b.N; i++ {
		Merge(list)
	}
}
