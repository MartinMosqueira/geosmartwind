package router

import (
	"github.com/MartinMosqueira/geosmartwind/internal/services/user"
	"github.com/gorilla/mux"
)

func RuterLink(rout *mux.Router) {
	user.Rutas(rout)
}
