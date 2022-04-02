package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"phonebook/database"
	"phonebook/entity"
	"strconv"

	"github.com/gorilla/mux"
)

//get all person data
func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	var phonebook []entity.Phonebook
	database.Connector.Find(&phonebook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(phonebook)
}

//delete specific person by ID
func DeletPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println(key, vars)

	var phonebook entity.Phonebook
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&phonebook)
	w.WriteHeader(http.StatusNoContent)
}

//create new person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var phonebook entity.Phonebook
	json.Unmarshal(requestBody, &phonebook)

	database.Connector.Create(phonebook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(phonebook)
}

//update person by ID
func UpdatePersonByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var phonebook entity.Phonebook
	json.Unmarshal(requestBody, &phonebook)
	database.Connector.Save(&phonebook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(phonebook)
}
