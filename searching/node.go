package searching

import "math/rand"

// Node ...
type Node struct {
	Score     int
	Data      interface{}
	Parent    *Node
	Childrens []*Node
}

// MakeTree ...
func MakeTree(depth, childrensLength, rangeScore int) *Node {
	node := &Node{}
	makeChildrens(node, depth, childrensLength, rangeScore)
	return node
}

func makeChildrens(node *Node, depth, childrensLength, rangeScore int) {
	if depth == 0 {
		node.Score = rand.Intn(rangeScore) + 1
		return
	}

	for i := 0; i < childrensLength; i++ {
		node.Childrens = append(node.Childrens, &Node{})
		node.Childrens[i].Parent = node
		makeChildrens(node.Childrens[i], depth-1, childrensLength, rangeScore)
	}
}
