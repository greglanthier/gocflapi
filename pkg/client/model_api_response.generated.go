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

type ApiResponse struct {
	Code int32 `json:"code,omitempty"`
	Type string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}
