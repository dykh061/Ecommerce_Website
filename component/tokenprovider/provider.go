package tokenprovider

import (
	"OpenMarket/common"
	"errors"
	"time"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("Token Not Found"),
		"token not found",
		"ErrNotFound",
	)
	ErrEncodingToken = common.NewCustomError(
		errors.New("Error Encoding The Token"),
		"error encoding the token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = common.NewCustomError(
		errors.New("Invalid Token provided	"),
		"invalid idtoken",
		"ErrInvalidToken",
	)
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserId int `json:"user_id"`
	// Role   string `json:"role"`
}
