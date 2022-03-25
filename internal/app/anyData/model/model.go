package model

type (
	AnyData struct {
		Id         uint64 `json:"id"`
		Address    string `json:"address"`
		Date       string `json:"date"`
		Amount     uint64 `json:"amount"`
		Data       string `json:"data"`
	}
)
