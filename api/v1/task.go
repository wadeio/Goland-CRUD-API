package v1

import (
	"goapi/model"
	"goapi/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetTasks 查詢任務列表
func GetTasks(c *gin.Context) {
	data, code := model.GetTasks()

	if (code == errmsg.SUCCSE) {
		c.JSON(
			http.StatusOK, gin.H{
				"result":  data,
			},
		)
	} else {
		c.JSON(
			http.StatusOK, gin.H{
				"result":  data,
				"message": errmsg.GetErrMsg(code),
			},
		)
	}
}

// AddTask 添加任務
func AddTask(c *gin.Context) {
	var data model.Task
	_ = c.ShouldBindJSON(&data)

	code, task:=model.CreateTask(&data)

	if code == errmsg.SUCCSE {
		c.JSON(
			http.StatusOK, gin.H{
				"result":  task,
			},
		)
	} else {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			},
		)
	}
}

// UpdateTask 更新任務
func UpdateTask(c *gin.Context) {
	var data model.Task
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.CheckTask(id)
	if code == errmsg.SUCCSE {
		_,task:=model.EditTask(id, &data)
		c.JSON(
			http.StatusOK, gin.H{
				"result":  task,
			},
		)
	} else {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			},
		)
	}

}

// DeleteTask 刪除任務
func DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteTask(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}