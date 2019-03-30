# \GamesApi

All URIs are relative to *https://api.cfl.ca*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGameById**](GamesApi.md#GetGameById) | **Get** /v1/games/{season}/game/{gameId} | Get data for a specific game
[**GetGames**](GamesApi.md#GetGames) | **Get** /v1/games/{season} | Get a list of all games in a particular season


# **GetGameById**
> Games GetGameById(ctx, season, gameId, optional)
Get data for a specific game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **season** | **int32**| Season to retrieve | 
  **gameId** | **int32**| Specific game to retrieve | 
 **optional** | ***GetGameByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetGameByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | [**optional.Interface of Include**](.md)| Additional game data can be optionally accessed through use of the include parameter. To get both the box score and play-by-play data in a single API request, pass both values separated by a comma: include&#x3D;boxscore,play_by_play  | 
 **sort** | [**optional.Interface of Sort**](.md)| To specify how to sort the games returned, use the sort parameter. Sorting by multiple parameters can be accomplished by specifying multiple fields separated by a comma. The direction of sorting is ascending by default; to sort in descending order, prefix the field with a minus (U+002D HYPHEN-MINUS, \&quot;-\&quot;) character.  | 
 **filter** | [**optional.Interface of Filter**](.md)|  | 
 **page** | [**optional.Interface of Page**](.md)| To specify how many and which games should be returned in an API request, use the page[number] and page[size] parameters. By default, 20 games from page 1 of the result set will be returned. If a specific season of games has been requested, all games are always returned.   | 

### Return type

[**Games**](Games.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGames**
> Games GetGames(ctx, season, optional)
Get a list of all games in a particular season

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **season** | **int32**| Season to retrieve | 
 **optional** | ***GetGamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetGamesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | [**optional.Interface of Include**](.md)| Additional game data can be optionally accessed through use of the include parameter. To get both the box score and play-by-play data in a single API request, pass both values separated by a comma: include&#x3D;boxscore,play_by_play  | 
 **sort** | [**optional.Interface of Sort**](.md)| To specify how to sort the games returned, use the sort parameter. Sorting by multiple parameters can be accomplished by specifying multiple fields separated by a comma. The direction of sorting is ascending by default; to sort in descending order, prefix the field with a minus (U+002D HYPHEN-MINUS, \&quot;-\&quot;) character.  | 
 **filter** | [**optional.Interface of Filter**](.md)|  | 
 **page** | [**optional.Interface of Page**](.md)| To specify how many and which games should be returned in an API request, use the page[number] and page[size] parameters. By default, 20 games from page 1 of the result set will be returned. If a specific season of games has been requested, all games are always returned.   | 

### Return type

[**Games**](Games.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

