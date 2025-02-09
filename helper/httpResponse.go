package helper

import (
	"encoding/json"
	"go-testing/constants"
	"net/http"
	"time"
)

type ResponsePublisher struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Code   int         `json:"code"`
}

type setResponse struct {
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
	Code       int         `json:"code"`
	AccessTime string      `json:"accessTime"`
}

func HttpResponseSuccess(w http.ResponseWriter, r *http.Request, data interface{}) {
	location, _ := time.LoadLocation(constants.TimeLocation)
	setResponse := setResponse{
		Status:     http.StatusText(200),
		AccessTime: time.Now().In(location).Format("02-01-2006 15:04:05"),
		Data:       data, Code: 200}
	response, _ := json.Marshal(setResponse)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(200)
	w.Write(response)
}

func HttpResponseError(w http.ResponseWriter, r *http.Request, data interface{}, code int) {
	setResponse := setResponse{
		Status:     http.StatusText(code),
		AccessTime: time.Now().Format("02-01-2006 15:04:05"),
		Data:       data,
		Code:       code}
	response, _ := json.Marshal(setResponse)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(code)
	w.Write(response)
}
