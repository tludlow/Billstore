package view_models

import "github.com/gofrs/uuid"

type UserDetails struct{
	Id uuid.UUID
	Username string
	ContactEmail string
	AgeVerified bool
	MatureContentFilter bool
}