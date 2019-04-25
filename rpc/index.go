package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kylelemons/go-gypsy/yaml"
)

var Config *yaml.File

func main() {

	Config, _ = yaml.ReadFile("conf.yaml")
	fmt.Println(Config.Get("sid"))

	router := gin.Default()

	types := router.Group("/types")
	{
		// Query string parameters are parsed using the existing underlying request object.
		// The request responds to a url matching:  /type?page=1&pageSize=20
		types.GET("", GetTypes)

		types.POST("", CreateType)

		types.PUT("", UpdateType)

		types.DELETE("", DeleteType)
	}

	questions := router.Group("/questions")
	{
		// Query string parameters are parsed using the existing underlying request object.
		// The request responds to a url matching:  /type?page=1&pageSize=20
		questions.GET("", GetQuestions)

		questions.POST("", CreateQuestion)

		questions.PUT("", UpdateQuestion)

		questions.DELETE("", DeleteQuestion)

		questions.POST("/search", SearchQuestions)
	}


	results := router.Group("/results")
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

	router.Run(":8080")
}
