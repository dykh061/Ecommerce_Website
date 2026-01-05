package common

const (
	CurrentUser = "user"
)

const (
	DbTypeUser    = 1
	DbTypeSeller  = 2
	DbTypeVariant = 3
)

type Requester interface {
	GetUserId() int
	GetUserEmail() string
}
