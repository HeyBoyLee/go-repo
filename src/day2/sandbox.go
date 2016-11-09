package main

import ("fmt"
	"time"
	"math/rand"
	"day1"
	"math/cmplx")

func swap (a string, b string) (string ,string){
	return b , a
}

func main() {
	var x int
	x = 10
	fmt.Println("x:",x)
	fmt.Println("welcome to the playgound!")
	fmt.Println("The time is" , time.Now())
	fmt.Println(time.Now())
	rand.Seed(int64(time.Now().Nanosecond()))
	for i:= 0;i < 5;i++{
		fmt.Println("My favorite number is" , rand.Intn(100))
	}
	a,b := swap("hello" , "world")
	fmt.Println(a , b)

	var z complex128 = cmplx.Sqrt(-5 + 12i)
	//常量不能使用 := 语法定义。
	const f = "%T(%v)\n"
	fmt.Printf(f , x, x)
	fmt.Printf(f , z , z)
	fmt.Printf("%s" , f)
	xn.HelloWorld()
}