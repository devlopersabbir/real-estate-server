package elastic

type IndexDefinition struct {
	Name    string
	Mapping map[string]any
}

// ── Register every index from every API module below ─────────────────────────
var AllIndexes = []IndexDefinition{
	UsersIndex,
	// add more here as new API modules are created, e.g.:
	// propertiesIndex,
}
