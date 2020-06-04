package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Person struct {
	Name string `form:"name"`
	Address string `form:"address"`
	Birthday time.Time `form:"birthday"`
}

func  main()  {
	r:=gin.Default()
	r.GET("/testing",testing)
	r.POST("/testing",testing)

}

func testing(c *gin.Context){
	var person Person
	//这里根据content-type来做不同binding操作
	if err:=c.ShouldBind(&person); err==nil {
		c.String(200,"%v",person)
	}else {
		c.String(200,"person bind error:%v",err)
	}
}
