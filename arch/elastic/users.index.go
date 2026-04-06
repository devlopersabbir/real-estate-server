package elastic

// usersIndex — index/mapping for the users module.
var UsersIndex = IndexDefinition{
	Name: "users",
	Mapping: map[string]any{
		"mappings": map[string]any{
			"properties": map[string]any{
				"id": map[string]any{
					"type": "integer",
				},
				"name": map[string]any{
					"type": "text",
					"fields": map[string]any{
						"keyword": map[string]any{"type": "keyword"},
					},
				},
				"email": map[string]any{
					"type": "keyword",
				},
				"role": map[string]any{
					"type": "keyword",
				},
			},
		},
	},
}
