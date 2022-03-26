package routes

import (
	"anyData/internal/app/anyData/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes() (router *gin.Engine) {
	router = gin.Default()
	err := router.SetTrustedProxies([]string{"192.168.0.1"})
	if err != nil {
		return nil
	}

	router.POST("/addData", controller.AddData)
	router.POST("/addDataset", controller.AddDataset)

	router.GET("/fetchDataset", controller.FetchDataset)
	router.GET("/fetchData", controller.FetchData)

	router.PATCH("/updateData", controller.UpdateData)
	router.DELETE("/deleteData", controller.DeleteData)

	return router
}
