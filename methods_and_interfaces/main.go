package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat64 float64

// receiver 引数があるだけで、実際のところただの関数。
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインタのレシーバーにする理由：
//  1. 構造体の持つvalueを変更できる
//  2. レシーバーの値をコピーしないので、メモリ効率が高まる
func (v *Vertex) Scale(n float64) {
	v.X *= n
	v.Y *= n
	// こうやって書くこともできる
	// (*v).X *= n
	// (*v).Y *= n
}

func (x MyFloat64) Abs() MyFloat64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	v1 := Vertex{3, 4}
	fmt.Println(v1.Abs())
	// 構造体でなくても、自身のパッケージのtypeに対してならなんでもメソッドを生やすことができる。
	var x1 MyFloat64 = -math.Sqrt2
	fmt.Println(x1.Abs())
	v1.Scale(2)
	// もちろんこうやって書くこともできる
	// (&v1).Scale(2)
	fmt.Println(v1)
}
