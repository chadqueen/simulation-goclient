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

// ArchiveBuildFileReader is a Reader for the ArchiveBuildFile structure.
type ArchiveBuildFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArchiveBuildFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewArchiveBuildFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewArchiveBuildFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewArchiveBuildFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewArchiveBuildFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewArchiveBuildFileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewArchiveBuildFileOK creates a ArchiveBuildFileOK with default headers values
func NewArchiveBuildFileOK() *ArchiveBuildFileOK {
	return &ArchiveBuildFileOK{}
}

/*ArchiveBuildFileOK handles this case with default header values.

Build file that was archived
*/
type ArchiveBuildFileOK struct {
	Payload *models.BuildFile
}

func (o *ArchiveBuildFileOK) Error() string {
	return fmt.Sprintf("[PUT /buildfiles/{id}/archive][%d] archiveBuildFileOK  %+v", 200, o.Payload)
}

func (o *ArchiveBuildFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveBuildFileUnauthorized creates a ArchiveBuildFileUnauthorized with default headers values
func NewArchiveBuildFileUnauthorized() *ArchiveBuildFileUnauthorized {
	return &ArchiveBuildFileUnauthorized{}
}

/*ArchiveBuildFileUnauthorized handles this case with default header values.

Not authorized
*/
type ArchiveBuildFileUnauthorized struct {
	Payload *models.Error
}

func (o *ArchiveBuildFileUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /buildfiles/{id}/archive][%d] archiveBuildFileUnauthorized  %+v", 401, o.Payload)
}

func (o *ArchiveBuildFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveBuildFileForbidden creates a ArchiveBuildFileForbidden with default headers values
func NewArchiveBuildFileForbidden() *ArchiveBuildFileForbidden {
	return &ArchiveBuildFileForbidden{}
}

/*ArchiveBuildFileForbidden handles this case with default header values.

Forbidden
*/
type ArchiveBuildFileForbidden struct {
	Payload *models.Error
}

func (o *ArchiveBuildFileForbidden) Error() string {
	return fmt.Sprintf("[PUT /buildfiles/{id}/archive][%d] archiveBuildFileForbidden  %+v", 403, o.Payload)
}

func (o *ArchiveBuildFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveBuildFileNotFound creates a ArchiveBuildFileNotFound with default headers values
func NewArchiveBuildFileNotFound() *ArchiveBuildFileNotFound {
	return &ArchiveBuildFileNotFound{}
}

/*ArchiveBuildFileNotFound handles this case with default header values.

build file not found
*/
type ArchiveBuildFileNotFound struct {
}

func (o *ArchiveBuildFileNotFound) Error() string {
	return fmt.Sprintf("[PUT /buildfiles/{id}/archive][%d] archiveBuildFileNotFound ", 404)
}

func (o *ArchiveBuildFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewArchiveBuildFileDefault creates a ArchiveBuildFileDefault with default headers values
func NewArchiveBuildFileDefault(code int) *ArchiveBuildFileDefault {
	return &ArchiveBuildFileDefault{
		_statusCode: code,
	}
}

/*ArchiveBuildFileDefault handles this case with default header values.

unexpected error
*/
type ArchiveBuildFileDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the archive build file default response
func (o *ArchiveBuildFileDefault) Code() int {
	return o._statusCode
}

func (o *ArchiveBuildFileDefault) Error() string {
	return fmt.Sprintf("[PUT /buildfiles/{id}/archive][%d] archiveBuildFile default  %+v", o._statusCode, o.Payload)
}

func (o *ArchiveBuildFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
