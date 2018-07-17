package terraform2d

// RoomID is a room ID
type RoomID int

// ConnectedRooms is used to configure adjacent rooms
type ConnectedRooms struct {
	Top    RoomID
	Right  RoomID
	Bottom RoomID
	Left   RoomID
}

// Room defines an API for room implementations
type Room interface {
	MapName() string
	ConnectedRooms() *ConnectedRooms
	SetConnectedRoom(Direction, RoomID)
}
