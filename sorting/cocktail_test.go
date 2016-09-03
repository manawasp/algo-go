package sort

import (
	"fmt"
	"testing"

	"github.com/manawasp/algo-go/utils"
)

func TestCocktailSort(t *testing.T) {
	list := utils.RandArray(100)

	Cocktail(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

func BenchmarkCocktailSort100(b *testing.B)    { benchmarkCocktailSort(100, b) }
func BenchmarkCocktailSort1000(b *testing.B)   { benchmarkCocktailSort(1000, b) }
func BenchmarkCocktailSort10000(b *testing.B)  { benchmarkCocktailSort(10000, b) }
func BenchmarkCocktailSort100000(b *testing.B) { benchmarkCocktailSort(100000, b) }

/// PRIVATE

func benchmarkCocktailSort(n int, b *testing.B) {
	list := utils.RandArray(n)
	for i := 0; i < b.N; i++ {
		Cocktail(list)
	}
}
