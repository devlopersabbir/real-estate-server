package core

import "time"

type SubscriptionPlan struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	Name          string  `json:"name" gorm:"not null"`
	Price         float64 `json:"price" gorm:"not null"`
	PropertyLimit int     `json:"property_limit" gorm:"not null"`
	DurationDays  int     `json:"duration_days" gorm:"not null"`
}

type AgentSubscription struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	AgentID   uint      `json:"agent_id" gorm:"not null"`
	PlanID    uint      `json:"plan_id" gorm:"not null"`
	StartDate time.Time `json:"start_date" gorm:"not null"`
	EndDate   time.Time `json:"end_date" gorm:"not null"`
	Status    string    `json:"status" gorm:"not null"` // e.g. ACTIVE, EXPIRED
}
