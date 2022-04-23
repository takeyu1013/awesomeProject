package main

import (
	"fmt"
	"os/user"
	"time"
)

const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

// var big int = 9223372036854775807 + 1
const big = 9223372036854775807 + 1

var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func foo() {
	xi := 1
	var xf32 float32 = 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf32, xs, xt, xf)
	fmt.Printf("%T\n", xf32)
	fmt.Printf("%T\n", xi)
}

func main() {
	fmt.Println("Hello, world!", time.Now())
	fmt.Println(user.Current())
	fmt.Println(i, f64, s, t, f)
	foo()
	fmt.Println(Pi, Username, Password)
	fmt.Println(big - 1)
	/*
		var (
			u8  uint8     = 255
			i8  int8      = 127
			f32 float32   = 0.2
			c64 complex64 = -5 + 12i
		)
		fmt.Println(u8, i8, f32, c64)
		fmt.Printf("type=%T value=%v\n", u8, u8)
	*/

	/*
		fmt.Println("1 + 1 =", 1+1)
		fmt.Println("10 - 1 =", 10-1)
		fmt.Println("10 / 2 =", 10/2)
		fmt.Println("10 / 3 =", 10/3)
		fmt.Println("10.0 / 3 =", 10.0/3)
		fmt.Println("10 / 3.0 =", 10/3.0)
		fmt.Println("10 % 2 =", 10%2)
		fmt.Println("10 % 3 =", 10%3)
	*/

	/*
		x := 0
		fmt.Println(x)
		// x = x + 1
		x++
		fmt.Println(x)
		// x = x - 1
		x--
		fmt.Println(x)
	*/

	fmt.Println(1 << 0) // 0001 0001
	fmt.Println(1 << 1) // 0001 0010
	fmt.Println(1 << 2) // 0001 0100
	fmt.Println(1 << 3) // 0001 1000
}
