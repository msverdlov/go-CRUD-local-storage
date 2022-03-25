package helper

import (
	"anyData/internal/app/anyData/localStorage"
	"errors"
	"fmt"
	"time"
)

func IndexOf(id uint64) (uint64, error) {
	for index, data := range localStorage.DataStorage {
		if data.Id == id {
			return uint64(index), nil
		}
	}
	errorMessage := fmt.Sprintf("ID %v: data not found.", id)

	return 0, errors.New(errorMessage)
}

func GetCurrentTime() string {
	loc, _ := time.LoadLocation("UTC")
	now := (time.Now().In(loc)).String()

	return now
}
