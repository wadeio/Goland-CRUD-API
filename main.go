package main

import (
	"fmt"
	//"goapi/model"
	"goapi/routes"
)

func main() {
	fmt.Println("載入路由组件")
    // 引入路由组件
	routes.InitRouter()

}
