package searching

import (
	"fmt"
	"testing"
)

func TestMinimaxSearch(t *testing.T) {
	node := MakeTree(3, 3, 20)

	Minimax(node, 4)
	checkMin(t, node)
}

func checkMin(t *testing.T, node *Node) {
	if len(node.Childrens) == 0 {
		return
	}

	for i := 0; i < len(node.Childrens); i++ {
	}

	minScore := node.Childrens[0].Score

	for i := 0; i < len(node.Childrens); i++ {
		if node.Childrens[i].Score > minScore {
			minScore = node.Childrens[i].Score
		}
		if node.Childrens[i].Score > node.Score {
			fmt.Printf("Max error : Node score %d / Children score %d\n",
				node.Score,
				node.Childrens[i].Score)
			t.Error()
		} else {
			checkMax(t, node.Childrens[i])
		}
	}

	if node.Score != minScore {
		t.Error()
	}
}

func checkMax(t *testing.T, node *Node) {
	if len(node.Childrens) == 0 {
		return
	}
	maxScore := node.Childrens[0].Score

	for i := 0; i < len(node.Childrens); i++ {
		if node.Childrens[i].Score < maxScore {
			maxScore = node.Childrens[i].Score
		}
		if node.Childrens[i].Score < node.Score {
			fmt.Printf("Max error : Node score %d / Children score %d\n",
				node.Score,
				node.Childrens[i].Score)
			t.Error()
		}
		checkMin(t, node.Childrens[i])
	}

	if node.Score != maxScore {
		t.Error()
	}
}

func BenchmarkMinimaxSearch3x3x20(b *testing.B)   { benchmarkMinimaxSearch(3, 3, 20, b) }
func BenchmarkMinimaxSearch3x10x20(b *testing.B)  { benchmarkMinimaxSearch(3, 10, 20, b) }
func BenchmarkMinimaxSearch30x200(b *testing.B)   { benchmarkMinimaxSearch(6, 3, 200, b) }
func BenchmarkMinimaxSearch100x2000(b *testing.B) { benchmarkMinimaxSearch(6, 10, 2000, b) }

/// PRIVATE

func benchmarkMinimaxSearch(depth, childrensLength, valueRange int, b *testing.B) {
	node := MakeTree(depth, childrensLength, valueRange)
	for i := 0; i < b.N; i++ {
		Minimax(node, -1)
	}
}
