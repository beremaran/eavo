package aotree

func ContainsNode(arr []*AoNode, n *AoNode) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			return true
		}
	}

	return false
}
