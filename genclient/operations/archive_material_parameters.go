// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewArchiveMaterialParams creates a new ArchiveMaterialParams object
// with the default values initialized.
func NewArchiveMaterialParams() *ArchiveMaterialParams {
	var ()
	return &ArchiveMaterialParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewArchiveMaterialParamsWithTimeout creates a new ArchiveMaterialParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewArchiveMaterialParamsWithTimeout(timeout time.Duration) *ArchiveMaterialParams {
	var ()
	return &ArchiveMaterialParams{

		timeout: timeout,
	}
}

// NewArchiveMaterialParamsWithContext creates a new ArchiveMaterialParams object
// with the default values initialized, and the ability to set a context for a request
func NewArchiveMaterialParamsWithContext(ctx context.Context) *ArchiveMaterialParams {
	var ()
	return &ArchiveMaterialParams{

		Context: ctx,
	}
}

// NewArchiveMaterialParamsWithHTTPClient creates a new ArchiveMaterialParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewArchiveMaterialParamsWithHTTPClient(client *http.Client) *ArchiveMaterialParams {
	var ()
	return &ArchiveMaterialParams{
		HTTPClient: client,
	}
}

/*ArchiveMaterialParams contains all the parameters to send to the API endpoint
for the archive material operation typically these are written to a http.Request
*/
type ArchiveMaterialParams struct {

	/*ID
	  ID of material to archive

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the archive material params
func (o *ArchiveMaterialParams) WithTimeout(timeout time.Duration) *ArchiveMaterialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the archive material params
func (o *ArchiveMaterialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the archive material params
func (o *ArchiveMaterialParams) WithContext(ctx context.Context) *ArchiveMaterialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the archive material params
func (o *ArchiveMaterialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the archive material params
func (o *ArchiveMaterialParams) WithHTTPClient(client *http.Client) *ArchiveMaterialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the archive material params
func (o *ArchiveMaterialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the archive material params
func (o *ArchiveMaterialParams) WithID(id int32) *ArchiveMaterialParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the archive material params
func (o *ArchiveMaterialParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ArchiveMaterialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
