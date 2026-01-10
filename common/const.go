package common

const (
	CurrentUser = "user"
)

const (
	DbTypeUser    = 1
	DbTypeSeller  = 2
	DbTypeVariant = 3
	DbTypeProduct = 4
	DbTypeAddress = 5
)

type Requester interface {
	GetUserId() int
	GetUserEmail() string
}
