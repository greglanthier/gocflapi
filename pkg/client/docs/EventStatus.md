# EventStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventStatusId** | **int32** | An integer value indicating the current status of the game. Possible values are: 1: Pre-Game 2: In-Progress 4: Final 9: Cancelled  | [optional] 
**Name** | **string** |  | [optional] 
**IsActive** | **bool** | A Boolean value indicating if play is currently live in this game. If play is live at the time of the API request, the value returned is true; else it is false.  | [optional] 
**Quarter** | **int32** | If this game is in progress, contains an integer value indicating the current quarter of play. Contains null if the game is not in play.  | [optional] 
**Minutes** | **int32** | If this game is in progress, contains an integer value indicating the last recorded minute value of the game clock. Contains null if the game is not in play.  | [optional] 
**Seconds** | **int32** | If this game is in progress, contains an integer value indicating the last recorded second value of the game clock. Contains null if the game is not in play.  | [optional] 
**Down** | **int32** | If this game is in progress, contains an integer value indicating the down of the last recorded play. Contains null if the game is not in play.  | [optional] 
**YardsToGo** | **int32** | If this game is in progress, contains an integer value indicating the last recorded number of yards to a first down. Contains null if the game is not in play.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


