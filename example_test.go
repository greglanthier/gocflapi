// Copyright 2019 Greg Lanthier
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gocflapi_test

import (
	"context"
	"fmt"
	"os"
	"time"

	cflapi "github.com/greglanthier/gocflapi/pkg/client"
	_ "github.com/joho/godotenv/autoload" // load environment variable
	// _ "github.com/motemen/go-loghttp/global"
)

// Here is an example of using the swagger-codegen generated client to retrieve game data for all the games in the 2018 season.
func Example() {
	// import (
	// 	"context"
	// 	"fmt"
	// 	"os"
	// 	"time"

	// 	cflapi "github.com/greglanthier/gocflapi/pkg/client"
	// 	_ "github.com/joho/godotenv/autoload" // load environment variable
	// 	// _ "github.com/motemen/go-loghttp/global" // Uncomment for degug logging of HTTP traffic.
	// )

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1000*time.Millisecond)
	defer cancel()

	ctx = context.WithValue(ctx, cflapi.ContextAPIKey, cflapi.APIKey{Key: os.Getenv("CFL_API_KEY")})
	config := cflapi.NewConfiguration()
	games, _, err := cflapi.NewAPIClient(config).GamesApi.GetGames(ctx, 2018)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(games.Data))
	// Output: 95
}
