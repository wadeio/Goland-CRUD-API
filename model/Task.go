package model

import (
	"fmt"
	"encoding/json"
	"goapi/utils/errmsg"
	//"gorm.io/gorm"
)

type Task struct {
	Id   int    `json:"id"`
	Name string `json:name`
	Status int `json:status`
}

type Tasks []Task
var Data = []byte(`[{"id" : 1 , "name" : "買早餐", "status" : 0},{"id" : 2 , "name" : "買晚餐", "status" : 0}]`)
var tasks Tasks
var FIRST_INIT=0 

// GetTasks 查詢任務列表
func GetTasks() ([]Task, int) {
	if FIRST_INIT==0 {
		// json.Unmarshal 把 json 字串轉成 struct
		json.Unmarshal(Data, &tasks)
		FIRST_INIT=1
	}
    fmt.Println("get tasks list") 
	//fmt.Println(tasks)
	// json.Marshal 把 struct 轉成 json 字串
	// jsontasks, _ := json.Marshal(tasks)

	return tasks, errmsg.SUCCSE
}

// CheckTask 查詢任務是否存在
func CheckTask(id int) (code int) {
	fmt.Println(tasks)
	for counter := 0; counter < len(tasks); counter++ {
		if tasks[counter].Id==id {
			return errmsg.SUCCSE
		}
	}

	return errmsg.ERROR
}

// CreateTask 新增任務
func CreateTask(data *Task) (int , Task){
	newId:=len(tasks)+1
	task:=Task{Id:newId,Name:data.Name,Status:0}
	tasks = append(tasks, task)
	fmt.Println("add task")
	
	return errmsg.SUCCSE, task
}

// EditTask 編輯任務
func EditTask(id int, data *Task) (int, Task) {
	var task Task
	for counter := 0; counter < len(tasks); counter++ {
		if tasks[counter].Id==id {
			tasks[counter].Name=data.Name
			tasks[counter].Status=data.Status
			task=tasks[counter]
		}
	}
	return errmsg.SUCCSE, task
}

// DeleteTask 刪除任務
func DeleteTask(id int) int {
	temp := tasks[:0]
	// fmt.Println(temp)
	for _, x := range tasks {
		if x.Id != id {
			temp = append(temp, x)
		}
	}
	tasks = temp
	return errmsg.SUCCSE
}

