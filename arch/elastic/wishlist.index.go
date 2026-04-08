package elastic

var WishlistIndex = IndexDefinition{
	Name: "wishlist",
	Mapping: map[string]any{
		"mappings": map[string]any{
			"properties": map[string]any{
				"id":          map[string]any{"type": "integer"},
				"user_id":     map[string]any{"type": "integer"},
				"property_id": map[string]any{"type": "integer"},
			},
		},
	},
}
