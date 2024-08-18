package main

import "fmt"

func pointerTest() {
	i, j := 10, 42
	p := &i

	fmt.Println(*p)
	*p += 1
	fmt.Println(*p)

	p = &j
	fmt.Println(*p)
	*p = *p * 2
	fmt.Println(*p)
	fmt.Println(i)
	fmt.Println(j)
}

func main() {
	// pointerTest();

}
