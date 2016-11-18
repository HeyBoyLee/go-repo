// hello.go
package xn

import (
	"fmt"
)

func HelloWorld() {
	var x int
	x = 10
	fmt.Println(x)
	fmt.Println("Hello World!")
	GetWorld()
	str:=`this is a string`
	fmt.Printf("type of str: %T " , str)
}
