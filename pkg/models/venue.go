// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Venue venue
// swagger:model Venue
type Venue struct {

	// The name of the stadium / venue the game is held within.
	//
	Name string `json:"name,omitempty"`

	// A unique numeric value assigned to the stadium / venue the game is held within.
	//
	VenueID int32 `json:"venue_id,omitempty"`
}

// Validate validates this venue
func (m *Venue) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Venue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Venue) UnmarshalBinary(b []byte) error {
	var res Venue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
