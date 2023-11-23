package handlers

import "github.com/julienschmidt/httprouter"

type Handler interface{
	Registr(router *httprouter.Router)
}