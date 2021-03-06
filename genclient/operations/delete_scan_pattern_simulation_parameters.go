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

// NewDeleteScanPatternSimulationParams creates a new DeleteScanPatternSimulationParams object
// with the default values initialized.
func NewDeleteScanPatternSimulationParams() *DeleteScanPatternSimulationParams {
	var ()
	return &DeleteScanPatternSimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteScanPatternSimulationParamsWithTimeout creates a new DeleteScanPatternSimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteScanPatternSimulationParamsWithTimeout(timeout time.Duration) *DeleteScanPatternSimulationParams {
	var ()
	return &DeleteScanPatternSimulationParams{

		timeout: timeout,
	}
}

// NewDeleteScanPatternSimulationParamsWithContext creates a new DeleteScanPatternSimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteScanPatternSimulationParamsWithContext(ctx context.Context) *DeleteScanPatternSimulationParams {
	var ()
	return &DeleteScanPatternSimulationParams{

		Context: ctx,
	}
}

// NewDeleteScanPatternSimulationParamsWithHTTPClient creates a new DeleteScanPatternSimulationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteScanPatternSimulationParamsWithHTTPClient(client *http.Client) *DeleteScanPatternSimulationParams {
	var ()
	return &DeleteScanPatternSimulationParams{
		HTTPClient: client,
	}
}

/*DeleteScanPatternSimulationParams contains all the parameters to send to the API endpoint
for the delete scan pattern simulation operation typically these are written to a http.Request
*/
type DeleteScanPatternSimulationParams struct {

	/*ID
	  ID of simulation

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete scan pattern simulation params
func (o *DeleteScanPatternSimulationParams) WithTimeout(timeout time.Duration) *DeleteScanPatternSimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete scan pattern simulation params
func (o *DeleteScanPatternSimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete scan pattern simulation params
func (o *DeleteScanPatternSimulationParams) WithContext(ctx context.Context) *DeleteScanPatternSimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete scan pattern simulation params
func (o *DeleteScanPatternSimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete scan pattern simulation params
func (o *DeleteScanPatternSimulationParams) WithHTTPClient(client *http.Client) *DeleteScanPatternSimulationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete scan pattern simulation params
func (o *DeleteScanPatternSimulationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete scan pattern simulation params
func (o *DeleteScanPatternSimulationParams) WithID(id int32) *DeleteScanPatternSimulationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete scan pattern simulation params
func (o *DeleteScanPatternSimulationParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteScanPatternSimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
