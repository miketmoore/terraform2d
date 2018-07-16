package world

import (
	"github.com/miketmoore/terraform2d/categories"
	"github.com/miketmoore/terraform2d/entities"
)

// System is an interface
type System interface {
	Update()
	AddEntity(entities.Entity)
}

// World is a world struct
type World struct {
	systems      []System
	SystemsMap   map[string]System
	lastEntityID entities.EntityID
}

// New returns a new World
func New() World {
	return World{
		lastEntityID: 0,
		SystemsMap:   map[string]System{},
	}
}

// AddSystem adds a System to the World
func (w *World) AddSystem(sys System) {
	w.systems = append(w.systems, sys)
}

// AddSystems adds a batch of systems to the world
func (w *World) AddSystems(all ...System) {
	for _, sys := range all {
		w.AddSystem(sys)
	}
}

// Update executes Update on all systems in this World
func (w *World) Update() {
	for _, sys := range w.systems {
		sys.Update()
	}
}

// Systems returns the systems in this World
func (w *World) Systems() []System {
	return w.systems
}

// NewEntityID generates and returns a new Entity ID
func (w *World) NewEntityID() entities.EntityID {
	w.lastEntityID++
	return w.lastEntityID
}

// AddEntityToSystem adds the entity to it's system
func (w *World) AddEntityToSystem(entity entities.Entity) {
	for _, system := range w.Systems() {
		system.AddEntity(entity)
	}
}

// AddEntitiesToSystem adds the entities to their system
func (w *World) AddEntitiesToSystem(entities ...entities.Entity) {
	for _, entity := range entities {
		w.AddEntityToSystem(entity)
	}
}

// EntityRemover is an interface to describe an API for removing
// an entity by category and ID
type EntityRemover interface {
	Remove(category categories.Category, id entities.EntityID)
}
