package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kylelemons/go-gypsy/yaml"
	"time"
)

var Config *yaml.File
var locZone = time.FixedZone("CST", 8*3600)

func main() {

	Config, _ = yaml.ReadFile("conf.yaml")
	fmt.Println(Config.Get("sms.sid"))

	router := gin.Default()

	types := router.Group("/types")
	types.Use(IsLogin())
	{
		// Query string parameters are parsed using the existing underlying request object.
		// The request responds to a url matching:  /type?page=1&pageSize=20
		types.GET("", GetTypes)
	}
	types.Use(IsAdmin())
	{
		types.POST("", CreateType)
		types.PUT("", UpdateType)
		types.DELETE("", DeleteType)
	}

	questions := router.Group("/questions")
	questions.Use(IsLogin())
	{
		// Query string parameters are parsed using the existing underlying request object.
		// The request responds to a url matching:  /type?page=1&pageSize=20
		questions.GET("", GetQuestions)
		questions.POST("/search", SearchQuestions)
	}
	questions.Use(IsAdmin())
	{
		questions.POST("", CreateQuestion)
		questions.PUT("", UpdateQuestion)
		questions.DELETE("", DeleteQuestion)
	}

	router.POST("/testing", IsLogin(), Testing)
	results := router.Group("/results")
	results.Use(IsLogin())
	{
		// Query string parameters are parsed using the existing underlying request object.
		// The request responds to a url matching:  /type?page=1&pageSize=20
		results.GET("", GetResults)
		results.POST("", CreateResult)
		results.PUT("", UpdateResult)
		results.DELETE("", DeleteResult)
	}

	sms := router.Group("/sms")
	{
		sms.POST("", Send)
	}

	users := router.Group("/users")
	{
		users.POST("", Register)
		users.POST("/login", Login)
	}

	router.Run(":8080")
}
