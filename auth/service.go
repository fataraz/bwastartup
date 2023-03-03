package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type JWTService struct {
}

func NewService() *JWTService {
	return &JWTService{}
}

var SECRET_KEY = []byte("BWASTARTUP_s3cr3T_k3Y")

func (s *JWTService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signetedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signetedToken, err
	}

	return signetedToken, nil
}
