package main

import (
	"fmt"
	"io"
	"log"
	"os"
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

func foo2(params ...int) {
	fmt.Println(len(params), params)
	for _, param := range params {
		fmt.Println(param)
	}
}

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func getOsName() string {
	return "dafdafad"
}

func foo3() {
	defer fmt.Println("world foo")

	fmt.Println("hello foo")
}

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func thirdPartyConnectDB() {
	panic("Unable to connect database")
}

func save() {
	thirdPartyConnectDB()
}

func one(x *int) {
	*x = 1
}

type Vertex struct {
	X, Y int
	S    string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 1000
}

type Vertex2 struct {
	x, y int
}

func (v Vertex2) Area() int {
	return v.x * v.y
}

func (v *Vertex2) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

func Area(v Vertex2) int {
	return v.x * v.y
}

func New(x, y int) *Vertex2 {
	return &Vertex2{x, y}
}

type Vertex3D struct {
	Vertex2
	z int
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func New3D(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex2{x, y}, z}
}

type MyInt int

func (i MyInt) Double() int {
	fmt.Printf("%T %v\n", i, i)
	fmt.Printf("%T %v\n", 1, 1)
	return int(i * 2)
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

	foo2()
	foo2(10, 20)
	foo2(10, 20, 30)

	s3 := []int{1, 2, 3}
	fmt.Println(s3)

	foo2(s3...)

	f3 := 1.11
	i2 := int(f3)
	fmt.Printf("%T %v\n", i2, i2)

	s4 := []int{1, 2, 5, 6, 2, 3, 1}
	fmt.Println(s4[2:4])

	m3 := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T %v\n", m3, m3)

	result := by2(10)

	if result == "ok" {
		fmt.Println("great")
	}
	fmt.Println(result)

	if result2 := by2(10); result2 == "ok" {
		fmt.Println("great 2")
	}
	// fmt.Println(result2)
	/*
		num := 6
		if num%2 == 0 {
			fmt.Println("by 2")
		} else if num%3 == 0 {
			fmt.Println("by 3")
		} else {
			fmt.Println("else")
		}
	*/

	x2, y2 := 11, 12
	if x2 == 10 && y2 == 10 {
		fmt.Println("&&")
	}

	if x2 == 10 || y2 == 10 {
		fmt.Println("||")
	}

	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}

		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}

	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)

	/*
		for {
			fmt.Println("hello")
		}
	*/

	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	for i, v := range l {
		fmt.Println(i, v)
	}

	for _, v := range l {
		fmt.Println(v)
	}

	m4 := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m4 {
		fmt.Println(k, v)
	}

	for k := range m4 {
		fmt.Println(k)
	}

	for _, v := range m4 {
		fmt.Println(v)
	}

	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows!!")
	default:
		fmt.Println("Default!!", os)
	}

	t2 := time.Now()
	fmt.Println(t2.Hour())
	switch {
	case t2.Hour() < 12:
		fmt.Println("Morning")
	case t2.Hour() < 17:
		fmt.Println("Afternoon")
	}

	/*
		defer fmt.Println("world")

		foo3()

		fmt.Println("hello")
	*/

	/*
		fmt.Println("run")
		defer fmt.Println(1)
		defer fmt.Println(2)
		defer fmt.Println(3)
		fmt.Println("success")
	*/
	file, _ := os.Open("./lesson.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))

	LoggingSettings("test.log")
	_, err := os.Open("fdafdsafa")
	if err != nil {
		// log.Fatalln("Exit", err)
	}
	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	// log.Fatalf("%T %v", "test", "test")
	// log.Fatalln("error!!")

	fmt.Println("ok!")

	file2, err := os.Open("./lesson.go")
	if err != nil {
		log.Fatal("Error!")
	}
	defer file2.Close()
	data2 := make([]byte, 100)
	count, err := file.Read(data2)
	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data2))

	if err = os.Chdir("test"); err != nil {
		// log.Fatalln("Error")
	}

	// save()
	fmt.Println("OK?")

	l2 := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	var min int

	for i, num := range l2 {
		if i == 0 {
			min = num
			continue
		}

		if min > num {
			min = num
		}
	}

	fmt.Println(min)

	m5 := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	sum2 := 0
	for _, v := range m5 {
		sum2 += v
	}

	fmt.Println(sum2)

	var n2 int = 100
	one(&n2)
	fmt.Println(n2)
	/*
		fmt.Println(n2)

		fmt.Println(&n2)

		var p *int = &n2

		fmt.Println(p)

		fmt.Println(*p)
	*/

	s5 := make([]int, 0)
	fmt.Printf("%T\n", s5)

	m6 := make(map[string]int)
	fmt.Printf("%T\n", m6)

	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	var p2 *int = new(int)
	fmt.Printf("%T\n", p2)

	var st = new(struct{})
	fmt.Printf("%T\n", st)

	/*
		var p2 *int = new(int)
		fmt.Println(*p2)
		*p2++
		fmt.Println(*p2)

		var p3 *int
		fmt.Println(p3)
		*p3++
		fmt.Println(p3)
	*/

	v3 := Vertex{1, 2, "test"}
	changeVertex(v3)
	fmt.Println(v3)

	v4 := &Vertex{1, 2, "test"}
	changeVertex2(v4)
	fmt.Println(*v4)

	/*
		v3 := Vertex{X: 1, Y: 2}
		fmt.Println(v3)
		fmt.Println(v3.X, v3.Y)

		v4 := Vertex{X: 1}
		fmt.Println(v4)

		v5 := Vertex{1, 2, "test"}
		fmt.Println(v5)

		v6 := Vertex{}
		fmt.Printf("%T %v\n", v6, v6)

		var v7 Vertex
		fmt.Printf("%T %v\n", v7, v7)

		v8 := new(Vertex)
		fmt.Printf("%T %v\n", v8, v8)

		v9 := &Vertex{}
		fmt.Printf("%T %v\n", v9, v9)

		s := make([]int, 0)
		s := []int{}
		fmt.Println(s)
	*/

	var i3 int = 10
	var p3 *int
	p3 = &i3
	var j int
	j = *p3
	fmt.Println(j)

	var i4 int = 100
	var j2 int = 200
	var p4 *int
	var p5 *int
	p4 = &i4
	p5 = &j2
	i4 = *p4 + *p5
	p5 = p4
	j2 = *p5 + i4
	fmt.Println(j2)

	// v5 := Vertex2{3, 4}
	// fmt.Println(Area(v5))
	v5 := New(3, 4)
	v5.Scale(10)
	fmt.Println(v5.Area())

	v6 := New3D(3, 4, 5)
	v6.Scale3D(10)
	// fmt.Println(v6.Area())
	fmt.Println(v6.Area3D())

	myInt := MyInt(10)
	fmt.Println(myInt.Double())
}
