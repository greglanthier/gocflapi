package gocflapi

import (
	"fmt"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	cflapi "github.com/greglanthier/gocflapi/pkg/client"
	"github.com/greglanthier/gocflapi/pkg/client/games"
	_ "github.com/joho/godotenv/autoload"
)

// Here is an example of using the go-swagger generated client to list all the games in the 2017 season.
func Example() {
	auth := httptransport.APIKeyAuth("key", "query", os.Getenv("CFL_API_KEY"))
	result, err := cflapi.Default.Games.GetGames(games.NewGetGamesParams().WithSeason(2017), auth)
	if err != nil {
		panic(err)
	}
	for _, game := range result.Payload.Data {
		fmt.Println(game)
	}
}
