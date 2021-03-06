// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3dsim/simulation-goclient/models"
)

// GetAssumedStrainSimulationsReader is a Reader for the GetAssumedStrainSimulations structure.
type GetAssumedStrainSimulationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAssumedStrainSimulationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAssumedStrainSimulationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAssumedStrainSimulationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAssumedStrainSimulationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetAssumedStrainSimulationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAssumedStrainSimulationsOK creates a GetAssumedStrainSimulationsOK with default headers values
func NewGetAssumedStrainSimulationsOK() *GetAssumedStrainSimulationsOK {
	return &GetAssumedStrainSimulationsOK{}
}

/*GetAssumedStrainSimulationsOK handles this case with default header values.

Successfully found the list of items
*/
type GetAssumedStrainSimulationsOK struct {
	Payload []*models.AssumedStrainSimulation
}

func (o *GetAssumedStrainSimulationsOK) Error() string {
	return fmt.Sprintf("[GET /assumedstrainsimulations][%d] getAssumedStrainSimulationsOK  %+v", 200, o.Payload)
}

func (o *GetAssumedStrainSimulationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssumedStrainSimulationsUnauthorized creates a GetAssumedStrainSimulationsUnauthorized with default headers values
func NewGetAssumedStrainSimulationsUnauthorized() *GetAssumedStrainSimulationsUnauthorized {
	return &GetAssumedStrainSimulationsUnauthorized{}
}

/*GetAssumedStrainSimulationsUnauthorized handles this case with default header values.

Not authorized
*/
type GetAssumedStrainSimulationsUnauthorized struct {
	Payload *models.Error
}

func (o *GetAssumedStrainSimulationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /assumedstrainsimulations][%d] getAssumedStrainSimulationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAssumedStrainSimulationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssumedStrainSimulationsForbidden creates a GetAssumedStrainSimulationsForbidden with default headers values
func NewGetAssumedStrainSimulationsForbidden() *GetAssumedStrainSimulationsForbidden {
	return &GetAssumedStrainSimulationsForbidden{}
}

/*GetAssumedStrainSimulationsForbidden handles this case with default header values.

Forbidden
*/
type GetAssumedStrainSimulationsForbidden struct {
	Payload *models.Error
}

func (o *GetAssumedStrainSimulationsForbidden) Error() string {
	return fmt.Sprintf("[GET /assumedstrainsimulations][%d] getAssumedStrainSimulationsForbidden  %+v", 403, o.Payload)
}

func (o *GetAssumedStrainSimulationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssumedStrainSimulationsDefault creates a GetAssumedStrainSimulationsDefault with default headers values
func NewGetAssumedStrainSimulationsDefault(code int) *GetAssumedStrainSimulationsDefault {
	return &GetAssumedStrainSimulationsDefault{
		_statusCode: code,
	}
}

/*GetAssumedStrainSimulationsDefault handles this case with default header values.

unexpected error
*/
type GetAssumedStrainSimulationsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get assumed strain simulations default response
func (o *GetAssumedStrainSimulationsDefault) Code() int {
	return o._statusCode
}

func (o *GetAssumedStrainSimulationsDefault) Error() string {
	return fmt.Sprintf("[GET /assumedstrainsimulations][%d] getAssumedStrainSimulations default  %+v", o._statusCode, o.Payload)
}

func (o *GetAssumedStrainSimulationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
