// Code generated by go-swagger; DO NOT EDIT.

package ipam_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1UpdateNetworkPolicyRequest v1 update network policy request
//
// swagger:model v1UpdateNetworkPolicyRequest
type V1UpdateNetworkPolicyRequest struct {

	// network policy
	NetworkPolicy *V1NetworkPolicy `json:"networkPolicy,omitempty"`
}

// Validate validates this v1 update network policy request
func (m *V1UpdateNetworkPolicyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1UpdateNetworkPolicyRequest) validateNetworkPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkPolicy) { // not required
		return nil
	}

	if m.NetworkPolicy != nil {
		if err := m.NetworkPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkPolicy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 update network policy request based on the context it is used
func (m *V1UpdateNetworkPolicyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetworkPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1UpdateNetworkPolicyRequest) contextValidateNetworkPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkPolicy != nil {
		if err := m.NetworkPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1UpdateNetworkPolicyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UpdateNetworkPolicyRequest) UnmarshalBinary(b []byte) error {
	var res V1UpdateNetworkPolicyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
