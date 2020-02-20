// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/terraform-providers/terraform-provider-stackpath/stackpath/api/workload/workload_models"
)

// GetMetricsReader is a Reader for the GetMetrics structure.
type GetMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetMetricsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMetricsOK creates a GetMetricsOK with default headers values
func NewGetMetricsOK() *GetMetricsOK {
	return &GetMetricsOK{}
}

/*GetMetricsOK handles this case with default header values.

GetMetricsOK get metrics o k
*/
type GetMetricsOK struct {
	Payload *workload_models.PrometheusMetrics
}

func (o *GetMetricsOK) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/metrics][%d] getMetricsOK  %+v", 200, o.Payload)
}

func (o *GetMetricsOK) GetPayload() *workload_models.PrometheusMetrics {
	return o.Payload
}

func (o *GetMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.PrometheusMetrics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetricsUnauthorized creates a GetMetricsUnauthorized with default headers values
func NewGetMetricsUnauthorized() *GetMetricsUnauthorized {
	return &GetMetricsUnauthorized{}
}

/*GetMetricsUnauthorized handles this case with default header values.

Returned when an unauthorized request is attempted.
*/
type GetMetricsUnauthorized struct {
	Payload *workload_models.StackpathapiStatus
}

func (o *GetMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/metrics][%d] getMetricsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetMetricsUnauthorized) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetricsInternalServerError creates a GetMetricsInternalServerError with default headers values
func NewGetMetricsInternalServerError() *GetMetricsInternalServerError {
	return &GetMetricsInternalServerError{}
}

/*GetMetricsInternalServerError handles this case with default header values.

Internal server error.
*/
type GetMetricsInternalServerError struct {
	Payload *workload_models.StackpathapiStatus
}

func (o *GetMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/metrics][%d] getMetricsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMetricsInternalServerError) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetricsDefault creates a GetMetricsDefault with default headers values
func NewGetMetricsDefault(code int) *GetMetricsDefault {
	return &GetMetricsDefault{
		_statusCode: code,
	}
}

/*GetMetricsDefault handles this case with default header values.

Default error structure.
*/
type GetMetricsDefault struct {
	_statusCode int

	Payload *workload_models.StackpathapiStatus
}

// Code gets the status code for the get metrics default response
func (o *GetMetricsDefault) Code() int {
	return o._statusCode
}

func (o *GetMetricsDefault) Error() string {
	return fmt.Sprintf("[GET /workload/v1/stacks/{stack_id}/metrics][%d] GetMetrics default  %+v", o._statusCode, o.Payload)
}

func (o *GetMetricsDefault) GetPayload() *workload_models.StackpathapiStatus {
	return o.Payload
}

func (o *GetMetricsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(workload_models.StackpathapiStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
