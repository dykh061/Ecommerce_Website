package userbusiness

type PasswordHasher interface {
	Hash(raw string) (string, error)
}
