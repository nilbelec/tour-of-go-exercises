package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(tree *tree.Tree, channel chan int) {
	defer close(channel)
	recursiveWalk(tree, channel)
}

func recursiveWalk(tree *tree.Tree, channel chan int) {
	if tree == nil {
		return
	}
	recursiveWalk(tree.Left, channel)
	channel <- tree.Value
	recursiveWalk(tree.Right, channel)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(tree1, tree2 *tree.Tree) bool {
	channel1 := make(chan int)
	channel2 := make(chan int)
	go Walk(tree1, channel1)
	go Walk(tree2, channel2)
	for {
		value1, ok1 := <-channel1
		value2, ok2 := <-channel2

		notEqualsOrAnyIsDone := (value1 != value2) || !ok1 || !ok2
		equalsAndBothDone := (value1 == value2) && !ok1 && !ok2

		if notEqualsOrAnyIsDone {
			return equalsAndBothDone
		}
	}
}

func testWalk(n int) {
	ch := make(chan int)
	go Walk(tree.New(n), ch)
	for v := range ch {
		fmt.Println(v)
	}
}

func exercise10() {
	testWalk(1)
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(2)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
