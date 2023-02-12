package handlers

import (
	"net/http"

	"github.com/Hughost0/awa-dentistaa/pkg/config"
	"github.com/Hughost0/awa-dentistaa/pkg/models"
	"github.com/Hughost0/awa-dentistaa/pkg/render"
)

// TemplateData holds data sent from handlers to templates

// Repo theadwadawd
var Repo *Repository

// Repository is the  repository type
type Repository struct {
	App *config.Appconfig
}

// NewRepo Create new repository
func NewRepo(a *config.Appconfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the new handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Index is the home page handler
func (m *Repository) Index(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About us the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	// send some data
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
