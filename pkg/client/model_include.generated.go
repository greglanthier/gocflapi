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
type Include string

// List of Include
const (
	BOXSCORE Include = "boxscore"
	PLAY_BY_PLAY Include = "play_by_play"
	ROSTERS Include = "rosters"
	PENALTIES Include = "penalties"
	PLAY_REVIEWS Include = "play_reviews"
)