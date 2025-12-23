package userbusiness

type PasswordHasher interface {
	Hash(raw string) (string, error)
	Compare(hashed, raw string) bool
}
