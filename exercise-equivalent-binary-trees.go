package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	readValue(t, ch)
	close(ch)
}

func readValue(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		readValue(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		readValue(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	isSame := false
	for v1 := range ch1 {
		v2 := <-ch2
		isSame = v1 == v2
		if !isSame {
			break;
		}
	}
	return isSame
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println("True:", Same(tree.New(1), tree.New(1)))
	fmt.Println("False:", Same(tree.New(1), tree.New(2)))
}
