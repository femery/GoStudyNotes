package main

import "fmt"

func main() {

	key := ""
	flag := true
	for {
		fmt.Println("-------------家庭收支记账软件📖-----------------")
		fmt.Println("1 收支明细")
		fmt.Println("2 登记收入")
		fmt.Println("3 登记支出")
		fmt.Println("4 退出软件")
		fmt.Println("please choose a number between 0ne and four")

		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("")
		case "2":
			fmt.Println("")
		case "3":
			fmt.Println("")
		case "4":
			fmt.Println("end")
			flag = false
		default:
			fmt.Println("please input right number")
		}

		if !flag {
			break
		}

	}

	fmt.Println("软件关闭😄")
}
