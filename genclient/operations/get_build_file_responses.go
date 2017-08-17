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

// GetBuildFileReader is a Reader for the GetBuildFile structure.
type GetBuildFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBuildFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetBuildFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetBuildFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetBuildFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetBuildFileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBuildFileOK creates a GetBuildFileOK with default headers values
func NewGetBuildFileOK() *GetBuildFileOK {
	return &GetBuildFileOK{}
}

/*GetBuildFileOK handles this case with default header values.

Retrieved build file
*/
type GetBuildFileOK struct {
	Payload *models.BuildFile
}

func (o *GetBuildFileOK) Error() string {
	return fmt.Sprintf("[GET /buildfiles/{id}][%d] getBuildFileOK  %+v", 200, o.Payload)
}

func (o *GetBuildFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildFileUnauthorized creates a GetBuildFileUnauthorized with default headers values
func NewGetBuildFileUnauthorized() *GetBuildFileUnauthorized {
	return &GetBuildFileUnauthorized{}
}

/*GetBuildFileUnauthorized handles this case with default header values.

Not authorized
*/
type GetBuildFileUnauthorized struct {
	Payload *models.Error
}

func (o *GetBuildFileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /buildfiles/{id}][%d] getBuildFileUnauthorized  %+v", 401, o.Payload)
}

func (o *GetBuildFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildFileForbidden creates a GetBuildFileForbidden with default headers values
func NewGetBuildFileForbidden() *GetBuildFileForbidden {
	return &GetBuildFileForbidden{}
}

/*GetBuildFileForbidden handles this case with default header values.

Forbidden
*/
type GetBuildFileForbidden struct {
	Payload *models.Error
}

func (o *GetBuildFileForbidden) Error() string {
	return fmt.Sprintf("[GET /buildfiles/{id}][%d] getBuildFileForbidden  %+v", 403, o.Payload)
}

func (o *GetBuildFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildFileNotFound creates a GetBuildFileNotFound with default headers values
func NewGetBuildFileNotFound() *GetBuildFileNotFound {
	return &GetBuildFileNotFound{}
}

/*GetBuildFileNotFound handles this case with default header values.

build file not found
*/
type GetBuildFileNotFound struct {
}

func (o *GetBuildFileNotFound) Error() string {
	return fmt.Sprintf("[GET /buildfiles/{id}][%d] getBuildFileNotFound ", 404)
}

func (o *GetBuildFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBuildFileDefault creates a GetBuildFileDefault with default headers values
func NewGetBuildFileDefault(code int) *GetBuildFileDefault {
	return &GetBuildFileDefault{
		_statusCode: code,
	}
}

/*GetBuildFileDefault handles this case with default header values.

unexpected error
*/
type GetBuildFileDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get build file default response
func (o *GetBuildFileDefault) Code() int {
	return o._statusCode
}

func (o *GetBuildFileDefault) Error() string {
	return fmt.Sprintf("[GET /buildfiles/{id}][%d] getBuildFile default  %+v", o._statusCode, o.Payload)
}

func (o *GetBuildFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
