// Code generated by go-swagger; DO NOT EDIT.

package workload_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1LogChunk A chunk of raw log data
// swagger:model v1LogChunk
type V1LogChunk struct {

	// Raw log contents
	// Format: byte
	Bytes strfmt.Base64 `json:"bytes,omitempty"`
}

// Validate validates this v1 log chunk
func (m *V1LogChunk) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1LogChunk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1LogChunk) UnmarshalBinary(b []byte) error {
	var res V1LogChunk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
