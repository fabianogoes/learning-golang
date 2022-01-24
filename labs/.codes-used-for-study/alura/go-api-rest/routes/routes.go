package routes

import (
	"log"
	"net/http"

	"github.com/fabianogoes/learning-golang/alura/go-api-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornarUmaPersonalidade).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
