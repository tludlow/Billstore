package domain

import (
	"github.com/gofrs/uuid"
)

type User struct{
	Id uuid.UUID
	Username string
}

type EmailUserIdentity struct{
	Id uuid.UUID
	Email string
	Password string
	UserId uuid.UUID
}

type RefreshToken struct{
	Id uuid.UUID
	UserId uuid.UUID
	UserIdentityId uuid.UUID
}