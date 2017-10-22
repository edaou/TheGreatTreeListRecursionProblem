package main

func makeCDL(data int) *Node {
	node := &Node{Data: data, Left: nil, Right: nil}
	node.Left = node
	node.Right = node
	return node
}

func Append(node1 *Node, node2 *Node) {
	node1.Left.Right = node2
	node2.Left.Right = node1
	node1.Left, node2.Left = node2.Left, node1.Left
}

func Traverse(root *Node, backwards bool, f func(*Node)) {
	iscycle := makeCycleDetector(root.Data)
	next := makeIterator(backwards)

	for temp := root; !iscycle(temp); temp = next(temp) {
		f(temp)
	}
}

func makeIterator(backwards bool) func(node *Node) *Node {
	if backwards {
		return func(node *Node) *Node { return node.Left }
	} else {
		return func(node *Node) *Node { return node.Right }
	}
}

func makeCycleDetector(start int) func(*Node) bool {
	count := 0
	return func(node *Node) bool {
		if node.Data == start {
			count++
		}
		return count > 1
	}
}
