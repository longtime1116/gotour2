package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	// 書き込み中にlockしないで読み込もうとすると、ランタイムエラーが起きる
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	key := "hoge"
	func() {
		for i := 0; i < 10000; i++ {
			go c.Inc(key)
		}
	}()
	fmt.Println(c.Value(key))
	time.Sleep(time.Second)
	fmt.Println(c.Value(key))

}
