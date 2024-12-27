package common

import "net/http"

type HandlerInput struct {
	Url     string
	Handler func(w http.ResponseWriter, r *http.Request)
}
