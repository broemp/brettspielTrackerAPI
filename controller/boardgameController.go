package controller

import (
	"errors"

	"github.com/broemp/brettspielTrackerAPI/entity"
	"github.com/broemp/brettspielTrackerAPI/service"
	"github.com/gin-gonic/gin"
)

type BoardgameController interface {
	RandomBoardgame(ctx *gin.Context) (entity.Boardgame, error)
}
type controller struct {
	service service.BoardgameService
}

func NewBoardgameService(service service.BoardgameService) BoardgameController {
	return &controller{
		service: service,
	}
}

func (c *controller) RandomBoardgame(ctx *gin.Context) (entity.Boardgame, error) {

	queryMap := make(map[string]string)

	if ctx.Query("username") == "" {
		return entity.Boardgame{}, errors.New("Empty username")
	}

	for _, s := range service.PARAMETERS {
		if val, exists := ctx.GetQuery(s); exists {
			queryMap[s] = val
		}
	}

	randomGame, err := c.service.RandomBoardgame(queryMap)

	if err != nil {
		return randomGame, err
	}

	return randomGame, nil
}
