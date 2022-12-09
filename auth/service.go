package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	GenerateToken(userID int) (string, error) //generate token dengan memasukkan userID
	ValidateToken(token string) (*jwt.Token, error)
}
type jwtService struct {
}
var SECRET_KEY = []byte("test_berhas1l")
// membuat new service
func NewService() *jwtService {
	return &jwtService{}
}

// fungsi untuk men-generate token
func (s *jwtService) GenerateToken(userID int) (string, error) {
	// membuat objek payload data / claims
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	// generate new token dengan menyisipkan data user id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// bubuhkan secret_key ke dalam token untuk memverifikasi
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	tokenValid, err := jwt.Parse(token, func (token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return tokenValid, err
	}

	return tokenValid, nil
}