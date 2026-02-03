package main

import (
	"fmt"
)

import "golang.org/x/tour/tree"

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		var x int = <-ch1
		var y int = <-ch2
		if x != y {
			return false
		}
	}
	return true
}

func PrintTree(t *tree.Tree) {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		close(ch)
	}()
	Walk(t, ch)
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	t3 := tree.New(2)
	fmt.Println("t1", t1)
	fmt.Println("t2", t2)
	fmt.Println("t3", t3)
	fmt.Println("t1 == t2 ?", Same(t1, t2))
	fmt.Println("t1 == t3 ?", Same(t1, t3))

	// PrintTree(t1)
}
