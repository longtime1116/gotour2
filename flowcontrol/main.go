package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	} else {
		fmt.Printf("%g > %g\n", v, lim)
	}
	return lim
}

func whenSaturday(today time.Weekday) string {
	// この例だとわからないが、特定のcaseに妥当すると、そのまま抜け出すのでbreakは不要。
	switch time.Saturday {
	case today:
		return "Today!"
	case today + 1:
		return "Tomorrow!!!"
	default:
		return "Too far away..."
	}
}

func main() {
	sum := 0
	// loop には for を使う。forしかない。
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// for {
	// 	fmt.Println("infinite loop")
	// }

	fmt.Println(sqrt(4), sqrt(2), sqrt(-2))

	fmt.Println(pow(2, 2, 10), pow(3, 3, 20))

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// メソッドチェーンも使える
	println(whenSaturday(time.Friday))
	println(whenSaturday(time.Now().Weekday()))

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	fmt.Println("defer test start")
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}
