package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func Testing(c *gin.Context) {
	db := Conn()
	questionIds := c.QueryMap("question_id")
	typeId, err := strconv.ParseUint(c.Query("type_id"), 10, 64)

	if typeId <= 0 {
		c.JSON(http.StatusOK, Response{RespCode: RespStatusArgs, RespDesc: err.Error()})
		c.Abort()
		return
	}

	var testingType Type
	if err := db.Model(&Question{}).Where("id=?", typeId).First(&testingType).Error; err != nil {
		c.JSON(http.StatusOK, Response{RespCode: RespStatusArgs, RespDesc: err.Error()})
		c.Abort()
		return
	}


	var questionIdArr []uint64
	for key, _ := range questionIds {
		questionId, _ := strconv.ParseUint(key, 10, 64)
		questionIdArr = append(questionIdArr, questionId)
	}

	// 读取
	var questions []Question
	db.Model(&Question{}).Where("id IN (?)", questionIdArr).Find(&questions)

	var result Result
	if err := db.Model(&Result{}).Where("type_id=?", testingType.ID).First(&result).Error; err != nil {
		c.JSON(http.StatusOK, Response{RespCode: RespStatusArgs, RespDesc: err.Error()})
		c.Abort()
		return
	}

	// 初始化计分
	var score uint = 0
	var totalScore uint = 0
	// 开启事务
	ts := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			ts.Rollback()
		}
	}()

	if err := ts.Error; err != nil {
		abort(c, ts, err)
		return
	}

	// 要么全新增，要么全更新
	if result.ID <= 0 {
		result.TypeId = testingType.ID
		result.UserId = global.Token.UserId

		if err := ts.Create(&result).Error; err != nil {
			abort(c, ts, err)
			return
		}

		for _, question := range questions {
			var resultItem ResultItem
			if value, ok := questionIds[strconv.Itoa(int(question.ID))]; ok && value == question.Answer {
				score++
				resultItem = ResultItem{UserId: global.Token.UserId, ResultId: result.ID, QuestionId: question.ID, UserAnswer: value, IsRight: 1}
			} else {
				resultItem = ResultItem{UserId: global.Token.UserId, ResultId: result.ID, QuestionId: question.ID, UserAnswer: value, IsRight: 0}
			}
			if err := ts.Create(&resultItem).Error; err != nil {
				abort(c, ts, err)
				return
			}
			totalScore++
		}
	} else {
		for _, question := range questions {
			var resultItem ResultItem
			if err := ts.Model(&ResultItem{}).
				Where("user_id=? AND result_id=? AND question_id=?", global.Token.UserId, result.ID, question.ID).
				First(&resultItem).Error; err != nil {
				abort(c, ts, err)
				return
			}

			if value, ok := questionIds[strconv.Itoa(int(question.ID))]; ok && value == question.Answer {
				score++
				resultItem.UserAnswer = value
				resultItem.IsRight = 1
			} else {
				resultItem.UserAnswer = value
				resultItem.IsRight = 0
			}
			if err := ts.Save(&resultItem).Error; err != nil {
				abort(c, ts, err)
				return
			}
			totalScore++
		}
	}

	// 更新得分
	result.Score = score
	result.TotalScore = totalScore
	if err := ts.Save(&result).Error; err != nil {
		abort(c, ts, err)
		return
	}

	// 返回最终结果
	ts.Commit()
	c.JSON(http.StatusOK, Response{RespCode: RespStatusOK, RespDesc: "success", RespData: TestResult{Score:score, TotalScore:totalScore,Answer:questionIds}})
}

func abort(c *gin.Context, ts *gorm.DB, err error) {
	ts.Rollback()
	c.JSON(http.StatusOK, Response{RespCode: RespStatusArgs, RespDesc: err.Error()})
	c.Abort()
}
