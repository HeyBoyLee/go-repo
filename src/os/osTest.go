package main

import(
	"fmt"
	"os")


//go run os/osTest.go 1 2
func main() {
	param1:=os.Args[1]
	param2:=os.Args[2]
	fmt.Println(param1,param2)
}
