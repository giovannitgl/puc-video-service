package manager

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/giovannitgl/video-services/user-service/internal/config"
	"github.com/giovannitgl/video-services/user-service/internal/entities"
	"github.com/giovannitgl/video-services/user-service/internal/service"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/pbkdf2"
	"time"
)

func RegisterNewUser(user *entities.User) error {
	user.Password = HashPassword(user.Password)
	return service.UserCreate(user)
}

func LoginUser(email, password string) (string, error) {
	password = HashPassword(password)
	usr := service.UserGetLogin(email, password)
	if usr == nil {
		return "", errors.New("Invalid credentials")
	}

	return GenerateJWT(usr.ID)
}

func HashPassword(pass string) string {
	hashed := pbkdf2.Key([]byte(pass), []byte(config.Salt()), 10000, 50, sha256.New)
	return hex.EncodeToString(hashed)
}

func GenerateJWT(userId uint) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["user_id"] = userId
	return token.SignedString(config.SigningKey())

}
