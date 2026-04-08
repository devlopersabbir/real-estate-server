package elastic

type IndexDefinition struct {
	Name    string
	Mapping map[string]any
}

// ── Register every index from every API module below ─────────────────────────
var AllIndexes = []IndexDefinition{
	UsersIndex,
	PropertiesIndex,
	SubscriptionsIndex,
	PlansIndex,
	ChatRoomsIndex,
	MessagesIndex,
	WishlistIndex,
}
