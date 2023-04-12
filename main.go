package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("123")

	var r = gin.Default() //创建一个默认引擎

	//用户注册
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		//获取参数
		var user = ctx.PostForm("username")
		var password = ctx.PostForm("password")
		var telephone = ctx.PostForm("telephone")
		//数据校验
		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "手机号必须为11位"})
		}

		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "密码不能少于6位"})
		}
		fmt.Println(len(user))
		if len(user) == 0 {
			user = randomString(9) //如果没有用户名，那么就生成一个是个字符串的随机用户名
		}
		//判断手机号是否存在

		//创建用户

		//返回结果
		ctx.JSON(http.StatusOK, gin.H{
			"user":      user,
			"password":  password,
			"telephone": telephone,
		})
	})
	r.Run()

}
func randomString(n int) string {
	var letters = "qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM" //随机用户名的字符种子
	var result = make([]byte, n)
	rand.Seed(time.Now().Unix()) //用于接收随机用户名的结果
	for i := range letters {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)

}
