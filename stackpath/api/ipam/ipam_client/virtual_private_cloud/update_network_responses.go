// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/stackpath/terraform-provider-stackpath/stackpath/api/ipam/ipam_models"
)

// UpdateNetworkReader is a Reader for the UpdateNetwork structure.
type UpdateNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateNetworkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateNetworkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateNetworkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateNetworkOK creates a UpdateNetworkOK with default headers values
func NewUpdateNetworkOK() *UpdateNetworkOK {
	return &UpdateNetworkOK{}
}

/* UpdateNetworkOK describes a response with status code 200, with default header values.

UpdateNetworkOK update network o k
*/
type UpdateNetworkOK struct {
	Payload *ipam_models.NetworkUpdateNetworkResponse
}

func (o *UpdateNetworkOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/v1/stacks/{stack_id}/networks/{network_id}][%d] updateNetworkOK  %+v", 200, o.Payload)
}
func (o *UpdateNetworkOK) GetPayload() *ipam_models.NetworkUpdateNetworkResponse {
	return o.Payload
}

func (o *UpdateNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.NetworkUpdateNetworkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNetworkUnauthorized creates a UpdateNetworkUnauthorized with default headers values
func NewUpdateNetworkUnauthorized() *UpdateNetworkUnauthorized {
	return &UpdateNetworkUnauthorized{}
}

/* UpdateNetworkUnauthorized describes a response with status code 401, with default header values.

Returned when an unauthorized request is attempted.
*/
type UpdateNetworkUnauthorized struct {
	Payload *ipam_models.APIStatus
}

func (o *UpdateNetworkUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /ipam/v1/stacks/{stack_id}/networks/{network_id}][%d] updateNetworkUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateNetworkUnauthorized) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *UpdateNetworkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNetworkInternalServerError creates a UpdateNetworkInternalServerError with default headers values
func NewUpdateNetworkInternalServerError() *UpdateNetworkInternalServerError {
	return &UpdateNetworkInternalServerError{}
}

/* UpdateNetworkInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type UpdateNetworkInternalServerError struct {
	Payload *ipam_models.APIStatus
}

func (o *UpdateNetworkInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /ipam/v1/stacks/{stack_id}/networks/{network_id}][%d] updateNetworkInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateNetworkInternalServerError) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *UpdateNetworkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNetworkDefault creates a UpdateNetworkDefault with default headers values
func NewUpdateNetworkDefault(code int) *UpdateNetworkDefault {
	return &UpdateNetworkDefault{
		_statusCode: code,
	}
}

/* UpdateNetworkDefault describes a response with status code -1, with default header values.

Default error structure.
*/
type UpdateNetworkDefault struct {
	_statusCode int

	Payload *ipam_models.APIStatus
}

// Code gets the status code for the update network default response
func (o *UpdateNetworkDefault) Code() int {
	return o._statusCode
}

func (o *UpdateNetworkDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/v1/stacks/{stack_id}/networks/{network_id}][%d] UpdateNetwork default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateNetworkDefault) GetPayload() *ipam_models.APIStatus {
	return o.Payload
}

func (o *UpdateNetworkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ipam_models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}