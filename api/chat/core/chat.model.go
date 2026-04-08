package core

import "time"

type ChatRoom struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	AgentID    uint      `json:"agent_id" gorm:"not null"`
	PropertyID uint      `json:"property_id" gorm:"not null"` // Negotiation with reference to property
	CreatedAt  time.Time `json:"created_at"`
}

type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	RoomID    uint      `json:"room_id" gorm:"not null"`
	SenderID  uint      `json:"sender_id" gorm:"not null"`
	Content   string    `json:"content" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
}
