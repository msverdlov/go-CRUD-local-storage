package main

import (
	"anyData/internal/app/anyData/helper"
	"anyData/internal/app/anyData/localStorage"
	"anyData/internal/app/anyData/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"math"
	"net/http"
	"strconv"
)

var (
	host string
	port string
	url  string
	dataIds helper.Counter
	idToData = make(map[uint64]model.AnyData)
)

func main() {
	getConfigData()
	router := gin.Default()

	router.POST("/addData", addData)
	router.POST("/addDataset", addDataset)

	router.GET("/fetchDataset", fetchDataset)
	router.GET("/fetchData", fetchData)

	router.PATCH("/updateData", updateData)
	router.DELETE("/deleteData", deleteData)

	err := router.Run(url)
	if err != nil {
		return 
	}
}

func getConfigData() {
	// Setting viper
	helper.SetConfig("../../config" /*path*/, "config" /*name*/, "yaml" /*file extension*/)

	host = viper.GetString("server.host")
	port = viper.GetString("server.port")
	url = host + ":" + port
}

func addData(c *gin.Context) {
	var newData model.AnyData
	newData.Id = dataIds.Increment()
	newData.Date = helper.GetCurrentTime()

	if err := c.BindJSON(&newData); err != nil {
		return
	}

	localStorage.DataStorage = append(localStorage.DataStorage, newData)

	idToData[newData.Id] = newData
	c.IndentedJSON(http.StatusCreated, idToData[newData.Id])
}

func addDataset(c *gin.Context) {
	var newDataset []model.AnyData

	if err := c.BindJSON(&newDataset); err != nil {
		return
	}

	for _, newData := range newDataset {
		newData.Id = dataIds.Increment()
		newData.Date = helper.GetCurrentTime()

		localStorage.DataStorage = append(localStorage.DataStorage, newData)

		idToData[newData.Id] = newData
		c.IndentedJSON(http.StatusCreated, idToData[newData.Id])
	}

}

func fetchDataset(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, localStorage.DataStorage)
}

func fetchData(c *gin.Context) {
	strId, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query param."})
		return
	}

	id, _ := strconv.ParseUint(strId, 10, 64)
	_, index, err := getDataById(id)

	if err != nil {
		errorMessage := fmt.Sprintf("ID %v: the data is invalid.", id)
		c.IndentedJSON(
			http.StatusBadRequest, gin.H{"message": errorMessage})
		return
	}

	c.IndentedJSON(http.StatusOK, localStorage.DataStorage[index])
}

func updateData(c *gin.Context) {
	strId, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query param."})
		return
	}

	id, _ := strconv.ParseUint(strId, 10, 64)
	_, index, err := getDataById(id)

	if err != nil {
		errorMessage := fmt.Sprintf("ID %v: the data is invalid.", id)
		c.IndentedJSON(
			http.StatusBadRequest, gin.H{"message": errorMessage})
		return
	}

	localStorage.DataStorage[index].Amount += uint64(0.0716 * math.Pow(10, 18))
	idToData[id] = localStorage.DataStorage[index]
	c.IndentedJSON(http.StatusOK, localStorage.DataStorage[index])
}

func deleteData(c *gin.Context) {
	strId, _ := c.GetQuery("id")
	id, _ := strconv.ParseUint(strId, 10, 64)

	_, index, err := getDataById(id)
	if err != nil {
		errorMessage := fmt.Sprintf("ID %v: the data is invalid.", id)
		c.IndentedJSON(
			http.StatusBadRequest, gin.H{"message": errorMessage})
		return
	}

	localStorage.DataStorage = append(localStorage.DataStorage[:index], localStorage.DataStorage[index+1:]...)

	message := fmt.Sprintf("ID %v: deleted.\n", id)
	c.IndentedJSON(http.StatusOK, message)
}

func getDataById(id uint64) (model.AnyData, uint64, error) {
	dataIndex, err := helper.IndexOf(id)

	isValid := localStorage.DataStorage[dataIndex] == idToData[id]
	if isValid == true {
		return localStorage.DataStorage[dataIndex], dataIndex, err
	} else {
		return localStorage.DataStorage[dataIndex], 0, errors.New("-1")
	}
}
