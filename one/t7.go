package main

import "github.com/gin-gonic/gin"

func main()  {
	r:=gin.Default()
	r.Use(IPCheck())

	r.GET("/test", func(c *gin.Context) {
		c.String(200,"hello test")

	})
	r.Run()

}

func IPCheck() gin.HandlerFunc {
	return func(c * gin.Context) {
		ipList :=[]string{
			"127.0.0.2",
		}
		flag :=false
		clientIp:=c.ClientIP()
		for _,host:=range ipList{
			if clientIp==host {
				flag=true
				break
			}
		}
		if !flag {
			c.String(401,"%s,not in iplist",clientIp)
			c.Abort()
		}

	}
}
