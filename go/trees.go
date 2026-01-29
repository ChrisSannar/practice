package main

import (
	"fmt"
)

import "golang.org/x/tour/tree"

func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// func Same(t1, t2 *tree.Tree)

func main() {
	n := 1
	t := tree.New(n)
	fmt.Println(t)
	fmt.Println()

	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		close(ch)
	}()
	Walk(t, ch)
}
