package helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

type (
	ParamsGenerateJWT struct {
		SecretKey string
		UserID    string
		RoleID    string
	}
	ParamsValidateJWT struct {
		Token     string
		SecretKey string
	}
	Claims struct {
		jwt.StandardClaims
		UserID string `json:"user_id,omitempty"`
		RoleID string `json:"role_id,omitempty"`
	}
)

const JWT_EXPIRED_IN_MINUTES = 480

func GenerateJWT(p *ParamsGenerateJWT) (string, error) {
	expiredAt := time.Now().Add(time.Duration(JWT_EXPIRED_IN_MINUTES) * time.Minute).Unix()
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
		},
		UserID: p.UserID,
		RoleID: p.RoleID,
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := token.SignedString([]byte(p.SecretKey))

	return signedToken, err
}

func ValidateJWT(p *ParamsValidateJWT) (jwt.MapClaims, error) {
	token, err := jwt.Parse(p.Token, func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok || method != JWT_SIGNING_METHOD {
			return nil, errors.New("token tidak valid")
		}

		return []byte(p.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
