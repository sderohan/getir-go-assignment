package routes

import (
	"net/http"

	"github.com/sderohan/getir-go-assignment/app/controller"
)

type Router struct {
	*controller.Controller
}

func (r *Router) InitRoutes() http.Handler {
	rts := http.NewServeMux()
	rts.HandleFunc("/fetch", r.FetchData)
	rts.HandleFunc("/echo-endpoint", r.EchoAPI)
	rts.HandleFunc("/print-endpoint", r.PrintParams)
	return rts
}
