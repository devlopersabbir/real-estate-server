package domain

type CreatePlanRequest struct {
	Name          string  `json:"name" validate:"required"`
	Price         float64 `json:"price" validate:"required"`
	PropertyLimit int     `json:"property_limit" validate:"required"`
	DurationDays  int     `json:"duration_days" validate:"required"`
}

type PurchaseSubscriptionRequest struct {
	PlanID uint `json:"plan_id" validate:"required"`
}
