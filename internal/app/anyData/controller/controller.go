package controller

import (
	"anyData/internal/app/anyData/helper"
	"anyData/internal/app/anyData/localStorage"
	"anyData/internal/app/anyData/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

var (
	dataIds    helper.Counter
	dataBuffer = make(map[uint64]model.AnyDataStruct)
)

func AddData(c *gin.Context) {
	var newData model.AnyDataStruct
	newData.Id = dataIds.Increment()
	newData.Date = helper.GetCurrentTime()

	if err := c.BindJSON(&newData); err != nil {
		return
	}

	localStorage.DataStorage = append(localStorage.DataStorage, newData)

	dataBuffer[newData.Id] = newData
	c.IndentedJSON(http.StatusCreated, dataBuffer[newData.Id])
}

func AddDataset(c *gin.Context) {
	var newDataset []model.AnyDataStruct

	if err := c.BindJSON(&newDataset); err != nil {
		return
	}

	for _, newData := range newDataset {
		newData.Id = dataIds.Increment()
		newData.Date = helper.GetCurrentTime()

		localStorage.DataStorage = append(localStorage.DataStorage, newData)

		dataBuffer[newData.Id] = newData
		c.IndentedJSON(http.StatusCreated, dataBuffer[newData.Id])
	}

}

func FetchDataset(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, localStorage.DataStorage)
}

func FetchData(c *gin.Context) {
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

func UpdateData(c *gin.Context) {
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
	dataBuffer[id] = localStorage.DataStorage[index]
	c.IndentedJSON(http.StatusOK, localStorage.DataStorage[index])
}

func DeleteData(c *gin.Context) {
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

func getDataById(id uint64) (model.AnyDataStruct, uint64, error) {
	dataIndex, err := helper.IndexOf(id)

	isValid := localStorage.DataStorage[dataIndex] == dataBuffer[id]
	if isValid == true {
		return localStorage.DataStorage[dataIndex], dataIndex, err
	} else {
		return localStorage.DataStorage[dataIndex], 0, errors.New("-1")
	}
}
