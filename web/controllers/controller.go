package controllers

import (
	"context"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	user "github.com/jaredwarren/rx/user-service/proto/user"
)

// Controller implements the home resource.
type Controller struct {
	Mux *mux.Router
	// service *app.Service
	client user.UserServiceClient
}

// Register ...
func Register(mux *mux.Router, client user.UserServiceClient) *Controller {
	c := &Controller{
		Mux: mux,
		// service: service,
		client: client,
	}
	mux.HandleFunc("/", c.Home).Methods("GET")
	// mux.HandleFunc("/host", c.HostHandler).Methods("POST")

	// service.Register(c)
	return c
}

// Home ...
func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("~ Home ~")

	var u user.User
	ur, _ := c.client.Get(context.Background(), &u)
	fmt.Printf("User:%+v\n", u)
	fmt.Printf("UR:%+v\n", ur)

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
