package users

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/devlopersabbir/juan_don82-server/api/users/core"
	"github.com/devlopersabbir/juan_don82-server/arch/elastic"
	"github.com/devlopersabbir/juan_don82-server/internal/database"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

func StoreElastic(ctx context.Context, user *core.Users) error {
	index := elastic.UsersIndex.Name
	if exists, err := database.ESClient.Exists(index, strconv.Itoa(int(user.ID))).IsSuccess(ctx); exists {
		log.Println("User already exists in ES")
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to check if user exists in ES: %w", err)
	}
	_, err := database.ESClient.Index(index).
		Id(strconv.Itoa(int(user.ID))).
		Request(user).
		Do(ctx)

	if err != nil {
		return fmt.Errorf("failed to store user in ES: %w", err)
	}
	return nil
}

func FetchElastic(ctx context.Context, id string) (*core.Users, error) {
	index := elastic.UsersIndex.Name
	res, err := database.ESClient.Get(index, id).Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user from ES: %w", err)
	}

	if !res.Found {
		return nil, fmt.Errorf("user not found in ES: %s", id)
	}

	var user core.Users
	if err := json.Unmarshal(res.Source_, &user); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user from ES: %w", err)
	}

	return &user, nil
}

func UpdateElastic(ctx context.Context, id string, user *core.Users) error {
	index := elastic.UsersIndex.Name
	_, err := database.ESClient.Update(index, id).
		Doc(user).
		Do(ctx)

	if err != nil {
		return fmt.Errorf("failed to update user in ES: %w", err)
	}

	return nil
}

func DeleteElastic(ctx context.Context, id string) error {
	index := elastic.UsersIndex.Name
	_, err := database.ESClient.Delete(index, id).Do(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete user from ES: %w", err)
	}

	return nil
}

func ListUsersElastic(ctx context.Context) ([]core.Users, error) {
	index := elastic.UsersIndex.Name
	res, err := database.ESClient.Search().
		Index(index).
		Query(&types.Query{MatchAll: &types.MatchAllQuery{}}).
		Size(100).
		Do(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to search users in ES: %w", err)
	}

	users := make([]core.Users, 0)
	for _, hit := range res.Hits.Hits {
		var u core.Users
		if err := json.Unmarshal(hit.Source_, &u); err != nil {
			log.Printf("Error unmarshaling user hit: %v", err)
			continue
		}
		users = append(users, u)
	}

	return users, nil
}
