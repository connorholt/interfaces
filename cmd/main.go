package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/connorholt/interfaces/internal/handler/add"
	"github.com/connorholt/interfaces/internal/handler/list"
	"github.com/connorholt/interfaces/internal/handler/upload/audio"
	"github.com/connorholt/interfaces/internal/storage"
)

func main() {

	s := storage.NewInMemory()
	handlerList := list.NewList(s)
	handlerCreate := add.NewAdd(s)
	handlerAudio := audio.NewUpload(s)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handlerList.GetAll)
	r.Post("/add", handlerCreate.Create)
	r.Post("/addPicture", handlerAudio.Create)

	http.ListenAndServe(":3000", r)
}
