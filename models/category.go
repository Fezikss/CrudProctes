package models

import "github.com/google/uuid"

type Category struct {
	Id        uuid.UUID
	Name      string
	CreatedAt string
	UpdatedAt string
}
