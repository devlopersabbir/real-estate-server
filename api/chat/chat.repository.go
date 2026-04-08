package chat

import (
	"github.com/devlopersabbir/juan_don82-server/api/chat/core"
	"github.com/devlopersabbir/juan_don82-server/internal/database"
)

func CreateRoom(room *core.ChatRoom) error {
	return database.DB.Create(room).Error
}

func FindRoom(userID, agentID, propertyID uint) (core.ChatRoom, error) {
	var room core.ChatRoom
	err := database.DB.Where("user_id = ? AND agent_id = ? AND property_id = ?", userID, agentID, propertyID).First(&room).Error
	return room, err
}

func CreateMessage(msg *core.Message) error {
	return database.DB.Create(msg).Error
}

func GetMessagesByRoom(roomID uint) ([]core.Message, error) {
	var msgs []core.Message
	err := database.DB.Where("room_id = ?", roomID).Order("created_at asc").Find(&msgs).Error
	return msgs, err
}

func GetRoomsByUser(userID uint) ([]core.ChatRoom, error) {
	var rooms []core.ChatRoom
	err := database.DB.Where("user_id = ? OR agent_id = ?", userID, userID).Find(&rooms).Error
	return rooms, err
}

func GetAllRooms() ([]core.ChatRoom, error) {
	var rooms []core.ChatRoom
	err := database.DB.Find(&rooms).Error
	return rooms, err
}
