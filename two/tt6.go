package main

import "fmt"

func main()  {
	var ch1 chan bool
	ch1 = make(chan bool)

	go func() {
		for i:=0;i<10;i++{
			fmt.Println(i)
		}
		fmt.Println("over")
	}()

	fmt.Println("main --- over")

}
