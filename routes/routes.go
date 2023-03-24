package routes

import (
	"github.com/felipeazsantos/go-rest-api/controllers"
	"github.com/felipeazsantos/go-rest-api/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleWare)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidade).Methods(http.MethodGet)
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods(http.MethodGet)
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods(http.MethodPost)
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods(http.MethodDelete)
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
