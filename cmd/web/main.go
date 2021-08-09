package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/kmulqueen/gobnb/internal/config"
	"github.com/kmulqueen/gobnb/internal/handlers"
	"github.com/kmulqueen/gobnb/internal/render"
	"github.com/kmulqueen/gobnb/models"
)

// Variables available to entire main package
const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting app on port %s", port)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	// What to store in session
	gob.Register(models.Reservation{})

	// Change to true when in production
	app.InProduction = false

	// Initialize a new session & set the lifetime to 24 hours
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true // Should session persist after browser window closed?
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // Ensure that cookie is encrypted & https connection. Prod = true, dev/localhost is not an encrypted connection

	// Set session in config to session we just initialized
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can't create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	return nil
}
