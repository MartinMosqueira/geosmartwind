package api

import (
	"fmt"
	"github.com/MartinMosqueira/geosmartwind/internal/router"
	"github.com/MartinMosqueira/geosmartwind/pkg/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func ConfigServer(port string, router *mux.Router) *http.Server {
	return &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func StartAPI() {
	r := mux.NewRouter()
	//link routes
	router.RuterLink(r)
	http.Handle("/", r)
	server := ConfigServer(config.PORT, r)
	fmt.Printf("Server listen http://localhost:%s\n", config.PORT)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error connect server")
	}
}
