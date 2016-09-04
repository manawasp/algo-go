package searching

import (
	"fmt"
	"testing"
)

func TestNegamaxSearch(t *testing.T) {
	node := MakeTree(3, 3, 20)

	Negamax(node, 4, 1)
	checkNegamax(t, node, -1)
}

// Return the highest score received
func checkNegamax(t *testing.T, node *Node, color int) {
	if len(node.Childrens) == 0 {
		return
	}
	maxScore := node.Childrens[0].Score

	for i := 0; i < len(node.Childrens); i++ {
		if node.Childrens[i].Score*-color > maxScore {
			maxScore = node.Childrens[i].Score * -color
		}
		if node.Childrens[i].Score*-color > node.Score {
			fmt.Printf("Max error : Node score %d / Children score %d\n",
				node.Score,
				node.Childrens[i].Score)
			t.Error()
		}
		checkNegamax(t, node.Childrens[i], -color)
	}
}

func BenchmarkNegamaxSearch3x3x20(b *testing.B)   { benchmarkNegamaxSearch(3, 3, 20, b) }
func BenchmarkNegamaxSearch3x10x20(b *testing.B)  { benchmarkNegamaxSearch(3, 10, 20, b) }
func BenchmarkNegamaxSearch30x200(b *testing.B)   { benchmarkNegamaxSearch(6, 3, 200, b) }
func BenchmarkNegamaxSearch100x2000(b *testing.B) { benchmarkNegamaxSearch(6, 10, 2000, b) }

/// PRIVATE

func benchmarkNegamaxSearch(depth, childrensLength, valueRange int, b *testing.B) {
	node := MakeTree(depth, childrensLength, valueRange)
	for i := 0; i < b.N; i++ {
		Negamax(node, -1, -1)
	}
}
