package jwt

import (
	"food-delivery/components/tokenprovider"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtProvider struct {
	secret string
}

func NewTokenJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	// Generate jwt token
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	// return token
	return &tokenprovider.Token{
		Token:   myToken,
		Expiry:  expiry,
		Created: time.Now(),
	}, nil
}

func (j *jwtProvider) Validate(token string) (*tokenprovider.TokenPayload, error) {
	return &tokenprovider.TokenPayload{
		0,
		"",
	}, nil
}
