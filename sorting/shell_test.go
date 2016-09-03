package sort

import (
	"fmt"
	"testing"

	"github.com/manawasp/algo-go/utils"
)

func TestShellSort(t *testing.T) {
	list := utils.RandArray(100)

	Shell(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

func BenchmarkShellSort100(b *testing.B)    { benchmarkShellSort(100, b) }
func BenchmarkShellSort1000(b *testing.B)   { benchmarkShellSort(1000, b) }
func BenchmarkShellSort10000(b *testing.B)  { benchmarkShellSort(10000, b) }
func BenchmarkShellSort100000(b *testing.B) { benchmarkShellSort(100000, b) }

/// PRIVATE

func benchmarkShellSort(n int, b *testing.B) {
	list := utils.RandArray(n)
	for i := 0; i < b.N; i++ {
		Shell(list)
	}
}
