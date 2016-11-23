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

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func newton(z, x float64) float64 {
	return z - (z*z-x)/2*z
}

func Sqrt(x float64) float64 {
	z := 1.0
	z1 := 100.0
	delta := 1000.0
	for delta > 0.01 {
		z1 = newton(z, x)
		delta = math.Abs(z1 - z)
		z = z1
	}
	return z
}

func switching() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OSX")
	case "linux":
		fmt.Println("LINUX")
	default:
		fmt.Printf("%s. \n", os)
	}
}

func tomorrow() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
		break
	case today + 1:
		fmt.Println("Tomorrow.")
		break
	case today + 2:
		fmt.Println("In two days.")
		break
	default:
		fmt.Println("Too far away.")
	}
}

func timeOfDay() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
func defered() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func deferedLoop() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 10),
	)
	fmt.Println(Sqrt(4.0))
	switching()
	tomorrow()
	defered()
	timeOfDay()
	deferedLoop()
}
