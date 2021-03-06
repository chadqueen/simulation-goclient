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

// NewStartThermalSimulationParams creates a new StartThermalSimulationParams object
// with the default values initialized.
func NewStartThermalSimulationParams() *StartThermalSimulationParams {
	var ()
	return &StartThermalSimulationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStartThermalSimulationParamsWithTimeout creates a new StartThermalSimulationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStartThermalSimulationParamsWithTimeout(timeout time.Duration) *StartThermalSimulationParams {
	var ()
	return &StartThermalSimulationParams{

		timeout: timeout,
	}
}

// NewStartThermalSimulationParamsWithContext creates a new StartThermalSimulationParams object
// with the default values initialized, and the ability to set a context for a request
func NewStartThermalSimulationParamsWithContext(ctx context.Context) *StartThermalSimulationParams {
	var ()
	return &StartThermalSimulationParams{

		Context: ctx,
	}
}

// NewStartThermalSimulationParamsWithHTTPClient creates a new StartThermalSimulationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStartThermalSimulationParamsWithHTTPClient(client *http.Client) *StartThermalSimulationParams {
	var ()
	return &StartThermalSimulationParams{
		HTTPClient: client,
	}
}

/*StartThermalSimulationParams contains all the parameters to send to the API endpoint
for the start thermal simulation operation typically these are written to a http.Request
*/
type StartThermalSimulationParams struct {

	/*ID
	  simulation identifier

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the start thermal simulation params
func (o *StartThermalSimulationParams) WithTimeout(timeout time.Duration) *StartThermalSimulationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start thermal simulation params
func (o *StartThermalSimulationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start thermal simulation params
func (o *StartThermalSimulationParams) WithContext(ctx context.Context) *StartThermalSimulationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start thermal simulation params
func (o *StartThermalSimulationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start thermal simulation params
func (o *StartThermalSimulationParams) WithHTTPClient(client *http.Client) *StartThermalSimulationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start thermal simulation params
func (o *StartThermalSimulationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the start thermal simulation params
func (o *StartThermalSimulationParams) WithID(id int32) *StartThermalSimulationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the start thermal simulation params
func (o *StartThermalSimulationParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *StartThermalSimulationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
