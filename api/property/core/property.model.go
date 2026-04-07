package core

type Property struct {
	ID            uint     `json:"id" gorm:"primaryKey"`
	UserID        uint     `json:"user_id" gorm:"not null"`
	Name          string   `json:"name" gorm:"not null"`
	Description   string   `json:"description" gorm:"not null"`
	Address       string   `json:"address" gorm:"not null"`
	City          string   `json:"city" gorm:"not null"`
	State         string   `json:"state" gorm:"not null"`
	ZipCode       string   `json:"zip_code" gorm:"not null"`
	Country       string   `json:"country" gorm:"not null"`
	Price         float64  `json:"price" gorm:"not null"`
	DiscountPrice float64  `json:"discount_price" gorm:"not null"`
	Bedrooms      int      `json:"bedrooms" gorm:"not null"`
	Bathrooms     int      `json:"bathrooms" gorm:"not null"`
	SquareFeet    int      `json:"square_feet" gorm:"not null"`
	PropertyType  string   `json:"property_type" gorm:"not null"` // e.g. SELL or RENT
	RentPeriod    string   `json:"rent_period"`                   // e.g. YEARLY or MONTHLY
	Status        string   `json:"status" gorm:"not null"`
	Images        []string `json:"images" gorm:"serializer:json"`
	Videos        []string `json:"videos" gorm:"serializer:json"`
}
