package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/devlopersabbir/juan_don82-server/api/chat/core"
	"github.com/devlopersabbir/juan_don82-server/arch/elastic"
	"github.com/devlopersabbir/juan_don82-server/internal/database"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

func StoreRoomElastic(ctx context.Context, room *core.ChatRoom) error {
	index := elastic.ChatRoomsIndex.Name
	if exists, err := database.ESClient.Exists(index, strconv.Itoa(int(room.ID))).IsSuccess(ctx); exists {
		log.Println("Chat room already exists in ES")
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to check if chat room exists in ES: %w", err)
	}
	_, err := database.ESClient.Index(index).
		Id(strconv.Itoa(int(room.ID))).
		Request(room).
		Do(ctx)

	if err != nil {
		return fmt.Errorf("failed to store chat room in ES: %w", err)
	}
	return nil
}

func ListRoomsElastic(ctx context.Context, userID uint) ([]core.ChatRoom, error) {
	index := elastic.ChatRoomsIndex.Name
	// Filter by user_id OR agent_id
	query := &types.Query{
		Bool: &types.BoolQuery{
			Should: []types.Query{
				{Term: map[string]types.TermQuery{"user_id": {Value: userID}}},
				{Term: map[string]types.TermQuery{"agent_id": {Value: userID}}},
			},
		},
	}
	res, err := database.ESClient.Search().
		Index(index).
		Query(query).
		Size(100).
		Do(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to search chat rooms in ES: %w", err)
	}

	rooms := make([]core.ChatRoom, 0)
	for _, hit := range res.Hits.Hits {
		var r core.ChatRoom
		if err := json.Unmarshal(hit.Source_, &r); err != nil {
			log.Printf("Error unmarshaling chat room hit: %v", err)
			continue
		}
		rooms = append(rooms, r)
	}

	return rooms, nil
}

func StoreMessageElastic(ctx context.Context, msg *core.Message) error {
	index := elastic.MessagesIndex.Name
	if exists, err := database.ESClient.Exists(index, strconv.Itoa(int(msg.ID))).IsSuccess(ctx); exists {
		log.Println("Message already exists in ES")
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to check if message exists in ES: %w", err)
	}
	_, err := database.ESClient.Index(index).
		Id(strconv.Itoa(int(msg.ID))).
		Request(msg).
		Do(ctx)

	if err != nil {
		return fmt.Errorf("failed to store message in ES: %w", err)
	}
	return nil
}

func ListMessagesElastic(ctx context.Context, roomID uint) ([]core.Message, error) {
	index := elastic.MessagesIndex.Name
	query := &types.Query{
		Term: map[string]types.TermQuery{"room_id": {Value: roomID}},
	}
	res, err := database.ESClient.Search().
		Index(index).
		Query(query).
		Size(500).
		Do(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to search messages in ES: %w", err)
	}

	msgs := make([]core.Message, 0)
	for _, hit := range res.Hits.Hits {
		var m core.Message
		if err := json.Unmarshal(hit.Source_, &m); err != nil {
			log.Printf("Error unmarshaling message hit: %v", err)
			continue
		}
		msgs = append(msgs, m)
	}

	// Sort in memory
	sort.Slice(msgs, func(i, j int) bool {
		return msgs[i].CreatedAt.Before(msgs[j].CreatedAt)
	})

	return msgs, nil
}
