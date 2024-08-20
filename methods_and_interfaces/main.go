package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

type MyFloat64 float64

// receiver 引数があるだけで、実際のところただの関数。
func (v *Vertex) Abs() float64 {
	// このように、nilのケースに対処する処理を素直に書くのが一般的
	if v == nil {
		return 0
	}
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインタのレシーバーにする理由：
//  1. 構造体の持つvalueを変更できる
//  2. レシーバーの値をコピーしないので、メモリ効率が高まる
//  3. 特定のinterfaceを満たすという性質に合わせるために、ポインタにしておきたくなる
func (v *Vertex) Scale(n float64) {
	v.X *= n
	v.Y *= n
	// こうやって書くこともできる
	// (*v).X *= n
	// (*v).Y *= n
}

func (x MyFloat64) Abs() float64 {
	if x < 0 {
		return float64(-x)
	}
	return float64(x)
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

	// Abser インタフェースを満たすもののみに代入可能
	// ここで、a は Abser インタフェースを満たす何らかのものであることを示す
	var a Abser
	a = &v1
	a = x1
	fmt.Println(a)
	// a = v1 // エラーになる。
	// こんなふうにもできる
	var b Abser = &Vertex{1, 2}
	fmt.Println(b)

	var p *Vertex
	b = p
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", p, p)
	fmt.Println(b.Abs())

	var c Abser
	fmt.Printf("%v, %T\n", c, c)
	// c.Abs() // SEGV
}
