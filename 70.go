package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

//type Tree struct {
//    Left  *Tree
//    Value int
//    Right *Tree
//}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk := func(t *tree.Tree, ch chan int) {}
	walk = func(t *tree.Tree, ch chan int) {

		if t.Left != nil {
			walk(t.Left, ch)
		}

		ch <- t.Value

		if t.Right != nil {
			walk(t.Right, ch)
		}

	}
	walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v1 := range ch1 {
		if v1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
