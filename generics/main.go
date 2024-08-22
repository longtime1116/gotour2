package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func printList(list List[int]) {
	cur := &list
	for {
		fmt.Printf("%v, ", cur.val)
		if cur.next == nil {
			break
		}
		cur = cur.next
	}
	fmt.Println()

}

func (list *List[T]) append(v T) {
	cur := list
	for {
		if cur.next == nil {
			break
		}
		cur = cur.next
	}
	l := List[T]{val: v}
	cur.next = &l
}

// ここで、TをType parametersと呼ぶ
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{1, 1, 2, 3, 5, 8, 13}
	fmt.Println(Index(si, 8))

	ss := []string{"Good Evening!", "Hogehoge", "Hello"}
	fmt.Println(Index(ss, "Hello"))

	list := List[int]{val: 1}
	list.append(20)
	list.append(10)
	printList(list)

}
