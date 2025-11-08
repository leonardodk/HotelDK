package handlers

import (
	"net/http"

	"github.com/leonardodk/HotelDK/pkg/config"
	"github.com/leonardodk/HotelDK/pkg/models"
	"github.com/leonardodk/HotelDK/pkg/render"
)

type Repository struct {
	App *config.AppConfig
	// apparantly more stuff will be added in here in the future
}

var PackageRepo *Repository

// returns a new repository
// important to note that this will be used in other packages, not here
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// SetPackageRepo sets the repisitory for the handlers
// this will be called in main but sets the repository variable for HERE
func SetPackageRepo(r *Repository) {
	PackageRepo = r
}

// handlers are now mehtods on the repsoitory type and not regular functions
// this is for reasons that escape me
// now would need to be called via handlers.PackageRepo.Home (use PackageRepo as
// that's the global var in this package)

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	// Store the remote IP in the session.
	// r.Context() tells SCS via cookies, who this session this data belongs to.
	remoteIP := r.RemoteAddr                              // get the remote IP of the client
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP) // store it in the session using the key remote_ip

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// Retrieve the remote IP from the current user’s session.
	// r.Context() checks to see in the cookies who we're talking about.
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip") // get the remote IP on of the session that was captured in Home

	// passing data by making a string map and giving it to the handler via the template data

	testStringMap := make(map[string]string)

	testStringMap["test"] = "hello, world!"
	testStringMap["remote_ip"] = remoteIP

	//send data to the tempalte?
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: testStringMap,
	})
}
