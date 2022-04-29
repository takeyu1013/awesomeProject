package main

import (
	"fmt"
	"os/user"
	"strconv"
	"strings"
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

func add(x, y int) (int, int) {
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	return
}

func convert(price int) float64 {
	return float64(price)
}

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
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
	fmt.Println(("Hello World"))
	fmt.Println("Hello " + "World")
	fmt.Println(string("Hello World"[0]))

	var s string = "Hello World"
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)
	fmt.Println(strings.Contains(s, "World"))
	fmt.Println(`Test
                        Test
Test`)

	fmt.Println("\"")
	fmt.Println(`"`)
	t, f := true, false
	fmt.Printf("%T, %v %t\n", t, 1, t)
	fmt.Printf("%T, %v %t\n", f, 0, f)

	// fmt.Println(true && true)
	fmt.Println(true && false)
	// fmt.Println(false && false)

	// fmt.Println(true || true)
	fmt.Println(true || false)
	// fmt.Println(false || false)

	fmt.Println(!true)
	fmt.Println(!false)

	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var s1 string = "14"
	i, _ := strconv.Atoi(s1)
	fmt.Printf("%T %v\n", i, i)

	h := "Hello World"
	fmt.Println(string(h[0]))

	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	/*
		var b [2]int = [2]int{100, 200}
		b = append(b, 300)
		fmt.Println(b)
	*/

	var b []int = []int{100, 200}
	b = append(b, 300)
	fmt.Println(b)

	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])
	fmt.Println(n[:])

	n[2] = 100
	fmt.Println(n)

	var board = [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	fmt.Println(board)
	n = append(n, 100, 200, 300, 400)
	fmt.Println(n)
	n1 := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n1), cap(n1), n1)
	n1 = append(n1, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n1), cap(n1), n1)
	n1 = append(n1, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n1), cap(n1), n1)

	a1 := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a1), cap(a1), a1)

	b1 := make([]int, 0)
	var c []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(b1), cap(b1), b1)
	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)

	// c = make([]int, 5)
	c = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)

	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)

	fmt.Println(m["nothing"])

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	/*
		var m3 map[string]int
		m3["pc"] = 5000
		fmt.Println(m3)
	*/

	var s2 []int
	if s2 == nil {
		fmt.Println("Nil")
	}

	b2 := []byte{72, 73}
	fmt.Println(b2)
	fmt.Println(string(b2))

	c2 := []byte("HI")
	fmt.Println(c2)
	fmt.Println(string(c2))

	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	f2 := func(x int) {
		fmt.Println("inner func", x)
	}
	f2(1)

	func(x int) {
		fmt.Println("inner func", x)
	}(1)

	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	c3 := circleArea(3.14)
	fmt.Println(c3(2))

	c4 := circleArea(3)
	fmt.Println(c4(2))
}
