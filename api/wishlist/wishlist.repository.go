package wishlist

import (
	"github.com/devlopersabbir/juan_don82-server/api/wishlist/core"
	"github.com/devlopersabbir/juan_don82-server/internal/database"
)

func AddToWishlist(item *core.Wishlist) error {
	return database.DB.Create(item).Error
}

func RemoveFromWishlist(userID, propertyID uint) error {
	return database.DB.Where("user_id = ? AND property_id = ?", userID, propertyID).Delete(&core.Wishlist{}).Error
}

func GetUserWishlist(userID uint) ([]core.Wishlist, error) {
	var list []core.Wishlist
	err := database.DB.Where("user_id = ?", userID).Find(&list).Error
	return list, err
}
