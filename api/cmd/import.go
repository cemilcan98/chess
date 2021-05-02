package cmd

import (
	"context"
	"fmt"
	"github.com/cemilcan98/chess/internal/chess"
	"github.com/cemilcan98/chess/pkg/log"
	mongohelper "github.com/cemilcan98/chess/pkg/mongoextentions"
	"github.com/cemilcan98/chess/pkg/pgn"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import played games to database",
}
var (
	dbconn  = "mongodb://root:example@127.0.0.1:27017/"
	dbName  = "chess"
	mongoDb *mongo.Database
)

func init() {
	rootCmd.AddCommand(importCmd)
	log.SetupLogger()
	file := ""
	importCmd.Flags().StringVarP(&file, "file", "f", "", "select file")
	importCmd.Run = func(cmd *cobra.Command, args []string) {
		var err error
		mongoDb, err = mongohelper.NewDatabase(dbconn, dbName)
		if err != nil {
			log.Logger.Fatalf("Failed to connect database. Error :%s", err.Error())
		}
		parse(file)

	}
}
func parse(file string) {

	game_collection := mongoDb.Collection("games")

	f, err := os.Open(file)
	if err != nil {
		log.Logger.Fatal(err)
	}
	ps := pgn.NewPGNScanner(f)
	for ps.Next() {

		game, err := ps.Scan()
		if err != nil {
			log.Logger.Fatal(err)
		}

		mGame := chess.Game{
			Id:          uuid.New().String(),
			Username:    "cemilcan",
			Site:        setValue(game.Tags, "Site"),
			BlackElo:    setValue(game.Tags, "BlackElo"),
			BlackTitle:  setValue(game.Tags, "BlackTitle"),
			ECO:         setValue(game.Tags, "ECO"),
			White:       setValue(game.Tags, "White"),
			Black:       setValue(game.Tags, "Black"),
			EventDate:   setValue(game.Tags, "EventDate"),
			Variation:   setValue(game.Tags, "Variation"),
			Date:        setValue(game.Tags, "Date"),
			BlackTeam:   setValue(game.Tags, "BlackTeam"),
			WhiteTeam:   setValue(game.Tags, "WhiteTeam"),
			Event:       setValue(game.Tags, "Event"),
			Round:       setValue(game.Tags, "Round"),
			Result:      setValue(game.Tags, "Result"),
			BlackFideID: setValue(game.Tags, "BlackFideID"),
			Opening:     setValue(game.Tags, "Opening"),
			WhiteElo:    setValue(game.Tags, "WhiteElo"),
			WhiteFideID: setValue(game.Tags, "WhiteFideID"),
			WhiteTitle:  setValue(game.Tags, "WhiteTitle"),
			Moves:       []chess.Move{},
		}
		fmt.Println(game.Tags)
		b := pgn.NewBoard()
		for _, move := range game.Moves {

			fmt.Printf("%v\n", move)
			mMove := chess.Move{
				From:    move.From.String(),
				To:      move.To.String(),
				Promote: byte(move.Promote),
				San:     move.San,
			}
			b.MakeMove(move)
			mMove.Fen = b.String()
			mGame.AddMove(mMove)
		}

		game_collection.InsertOne(context.Background(), mGame)
	}
}

func setValue(m map[string]string, f string) string {

	if value, ok := m[f]; ok {
		return value
	}
	return ""
}
