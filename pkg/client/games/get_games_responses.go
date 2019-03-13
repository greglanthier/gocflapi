// Code generated by go-swagger; DO NOT EDIT.

package games

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/greglanthier/gocflapi/pkg/models"
)

// GetGamesReader is a Reader for the GetGames structure.
type GetGamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetGamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetGamesOK creates a GetGamesOK with default headers values
func NewGetGamesOK() *GetGamesOK {
	return &GetGamesOK{}
}

/*GetGamesOK handles this case with default header values.

Games from seaon
*/
type GetGamesOK struct {
	Payload *models.Games
}

func (o *GetGamesOK) Error() string {
	return fmt.Sprintf("[GET /games/{season}][%d] getGamesOK  %+v", 200, o.Payload)
}

func (o *GetGamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Games)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
