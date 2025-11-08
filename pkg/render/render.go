package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/leonardodk/HotelDK/pkg/config"
	"github.com/leonardodk/HotelDK/pkg/models"
)

var app *config.AppConfig

// sets the address for app to the address of the config struct made in main.go
// now we're all looking at the same data
func RenderSetApp(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// renders the templates, must be done each time the web page is loaded
// args are the writer and request, plus the struct we use to send the template data from the handlers
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	// if we're using the cache of templates in the app config
	if app.UseCache {

		// get the cache of templates fromn config struct made in main.go
		tc = app.TemplateCache
	} else {
		//else make it anew
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]

	// if there isn't a template in teh cache under that name then exit?
	if !ok {
		log.Fatal("no template in cache under taht name")
		return
	}

	// make a space to load the ececuted template into
	// we do this because if there's an issue executing the template we don't want half of it sent to the
	// client as they'd get half a web page
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)

	// if there's no errorrs in the completely executed web page then continue
	if err != nil {
		log.Println(err)
	}

	// buffers are also counted as writers so we're passing the contents of the buffer (quanrantine zone)
	// to the response writer (send it to the client, finally)
	_, err = buf.WriteTo(w)

	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	fmt.Println("making the template cache")
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all files named *.page.tmpl first
	// returns a slice of strings

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	// look for layout template files, and store them in a slice of strings.
	// we're doing this now so we don't have to do it every time later on.
	matches, err := filepath.Glob("./templates/*.layout.tmpl")

	if err != nil {
		return myCache, err
	}

	// range through templates in the slice of strings returned
	// by Glob and get their names

	for _, page := range pages {

		// get the file name for the page
		name := filepath.Base(page)

		// create a new template set. All templates can be template sets, you can essentially hold multiple templates
		// in one template. A bit like a bag that has other bags inside it.

		// Give that template set a name, and then parse the page template into it
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// if there are layout files then parse them all and shove them into the template set
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
		}

		if err != nil {
			return myCache, err
		}

		// add the template set (rememeber it's basically just one template) to the cache
		myCache[name] = ts
	}

	return myCache, nil
}
