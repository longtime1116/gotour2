package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

var rust, python, java bool
var hoge int = 1

// まとめることも可能
var (
	// 最初が大文字だと、package外でも利用可能。
	IsValid bool   = false
	MaxInt  uint64 = 1<<64 - 1
	// エラーになる
	// neverOverFlowed uint64 = 1 << 64
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

// https://go.dev/blog/declaration-syntax
func add(x int, y int) int {
	return x + y
}

// 続く場合は省略可能
func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// 短い関数ならこういう書き方をしても可読性が失われない
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// type conversions
func leftEdgeLengh(x, y int) int {
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	return z
}

func needInt(x int) int           { return x * 10 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	if false {
		fmt.Println("My favorit number is", rand.Intn(10))
		fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
		// export されるものは必ず大文字始まり。定数も関数も。
		fmt.Println(math.Pi)

		fmt.Println(add(11, 23))
		fmt.Println(add2(11, 23))
		fmt.Println(swap("aaa", "b"))
		fmt.Println(split(100))
		var fuga int = 10
		// 省略可能
		var ruby, cobol = true, "no!"
		fmt.Println(hoge, fuga, rust, python, java, ruby, cobol)

		fmt.Printf("Type: %T Value: %v\n", IsValid, IsValid)
		fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
		fmt.Printf("Type: %T Value: %v\n", z, z)

		fmt.Println(leftEdgeLengh(5, 11))

		// 型を推論して変数定義して代入をしてくれる書き方。
		v := 32
		fmt.Println(v)

		// 定数は := では定義できない
		const World = "World"
		fmt.Println("Hello", World)
	}

	// 定数は無限精度の数値であり、BigやSmallはintなどの型を持たないので、内部的には数値を保持できている。
	// 具体的な型にキャストしたときに、オーバーフローするかどうかが決まる
	// 引数に渡すときにintなどに変換される
	fmt.Println(Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	// fmt.Println(needInt(Big)) // Error!
	fmt.Println(needFloat(Big))
}
