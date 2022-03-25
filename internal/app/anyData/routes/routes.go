package routes

import (
	"anyData/internal/app/anyData/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes() (router *gin.Engine) {
	router = gin.Default()

	router.POST("/addData", controller.AddData)
	router.POST("/addDataset", controller.AddDataset)

	router.GET("/fetchDataset", controller.FetchDataset)
	router.GET("/fetchData", controller.FetchData)

	router.PATCH("/updateData", controller.UpdateData)
	router.DELETE("/deleteData", controller.DeleteData)

	return router
}
