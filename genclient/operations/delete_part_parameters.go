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

// NewDeletePartParams creates a new DeletePartParams object
// with the default values initialized.
func NewDeletePartParams() *DeletePartParams {
	var ()
	return &DeletePartParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePartParamsWithTimeout creates a new DeletePartParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePartParamsWithTimeout(timeout time.Duration) *DeletePartParams {
	var ()
	return &DeletePartParams{

		timeout: timeout,
	}
}

// NewDeletePartParamsWithContext creates a new DeletePartParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePartParamsWithContext(ctx context.Context) *DeletePartParams {
	var ()
	return &DeletePartParams{

		Context: ctx,
	}
}

// NewDeletePartParamsWithHTTPClient creates a new DeletePartParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePartParamsWithHTTPClient(client *http.Client) *DeletePartParams {
	var ()
	return &DeletePartParams{
		HTTPClient: client,
	}
}

/*DeletePartParams contains all the parameters to send to the API endpoint
for the delete part operation typically these are written to a http.Request
*/
type DeletePartParams struct {

	/*ID
	  ID of a part

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete part params
func (o *DeletePartParams) WithTimeout(timeout time.Duration) *DeletePartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete part params
func (o *DeletePartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete part params
func (o *DeletePartParams) WithContext(ctx context.Context) *DeletePartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete part params
func (o *DeletePartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete part params
func (o *DeletePartParams) WithHTTPClient(client *http.Client) *DeletePartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete part params
func (o *DeletePartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete part params
func (o *DeletePartParams) WithID(id int32) *DeletePartParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete part params
func (o *DeletePartParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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