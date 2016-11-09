package main

import ("fmt"
	"time"
	"math/rand")
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

}