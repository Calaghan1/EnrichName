package helpers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Calaghan1/EnrichName/iternal/schemas"
)



func CheckErrorFatal(err error, msg string, msg1 string) {
	if err != nil {
		log.Fatal(msg,"Error:", err)
	} else {
		log.Print(msg1)
	}
}

type age_resp struct {
	Count int64 `json:"count"`
	Name string	`json:"name"`
	Age uint8	`json:"age"`
}
type gen_resp struct {
	Count int64	`json:"count"`
	Name string	`json:"name"`
	Gender string	`json:"gender"`
	Probability float64	`json:"probability"`
}

type national_resp struct {
	Count int64	`json:"count"`
	Name string	`json:"name"`
	Country []Country `json:"country"`
}

type Country struct {
	Id string	`json:"country_id"`
	Probability float64	`json:"probability"`
}


func Enrase_data(person schemas.Person) schemas.Person {
	resp, err := http.Get("https://api.agify.io/?name=" + person.Name)
	CheckErrorFatal(err, "", "")
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %d", resp.StatusCode)
	}
	var p age_resp
	var gen_r gen_resp
	var nat_r national_resp
	err = json.NewDecoder(resp.Body).Decode(&p)
	CheckErrorFatal(err, "", "")

	person.Age = p.Age
	resp, err = http.Get("https://api.genderize.io/?name=" + person.Name)
	CheckErrorFatal(err, "", "")
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %d", resp.StatusCode)
	}
	err = json.NewDecoder(resp.Body).Decode(&gen_r)
	CheckErrorFatal(err, "", "")

	person.Gender = gen_r.Gender

	resp, err = http.Get("https://api.nationalize.io/?name=" + person.Name)
	CheckErrorFatal(err, "", "")
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %d", resp.StatusCode)
	}
	err = json.NewDecoder(resp.Body).Decode(&nat_r)
	CheckErrorFatal(err, "", "")
	buff := nat_r.Country[0]
	for i := 1; i < len(nat_r.Country); i++ {
		if nat_r.Country[i].Probability > buff.Probability {
			buff = nat_r.Country[i]
		}
	}
	log.Println("Enrase data complite")
	person.Nationality = buff.Id
	return person
}