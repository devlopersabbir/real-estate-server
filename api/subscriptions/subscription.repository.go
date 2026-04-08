package subscriptions

import (
	"github.com/devlopersabbir/juan_don82-server/api/subscriptions/core"
	"github.com/devlopersabbir/juan_don82-server/internal/database"
)

func CreatePlan(plan *core.SubscriptionPlan) error {
	return database.DB.Create(plan).Error
}

func GetPlans() ([]core.SubscriptionPlan, error) {
	var plans []core.SubscriptionPlan
	err := database.DB.Find(&plans).Error
	return plans, err
}

func FindPlanByID(id uint) (core.SubscriptionPlan, error) {
	var plan core.SubscriptionPlan
	err := database.DB.First(&plan, id).Error
	return plan, err
}

func CreateAgentSubscription(sub *core.AgentSubscription) error {
	return database.DB.Create(sub).Error
}

func FindActiveSubscription(agentID uint) (core.AgentSubscription, error) {
	var sub core.AgentSubscription
	err := database.DB.Where("agent_id = ? AND status = 'ACTIVE' AND end_date > NOW()", agentID).First(&sub).Error
	return sub, err
}
