package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main()  {

	//var a int
	//var b bool
	//fmt.Scan(&a,&b)
	//fmt.Printf("a:%d  b:%t",a,b)
	//fmt.Println(a,b)


	//readfromIO()

	randomNum()
}
func readfromIO(){
	fmt.Println("please input a str")
	reader := bufio.NewReader(os.Stdin)
	s1,_ := reader.ReadString('\n')
	fmt.Println("read data",s1)

}

func randomNum () {
	for i:=0;i<10;i++ {
		num :=rand.Intn(20)
		fmt.Printf("%d \t",num)
	}

}
