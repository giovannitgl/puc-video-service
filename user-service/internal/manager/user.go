package manager

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/giovannitgl/video-services/user-service/internal/config"
	"github.com/giovannitgl/video-services/user-service/internal/entities"
	"github.com/giovannitgl/video-services/user-service/internal/service"
	"golang.org/x/crypto/pbkdf2"
)

func RegisterNewUser(user *entities.User) error {
	user.Password = HashPassword(user.Password)
	return service.UserCreate(user)
}

func LoginUser(email, password string) error {
	password = HashPassword(password)
	usr := service.UserGetLogin(email, password)
	if usr == nil {
		return errors.New("Invalid credentials")
	}
	return nil
}

func HashPassword(pass string) string {
	hashed := pbkdf2.Key([]byte(pass), []byte(config.Salt()), 10000, 50, sha256.New)
	return hex.EncodeToString(hashed)
}
