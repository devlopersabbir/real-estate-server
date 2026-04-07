package domain

type Property struct {
	ID            int      `json:"id"`
	UserID        int      `json:"user_id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Address       string   `json:"address"`
	City          string   `json:"city"`
	State         string   `json:"state"`
	ZipCode       string   `json:"zip_code"`
	Country       string   `json:"country"`
	Price         float64  `json:"price"`
	DiscountPrice float64  `json:"discount_price"`
	Bedrooms      int      `json:"bedrooms"`
	Bathrooms     int      `json:"bathrooms"`
	SquareFeet    int      `json:"square_feet"`
	PropertyType  string   `json:"property_type"`
	RentPeriod    string   `json:"rent_period"`
	Status        string   `json:"status"`
	Images        []string `json:"images"`
	Videos        []string `json:"videos"`
}

type PropertyRequest struct {
	Name          string   `json:"name" validate:"required"`
	Description   string   `json:"description" validate:"required"`
	Address       string   `json:"address" validate:"required"`
	City          string   `json:"city" validate:"required"`
	State         string   `json:"state" validate:"required"`
	ZipCode       string   `json:"zip_code" validate:"required"`
	Country       string   `json:"country" validate:"required"`
	Price         float64  `json:"price" validate:"required"`
	DiscountPrice float64  `json:"discount_price"`
	Bedrooms      int      `json:"bedrooms" validate:"required"`
	Bathrooms     int      `json:"bathrooms" validate:"required"`
	SquareFeet    int      `json:"square_feet" validate:"required"`
	PropertyType  string   `json:"property_type" validate:"required,oneof=SELL RENT"`
	RentPeriod    string   `json:"rent_period" validate:"omitempty,oneof=YEARLY MONTHLY"`
	Status        string   `json:"status" validate:"required"`
	Images        []string `json:"images" validate:"required"`
	Videos        []string `json:"videos"`
}
