package user

import (
	domainUser "github.com/chaiyawatkit/clean-architecture-gorm-mysql/src/domain/user"
)

type PaginationResultUser struct {
	Data       []domainUser.User
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}
