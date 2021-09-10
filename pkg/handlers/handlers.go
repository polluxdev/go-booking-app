package handlers

import (
	"net/http"

	"github.com/polluxdev/go-booking-app/pkg/config"
	"github.com/polluxdev/go-booking-app/pkg/models"
	"github.com/polluxdev/go-booking-app/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}
