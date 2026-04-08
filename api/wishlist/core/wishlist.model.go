package core

type Wishlist struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	UserID     uint `json:"user_id" gorm:"not null"`
	PropertyID uint `json:"property_id" gorm:"not null"`
}
