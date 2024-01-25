package database

import (
	"os"

	"github.com/Calaghan1/EnrichName/helpers"
	"github.com/Calaghan1/EnrichName/iternal/schemas"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type DB struct {
	Client gorm.DB
}

func (d* DB)Init_database() {
	err := godotenv.Load()
	helpers.CheckErrorFatal(err, "Failed to load accsess data to database", "Getting acces data")
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTRGRES_CONNECTION")), &gorm.Config{})
	helpers.CheckErrorFatal(err, "Failed to connect to databse", "Connect to database")
	err = db.AutoMigrate(&schemas.Person{})
	helpers.CheckErrorFatal(err, "Failed to process migrations", "Migrations complite")
	d.Client = *db
}

func (d *DB)Create_person(person schemas.Person) schemas.Person {
	res := d.Client.Create(&person)
	helpers.CheckErrorFatal(res.Error, "Failed to create person", "Created a person in database")
	return person
} 


func (d *DB)Show_all(page, pageSize int, filters map[string]string) []schemas.Person{
	var persons []schemas.Person
	offset := (page - 1) * pageSize
	
	query := d.Client
	for key, value := range filters {
		query = *query.Where(key, value)
	}

	res := query.Limit(pageSize).Offset(offset).Find(&persons)
	helpers.CheckErrorFatal(res.Error, "Failed to find persons", "Find all persons in database")
	return persons
}

func (d *DB)DeleteByID(id string) int64{
	res := d.Client.Where("id = ?", id).Delete(&schemas.Person{})
	helpers.CheckErrorFatal(res.Error, "Failed to delete person", "Person with id: " + id + "deleted")
	return	res.RowsAffected
}

func (d *DB) FindByID(id string) (schemas.Person, error) {
    var person schemas.Person
    res := d.Client.Where("id = ?", id).First(&person)
    helpers.CheckErrorFatal(res.Error, "Failed to find person with id:" + id, "Person with id: "+ id + "find")
    return person, nil
}

func (d *DB)UpdateById(id string, NewData schemas.Person) (schemas.Person, error) {
	d.FindByID(id)
	id_new, err := uuid.Parse(id)
	helpers.CheckErrorFatal(err, "", "")
	NewData.ID = id_new
	res := d.Client.Save(&NewData)
	helpers.CheckErrorFatal(res.Error, "Failed to update person with id" + id, "Person with id: "+ id + "updated")
	return NewData, nil
}
	