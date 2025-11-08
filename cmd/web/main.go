package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/leonardodk/wedding_site/pkg/config"
	"github.com/leonardodk/wedding_site/pkg/handlers"
	"github.com/leonardodk/wedding_site/pkg/render"
)

const portNumber = ":8080"

// Create an app config variable
var app = config.AppConfig{}

// make a session manager
var MainSession *scs.SessionManager

func main() {

	// change to true when in production
	app.InProduction = false

	MainSession = scs.New()                            // start a new session
	MainSession.Lifetime = 24 * time.Hour              // let it last for a day
	MainSession.Cookie.Persist = true                  // remember the personn after the browser is closed
	MainSession.Cookie.SameSite = http.SameSiteLaxMode // set samesite mode
	MainSession.Cookie.Secure = app.InProduction       // set production mode
	app.Session = MainSession                          // passs session manager to config variable

	// create a cache of templates, and check for errors
	tc, err := render.CreateTemplateCache()

	// check for errors
	if err != nil {
		log.Fatal("couldn't load template cache")
	}

	// give template cache to app config variable
	app.TemplateCache = tc
	app.UseCache = false

	// create a new Repository variable (just has a pointer to app in it atm but there will be more)
	repo := handlers.NewRepo(&app)

	// pass the repo variable back to the handlers package
	handlers.SetPackageRepo(repo)

	// gives the address of the app variable to the app variable in render package
	// the render package now has access to the cache of templates
	render.RenderSetApp(&app)

	var msg2go string = fmt.Sprintf("running on port: %s", portNumber)
	fmt.Println(msg2go)

	// make a new server variable
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app), // routes() returns a *pat.PatternServeMux
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
