package controllers

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

// Controller implements the home resource.
type Controller struct {
	Mux *mux.Router
	// service *app.Service
}

// Register ...
func Register(mux *mux.Router) *Controller {
	c := &Controller{
		Mux: mux,
		// service: service,
	}
	mux.HandleFunc("/", c.Home).Methods("GET")
	// mux.HandleFunc("/host", c.HostHandler).Methods("POST")

	// service.Register(c)
	return c
}

// Home ...
func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	// parse every time to make updates easier, and save memory
	tpl := template.Must(template.New("base").ParseFiles("templates/home.html", "templates/base.html"))
	tpl.ExecuteTemplate(w, "base", &struct {
		Title string
		// Hosts    map[string]*HostEntry
		// HostFile string
	}{
		Title: "Home",
		// Hosts:    hosts,
		// HostFile: hostFilePath,
	})
}
