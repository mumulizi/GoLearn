package main

import "fmt"

type Student struct {
	Name string
	Age int
	Score float32
}

func main()  {
	var str = "stu01 10 99"
	var stu Student

	fmt.Sscanf(str,"%s %d %f",&stu.Name,&stu.Age,&stu.Score)
	fmt.Println(stu)
}