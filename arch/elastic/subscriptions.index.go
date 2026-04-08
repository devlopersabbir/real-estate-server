package elastic

var SubscriptionsIndex = IndexDefinition{
	Name: "subscriptions",
	Mapping: map[string]any{
		"mappings": map[string]any{
			"properties": map[string]any{
				"id":         map[string]any{"type": "integer"},
				"agent_id":   map[string]any{"type": "integer"},
				"plan_id":    map[string]any{"type": "integer"},
				"start_date": map[string]any{"type": "date"},
				"end_date":   map[string]any{"type": "date"},
				"status":     map[string]any{"type": "keyword"},
			},
		},
	},
}

var PlansIndex = IndexDefinition{
	Name: "subscription_plans",
	Mapping: map[string]any{
		"mappings": map[string]any{
			"properties": map[string]any{
				"id":             map[string]any{"type": "integer"},
				"name":           map[string]any{"type": "text"},
				"price":          map[string]any{"type": "float"},
				"property_limit": map[string]any{"type": "integer"},
				"duration_days":  map[string]any{"type": "integer"},
			},
		},
	},
}
