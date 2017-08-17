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

// NewGetThermalSimulationsParams creates a new GetThermalSimulationsParams object
// with the default values initialized.
func NewGetThermalSimulationsParams() *GetThermalSimulationsParams {
	var ()
	return &GetThermalSimulationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetThermalSimulationsParamsWithTimeout creates a new GetThermalSimulationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetThermalSimulationsParamsWithTimeout(timeout time.Duration) *GetThermalSimulationsParams {
	var ()
	return &GetThermalSimulationsParams{

		timeout: timeout,
	}
}

// NewGetThermalSimulationsParamsWithContext creates a new GetThermalSimulationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetThermalSimulationsParamsWithContext(ctx context.Context) *GetThermalSimulationsParams {
	var ()
	return &GetThermalSimulationsParams{

		Context: ctx,
	}
}

// NewGetThermalSimulationsParamsWithHTTPClient creates a new GetThermalSimulationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetThermalSimulationsParamsWithHTTPClient(client *http.Client) *GetThermalSimulationsParams {
	var ()
	return &GetThermalSimulationsParams{
		HTTPClient: client,
	}
}

/*GetThermalSimulationsParams contains all the parameters to send to the API endpoint
for the get thermal simulations operation typically these are written to a http.Request
*/
type GetThermalSimulationsParams struct {

	/*Limit
	  number of items to return within the query

	*/
	Limit *int32
	/*Offset
	  starting paging count; ex. offset of 60 will skip the first 60 items in the list

	*/
	Offset *int32
	/*OrganizationID
	  the organization id to get items for.  Must be provided as API callers only have access to items belonging to their organization.

	*/
	OrganizationID int32
	/*Sort
	  key:direction pairs for one or multiple field sort orders.  e.g. sort=key1:desc,key2:asc

	*/
	Sort []string
	/*Status
	  simulation status for items retrieved.  If an array of items is sent, they are treated as "OR" operations. e.g. status=InProgress,Requested would yield a list of simulations that are in either state.

	*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get thermal simulations params
func (o *GetThermalSimulationsParams) WithTimeout(timeout time.Duration) *GetThermalSimulationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get thermal simulations params
func (o *GetThermalSimulationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get thermal simulations params
func (o *GetThermalSimulationsParams) WithContext(ctx context.Context) *GetThermalSimulationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get thermal simulations params
func (o *GetThermalSimulationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get thermal simulations params
func (o *GetThermalSimulationsParams) WithHTTPClient(client *http.Client) *GetThermalSimulationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get thermal simulations params
func (o *GetThermalSimulationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get thermal simulations params
func (o *GetThermalSimulationsParams) WithLimit(limit *int32) *GetThermalSimulationsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get thermal simulations params
func (o *GetThermalSimulationsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get thermal simulations params
func (o *GetThermalSimulationsParams) WithOffset(offset *int32) *GetThermalSimulationsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get thermal simulations params
func (o *GetThermalSimulationsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the get thermal simulations params
func (o *GetThermalSimulationsParams) WithOrganizationID(organizationID int32) *GetThermalSimulationsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get thermal simulations params
func (o *GetThermalSimulationsParams) SetOrganizationID(organizationID int32) {
	o.OrganizationID = organizationID
}

// WithSort adds the sort to the get thermal simulations params
func (o *GetThermalSimulationsParams) WithSort(sort []string) *GetThermalSimulationsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get thermal simulations params
func (o *GetThermalSimulationsParams) SetSort(sort []string) {
	o.Sort = sort
}

// WithStatus adds the status to the get thermal simulations params
func (o *GetThermalSimulationsParams) WithStatus(status []string) *GetThermalSimulationsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get thermal simulations params
func (o *GetThermalSimulationsParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetThermalSimulationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	// query param organizationId
	qrOrganizationID := o.OrganizationID
	qOrganizationID := swag.FormatInt32(qrOrganizationID)
	if qOrganizationID != "" {
		if err := r.SetQueryParam("organizationId", qOrganizationID); err != nil {
			return err
		}
	}

	valuesSort := o.Sort

	joinedSort := swag.JoinByFormat(valuesSort, "csv")
	// query array param sort
	if err := r.SetQueryParam("sort", joinedSort...); err != nil {
		return err
	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
