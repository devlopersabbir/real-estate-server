package wishlist

import (
	"github.com/devlopersabbir/juan_don82-server/api/wishlist/core"
	"github.com/devlopersabbir/juan_don82-server/api/wishlist/domain"
	"github.com/devlopersabbir/juan_don82-server/arch/networks"
	v "github.com/devlopersabbir/juan_don82-server/internal/pkg/validator"
	"github.com/gin-gonic/gin"
)

// Add handles adding property to wishlist
func Add(c *gin.Context) {
	var body domain.WishlistRequest
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

	item := &core.Wishlist{
		UserID:     userID.(uint),
		PropertyID: body.PropertyID,
	}

	if err := AddToWishlist(item); err != nil {
		res.InternalServerError("Failed to add to wishlist", err)
		return
	}

	res.SuccessMsgResponse("Added to wishlist")
}

// Remove handles removing property from wishlist
func Remove(c *gin.Context) {
	var body domain.WishlistRequest
	res := networks.Send(c)

	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequestError("Invalid request body", err)
		return
	}

	userID, _ := c.Get("userID")

	if err := RemoveFromWishlist(userID.(uint), body.PropertyID); err != nil {
		res.InternalServerError("Failed to remove from wishlist", err)
		return
	}

	res.SuccessMsgResponse("Removed from wishlist")
}

// List handles fetching user wishlist
func List(c *gin.Context) {
	userID, _ := c.Get("userID")
	list, err := GetUserWishlist(userID.(uint))
	if err != nil {
		networks.Send(c).InternalServerError("Failed to fetch wishlist", err)
		return
	}
	networks.Send(c).SuccessDataResponse("Wishlist fetched", list)
}
