package admin

import (
	"github.com/devlopersabbir/juan_don82-server/api/chat"
	"github.com/devlopersabbir/juan_don82-server/api/property"
	"github.com/devlopersabbir/juan_don82-server/api/users"
	"github.com/devlopersabbir/juan_don82-server/arch/networks"
	"github.com/gin-gonic/gin"
)

// ManageUsers handles fetching all users/agents
//
//	@Summary		Get all users
//	@Description	Admin can fetch all users and agents
//	@Tags			Admin
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{array}	interface{}
//	@Router			/api/v1/admin/users [get]
func ManageUsers(c *gin.Context) {
	u, err := users.FindAllUsers()
	if err != nil {
		networks.Send(c).InternalServerError("Failed to fetch users", err)
		return
	}
	networks.Send(c).SuccessDataResponse("Users fetched successfully", u)
}

// ManageProperties handles fetching/modifying any property
//
//	@Summary		Get all properties
//	@Description	Admin can fetch all properties
//	@Tags			Admin
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{array}	interface{}
//	@Router			/api/v1/admin/properties [get]
func ManageProperties(c *gin.Context) {
	p, err := property.FindAll()
	if err != nil {
		networks.Send(c).InternalServerError("Failed to fetch properties", err)
		return
	}
	networks.Send(c).SuccessDataResponse("Properties fetched successfully", p)
}

// ViewAllChats allows admin to see all conversations
//
//	@Summary		Get all chats
//	@Description	Admin can fetch all chat rooms
//	@Tags			Admin
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{array}	interface{}
//	@Router			/api/v1/admin/chats [get]
func ViewAllChats(c *gin.Context) {
	rooms, err := chat.GetAllRooms()
	if err != nil {
		networks.Send(c).InternalServerError("Failed to fetch chats", err)
		return
	}
	networks.Send(c).SuccessDataResponse("All chats fetched", rooms)
}
