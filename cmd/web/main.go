package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/awa-nashe/paginadepapa/pkg/config"
	"github.com/awa-nashe/paginadepapa/pkg/handlers"
	"github.com/awa-nashe/paginadepapa/pkg/render"
)

const portnumber = ":8080"

var app config.Appconfig
var session *scs.SessionManager

// main is the main application function
func main() {
	//change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create tempalte cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)

	fmt.Println(fmt.Sprintf("Starting application on port: %s", portnumber))

	srv := &http.Server{
		Addr:    portnumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
