package subscriptions

import (
	"time"

	"github.com/devlopersabbir/juan_don82-server/api/subscriptions/core"
	"github.com/devlopersabbir/juan_don82-server/api/subscriptions/domain"
	"github.com/devlopersabbir/juan_don82-server/arch/networks"
	v "github.com/devlopersabbir/juan_don82-server/internal/pkg/validator"
	"github.com/gin-gonic/gin"
)

// GetAllPlans fetches all subscription plans
//
//	@Summary		Get all subscription plans
//	@Description	Fetches a list of all available subscription plans
//	@Tags			Subscriptions
//	@Produce		json
//	@Success		200	{array}	core.SubscriptionPlan
//	@Router			/api/v1/subscriptions/plans [get]
func GetAllPlans(c *gin.Context) {
	plans, err := ListPlansElastic(c)
	if err != nil || len(plans) == 0 {
		// Fallback to DB if ES is empty or errors
		plans, err = GetPlans()
		if err != nil {
			networks.Send(c).InternalServerError("Failed to fetch plans", err)
			return
		}
	}
	networks.Send(c).SuccessDataResponse("Plans fetched successfully from search index", plans)
}

// PurchaseSubscription handles subscription purchase for agents
//
//	@Summary		Purchase a subscription plan
//	@Description	Allows an agent to purchase a subscription plan
//	@Tags			Subscriptions
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		domain.PurchaseSubscriptionRequest	true	"Purchase Details"
//	@Success		200		{object}	map[string]string					"Subscription purchased successfully"
//	@Router			/api/v1/subscriptions/purchase [post]
func PurchaseSubscription(c *gin.Context) {
	var body domain.PurchaseSubscriptionRequest
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
	role, _ := c.Get("role")

	if role != "agent" {
		res.ForbiddenError("Only agents can purchase subscriptions", nil)
		return
	}

	plan, err := FindPlanByID(body.PlanID)
	if err != nil {
		res.NotFoundError("Plan not found", err)
		return
	}

	sub := &core.AgentSubscription{
		AgentID:   userID.(uint),
		PlanID:    plan.ID,
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 0, plan.DurationDays),
		Status:    "ACTIVE",
	}

	if err := CreateAgentSubscription(sub); err != nil {
		res.InternalServerError("Failed to purchase subscription", err)
		return
	}

	// Sync to Elastic
	if err := StoreElastic(c, sub); err != nil {
		res.InternalServerError("Failed to sync subscription to search index", err)
		return
	}

	res.SuccessMsgResponse("Subscription purchased successfully")
}

// AddPlan handles creation of new subscription plans (Admin only)
func AddPlan(c *gin.Context) {
	var body domain.CreatePlanRequest
	res := networks.Send(c)

	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequestError("Invalid request body", err)
		return
	}

	if errs := v.Validate(body); errs != nil {
		res.ValidationError("Validation failed", errs)
		return
	}

	plan := &core.SubscriptionPlan{
		Name:          body.Name,
		Price:         body.Price,
		PropertyLimit: body.PropertyLimit,
		DurationDays:  body.DurationDays,
	}

	if err := CreatePlan(plan); err != nil {
		res.InternalServerError("Failed to create plan", err)
		return
	}

	// Sync to Elastic
	if err := StorePlanElastic(c, plan); err != nil {
		res.InternalServerError("Failed to sync plan to search index", err)
		return
	}

	res.SuccessMsgResponse("Plan created successfully")
}
