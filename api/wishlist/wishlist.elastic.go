package wishlist

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/devlopersabbir/juan_don82-server/api/wishlist/core"
	"github.com/devlopersabbir/juan_don82-server/arch/elastic"
	"github.com/devlopersabbir/juan_don82-server/internal/database"
)

func StoreElastic(ctx context.Context, item *core.Wishlist) error {
	index := elastic.WishlistIndex.Name
	if exists, err := database.ESClient.Exists(index, strconv.Itoa(int(item.ID))).IsSuccess(ctx); exists {
		log.Println("Wishlist item already exists in ES")
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to check if wishlist item exists in ES: %w", err)
	}
	_, err := database.ESClient.Index(index).
		Id(strconv.Itoa(int(item.ID))).
		Request(item).
		Do(ctx)

	if err != nil {
		return fmt.Errorf("failed to store wishlist item in ES: %w", err)
	}
	return nil
}

func DeleteElastic(ctx context.Context, id string) error {
	index := elastic.WishlistIndex.Name
	_, err := database.ESClient.Delete(index, id).Do(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete wishlist item from ES: %w", err)
	}

	return nil
}
