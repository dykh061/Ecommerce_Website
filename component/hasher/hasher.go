package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct {
	cost int
}

func NewBcryptHasher(cost int) *BcryptHasher {
	return &BcryptHasher{cost: cost}
}

func (h *BcryptHasher) Hash(raw string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(raw),
		h.cost,
	)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
func (h *BcryptHasher) Compare(hashed, raw string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(raw),
	)
	return err == nil
}
