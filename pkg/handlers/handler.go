package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Calaghan1/EnrichName/helpers"
	"github.com/Calaghan1/EnrichName/iternal/database"
	"github.com/Calaghan1/EnrichName/iternal/schemas"
	"github.com/Calaghan1/EnrichName/settings"
)

type Person_handlers struct {
	Db *database.DB
}
func (p *Person_handlers)Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	var person schemas.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	helpers.CheckErrorFatal(err, "Failed to decode json", "")

	person = helpers.Enrase_data(person)
	res := p.Db.Create_person(person)
	json_data, err := json.Marshal(res)
	helpers.CheckErrorFatal(err, "", "")
	w.Write(json_data)
}

func (p *Person_handlers)Show_all(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return		
	}
	filter := make(map[string]string)
	if agefilter := r.URL.Query().Get("age"); agefilter != "" {
		filter["age"] = agefilter
	}
	if namefilter := r.URL.Query().Get("name"); namefilter != "" {
		filter["name"] = namefilter
	}
	if surnamefilter := r.URL.Query().Get("surname"); surnamefilter != "" {
		filter["age"] = surnamefilter
	}
	if genderfilter := r.URL.Query().Get("gender"); genderfilter != "" {
		filter["gender"] = genderfilter
	}
	if nationalityfilter := r.URL.Query().Get("nationality"); nationalityfilter != "" {
		filter["age"] = nationalityfilter
	}
	if settings.LogLevel == "debug" {
		log.Println("Filters:\n", filter)
	}
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	page, err := strconv.Atoi(pageStr)
	helpers.CheckErrorFatal(err, "", "")
	limit, err := strconv.Atoi(limitStr)
	helpers.CheckErrorFatal(err, "", "")


	data := p.Db.Show_all(page, limit, filter)
	log.Print(data)
	json_data, err := json.Marshal(data)
	helpers.CheckErrorFatal(err, "Failed to Marshal json", "")
	w.Write(json_data)
}

func (p *Person_handlers)DeleteByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return		
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	res := p.Db.DeleteByID(id)
	if res == 0 {
		http.Error(w, "No record found to delete", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Record was deleted"))
}

func (p *Person_handlers)UpdateById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	var person schemas.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	helpers.CheckErrorFatal(err, "Failed to decode json", "")
	id := r.URL.Query().Get("id")
	res, err := p.Db.UpdateById(id, person)
	helpers.CheckErrorFatal(err, "", "")
	data, err := json.Marshal(res)
	helpers.CheckErrorFatal(err, "", "")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}