package jwt

import (
	"OpenMarket/component/tokenprovider"
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
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		Payload: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}
	return &tokenprovider.Token{
		Token:   myToken,
		Expiry:  expiry,
		Created: time.Now(),
	}, nil
}

func (j *jwtProvider) Validate(token string) (*tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(
		token,
		&myClaims{},
		func(t *jwt.Token) (interface{}, error) { return []byte(j.secret), nil },
	)
	if err != nil {
		return nil, err
	}
	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	Claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}
	return &Claims.Payload, nil
}

func (j *jwtProvider) String() string {
	return "JWT implement Provider"
}
