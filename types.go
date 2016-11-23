package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"strings"
)

func first() {
	i, j := 42, 2710
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 3}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	pV = &Vertex{1, 2}
	a  [2]string
)

func slices() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func slice2() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func slice3() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[1:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func makeArray() {
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func slicesSlices() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func sliceAppend() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1, 2, 3)
	a := []int{4, 5}
	s = append(s, a...)
	printSlice(s)
}

func doRange() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func noIndexRange() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func countFunc(x, y int) uint8 {
	return uint8((x + y) / 2)
}

func Pic(dx, dy int) [][]uint8 {
	resArray := make([][]uint8, dx)
	//resArray := [dy][dx]uint8{}
	//println(resArray)
	for i, y := range resArray {
		resArray[i] = make([]uint8, dy)
		//println(len(y))
		for j := range y {
			//println(i + j)
			//println(i, j)
			resArray[i][j] = countFunc(j, i)
		}
	}
	return resArray
}

func main() {
	first()
	v := Vertex{1, 2}
	v.X = 4
	p := &v
	p.Y = 1e9 // (*p).Y
	a[0] = "hellow"
	a[1] = "world"
	fmt.Println(v, v.X)
	fmt.Println(v1, pV, v2, v3)
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	s := primes[1:4]
	var s1 []int = primes[1:4]
	fmt.Println(primes, s, s1)
	slices()
	slice2()
	fmt.Println("---------")
	slice3()
	var ss []int
	fmt.Println(ss, len(ss), cap(ss))
	if ss == nil {
		fmt.Println("nil!")
	}
	fmt.Println("---------")
	makeArray()
	fmt.Println("------------------------")
	slicesSlices()
	fmt.Println("------------------------")
	sliceAppend()
	fmt.Println("------------------------")
	doRange()
	noIndexRange()
	fmt.Println("------------------------")
	pic.Show(Pic)
}
