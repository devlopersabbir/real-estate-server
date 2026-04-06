package migrations

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/devlopersabbir/juan_don82-server/arch/elastic"
	"github.com/elastic/go-elasticsearch/v9"
)

// ── Public entrypoint ─────────────────────────────────────────────────────────
// CreateAllIndexes iterates over every registered index definition and creates
// the index if it does not already exist. Call this on startup after the ES
// client is connected — mirrors migrations.Automigrate(db) for Postgres.
func CreateAllIndexes(es *elasticsearch.TypedClient) {
	for _, idx := range elastic.AllIndexes {
		createIndex(es, idx)
	}
	log.Println("✅ Elasticsearch index migration completed")
}

// ── Internal helpers ──────────────────────────────────────────────────────────
func createIndex(es *elasticsearch.TypedClient, idx elastic.IndexDefinition) {
	// Check if index already exists.
	exists, err := es.Indices.Exists(idx.Name).Do(context.Background())
	if err != nil {
		log.Fatalf("ES: failed to check index %q existence: %v", idx.Name, err)
	}

	if exists {
		log.Printf("ES: index %q already exists — skipping", idx.Name)
		return
	}

	// Create the index with mapping.
	// We use the raw mapping from idx.Mapping.
	// Since idx.Mapping is map[string]any, we can't easily use the fully typed CreateIndexRequest
	// unless we want to rebuild it. But we can use the low-level body if needed,
	// or just use the Typed API with a raw request body.
	
	// Actually, the Typed API for CreateIndex allows setting the body.
	mappingJSON, _ := json.Marshal(idx.Mapping)
	
	_, err = es.Indices.Create(idx.Name).
		Raw(bytes.NewReader(mappingJSON)).
		Do(context.Background())

	if err != nil {
		log.Fatalf("ES: failed to create index %q: %v", idx.Name, err)
	}

	log.Printf("✅ ES: index %q created", idx.Name)
}
