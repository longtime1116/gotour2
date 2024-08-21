package main

import (
	"fmt"
	"image"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

type Person struct {
	Name string
	Age  int
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

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v %s", e.When, e.What)
}

// error interface は Error() メソッドを持つことを求める
func run() error {
	return &MyError{time.Now(), "it didnt work"}
}

type rot13Reader struct {
	r io.Reader
}

// s := strings.NewReader("Lbh penpxrq gur pbqr!")
// strings.NewReader() の戻り値をそのまま使うと、io.Readerが s.Read()して、引数のbufに読んだものを格納して n と error を返す。
// ここではそれをwrapして、bufの中の値を適切に変更して格納し返してあげている
func (r13r *rot13Reader) Read(buf []byte) (int, error) {
	n, err := r13r.r.Read(buf)
	// fmt.Printf("n = %v err = %v \n", n, err)
	// fmt.Printf("b[:n] = %q\n", buf[:n])
	if err == io.EOF {
		return n, err
	}
	for i, c := range buf[:n] {
		switch {
		case c >= 'A' && c <= 'Z':
			buf[i] = 'A' + (c-'A'+13)%26
		case c >= 'a' && c <= 'z':
			buf[i] = 'a' + (c-'a'+13)%26
		}
	}
	return n, nil
}

func main() {
	if false {
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

		// どんなものでも受け付けるよ、というシグニチャ表現のために用いる
		// fmt.Printf もそう。
		// 		type any = interface{}
		var i interface{}
		describe(i)
		i = 10
		describe(i)
		num, err := i.(int)
		if err {
			describe(num)
		} else {
			fmt.Println("i is not an integer")
		}
		i = "aaa"
		num2, err := i.(int)
		if err {
			describe(num2)
		} else {
			fmt.Println("i is not an integer")
		}
		// err を取ろうとしないとpanicになる
		// num3 := i.(int)
		// fmt.Println(num3)

		switch i.(type) {
		case int:
			fmt.Println("i is int")
		case string:
			fmt.Println("i is string")
		default:
			fmt.Println("I dont know this type.")
		}

		// Person は Stringer interface を満たしている。独自のString()メソッドが呼ばれる
		p1 := Person{"Taro", 21}
		fmt.Println(p1)
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}

	// io.Reader interface は以下の通り
	// type Reader interface {
	// 	Read(p []byte) (n int, err error)
	// }
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r13r := rot13Reader{s}
	io.Copy(os.Stdout, &r13r)

	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
