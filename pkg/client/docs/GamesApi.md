# \GamesApi

All URIs are relative to *https://api.cfl.ca/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGames**](GamesApi.md#GetGames) | **Get** /games/{season} | Get a list of all games in a particular season


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
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

