package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fernandomocrosky/mytemplates/pkg/config"
	"github.com/fernandomocrosky/mytemplates/pkg/handlers"
	"github.com/fernandomocrosky/mytemplates/pkg/render"
)

const portNumber = ":8000"

func main() {
	
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Listen on port %s\n", portNumber[1:])
	http.ListenAndServe(portNumber, nil)
}