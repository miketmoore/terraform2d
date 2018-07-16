package categories

// Category is used to group entities
type Category uint

const (
	Player = Category(1 << iota)
	Enemy
	Obstacle
	MovableObstacle
	CollisionSwitch
	Warp
)
