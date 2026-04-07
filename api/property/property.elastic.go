package property

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/devlopersabbir/juan_don82-server/api/property/core"
	"github.com/devlopersabbir/juan_don82-server/arch/elastic"
	"github.com/devlopersabbir/juan_don82-server/internal/database"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

func StoreElastic(ctx context.Context, property *core.Property) error {
	index := elastic.PropertiesIndex.Name
	if exists, err := database.ESClient.Exists(index, strconv.Itoa(int(property.ID))).IsSuccess(ctx); exists {
		log.Println("Property already exists in ES")
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to check if property exists in ES: %w", err)
	}
	_, err := database.ESClient.Index(index).
		Id(strconv.Itoa(int(property.ID))).
		Request(property).
		Do(ctx)

	if err != nil {
		return fmt.Errorf("failed to store property in ES: %w", err)
	}
	return nil
}

func FetchElastic(ctx context.Context, id string) (*core.Property, error) {
	index := elastic.PropertiesIndex.Name
	res, err := database.ESClient.Get(index, id).Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch property from ES: %w", err)
	}

	if !res.Found {
		return nil, fmt.Errorf("property not found in ES: %s", id)
	}

	var property core.Property
	if err := json.Unmarshal(res.Source_, &property); err != nil {
		return nil, fmt.Errorf("failed to unmarshal property from ES: %w", err)
	}

	return &property, nil
}

func UpdateElastic(ctx context.Context, id string, property *core.Property) error {
	index := elastic.PropertiesIndex.Name
	_, err := database.ESClient.Update(index, id).
		Doc(property).
		Do(ctx)

	if err != nil {
		return fmt.Errorf("failed to update property in ES: %w", err)
	}

	return nil
}

func DeleteElastic(ctx context.Context, id string) error {
	index := elastic.PropertiesIndex.Name
	_, err := database.ESClient.Delete(index, id).Do(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete property from ES: %w", err)
	}

	return nil
}

type PropertyFilter struct {
	City         string   `form:"city"`
	PropertyType string   `form:"property_type"`
	MinPrice     *float64 `form:"min_price"`
	MaxPrice     *float64 `form:"max_price"`
	Status       string   `form:"status"`
}

func SearchPropertiesElastic(ctx context.Context, filter PropertyFilter) ([]core.Property, error) {
	index := elastic.PropertiesIndex.Name

	// Build the boolean query using Typed API structs
	boolQuery := &types.BoolQuery{
		Must: []types.Query{},
	}

	if filter.City != "" {
		boolQuery.Must = append(boolQuery.Must, types.Query{
			Term: map[string]types.TermQuery{
				"city": {Value: filter.City},
			},
		})
	}
	if filter.PropertyType != "" {
		boolQuery.Must = append(boolQuery.Must, types.Query{
			Term: map[string]types.TermQuery{
				"property_type": {Value: filter.PropertyType},
			},
		})
	}
	if filter.Status != "" {
		boolQuery.Must = append(boolQuery.Must, types.Query{
			Term: map[string]types.TermQuery{
				"status": {Value: filter.Status},
			},
		})
	}

	// Price range filtering
	if filter.MinPrice != nil || filter.MaxPrice != nil {
		rangeQuery := types.NumberRangeQuery{}
		if filter.MinPrice != nil {
			minVal := types.Float64(*filter.MinPrice)
			rangeQuery.Gte = &minVal
		}
		if filter.MaxPrice != nil {
			maxVal := types.Float64(*filter.MaxPrice)
			rangeQuery.Lte = &maxVal
		}
		boolQuery.Must = append(boolQuery.Must, types.Query{
			Range: map[string]types.RangeQuery{
				"price": rangeQuery,
			},
		})
	}

	// If no filters, use match_all
	var query *types.Query
	if len(boolQuery.Must) == 0 {
		query = &types.Query{MatchAll: &types.MatchAllQuery{}}
	} else {
		query = &types.Query{Bool: boolQuery}
	}

	res, err := database.ESClient.Search().
		Index(index).
		Query(query).
		Size(100).
		Do(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to search properties in ES: %w", err)
	}

	properties := make([]core.Property, 0)
	for _, hit := range res.Hits.Hits {
		var p core.Property
		if err := json.Unmarshal(hit.Source_, &p); err != nil {
			log.Printf("Error unmarshaling property hit: %v", err)
			continue
		}
		properties = append(properties, p)
	}

	return properties, nil
}
