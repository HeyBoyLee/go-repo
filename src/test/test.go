package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type s_t struct{

}

var strArr = []string{"ok" , "hi"} // slice 类型
var strArr1 = [...]string{"ok" , "hi"}	// array 类型
func (s s_t) getName(name string){
	//strArr := []string{"ok" , "hi"}
	fmt.Printf("%T\n" , strArr)
	fmt.Printf("%T\n" , strArr1)
	s.getDetail(name , strArr)
}

func (s s_t) getDetail(elems ...interface{}){
	fmt.Printf("%T" , elems)
	s.getMore(elems...)
	//for _, v := range elems {
	//	fmt.Printf(string(v))
	//}
}

func (s s_t) getMore(e ...interface{}){
	fmt.Printf("%s\n" , e);
	for _, v := range e {
		fmt.Printf("%s\n" ,v)
	}
}

func main() {
	h := "hello world!"
	s := s_t{}
	s.getName(h)
	f := MyFloat(-math.Sqrt2)
	fmt.Println(-math.Sqrt2)
	fmt.Println(f.Abs())
}

