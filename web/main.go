package main

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/gorilla/mux"
)

func main() {

	// // load config from file
	// if conf == nil {
	// 	resourceDir := getResourceDir()
	// 	// setting config and payt is redundant if already set in web service
	// 	viper.SetConfigName("config_" + runtime.GOOS)
	// 	viper.AddConfigPath(resourceDir)
	// 	if err := viper.ReadInConfig(); err != nil {
	// 		log.Fatalf("Error reading config file, %s", err)
	// 	}
	// }

	// // load web service config
	// var (
	// 	serverConfig *WebConfig
	// )
	// {
	// 	if conf == nil {
	// 		viper.UnmarshalKey("server", &serverConfig)
	// 	} else {
	// 		serverConfig = conf
	// 	}
	// }

	// TODO: validate values
	// var addr string
	// if serverConfig.Host == "" {
	// 	if serverConfig.Port > 0 {
	// 		// port given, but host isn't assume local host
	// 		serverConfig.Host = "127.0.0.1"
	// 		addr = fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
	// 	}
	// } else if serverConfig.Port <= 0 {
	// 	addr = serverConfig.Host
	// } else {
	// 	addr = fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
	// }
	// app := &Service{
	// 	Name:   serverConfig.Name,
	// 	Exit:   make(chan error),
	// 	Config: conf,
	// }
	// if addr != "" {
	// 	app.Home, _ = url.Parse(fmt.Sprintf("http://%s", addr))
	// }
	exit := make(chan error)

	// Interrupt handler (ctrl-c)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		done := <-signalChan
		exit <- fmt.Errorf("%s", done)
	}()

	// Create Router and add default paths
	mux := mux.NewRouter()
	mux.HandleFunc("/static/{filename:[a-zA-Z0-9\\.\\-\\_\\/]*}", FileServer)
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		// brew install imagemagick
		// convert -background none static/db.svg -define icon:auto-resize static/favicon.ico
		// if !fileExists("static/favicon.ico") {
		// 	// TODO: spit out generic ico
		// 	fmt.Println("no favicon.ico found")
		// }
		w.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext("static/favicon.ico")))
		http.ServeFile(w, r, "static/favicon.ico")
	})
	mux.HandleFunc("/health-check", HealthCheck).Methods("GET", "HEAD")

	// Start Server
	server := &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: mux,
	}
	go func() {
		// TODO: add https, stuff...
		fmt.Printf("HTTP server listening on %q\n", server.Addr)
		exit <- server.ListenAndServe()
	}()

	// Wait to finish
	d := <-exit

	fmt.Println("DONE!!!!", d)

}

// FileServer serves a file with mime type header
func FileServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file := vars["filename"]
	w.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(file)))
	http.ServeFile(w, r, "./static/"+file)
}

// HealthCheck return ok
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
