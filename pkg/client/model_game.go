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

import (
	"time"
)

type Game struct {
	GameId int32 `json:"game_id,omitempty"`
	DateStart time.Time `json:"date_start,omitempty"`
	GameNumber int32 `json:"game_number,omitempty"`
	Week int32 `json:"week,omitempty"`
	Season int32 `json:"season,omitempty"`
	Attendance int32 `json:"attendance,omitempty"`
	EventType *EventType `json:"event_type,omitempty"`
	EventStatus *EventStatus `json:"event_status,omitempty"`
	Venue *Venue `json:"venue,omitempty"`
	Weather *Weather `json:"weather,omitempty"`
	CoinToss *CoinToss `json:"coin_toss,omitempty"`
	TicketsUrl string `json:"tickets_url,omitempty"`
}
