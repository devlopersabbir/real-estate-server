package database

import (
	"context"
	"log"

	"github.com/devlopersabbir/juan_don82-server/internal/migrations"
	"github.com/devlopersabbir/juan_don82-server/internal/pkg/config"
	"github.com/elastic/go-elasticsearch/v9"
)

var ESClient *elasticsearch.TypedClient

func ESClientConnection(config *config.ElasticConfig) {
	esConfig := elasticsearch.Config{
		Addresses: config.ESAddresses,
		Username:  config.ESUsername,
		Password:  config.ESPassword,
	}
	es, err := elasticsearch.NewTypedClient(esConfig)
	if err != nil {
		log.Fatalf("Fail to create elastic client: %v", err)
	}

	_, err = es.Info().Do(context.Background())
	if err != nil {
		log.Fatalf("Elastic ping fail: %v", err)
	}

	ESClient = es
	log.Println("✅ Elasticsearch connected")

	// Run index migrations — mirrors migrations.Automigrate(db) for Postgres
	migrations.CreateAllIndexes(ESClient)
}
