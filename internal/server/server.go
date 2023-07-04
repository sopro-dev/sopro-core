package server

import "github.com/pablodz/sopro/internal/api"

func Serve() {
	api.HandleRequest()
}
