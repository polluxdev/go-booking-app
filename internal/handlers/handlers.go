package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/polluxdev/go-booking-app/internal/config"
	"github.com/polluxdev/go-booking-app/internal/forms"
	"github.com/polluxdev/go-booking-app/internal/helpers"
	"github.com/polluxdev/go-booking-app/internal/models"
	"github.com/polluxdev/go-booking-app/internal/render"
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
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}

func (m *Repository) Search(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search.page.html", &models.TemplateData{})
}

func (m *Repository) PostSearch(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	w.Write([]byte(start))
}

type jsonData struct {
	Status  bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) PostSearchJSON(w http.ResponseWriter, r *http.Request) {
	res := jsonData{
		Status:  true,
		Message: "JSON Data was here",
	}

	jsonRes, err := json.MarshalIndent(res, "", "     ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func (m *Repository) General(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "general.page.html", &models.TemplateData{})
}

func (m *Repository) Major(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "major.page.html", &models.TemplateData{})
}

func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation

	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

func (m *Repository) PostMakeReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	err = errors.New("this is an error message")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	reservation := models.Reservation{
		FirstName:   r.Form.Get("first_name"),
		LastName:    r.Form.Get("last_name"),
		Email:       r.Form.Get("email"),
		PhoneNumber: r.Form.Get("phone_number"),
	}

	form := forms.New(r.PostForm)
	form.Required("first_name", "last_name", "email", "phone_number")
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("Can't get error from session")
		m.App.Session.Put(r.Context(), "error", "cannot get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

		return
	}

	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "reservation-summary.page.html", &models.TemplateData{
		Data: data,
	})
}
