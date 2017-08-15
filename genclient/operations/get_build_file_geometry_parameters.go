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

// NewGetBuildFileGeometryParams creates a new GetBuildFileGeometryParams object
// with the default values initialized.
func NewGetBuildFileGeometryParams() *GetBuildFileGeometryParams {
	var ()
	return &GetBuildFileGeometryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBuildFileGeometryParamsWithTimeout creates a new GetBuildFileGeometryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBuildFileGeometryParamsWithTimeout(timeout time.Duration) *GetBuildFileGeometryParams {
	var ()
	return &GetBuildFileGeometryParams{

		timeout: timeout,
	}
}

// NewGetBuildFileGeometryParamsWithContext creates a new GetBuildFileGeometryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBuildFileGeometryParamsWithContext(ctx context.Context) *GetBuildFileGeometryParams {
	var ()
	return &GetBuildFileGeometryParams{

		Context: ctx,
	}
}

// NewGetBuildFileGeometryParamsWithHTTPClient creates a new GetBuildFileGeometryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBuildFileGeometryParamsWithHTTPClient(client *http.Client) *GetBuildFileGeometryParams {
	var ()
	return &GetBuildFileGeometryParams{
		HTTPClient: client,
	}
}

/*GetBuildFileGeometryParams contains all the parameters to send to the API endpoint
for the get build file geometry operation typically these are written to a http.Request
*/
type GetBuildFileGeometryParams struct {

	/*ID
	  ID of build file to retrieve geometry of

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get build file geometry params
func (o *GetBuildFileGeometryParams) WithTimeout(timeout time.Duration) *GetBuildFileGeometryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get build file geometry params
func (o *GetBuildFileGeometryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get build file geometry params
func (o *GetBuildFileGeometryParams) WithContext(ctx context.Context) *GetBuildFileGeometryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get build file geometry params
func (o *GetBuildFileGeometryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get build file geometry params
func (o *GetBuildFileGeometryParams) WithHTTPClient(client *http.Client) *GetBuildFileGeometryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get build file geometry params
func (o *GetBuildFileGeometryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get build file geometry params
func (o *GetBuildFileGeometryParams) WithID(id int32) *GetBuildFileGeometryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get build file geometry params
func (o *GetBuildFileGeometryParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetBuildFileGeometryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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