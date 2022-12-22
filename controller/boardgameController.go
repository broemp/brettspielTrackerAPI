package controller

import (
	"github.com/broemp/brettspielTrackerAPI/entity"
	"github.com/broemp/brettspielTrackerAPI/service"
	"github.com/gin-gonic/gin"
)

type BoardgameController interface {
	RandomBoardgameFromBGGCollection(ctx *gin.Context) (entity.Boardgame, error)
}
type controller struct {
	service service.BoardgameService
}

func NewBoardgameService(service service.BoardgameService) BoardgameController {
	return &controller{
		service: service,
	}
}

func (c *controller) RandomBoardgameFromBGGCollection(ctx *gin.Context) (entity.Boardgame, error) {
	username := ctx.Param("username")

	randomGame, err := c.service.RandomBoardgameFromBGGCollection(username)

	if err != nil {
		return randomGame, err
	}

	return randomGame, nil
}
