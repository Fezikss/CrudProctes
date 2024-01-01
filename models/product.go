package models

import (
	"github.com/google/uuid"
)

type Product struct {
	Id         uuid.UUID
	Name       string
	Price      int
	CategoryID uuid.UUID
	CreatedAt  string
	UpdatedAt  string
}
