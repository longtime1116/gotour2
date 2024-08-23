package main

import (
	"fmt"
	"sort"
	"time"
)

type Tree struct {
	right *Tree
	left  *Tree
	value int
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println((s))
	}
}
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // 呼び出し元がcから値を要求したタイミングでこれが実行される
			fmt.Printf("x added: %v\n", x)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func fibonacci3(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // 呼び出し元がcから値を要求したタイミングでこれが実行される
			fmt.Printf("x added: %v\n", x)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println(".....")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func insert(parent **Tree, v int) {
	if *parent == nil {
		*parent = &Tree{value: v}
		fmt.Printf("%v is inserted\n", v)
		return
	}
	if (*parent).value >= v {
		fmt.Printf("%v is smaller than parent(%v)\n", v, (*parent).value)
		insert(&(*parent).left, v)
	} else {
		fmt.Printf("%v is larger than parent(%v)\n", v, (*parent).value)
		insert(&(*parent).right, v)
	}
}

func buildTree(s []int, root **Tree) *Tree {
	for _, v := range s {
		insert(root, v)
	}
	return *root
}

func walk(t *Tree, c chan int) {
	if t != nil {
		c <- t.value
		walk(t.left, c)
		walk(t.right, c)
	}
}

func tree2slice(t *Tree) []int {
	var s []int
	c := make(chan int)
	go func() {
		walk(t, c)
		close(c)
	}()

	for v := range c {
		s = append(s, v)
	}

	sort.Ints(s)
	return s
}

func same(t1, t2 *Tree) bool {
	s1 := tree2slice(t1)
	s2 := tree2slice(t2)

	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func main() {
	if false {
		fmt.Println("a")
		go say("world")
		say("hello")

		s := []int{7, 2, 8, -9, 4, 0}
		c := make(chan int)
		go sum(s[len(s)/2:], c)
		go sum(s[:len(s)/2], c)
		x, y := <-c, <-c
		fmt.Println(x, y, x+y)

		c2 := make(chan int, 2)
		c2 <- 1
		c2 <- 2
		// c2 <- 3 // 取り出される前に入れようとしているので、dead lock エラー！
		fmt.Println(<-c2)
		fmt.Println(<-c2)
		// fmt.Println(<-c3) // 2つまでしか入らないので、3回目を呼び出すと dead lock エラー！

		c3 := make(chan int, 10)
		go fibonacci(cap(c3), c3)
		v, ok := <-c3
		fmt.Println(v, ok)
		for i := range c3 {
			fmt.Println(i)
		}
		v, ok = <-c3
		fmt.Println(v, ok)

		c4 := make(chan int)
		quit := make(chan int)
		go func() {
			for i := 0; i < 10; i++ {
				fmt.Println(<-c4)
			}
			quit <- 0
		}()
		// fibonacci2(c4, quit)
		// fibonacci3(c4, quit)

	}
	s1 := []int{8, 13, 1, 2, 1, 5, 3}
	s2 := []int{1, 1, 2, 3, 5, 8, 13}
	s3 := []int{1, 1, 2, 3, 5, 8, 12}
	s4 := []int{1, 1, 2, 3, 5}
	//s3 := []int{3, 1, 1, 2, 8, 5, 13}
	var t1, t2, t3, t4 *Tree
	buildTree(s1, &t1)
	buildTree(s2, &t2)
	buildTree(s3, &t3)
	buildTree(s4, &t4)

	if same(t1, t2) {
		fmt.Println("same")
	} else {

		fmt.Println("different")
	}
	if same(t1, t3) {
		fmt.Println("same")
	} else {

		fmt.Println("different")
	}
	if same(t1, t4) {
		fmt.Println("same")
	} else {

		fmt.Println("different")
	}

}
