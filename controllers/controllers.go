package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/felipeazsantos/go-rest-api/database"
	"github.com/felipeazsantos/go-rest-api/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func TodasPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades)
	json.NewEncoder(w).Encode(personalidades)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id, err := strconv.Atoi(parametros["id"])
	if err != nil {
		json.NewEncoder(w).Encode(errors.New("Argumento inválido!"))
		return
	}

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	database.DB.Create(&novaPersonalidade)
	json.NewEncoder(w).Encode(novaPersonalidade)
}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id, err := strconv.Atoi(parametros["id"])
	if err != nil {
		json.NewEncoder(w).Encode(errors.New("Erro ao ler o parametro id"))
		return
	}

	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id, err := strconv.Atoi(parametros["id"])
	if err != nil {
		json.NewEncoder(w).Encode(errors.New("Erro ao ler o parametro id"))
		return
	}

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)

	json.NewEncoder(w).Encode(personalidade)
}
