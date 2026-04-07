package elastic

var PropertiesIndex = IndexDefinition{
	Name: "properties",
	Mapping: map[string]any{
		"mappings": map[string]any{
			"properties": map[string]any{
				"id": map[string]any{
					"type": "integer",
				},
				"user_id": map[string]any{
					"type": "integer",
				},
				"name": map[string]any{
					"type": "text",
					"fields": map[string]any{
						"keyword": map[string]any{"type": "keyword"},
					},
				},
				"description": map[string]any{
					"type": "text",
				},
				"address":        map[string]any{"type": "text"},
				"city":           map[string]any{"type": "keyword"},
				"state":          map[string]any{"type": "keyword"},
				"zip_code":       map[string]any{"type": "keyword"},
				"country":        map[string]any{"type": "keyword"},
				"price":          map[string]any{"type": "double"},
				"discount_price": map[string]any{"type": "double"},
				"bedrooms":       map[string]any{"type": "integer"},
				"bathrooms":      map[string]any{"type": "integer"},
				"square_feet":    map[string]any{"type": "integer"},
				"property_type":  map[string]any{"type": "keyword"},
				"rent_period":    map[string]any{"type": "keyword"},
				"status":         map[string]any{"type": "keyword"},
			},
		},
	},
}
