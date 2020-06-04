package main

import "fmt"

func main()  {

	go printNum()

	for i:=1;i<1000;i++{
		fmt.Println("打印字母A----%c")
	}

}

func printNum()  {
	for i:=1;i<1000;i++{
		fmt.Println("打印数字----%d",i)
	}

}
