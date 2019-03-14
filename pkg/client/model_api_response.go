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

type ModelApiResponse struct {
	Code int32 `json:"code,omitempty"`
	Type_ string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}