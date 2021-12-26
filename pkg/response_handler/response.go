package response_handler

import (
	"encoding/json"
	"net/http"

	err_d "github.com/sderohan/getir-go-assignment/app/error"
)

const (
	// CharsetUTF8 utf8 character set
	CharsetUTF8 = "charset=utf-8"
)

// Media types
const (
	ApplicationJSON            = "application/json"
	ApplicationJSONCharsetUTF8 = ApplicationJSON + "; " + CharsetUTF8
)

// Headers
const (
	ContentType = "Content-Type"
)

type CollectionResponse struct {
	Code    uint        `json:"code"`
	Message string      `json:"msg"`
	Records interface{} `json:"records"`
}

type IMData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func _json(w http.ResponseWriter, code int, data []byte) {
	w.Header().Set(ContentType, ApplicationJSONCharsetUTF8)
	w.WriteHeader(code)
	w.Write(data)
}

func SendResponse(w http.ResponseWriter, code int, data interface{}) {
	marshalled_data, err := json.Marshal(data)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err_d.ErrMarshalData.Error())
	}
	_json(w, code, marshalled_data)
}

func SendErrorResponse(w http.ResponseWriter, code int, msg string) {
	err := struct {
		Message string `json:"msg"`
	}{
		Message: msg,
	}
	response, _ := json.Marshal(err)
	_json(w, code, response)
}

func SendEcho(w http.ResponseWriter, code int) {
	echo := struct {
		Status int `json:"status"`
	}{
		Status: code,
	}
	response, _ := json.Marshal(echo)
	_json(w, code, response)
}
