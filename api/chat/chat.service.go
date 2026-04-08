package chat

import (
	"strconv"
	"time"

	"github.com/devlopersabbir/juan_don82-server/api/chat/core"
	"github.com/devlopersabbir/juan_don82-server/api/chat/domain"
	"github.com/devlopersabbir/juan_don82-server/arch/networks"
	v "github.com/devlopersabbir/juan_don82-server/internal/pkg/validator"
	"github.com/gin-gonic/gin"
)

// StartChat initializes a chat room for a property
//
//	@Summary		Start a chat with an agent
//	@Description	Creates a new chat room or returns existing one for a specific property
//	@Tags			Chat
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		domain.CreateRoomRequest	true	"Chat Details"
//	@Success		200		{object}	core.ChatRoom
//	@Router			/api/v1/chats/start [post]
func StartChat(c *gin.Context) {
	var body domain.CreateRoomRequest
	res := networks.Send(c)

	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequestError("Invalid request body", err)
		return
	}

	if errs := v.Validate(body); errs != nil {
		res.ValidationError("Validation failed", errs)
		return
	}

	userID, _ := c.Get("userID")

	room, err := FindRoom(userID.(uint), body.AgentID, body.PropertyID)
	if err == nil {
		res.SuccessDataResponse("Room redirected", room)
		return
	}

	newRoom := &core.ChatRoom{
		UserID:     userID.(uint),
		AgentID:    body.AgentID,
		PropertyID: body.PropertyID,
		CreatedAt:  time.Now(),
	}

	if err := CreateRoom(newRoom); err != nil {
		res.InternalServerError("Failed to create chat room", err)
		return
	}

	// Sync to Elastic
	if err := StoreRoomElastic(c, newRoom); err != nil {
		res.InternalServerError("Failed to sync chat room to search index", err)
		return
	}

	res.SuccessDataResponse("Chat room created", newRoom)
}

// SendMsg sends a message in a chat room
//
//	@Summary		Send a message
//	@Description	Sends a message to a specific chat room
//	@Tags			Chat
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		domain.SendMessageRequest	true	"Message Content"
//	@Success		200		{object}	core.Message
//	@Router			/api/v1/chats/message [post]
func SendMsg(c *gin.Context) {
	var body domain.SendMessageRequest
	res := networks.Send(c)

	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequestError("Invalid request body", err)
		return
	}

	if errs := v.Validate(body); errs != nil {
		res.ValidationError("Validation failed", errs)
		return
	}

	userID, _ := c.Get("userID")

	msg := &core.Message{
		RoomID:    body.RoomID,
		SenderID:  userID.(uint),
		Content:   body.Content,
		CreatedAt: time.Now(),
	}

	if err := CreateMessage(msg); err != nil {
		res.InternalServerError("Failed to send message", err)
		return
	}

	// Sync to Elastic
	if err := StoreMessageElastic(c, msg); err != nil {
		res.InternalServerError("Failed to sync message to search index", err)
		return
	}

	res.SuccessDataResponse("Message sent", msg)
}

// GetRooms fetches all chat rooms for the logged-in user
func GetRooms(c *gin.Context) {
	userID, _ := c.Get("userID")
	rooms, err := ListRoomsElastic(c, userID.(uint))
	if err != nil || len(rooms) == 0 {
		rooms, err = GetRoomsByUser(userID.(uint))
		if err != nil {
			networks.Send(c).InternalServerError("Failed to fetch rooms", err)
			return
		}
	}
	networks.Send(c).SuccessDataResponse("Rooms fetched from search index", rooms)
}

// GetRoomMessages fetches all messages for a room
func GetRoomMessages(c *gin.Context) {
	roomIDStr := c.Param("id")
	roomID, err := strconv.ParseUint(roomIDStr, 10, 32)
	if err != nil {
		networks.Send(c).BadRequestError("Invalid room ID", err)
		return
	}

	msgs, err := ListMessagesElastic(c, uint(roomID))
	if err != nil || len(msgs) == 0 {
		msgs, err = GetMessagesByRoom(uint(roomID))
		if err != nil {
			networks.Send(c).InternalServerError("Failed to fetch messages", err)
			return
		}
	}
	networks.Send(c).SuccessDataResponse("Messages fetched from search index", msgs)
}
