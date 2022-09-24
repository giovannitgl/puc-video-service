package entities

import (
	"encoding/json"
	"gorm.io/gorm"
	"strings"
	"time"
)

type JsonBirthDate time.Time

func (j *JsonBirthDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonBirthDate(t)
	return nil
}

func (j JsonBirthDate) MarshalJSON() ([]byte, error) {
	dateString := time.Time(j).Format("2006-01-02")
	return json.Marshal(dateString)
}

type User struct {
	gorm.Model
	Username  string        `json:"username" gorm:"unique"`
	FirstName string        `json:"first_name" validate:"required"`
	LastName  string        `json:"last_name" validate:"required"`
	Email     string        `json:"email" gorm:"unique;" validate:"required,email"`
	BirthDate JsonBirthDate `json:"birth_date" validate:"required"`
	Password  string        `json:"password" gorm:"type:text;"`
}
