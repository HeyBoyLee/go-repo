package main

import (
	"fmt"
)
func catch(){
	if err := recover(); err != nil {
		fmt.Println("err:",err)
	}
}

func main(){
	defer catch()
	x := 10/0
	fmt.Println(x)
}
