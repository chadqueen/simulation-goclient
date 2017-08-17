// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuildFilePost container for parameters used when posting a new build file
// swagger:model BuildFilePost
type BuildFilePost struct {

	// a description of this build file
	Description string `json:"description,omitempty"`

	// path of ZIP file containing build file elements
	// Required: true
	FileUploadLocation *string `json:"fileUploadLocation"`

	// type of machine associated with this build file, see BuildFile definition for valid values
	// Required: true
	MachineType *string `json:"machineType"`

	// a name given to this build file
	// Required: true
	Name *string `json:"name"`

	// identifier for organization this build file belongs to
	// Required: true
	OrganizationID *int64 `json:"organizationId"`

	// a list of tags assigned to this build file
	Tags []string `json:"tags"`
}

// Validate validates this build file post
func (m *BuildFilePost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileUploadLocation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMachineType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildFilePost) validateFileUploadLocation(formats strfmt.Registry) error {

	if err := validate.Required("fileUploadLocation", "body", m.FileUploadLocation); err != nil {
		return err
	}

	return nil
}

func (m *BuildFilePost) validateMachineType(formats strfmt.Registry) error {

	if err := validate.Required("machineType", "body", m.MachineType); err != nil {
		return err
	}

	return nil
}

func (m *BuildFilePost) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *BuildFilePost) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationId", "body", m.OrganizationID); err != nil {
		return err
	}

	return nil
}

func (m *BuildFilePost) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildFilePost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildFilePost) UnmarshalBinary(b []byte) error {
	var res BuildFilePost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
