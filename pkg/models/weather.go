// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Weather weather
// swagger:model Weather
type Weather struct {

	// field conditions
	FieldConditions string `json:"field_conditions,omitempty"`

	// sky
	Sky string `json:"sky,omitempty"`

	// An integer value indicating the temperature taken at the stadium / venue the game is held at. Note that this value can be zero or negative.
	//
	Temperature int32 `json:"temperature,omitempty"`

	// wind direction
	WindDirection string `json:"wind_direction,omitempty"`

	// wind speed
	WindSpeed string `json:"wind_speed,omitempty"`
}

// Validate validates this weather
func (m *Weather) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Weather) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Weather) UnmarshalBinary(b []byte) error {
	var res Weather
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
