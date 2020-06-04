package main

import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GC)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
}
