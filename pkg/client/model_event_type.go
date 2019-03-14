/*
 * CFL API
 *
 * This is an attempt to encode the CFL API using Swagger Spec.
 *
 * API version: 1.32
 * Contact: tech@cfl.ca
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type EventType struct {
	// An integer value indicating the type of game played. Possible values are: 0: Preseason 1: Regular Season 2: Playoffs 3: Grey Cup 
	EventTypeId int32 `json:"event_type_id,omitempty"`
	Name string `json:"name,omitempty"`
	// An string value indicating the formal title of the game. Possible values can include: Blank / No value Eastern Semi-Final Western Semi-Final Eastern Final Western Final XXXth Grey Cup 
	Title string `json:"title,omitempty"`
}
