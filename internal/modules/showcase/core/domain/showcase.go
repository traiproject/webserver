package domain

import "github.com/google/uuid"

// ShowcaseItem represents a showcase item.
type ShowcaseItem struct {
	ID    uuid.UUID
	Title string
}
