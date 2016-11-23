package main

import "fmt"

type Info_t struct{
	Name string `name`
	Password string	`password`
	Age string	`age`
}

func test(s interface{}){
	//r := s.(map[string]interface{})
	//fmt.Println(r["name"])
	//fmt.Println(s)
	r := s.(*Info_t)
	fmt.Println(r.Name)
}

func main() {
	v := Info_t{
		Name: "huifeng",
		Password: "12",
		Age : "22"}
	test(&v)
}
