// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewSendViewProductParams creates a new SendViewProductParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSendViewProductParams() *SendViewProductParams {
	return &SendViewProductParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSendViewProductParamsWithTimeout creates a new SendViewProductParams object
// with the ability to set a timeout on a request.
func NewSendViewProductParamsWithTimeout(timeout time.Duration) *SendViewProductParams {
	return &SendViewProductParams{
		timeout: timeout,
	}
}

// NewSendViewProductParamsWithContext creates a new SendViewProductParams object
// with the ability to set a context for a request.
func NewSendViewProductParamsWithContext(ctx context.Context) *SendViewProductParams {
	return &SendViewProductParams{
		Context: ctx,
	}
}

// NewSendViewProductParamsWithHTTPClient creates a new SendViewProductParams object
// with the ability to set a custom HTTPClient for a request.
func NewSendViewProductParamsWithHTTPClient(client *http.Client) *SendViewProductParams {
	return &SendViewProductParams{
		HTTPClient: client,
	}
}

/* SendViewProductParams contains all the parameters to send to the API endpoint
   for the send view product operation.

   Typically these are written to a http.Request.
*/
type SendViewProductParams struct {

	// Body.
	Body SendViewProductBody

	// SiteID.
	SiteID string

	// XCqClientID.
	XCqClientID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the send view product params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SendViewProductParams) WithDefaults() *SendViewProductParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the send view product params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SendViewProductParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the send view product params
func (o *SendViewProductParams) WithTimeout(timeout time.Duration) *SendViewProductParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send view product params
func (o *SendViewProductParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send view product params
func (o *SendViewProductParams) WithContext(ctx context.Context) *SendViewProductParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send view product params
func (o *SendViewProductParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send view product params
func (o *SendViewProductParams) WithHTTPClient(client *http.Client) *SendViewProductParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send view product params
func (o *SendViewProductParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the send view product params
func (o *SendViewProductParams) WithBody(body SendViewProductBody) *SendViewProductParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the send view product params
func (o *SendViewProductParams) SetBody(body SendViewProductBody) {
	o.Body = body
}

// WithSiteID adds the siteID to the send view product params
func (o *SendViewProductParams) WithSiteID(siteID string) *SendViewProductParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the send view product params
func (o *SendViewProductParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithXCqClientID adds the xCqClientID to the send view product params
func (o *SendViewProductParams) WithXCqClientID(xCqClientID string) *SendViewProductParams {
	o.SetXCqClientID(xCqClientID)
	return o
}

// SetXCqClientID adds the xCqClientId to the send view product params
func (o *SendViewProductParams) SetXCqClientID(xCqClientID string) {
	o.XCqClientID = xCqClientID
}

// WriteToRequest writes these params to a swagger request
func (o *SendViewProductParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param siteId
	if err := r.SetPathParam("siteId", o.SiteID); err != nil {
		return err
	}

	// header param x-cq-client-id
	if err := r.SetHeaderParam("x-cq-client-id", o.XCqClientID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
