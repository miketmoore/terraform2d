package terraform2d

// System is an interface
type System interface {
	Update()
	AddEntity(Entity)
}

// World is a world struct
type World struct {
	Systems       []System
	SystemsMap    map[string]System
	lastEntityID  EntityID
	EntityRemover EntityRemover
}

// NewWorld returns a new World
func NewWorld(entityRemover EntityRemover) World {
	return World{
		lastEntityID:  0,
		SystemsMap:    map[string]System{},
		EntityRemover: entityRemover,
	}
}

// AddSystem adds a System to the World
func (w *World) AddSystem(sys System) {
	w.Systems = append(w.Systems, sys)
}

// AddSystems adds a batch of Systems to the world
func (w *World) AddSystems(all ...System) {
	for _, sys := range all {
		w.AddSystem(sys)
	}
}

// Update executes Update on all Systems in this World
func (w *World) Update() {
	for _, sys := range w.Systems {
		sys.Update()
	}
}

// NewEntityID generates and returns a new Entity ID
func (w *World) NewEntityID() EntityID {
	w.lastEntityID++
	return w.lastEntityID
}

// AddEntityToSystem adds the entity to it's system
func (w *World) AddEntityToSystem(entity Entity) {
	for _, system := range w.Systems {
		system.AddEntity(entity)
	}
}

// AddEntitiesToSystem adds the terraform2d to their system
func (w *World) AddEntitiesToSystem(terraform2d ...Entity) {
	for _, entity := range terraform2d {
		w.AddEntityToSystem(entity)
	}
}

// EntityRemover is an interface to describe an API for removing
// an entity by category and ID
type EntityRemover interface {
	Remove(*World, EntityCategory, EntityID)
	RemoveAllEntities(*World)
}

// Remove is a proxy to a user implemented EntityRemover.Remove function
func (w *World) Remove(category EntityCategory, id EntityID) {
	w.EntityRemover.Remove(w, category, id)
}
