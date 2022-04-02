package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"phonebook/database"
	"phonebook/entity"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	var phonebook []entity.Phonebook
	database.Connector.Find(&phonebook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(phonebook)
}

func DeletPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println(key, vars)

	var phonebook entity.Phonebook
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&phonebook)
	w.WriteHeader(http.StatusNoContent)
}
