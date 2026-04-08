package domain

type WishlistRequest struct {
	PropertyID uint `json:"property_id" validate:"required"`
}
