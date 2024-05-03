package binary_tree

func Morris_inorder[T any](root Node[T]) []T {
	inorder := make([]T, 0)
	var current *Node[T] = &root
outerloop:
	for current != nil {
		if current.Left == nil {
			inorder = append(inorder, current.Val)
			current = current.Right
		} else {
			right_most := current.Left
			for right_most.Right != nil {
				if right_most.Right == current {
					//Loop connection: Left already visited
					right_most.Right = nil
					inorder = append(inorder, current.Val)
					current = current.Right
					continue outerloop
				}
				right_most = right_most.Right
			}
			right_most.Right = current
			current = current.Left
		}
	}
	return inorder
}
