package model

type WebResponse struct {
	Code   uint        `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
