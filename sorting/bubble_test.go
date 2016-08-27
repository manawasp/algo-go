package sort

import (
	"fmt"
	"testing"

	"github.com/manawasp/algo-go/utils"
)

func TestBubbleSort(t *testing.T) {
	list := utils.RandArray(100)

	Bubble(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

func BenchmarkBubbleSort100(b *testing.B)    { benchmarkBubbleSort(100, b) }
func BenchmarkBubbleSort1000(b *testing.B)   { benchmarkBubbleSort(1000, b) }
func BenchmarkBubbleSort10000(b *testing.B)  { benchmarkBubbleSort(10000, b) }
func BenchmarkBubbleSort100000(b *testing.B) { benchmarkBubbleSort(100000, b) }

/// PRIVATE

func benchmarkBubbleSort(n int, b *testing.B) {
	list := utils.RandArray(n)
	for i := 0; i < b.N; i++ {
		Bubble(list)
	}
}
