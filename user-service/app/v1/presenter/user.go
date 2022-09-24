package presenter

import (
	"github.com/giovannitgl/video-services/user-service/internal/entities"
)

type User struct {
	ID        uint                   `json:"id"`
	FirstName string                 `json:"first_name"`
	LastName  string                 `json:"last_name"`
	Email     string                 `json:"email"`
	BirthDate entities.JsonBirthDate `json:"birth_date"`
}

func UserResponse(user entities.User) *User {
	return &User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		BirthDate: user.BirthDate,
	}
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
