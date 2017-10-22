package main

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func makeBST(root int) *Node {
	return &Node{Data: root, Left: nil, Right: nil}
}

func makeBSTFromSlice(nums []int) *Node {
	root := makeBST(nums[0])
	for _, val := range nums[1:] {
		root.Append(val)
	}
	return root
}

func deepCopy(root *Node) *Node {
	var copy *Node = nil
	root.Walk(func(n *Node) {
		if copy == nil {
			copy = makeBST(n.Data)
		} else {
			copy.Append(n.Data)
		}
	})
	return copy
}

func (bst *Node) Append(data int) {
	if data <= bst.Data {
		if bst.Left == nil {
			bst.Left = &Node{Data: data, Left: nil, Right: nil}
		} else {
			bst.Left.Append(data)
		}
	}

	if data > bst.Data {
		if bst.Right == nil {
			bst.Right = &Node{Data: data, Left: nil, Right: nil}
		} else {
			bst.Right.Append(data)
		}
	}
}

func (bst *Node) Walk(f func(*Node)) {
	if bst == nil {
		return
	}

	bst.Left.Walk(f)
	f(bst)
	bst.Right.Walk(f)
}
