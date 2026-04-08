package domain

type CreateRoomRequest struct {
	PropertyID uint `json:"property_id" validate:"required"`
	AgentID    uint `json:"agent_id" validate:"required"`
}

type SendMessageRequest struct {
	RoomID  uint   `json:"room_id" validate:"required"`
	Content string `json:"content" validate:"required"`
}
