package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

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

func fillCoord(v *Vertex) {
	v.X = 10
	v.Y = 200
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printTicTacToe() {
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}
	for i := 0; i < len(board); i++ {
		fmt.Println(board[i])
	}
	board[1][1] = "O"
	board[0][0] = "X"
	board[2][0] = "O"
	board[0][2] = "X"
	board[0][1] = "O"
	board[2][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Println(board[i])
	}
}

func addr() func(int) int {
	sum := 0

	// addr 関数内で作成されたこの関数は、sumという変数を参照する。
	// sum は addr 関数内で定義されているが、addrの実行が完了した後もsum変数は関数内に閉じ込められた状態で保持される
	// この無名関数をクロージャーと呼ぶ。クロージャーの定義は、外部スコープの変数を参照する関数。クロージャーは外部スコープの変数をキャプチャする。
	// ニュアンス的には、外部スコープの変数を、関数の外にあるにも関わらず、自らのうちに close(閉じ込める) してしまうから、closure と呼ぶ。
	// また、関数が実行される場所に関係なく、この関数と閉じ込められたこのへ数が一緒にパッケージされ、変数が存在するスコープが閉じられて、関数内で使用できる状態が維持される。
	// その感覚から、closure(閉じ込め)という名前がついている
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		current := first
		first, second = second, first+second
		return current
	}
}

// A struct is a collection of fields.
type Vertex struct {
	X int
	Y int
}

type Location struct {
	Lat, Long float64
}

func main() {
	if false {
		pointerTest()
		v1 := Vertex{1, 2}
		var v2 Vertex
		v2.X = 4
		v2.Y = 8
		fmt.Println(v1, v2, v2.X)
		p1 := &v1
		fmt.Printf("%T, %v\n", p1, p1)
		p1.X = 1e9
		// どっちの書き方もできる
		fmt.Println(p1.X, (*p1).X)

		// 指定しなかったら暗黙的に0が入る。
		var (
			v3 = Vertex{X: 2}
			v4 = Vertex{}
		)
		fmt.Println(v3, v4)

		// 参照渡しして変更できることを確認
		fillCoord(&v4)
		fmt.Println(v4)

		// Arrays
		// An array's length is part of its type, so arrays cannot be resized.
		var a [2]string
		a[0] = "hello"
		a[1] = "world"
		fmt.Println(a)
		a[0] = "Hello"
		fmt.Println(a)

		// An array has a fixed size. A slice, on the other hand, is a dynamically-sized,
		// flexible view into the elements of an array. In practice, slices are much more common than arrays.
		primes := [6]int{2, 3, 5, 7, 11, 13}
		var slice1 []int = primes[1:4]
		fmt.Println(slice1) // 3,5,7
		fmt.Printf("%T, %T\n", primes, slice1)

		// Slices are like references to arrays. A slice does not store any data, it just describes a section of an underlying array.
		names := [4]string{"James", "CP3", "KD", "Curry"}
		nameSlice1 := names[0:2]
		nameSlice2 := names[1:3]
		fmt.Println(nameSlice1, nameSlice2)
		// CP3 is replaced by Rui!
		nameSlice1[1] = "Rui"
		fmt.Println(nameSlice1, nameSlice2)
		// 同じ参照先を示していることが確認できる。
		// 0x1400011e040, 0x1400011e050, 0x1400011e060, 0x1400011e070,
		// 0x1400011e040, 0x1400011e050,
		fmt.Printf("%p, %p, %p, %p, \n", &names[0], &names[1], &names[2], &names[3])
		fmt.Printf("%p, %p, \n", &nameSlice1[0], &nameSlice1[1])
		// 参考；↑Goの string 型は実際にはデータそのものを保持するのではなく、次の2つのフィールドを持つ構造体。なので16byteずつずれている
		// ・データへのポインタ（8バイト）
		// ・長さ（8バイト）
		stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&names[0]))
		fmt.Printf("Data Pointer: %p\n", unsafe.Pointer(stringHeader.Data))
		fmt.Printf("Length: %d\n", stringHeader.Len)

		// そもそも最初に確保する数を指定しなければ、Arrayを作った上でその参照からなるsliceを返してくれる。
		q := []int{2, 3, 5, 7, 11, 13}
		fmt.Println(q)
		s := []struct {
			val  int
			flag bool
		}{
			{1, true},
			{2, false},
			{5, false},
		}
		fmt.Printf("%T, %v\n", s, s)
		fmt.Println(q[1:])
		fmt.Println(q[:2])
		fmt.Println(q[:])

		printSlice(q)
		q = q[:0]
		printSlice(q)
		q = q[:4]
		printSlice(q)
		// スライスの容量は、スライスの開始位置から基になる配列の終わりまでの要素数を示す。
		// 従って、↓の操作で最初の2つがdropされ、capが4になってしまう。
		// すなわち、スライスの開始位置と終了位置は対称性を持たない。
		// 開始位置を変更することは容量に影響を与えるが、終了位置を変更しても容量には影響を与えない。
		q = q[2:]
		printSlice(q)

		var nilSlice []int
		fmt.Println(len(nilSlice), cap(nilSlice), nilSlice)
		if nilSlice == nil {
			fmt.Println("nil!!")
		}

		// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

		aa := make([]int, 10)
		printSlice(aa)
		b := make([]int, 0, 10)
		printSlice(b)
		b2 := b[:cap(b)]
		printSlice(b2)
		c := aa[8:]
		printSlice(c)

		printTicTacToe()

		// 足りなければ新たにallocateしてくれる
		var s2 []int
		printSlice((s2))
		s2 = append(s2, 0)
		printSlice((s2))
		s2 = append(s2, 1, 2, 3)
		printSlice((s2))

		var powedNum = []int{1, 2, 4, 8}
		for i, v := range powedNum {
			fmt.Printf("2**%d = %d\n", i, v)
		}
		for i := range powedNum {
			fmt.Println(i)
		}
		for _, v := range powedNum {
			// fmt.Println(_) // これはエラーになる
			fmt.Println(v)
		}

		// make しないと、ゼロ値である nil を指していて、キー設定できない
		var m map[string]Location
		fmt.Printf("map address: %p, content: %v\n", m, m)

		if m == nil {
			fmt.Println("The map is not initialized.")
		} else {
			fmt.Println("The map is initialized.")
		}
		// make で、map 内部で key と value を保存するためのハッシュテーブル用のメモリを割り当てている
		m = make(map[string]Location)
		fmt.Printf("map address: %p, content: %v\n", m, m)

		fmt.Println(m)
		m["Bell Labs"] = Location{
			40.68433, -74.39967,
		}
		fmt.Println(m["Bell Labs"])
		fmt.Println(m)

		// とはいえ、makeしなくても、初期値を設定すれば勝手にハッシュテーブル用のメモリは割り当てられる
		var m2 = map[string]Location{
			"a": Location{
				1.1, 2.2,
			},
			// 省略可能
			"b": {
				-3.343, -5.43,
			},
		}
		fmt.Printf("map address: %p, content: %v\n", m2, m2)

		m3 := make(map[string]int)
		m3["A"] = 1
		fmt.Println(m3["A"])
		m3["A"] = 3
		fmt.Println(m3["A"])
		m3V1, ok := m3["A"]
		fmt.Println(m3V1, ok)
		delete(m3, "A")
		fmt.Println(m3["A"])
		m3V2, ok := m3["A"]
		fmt.Println(m3V2, ok)
	}

	pos, neg := addr(), addr()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-i),
		)
	}
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}
