package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
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

	router.Run(":8081")
}
