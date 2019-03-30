/*
 * CFL API
 *
 * This is an attempt to encode the CFL API using Swagger Spec.
 *
 * API version: 1.32
 * Contact: tech@cfl.ca
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type GamesApiService service

/*
GamesApiService Get data for a specific game
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param season Season to retrieve
 * @param gameId Specific game to retrieve
 * @param optional nil or *GetGameByIdOpts - Optional Parameters:
 * @param "Include" (optional.Interface of Include) -  Additional game data can be optionally accessed through use of the include parameter. To get both the box score and play-by-play data in a single API request, pass both values separated by a comma: include=boxscore,play_by_play
 * @param "Sort" (optional.Interface of Sort) -  To specify how to sort the games returned, use the sort parameter. Sorting by multiple parameters can be accomplished by specifying multiple fields separated by a comma. The direction of sorting is ascending by default; to sort in descending order, prefix the field with a minus (U+002D HYPHEN-MINUS, \"-\") character.
 * @param "Filter" (optional.Interface of Filter) -
 * @param "Page" (optional.Interface of Page) -  To specify how many and which games should be returned in an API request, use the page[number] and page[size] parameters. By default, 20 games from page 1 of the result set will be returned. If a specific season of games has been requested, all games are always returned.
@return Games
*/

type GetGameByIdOpts struct {
	Include optional.Interface
	Sort    optional.Interface
	Filter  optional.Interface
	Page    optional.Interface
}

func (a *GamesApiService) GetGameById(ctx context.Context, season int32, gameId int32, localVarOptionals *GetGameByIdOpts) (Games, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Games
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/games/{season}/game/{gameId}"
	localVarPath = strings.Replace(localVarPath, "{"+"season"+"}", fmt.Sprintf("%v", season), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"gameId"+"}", fmt.Sprintf("%v", gameId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Include.IsSet() {
		localVarQueryParams.Add("include", parameterToString(localVarOptionals.Include.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), ""))
	}

	// -- begin manually customized code ----------------------------
	if localVarOptionals != nil && localVarOptionals.Filter.IsSet() {
		if v, ok := localVarOptionals.Filter.Value().(Filter); ok {
			localVarQueryParams.Add(fmt.Sprintf("filter[%s][%s]", v.Field, v.Operator), v.Operator)
		}
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		if v, ok := localVarOptionals.Page.Value().(Page); ok {
			if v.Number != 0 {
				localVarQueryParams.Add("page[number]", fmt.Sprintf("%d", v.Number))
			}
			if v.Size != 0 {
				localVarQueryParams.Add("page[size]", fmt.Sprintf("%d", v.Size))
			}
		}
	}
	// -- end manually customized code ------------------------------

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarQueryParams.Add("key", key)
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v Games
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
GamesApiService Get a list of all games in a particular season
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param season Season to retrieve
 * @param optional nil or *GetGamesOpts - Optional Parameters:
 * @param "Include" (optional.Interface of Include) -  Additional game data can be optionally accessed through use of the include parameter. To get both the box score and play-by-play data in a single API request, pass both values separated by a comma: include=boxscore,play_by_play
 * @param "Sort" (optional.Interface of Sort) -  To specify how to sort the games returned, use the sort parameter. Sorting by multiple parameters can be accomplished by specifying multiple fields separated by a comma. The direction of sorting is ascending by default; to sort in descending order, prefix the field with a minus (U+002D HYPHEN-MINUS, \"-\") character.
 * @param "Filter" (optional.Interface of Filter) -
 * @param "Page" (optional.Interface of Page) -  To specify how many and which games should be returned in an API request, use the page[number] and page[size] parameters. By default, 20 games from page 1 of the result set will be returned. If a specific season of games has been requested, all games are always returned.
@return Games
*/

type GetGamesOpts struct {
	Include optional.Interface
	Sort    optional.Interface
	Filter  optional.Interface
	Page    optional.Interface
}

func (a *GamesApiService) GetGames(ctx context.Context, season int32, localVarOptionals *GetGamesOpts) (Games, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Games
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/games/{season}"
	localVarPath = strings.Replace(localVarPath, "{"+"season"+"}", fmt.Sprintf("%v", season), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Include.IsSet() {
		localVarQueryParams.Add("include", parameterToString(localVarOptionals.Include.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), ""))
	}

	// -- begin manually customized code ----------------------------
	if localVarOptionals != nil && localVarOptionals.Filter.IsSet() {
		if v, ok := localVarOptionals.Filter.Value().(Filter); ok {
			localVarQueryParams.Add(fmt.Sprintf("filter[%s][%s]", v.Field, v.Operator), v.Operator)
		}
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		if v, ok := localVarOptionals.Page.Value().(Page); ok {
			if v.Number != 0 {
				localVarQueryParams.Add("page[number]", fmt.Sprintf("%d", v.Number))
			}
			if v.Size != 0 {
				localVarQueryParams.Add("page[size]", fmt.Sprintf("%d", v.Size))
			}
		}
	}
	// -- end manually customized code ------------------------------

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarQueryParams.Add("key", key)
		}
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v Games
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
