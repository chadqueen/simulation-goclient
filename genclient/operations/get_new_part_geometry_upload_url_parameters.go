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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNewPartGeometryUploadURLParams creates a new GetNewPartGeometryUploadURLParams object
// with the default values initialized.
func NewGetNewPartGeometryUploadURLParams() *GetNewPartGeometryUploadURLParams {

	return &GetNewPartGeometryUploadURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNewPartGeometryUploadURLParamsWithTimeout creates a new GetNewPartGeometryUploadURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNewPartGeometryUploadURLParamsWithTimeout(timeout time.Duration) *GetNewPartGeometryUploadURLParams {

	return &GetNewPartGeometryUploadURLParams{

		timeout: timeout,
	}
}

// NewGetNewPartGeometryUploadURLParamsWithContext creates a new GetNewPartGeometryUploadURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNewPartGeometryUploadURLParamsWithContext(ctx context.Context) *GetNewPartGeometryUploadURLParams {

	return &GetNewPartGeometryUploadURLParams{

		Context: ctx,
	}
}

// NewGetNewPartGeometryUploadURLParamsWithHTTPClient creates a new GetNewPartGeometryUploadURLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNewPartGeometryUploadURLParamsWithHTTPClient(client *http.Client) *GetNewPartGeometryUploadURLParams {

	return &GetNewPartGeometryUploadURLParams{
		HTTPClient: client,
	}
}

/*GetNewPartGeometryUploadURLParams contains all the parameters to send to the API endpoint
for the get new part geometry upload Url operation typically these are written to a http.Request
*/
type GetNewPartGeometryUploadURLParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get new part geometry upload Url params
func (o *GetNewPartGeometryUploadURLParams) WithTimeout(timeout time.Duration) *GetNewPartGeometryUploadURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get new part geometry upload Url params
func (o *GetNewPartGeometryUploadURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get new part geometry upload Url params
func (o *GetNewPartGeometryUploadURLParams) WithContext(ctx context.Context) *GetNewPartGeometryUploadURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get new part geometry upload Url params
func (o *GetNewPartGeometryUploadURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get new part geometry upload Url params
func (o *GetNewPartGeometryUploadURLParams) WithHTTPClient(client *http.Client) *GetNewPartGeometryUploadURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get new part geometry upload Url params
func (o *GetNewPartGeometryUploadURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetNewPartGeometryUploadURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
