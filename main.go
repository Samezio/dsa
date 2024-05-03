package main

import (
	binary_tree "com/dsa/binary_tree"
	"fmt"
	"io/ioutil"
)

func main() {
	root := binary_tree.Node[int]{
		Val: 0,
		Left: &binary_tree.Node[int]{
			Val: 1,
			Left: &binary_tree.Node[int]{
				Val:  2,
				Left: nil,
				Right: &binary_tree.Node[int]{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &binary_tree.Node[int]{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &binary_tree.Node[int]{
			Val: 5,
			Left: &binary_tree.Node[int]{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &binary_tree.Node[int]{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	if err := ioutil.WriteFile("tree.html", []byte(binary_tree.BuildTreeHtml(root)), 7666); err != nil {
		panic(err)
	}
	inorder := binary_tree.Morris_inorder(root)
	fmt.Println(inorder)
}
