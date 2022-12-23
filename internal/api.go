package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/broemp/brettspielTrackerAPI/controller"
	"github.com/broemp/brettspielTrackerAPI/service"
)

var (
	boardgameService service.BoardgameService = service.NewBoardgameRepository()
	jwtService       service.JWTService       = service.NewJWTService()

	boardgameController controller.BoardgameController = controller.NewBoardgameService(boardgameService)
)

func StartServer() {

	server := gin.New()
	server.Use(gin.Logger(), gin.Recovery())

	// Set config production to true to change to production mode
	if viper.GetBool("production") {
		gin.SetMode(gin.ReleaseMode)
	}

	// apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	apiRoutes := server.Group("/api")

	createBoardgameAPI(*apiRoutes)

	server.Run("localhost:" + viper.GetString("apiPort"))
}

func createBoardgameAPI(apiRoutes gin.RouterGroup) {

	// Random Boardgame
	apiRoutes.GET("/boardgames/random", func(ctx *gin.Context) {

		game, err := boardgameController.RandomBoardgame(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, game)
		}
	})

}

func response(err error, ctx *gin.Context) {
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Valid"})
	}

}
