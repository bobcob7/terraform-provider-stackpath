// Code generated by go-swagger; DO NOT EDIT.

package instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new instances API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for instances API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetWorkloadInstances(params *GetWorkloadInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWorkloadInstancesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetWorkloadInstances retrieves a workload s instances
*/
func (a *Client) GetWorkloadInstances(params *GetWorkloadInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWorkloadInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkloadInstancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetWorkloadInstances",
		Method:             "GET",
		PathPattern:        "/workload/v1/stacks/{stack_id}/workloads/{workload_id}/instances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWorkloadInstancesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWorkloadInstancesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetWorkloadInstancesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
