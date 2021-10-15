package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

type customClaims struct {
	jwt.StandardClaims
}

func Generate(id string, expires int64) string{
	claims := customClaims{
		StandardClaims: jwt.StandardClaims{
			Id: id,
			ExpiresAt: expires,
			Issuer:    "localhost",
			Audience: "localhost",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secureSecretText"))
	if err != nil {fmt.Println("Error signing token")}
	return signedToken
}

func Validate(token string) (uuid.UUID, error){
	tokend, err := jwt.ParseWithClaims(
		token,
		&customClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secureSecretText"), nil
		},
	)
	if err != nil {return uuid.Nil, err}
	validToken := tokend.Claims.(*customClaims)

	//if !validToken.Valid {return uuid.Nil, errors.New("invalid token")}
	id, _ := uuid.FromString(validToken.Id)
	return id, nil
}