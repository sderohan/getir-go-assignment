package controller

import (
	"fmt"
	"net/http"

	"github.com/sderohan/getir-go-assignment/app/models"
)

type Controller struct {
	*models.DBWriter
}

func (c *Controller) FetchData(w http.ResponseWriter, req *http.Request) {
	// c.ReadCollections()
	fmt.Fprintf(w, "hello from fetchdata")
}

func (c *Controller) EchoAPI(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello from echoapi")
}

func (c *Controller) PrintParams(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello from printparams")
}
