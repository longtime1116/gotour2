package main

import (
	"fmt"
	"time"
)

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

func main() {
	if false {
		fmt.Println("a")
		go say("world")
		say("hello")
	}
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
