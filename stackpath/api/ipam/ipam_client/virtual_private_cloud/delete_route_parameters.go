// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteRouteParams creates a new DeleteRouteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRouteParams() *DeleteRouteParams {
	return &DeleteRouteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRouteParamsWithTimeout creates a new DeleteRouteParams object
// with the ability to set a timeout on a request.
func NewDeleteRouteParamsWithTimeout(timeout time.Duration) *DeleteRouteParams {
	return &DeleteRouteParams{
		timeout: timeout,
	}
}

// NewDeleteRouteParamsWithContext creates a new DeleteRouteParams object
// with the ability to set a context for a request.
func NewDeleteRouteParamsWithContext(ctx context.Context) *DeleteRouteParams {
	return &DeleteRouteParams{
		Context: ctx,
	}
}

// NewDeleteRouteParamsWithHTTPClient creates a new DeleteRouteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRouteParamsWithHTTPClient(client *http.Client) *DeleteRouteParams {
	return &DeleteRouteParams{
		HTTPClient: client,
	}
}

/* DeleteRouteParams contains all the parameters to send to the API endpoint
   for the delete route operation.

   Typically these are written to a http.Request.
*/
type DeleteRouteParams struct {

	/* NetworkID.

	   A VPC network ID or slug
	*/
	NetworkID string

	/* RouteID.

	   A VPC route ID
	*/
	RouteID string

	/* StackID.

	   A stack ID or slug
	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRouteParams) WithDefaults() *DeleteRouteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete route params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRouteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete route params
func (o *DeleteRouteParams) WithTimeout(timeout time.Duration) *DeleteRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete route params
func (o *DeleteRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete route params
func (o *DeleteRouteParams) WithContext(ctx context.Context) *DeleteRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete route params
func (o *DeleteRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete route params
func (o *DeleteRouteParams) WithHTTPClient(client *http.Client) *DeleteRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete route params
func (o *DeleteRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the delete route params
func (o *DeleteRouteParams) WithNetworkID(networkID string) *DeleteRouteParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete route params
func (o *DeleteRouteParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRouteID adds the routeID to the delete route params
func (o *DeleteRouteParams) WithRouteID(routeID string) *DeleteRouteParams {
	o.SetRouteID(routeID)
	return o
}

// SetRouteID adds the routeId to the delete route params
func (o *DeleteRouteParams) SetRouteID(routeID string) {
	o.RouteID = routeID
}

// WithStackID adds the stackID to the delete route params
func (o *DeleteRouteParams) WithStackID(stackID string) *DeleteRouteParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the delete route params
func (o *DeleteRouteParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	// path param route_id
	if err := r.SetPathParam("route_id", o.RouteID); err != nil {
		return err
	}

	// path param stack_id
	if err := r.SetPathParam("stack_id", o.StackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}