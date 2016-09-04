package searching

// Minimax decision algorithm - https://en.wikipedia.org/wiki/Minimax
func Minimax(node *Node, depth int) (bestMove *Node) {
	score := 0
	if len(node.Childrens) > 0 {
		node.Score = min(node.Childrens[0], depth)
		bestMove = node.Childrens[0]
	}

	for i := 0; i < len(node.Childrens); i++ {
		score = min(node.Childrens[i], depth)
		if score > node.Score {
			bestMove = node.Childrens[i]
			node.Score = score
		}
	}
	return bestMove
}

// PRIVATE

// Return the lowest score received
func min(node *Node, depth int) int {
	if len(node.Childrens) == 0 || depth == 0 {
		return node.Score
	}

	// Init the lowestscore / score with the first node
	score := max(node.Childrens[0], depth-1)
	node.Score = score

	// Scope of 1..n solutions to get the lowest score
	for i := 1; i < len(node.Childrens); i++ {
		score := max(node.Childrens[i], depth-1)
		if score < node.Score {
			node.Score = score
		}
	}
	return node.Score
}

// Return the highest score received
func max(node *Node, depth int) int {
	if len(node.Childrens) == 0 || depth == 0 {
		return node.Score
	}

	// Init the highestscore / score with the first node
	score := max(node.Childrens[0], depth-1)
	node.Score = score

	// Scope of 1..n solutions to get the highest score
	for i := 1; i < len(node.Childrens); i++ {
		score := max(node.Childrens[i], depth-1)
		if score > node.Score {
			node.Score = score
		}
	}
	return node.Score
}
