// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1TargetSpec The specification for a target
// swagger:model v1TargetSpec
type V1TargetSpec struct {

	// The scope at which a deployment should be created. Valid values are: "cityCode"
	DeploymentScope string `json:"deploymentScope,omitempty"`

	// deployments
	Deployments *V1DeploymentSpec `json:"deployments,omitempty"`
}

// Validate validates this v1 target spec
func (m *V1TargetSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeployments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1TargetSpec) validateDeployments(formats strfmt.Registry) error {

	if swag.IsZero(m.Deployments) { // not required
		return nil
	}

	if m.Deployments != nil {
		if err := m.Deployments.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployments")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1TargetSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1TargetSpec) UnmarshalBinary(b []byte) error {
	var res V1TargetSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}