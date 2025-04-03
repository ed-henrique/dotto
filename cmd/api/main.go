package main

import (
	"embed"

	"github.com/ed-henrique/dotto/internal/models"
	"github.com/ed-henrique/dotto/internal/server"
)

//go:embed views/*
var views embed.FS

func main() {
	s := server.New(views, models.ToDos)
	s.Routes()
	s.Run()
}
