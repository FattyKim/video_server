package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter,h *http.Request,p httprouter.Params) {
	io.WriteString(w, "Create User Handler")
}
