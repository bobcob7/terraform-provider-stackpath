// Code generated by go-swagger; DO NOT EDIT.

package workload_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1InstanceVolumeMount Describes mounting a volume on containers or vms in an instance
//
// swagger:model v1InstanceVolumeMount
type V1InstanceVolumeMount struct {

	// The path in an instance to mount a volume
	MountPath string `json:"mountPath,omitempty"`

	// The slug of the volume claim to mount
	Slug string `json:"slug,omitempty"`
}

// Validate validates this v1 instance volume mount
func (m *V1InstanceVolumeMount) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1InstanceVolumeMount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1InstanceVolumeMount) UnmarshalBinary(b []byte) error {
	var res V1InstanceVolumeMount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
