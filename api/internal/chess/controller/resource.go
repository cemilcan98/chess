package controller

import (
	"fmt"
	"github.com/cemilcan98/chess/internal/chess"
	"github.com/cemilcan98/chess/pkg/pagination"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

type resource struct {
	gameCollection *mongo.Collection
}

func NewResource(db *mongo.Database) *resource {

	return &resource{gameCollection: db.Collection("games")}

}

func (receiver *resource) getGame(ctx echo.Context) error {

	id := ctx.Param("id")
	if len(id) == 0 {
		return ctx.JSON(http.StatusBadGateway, "Id required.")
	}
	game := chess.Game{}
	receiver.gameCollection.FindOne(ctx.Request().Context(),
		bson.M{"_id": id}).
		Decode(&game)

	return ctx.JSON(http.StatusOK, game)

}

func (receiver *resource) helloApi(c echo.Context) error {

	return c.JSON(http.StatusOK, "API is running.")
}

func (receiver *resource) postGame(c echo.Context) error {
	game := c.Param("game")
	fen := c.QueryParam("fen")
	pgn := c.QueryParam("pgn")
	fmt.Printf("Game Number: %s, Fen:%s\n", game, fen)
	fmt.Printf("%s\n", pgn)

	//TODO: post game

	return c.JSON(http.StatusOK, "ok")
}

func (receiver *resource) getGameByUsername(ctx echo.Context) error {
	c := ctx.Request().Context()
	pg := pagination.NewFromRequest(ctx.Request(), -1)

	opt := &options.FindOptions{}
	if pg != nil {

		opt.SetBatchSize(int32(pg.PerPage))
		opt.SetSkip(int64(pg.Offset()))
		opt.SetLimit(int64(pg.Limit()))
	}
	if pg != nil && len(pg.Sort) > 0 {
		intSortBy := 1
		if pg.SortBy == "desc" {
			intSortBy = -1
		}
		opt.SetSort(bson.D{{pg.Sort, intSortBy}})
	}

	username := ctx.Param("user")
	if len(username) == 0 {
		return ctx.JSON(http.StatusBadGateway, "Username required.")
	}
	result := []*chess.Game{}
	cur, err := receiver.gameCollection.Find(c, bson.M{"username": username}, opt)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	defer cur.Close(c)
	if err = cur.All(c, &result); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := cur.Err(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := []UserGames{}
	for _, item := range result {
		response = append(response, UserGames{
			Id:       item.Id,
			Username: item.Username,
			White:    item.White,
			Black:    item.Black,
			Date:     item.Date,
			Result:   item.Result,
		})
	}
	return ctx.JSON(http.StatusOK, response)
}
