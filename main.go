package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	SIZE = 100
)

func treeToList(node *Node) *Node {
	if node == nil {
		return node
	}

	ret := makeCDL(node.Data)
	if left := treeToList(node.Left); left != nil {
		Append(left, ret)
		ret = left
	}
	if right := treeToList(node.Right); right != nil {
		Append(ret, right)
	}

	return ret
}

func genslice(n int) []int {
	lo := -n / 2
	s := make([]int, n)
	for i := range s {
		s[i] = i + lo
	}
	return s
}

func shuffle(slice []int) []int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(slice); n > 0; n-- {
		ridx := r.Intn(n)
		slice[n-1], slice[ridx] = slice[ridx], slice[n-1]
	}
	return slice
}

func main() {
	randslice := shuffle(genslice(SIZE))
	fmt.Println("Unordered sequence of numbers:")
	fmt.Println(randslice)

	root := makeBSTFromSlice(randslice)
	slicemaker := func(s *[]int) func(n *Node) {
		return func(n *Node) {
			*s = append(*s, n.Data)
		}
	}

	fmt.Println("\nBinary Search Tree:")
	bstslice := []int{}
	root.Walk(slicemaker(&bstslice))
	fmt.Println(bstslice)

	// Convert a copy of the BST
	cdl := treeToList(deepCopy(root))

	fmt.Println("\nCircular Doubly Linked-List:")
	forwardslice := []int{}
	Traverse(cdl, false, slicemaker(&forwardslice))
	fmt.Println(forwardslice)

	fmt.Println("\nCircular Doubly Linked-List (backwards):")
	backslice := []int{}
	Traverse(cdl.Left, true, slicemaker(&backslice))
	fmt.Println(backslice)
}
