package main

import (
	"fmt"
)

type empty interface{}

func MyPrint(str ...string){
	for _, s := range str {
		fmt.Println(s)
	}
}

func main(){
	var x empty = 1.23

	switch x.(type){
	case int:
		fmt.Printf("it's int type. %d\n" , x);
	case string:
		fmt.Println("it's string type. " , x)
	default:
		fmt.Println("I don't know")
	}
	MyPrint("xxx" , "yyy" , "zzz")
}