package routes

import (
	"goapi/api/v1"
	//"goapi/middleware"
	"goapi/utils"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	//p.AddFromFiles("admin", "web/admin/dist/index.html")
	//p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.HTMLRender = createMyRender()
	//r.Use(middleware.Log())
	//r.Use(gin.Recovery())
	//Recovery 中間件會恢復任何(panics) 
	//r.Use(middleware.Cors())
	
	// r.Static("/static", "./web/front/dist/static")
	// r.Static("/admin", "./web/admin/dist")
	// r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(200, "front", nil)
	// })

	// r.GET("/admin", func(c *gin.Context) {
	// 	c.HTML(200, "admin", nil)
	// })

	/*
	api
	 */
	router := r.Group("api/v1")
	{
		router.GET("tasks", v1.GetTasks)
		router.POST("task", v1.AddTask)
		router.PUT("task/:id", v1.UpdateTask)
		router.DELETE("task/:id", v1.DeleteTask)
	}

	_ = r.Run(utils.HttpPort)

}
