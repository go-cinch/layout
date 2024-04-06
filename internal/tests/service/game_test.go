package service

import (
	"context"
	"strconv"
	"strings"
	"testing"

	"github.com/go-cinch/common/proto/params"
	"github.com/go-cinch/layout/api/game"
	"github.com/go-cinch/layout/internal/tests/mock"
	"github.com/google/uuid"
)

func TestGameService_CreateGame(t *testing.T) {
	s := mock.GameService()
	ctx := context.Background()
	userID := uuid.NewString()
	ctx = mock.NewContextWithUserId(ctx, userID)

	_, err := s.CreateGame(ctx, &game.CreateGameRequest{
		Name: "game1",
	})
	if err != nil {
		t.Error(err)
		return
	}
	_, err = s.CreateGame(ctx, &game.CreateGameRequest{
		Name: "game2",
	})
	if err != nil {
		t.Error(err)
		return
	}
	res1, _ := s.FindGame(ctx, &game.FindGameRequest{
		Page: &params.Page{
			Disable: true,
		},
	})
	if res1 == nil || len(res1.List) != 2 {
		t.Error("res1 len must be 2")
		return
	}
	res2, err := s.GetGame(ctx, &game.GetGameRequest{
		Id: res1.List[0].Id,
	})
	if err != nil {
		t.Error(err)
		return
	}
	if res2.Name != res1.List[0].Name {
		t.Errorf("res2.Name must be %s", res1.List[0].Name)
		return
	}
	_, err = s.DeleteGame(ctx, &params.IdsRequest{
		Ids: strings.Join([]string{
			strconv.FormatUint(res1.List[0].Id, 10),
			strconv.FormatUint(res1.List[1].Id, 10),
		}, ","),
	})
	if err != nil {
		t.Error(err)
		return
	}
	res3, _ := s.FindGame(ctx, &game.FindGameRequest{
		Page: &params.Page{
			Disable: true,
		},
	})
	if res3 == nil || len(res3.List) != 0 {
		t.Error("res3 len must be 0")
		return
	}
}
