package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tatyanayagolnikov/go-website/pkg/config"
	"github.com/tatyanayagolnikov/go-website/pkg/handlers"
	"github.com/tatyanayagolnikov/go-website/pkg/render"
)
const portNumber = ":8080"

// main - is the main application function
func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false 

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
		
	fmt.Println("starting application on port:", portNumber)
	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber),"using fmt.Sprintln")

	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)


	

}