package entities

import (
	"github.com/miketmoore/terraform2d/categories"
)

// Entity is an interface for implementing concrete "things" in the game
type Entity interface {
	ID() EntityID
	Category() categories.Category
}

// EntityID represents an entity ID
type EntityID int
