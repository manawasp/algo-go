package sort

import (
	"fmt"
	"testing"

	"github.com/manawasp/algo-go/utils"
)

func TestCountingSort(t *testing.T) {
	list := utils.RandArray(100)

	Counting(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

func BenchmarkCountingSort100(b *testing.B)    { benchmarkCountingSort(100, b) }
func BenchmarkCountingSort1000(b *testing.B)   { benchmarkCountingSort(1000, b) }
func BenchmarkCountingSort10000(b *testing.B)  { benchmarkCountingSort(10000, b) }
func BenchmarkCountingSort100000(b *testing.B) { benchmarkCountingSort(100000, b) }

/// PRIVATE

func benchmarkCountingSort(n int, b *testing.B) {
	list := utils.RandArray(n)
	for i := 0; i < b.N; i++ {
		Counting(list)
	}
}
