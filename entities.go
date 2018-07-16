package terraform2d

// Entity is an interface for implementing concrete "things" in the game
type Entity interface {
	ID() EntityID
	Category() EntityCategory
}

// EntityCategory is used to group entities
type EntityCategory uint

// EntityID represents an entity ID
type EntityID int
