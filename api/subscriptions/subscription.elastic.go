package subscriptions

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/devlopersabbir/juan_don82-server/api/subscriptions/core"
	"github.com/devlopersabbir/juan_don82-server/arch/elastic"
	"github.com/devlopersabbir/juan_don82-server/internal/database"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

func StoreElastic(ctx context.Context, sub *core.AgentSubscription) error {
	index := elastic.SubscriptionsIndex.Name
	if exists, err := database.ESClient.Exists(index, strconv.Itoa(int(sub.ID))).IsSuccess(ctx); exists {
		log.Println("Subscription already exists in ES")
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to check if subscription exists in ES: %w", err)
	}
	_, err := database.ESClient.Index(index).
		Id(strconv.Itoa(int(sub.ID))).
		Request(sub).
		Do(ctx)

	if err != nil {
		return fmt.Errorf("failed to store subscription in ES: %w", err)
	}
	return nil
}

func StorePlanElastic(ctx context.Context, plan *core.SubscriptionPlan) error {
	index := elastic.PlansIndex.Name
	if exists, err := database.ESClient.Exists(index, strconv.Itoa(int(plan.ID))).IsSuccess(ctx); exists {
		log.Println("Plan already exists in ES")
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to check if plan exists in ES: %w", err)
	}
	_, err := database.ESClient.Index(index).
		Id(strconv.Itoa(int(plan.ID))).
		Request(plan).
		Do(ctx)

	if err != nil {
		return fmt.Errorf("failed to store plan in ES: %w", err)
	}
	return nil
}

func ListPlansElastic(ctx context.Context) ([]core.SubscriptionPlan, error) {
	index := elastic.PlansIndex.Name
	res, err := database.ESClient.Search().
		Index(index).
		Query(&types.Query{MatchAll: &types.MatchAllQuery{}}).
		Size(100).
		Do(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to search plans in ES: %w", err)
	}

	plans := make([]core.SubscriptionPlan, 0)
	for _, hit := range res.Hits.Hits {
		var p core.SubscriptionPlan
		if err := json.Unmarshal(hit.Source_, &p); err != nil {
			log.Printf("Error unmarshaling plan hit: %v", err)
			continue
		}
		plans = append(plans, p)
	}

	return plans, nil
}
