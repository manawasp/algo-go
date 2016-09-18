package searching

// Negamax decision algorithm - https://en.wikipedia.org/wiki/Alpha%E2%80%93beta_pruning
// func Negamax(node *Node, depth int, alpha, beta, color int) int {
// 	if len(node.Childrens) == 0 || depth == 0 {
// 		return node.Score * color
// 	}
//
// 	// Init the bestScore / score with the first node
// 	score := -Negamax(node.Childrens[0], depth-1, -color)
// 	node.Score = score
// 	for i := 1; i < mm.Childrens.length; i++ {
// 		score = -Negamax(node.Childrens[i], depth-1, -color)
// 		if score > node.Score {
// 			node.Score = score
// 		}
// 	}
// 	return node.Score
// }
