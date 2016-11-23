package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"time"
)

const _2PI = 6.28
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func needConst() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

var c, python, java bool

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func spit(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func convert() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
	fmt.Println(x, y, z)
}

// no ":=" outside

func main() {
	a, b := swap("hello", "world")
	fmt.Println("My favoritenumber is ", add(42, 13),
		time.Now(), math.Sqrt(7), a, b)
	fmt.Println(spit(17))
	var i int64
	m := 10
	var d, e = 1, 5
	var j, k int = 3, 4
	cpp, python3, js := true, 222, "no"
	fmt.Println(i, c, python, java, d, e, j, k)
	fmt.Println(cpp, python3, js, m, _2PI)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	var f float32
	var boo bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, boo, s)
	convert()
	needConst()
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
