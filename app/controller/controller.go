package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	err_d "github.com/sderohan/getir-go-assignment/app/error"
	"github.com/sderohan/getir-go-assignment/app/models"
	"github.com/sderohan/getir-go-assignment/pkg/response_handler"
)

type Controller struct {
	*models.DBWriter
}

func (c *Controller) FetchData(w http.ResponseWriter, req *http.Request) {
	// c.ReadCollections()
	fmt.Fprintf(w, "hello from fetchdata")
}

func (c *Controller) EchoAPI(w http.ResponseWriter, req *http.Request) {
	request := response_handler.IMData{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&request)
	if err != nil {
		response_handler.SendErrorResponse(w, http.StatusInternalServerError, err_d.ErrMarshalData.Error())
	}
	if request.Key != "" && request.Value != "" {
		response_handler.SendEcho(w, http.StatusOK)
		return
	}
	response_handler.SendEcho(w, http.StatusBadRequest)
}

func (c *Controller) PrintParams(w http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()
	if len(queryParams) == 1 {
		response := &response_handler.IMData{}
		for key, value := range queryParams {
			response.Key = key
			if len(value) == 1 && value[0] != "" {
				response.Value = value[0]
			} else {
				response_handler.SendErrorResponse(w, http.StatusBadRequest, err_d.ErrInvalid_QUERY_PARAMS.Error())
				return
			}
		}
		response_handler.SendResponse(w, http.StatusOK, response)
		return
	}
	response_handler.SendErrorResponse(w, http.StatusBadRequest, err_d.ErrInvalid_QUERY_PARAMS.Error())
}
