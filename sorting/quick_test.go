package sort

import (
	"fmt"
	"testing"

	"github.com/manawasp/algo-go/utils"
)

func TestQuickSort(t *testing.T) {
	list := utils.RandArray(100)

	list = Quick(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

func BenchmarkQuickSort100(b *testing.B)    { benchmarkQuickSort(100, b) }
func BenchmarkQuickSort1000(b *testing.B)   { benchmarkQuickSort(1000, b) }
func BenchmarkQuickSort10000(b *testing.B)  { benchmarkQuickSort(10000, b) }
func BenchmarkQuickSort100000(b *testing.B) { benchmarkQuickSort(100000, b) }

/// PRIVATE

func benchmarkQuickSort(n int, b *testing.B) {
	list := utils.RandArray(n)
	for i := 0; i < b.N; i++ {
		Quick(list)
	}
}
