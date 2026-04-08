package elastic

var ChatRoomsIndex = IndexDefinition{
	Name: "chat_rooms",
	Mapping: map[string]any{
		"mappings": map[string]any{
			"properties": map[string]any{
				"id":          map[string]any{"type": "integer"},
				"user_id":     map[string]any{"type": "integer"},
				"agent_id":    map[string]any{"type": "integer"},
				"property_id": map[string]any{"type": "integer"},
				"created_at":  map[string]any{"type": "date"},
			},
		},
	},
}

var MessagesIndex = IndexDefinition{
	Name: "messages",
	Mapping: map[string]any{
		"mappings": map[string]any{
			"properties": map[string]any{
				"id":         map[string]any{"type": "integer"},
				"room_id":    map[string]any{"type": "integer"},
				"sender_id":  map[string]any{"type": "integer"},
				"content":    map[string]any{"type": "text"},
				"created_at": map[string]any{"type": "date"},
			},
		},
	},
}
