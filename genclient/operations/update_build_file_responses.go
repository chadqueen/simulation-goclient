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

// UpdateBuildFileReader is a Reader for the UpdateBuildFile structure.
type UpdateBuildFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBuildFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateBuildFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewUpdateBuildFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateBuildFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateBuildFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateBuildFileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateBuildFileOK creates a UpdateBuildFileOK with default headers values
func NewUpdateBuildFileOK() *UpdateBuildFileOK {
	return &UpdateBuildFileOK{}
}

/*UpdateBuildFileOK handles this case with default header values.

Updated build file
*/
type UpdateBuildFileOK struct {
	Payload *models.BuildFile
}

func (o *UpdateBuildFileOK) Error() string {
	return fmt.Sprintf("[PUT /buildfiles/{id}][%d] updateBuildFileOK  %+v", 200, o.Payload)
}

func (o *UpdateBuildFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBuildFileUnauthorized creates a UpdateBuildFileUnauthorized with default headers values
func NewUpdateBuildFileUnauthorized() *UpdateBuildFileUnauthorized {
	return &UpdateBuildFileUnauthorized{}
}

/*UpdateBuildFileUnauthorized handles this case with default header values.

Not authorized
*/
type UpdateBuildFileUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateBuildFileUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /buildfiles/{id}][%d] updateBuildFileUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateBuildFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBuildFileForbidden creates a UpdateBuildFileForbidden with default headers values
func NewUpdateBuildFileForbidden() *UpdateBuildFileForbidden {
	return &UpdateBuildFileForbidden{}
}

/*UpdateBuildFileForbidden handles this case with default header values.

Forbidden
*/
type UpdateBuildFileForbidden struct {
	Payload *models.Error
}

func (o *UpdateBuildFileForbidden) Error() string {
	return fmt.Sprintf("[PUT /buildfiles/{id}][%d] updateBuildFileForbidden  %+v", 403, o.Payload)
}

func (o *UpdateBuildFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBuildFileNotFound creates a UpdateBuildFileNotFound with default headers values
func NewUpdateBuildFileNotFound() *UpdateBuildFileNotFound {
	return &UpdateBuildFileNotFound{}
}

/*UpdateBuildFileNotFound handles this case with default header values.

build file not found
*/
type UpdateBuildFileNotFound struct {
}

func (o *UpdateBuildFileNotFound) Error() string {
	return fmt.Sprintf("[PUT /buildfiles/{id}][%d] updateBuildFileNotFound ", 404)
}

func (o *UpdateBuildFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBuildFileDefault creates a UpdateBuildFileDefault with default headers values
func NewUpdateBuildFileDefault(code int) *UpdateBuildFileDefault {
	return &UpdateBuildFileDefault{
		_statusCode: code,
	}
}

/*UpdateBuildFileDefault handles this case with default header values.

unexpected error
*/
type UpdateBuildFileDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update build file default response
func (o *UpdateBuildFileDefault) Code() int {
	return o._statusCode
}

func (o *UpdateBuildFileDefault) Error() string {
	return fmt.Sprintf("[PUT /buildfiles/{id}][%d] updateBuildFile default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateBuildFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
