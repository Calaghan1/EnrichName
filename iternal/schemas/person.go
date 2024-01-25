package schemas

import "github.com/google/uuid"

type Person struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key" json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Paronymic string `json:"patronymic"`
	Age uint8 `json:"age"`
	Gender string `json:"gender"`
	Nationality string `json:"nationality"`
}