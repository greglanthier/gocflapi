# \GamesApi

All URIs are relative to *https://api.cfl.ca*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGames**](GamesApi.md#GetGames) | **Get** /v1/games/{season} | Get a list of all games in a particular season
[**GetGamesById**](GamesApi.md#GetGamesById) | **Get** /v1/games/{season}/game/{gameId}{?include} | Get data for a specific game


# **GetGames**
> Games GetGames(ctx, season)
Get a list of all games in a particular season

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **season** | **int32**| Season to retrieve | 

### Return type

[**Games**](Games.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGamesById**
> Game GetGamesById(ctx, season, gameId, optional)
Get data for a specific game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **season** | **int32**| Season to retrieve | 
  **gameId** | **int32**| Specific game to retrieve | 
 **optional** | ***GetGamesByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetGamesByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | **optional.String**| Additional game data can be optionally accessed through use of the include parameter. To get both the box score and play-by-play data in a single API request, pass both values separated by a comma: include&#x3D;boxscore,play_by_play  | 

### Return type

[**Game**](Game.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

