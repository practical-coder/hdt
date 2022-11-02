// Code generated by go-swagger; DO NOT EDIT.

package spoe_transactions

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

// NewGetSpoeTransactionParams creates a new GetSpoeTransactionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSpoeTransactionParams() *GetSpoeTransactionParams {
	return &GetSpoeTransactionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSpoeTransactionParamsWithTimeout creates a new GetSpoeTransactionParams object
// with the ability to set a timeout on a request.
func NewGetSpoeTransactionParamsWithTimeout(timeout time.Duration) *GetSpoeTransactionParams {
	return &GetSpoeTransactionParams{
		timeout: timeout,
	}
}

// NewGetSpoeTransactionParamsWithContext creates a new GetSpoeTransactionParams object
// with the ability to set a context for a request.
func NewGetSpoeTransactionParamsWithContext(ctx context.Context) *GetSpoeTransactionParams {
	return &GetSpoeTransactionParams{
		Context: ctx,
	}
}

// NewGetSpoeTransactionParamsWithHTTPClient creates a new GetSpoeTransactionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSpoeTransactionParamsWithHTTPClient(client *http.Client) *GetSpoeTransactionParams {
	return &GetSpoeTransactionParams{
		HTTPClient: client,
	}
}

/* GetSpoeTransactionParams contains all the parameters to send to the API endpoint
   for the get spoe transaction operation.

   Typically these are written to a http.Request.
*/
type GetSpoeTransactionParams struct {

	/* ID.

	   Transaction id
	*/
	ID string

	/* Spoe.

	   Spoe file name
	*/
	Spoe string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get spoe transaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSpoeTransactionParams) WithDefaults() *GetSpoeTransactionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get spoe transaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSpoeTransactionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get spoe transaction params
func (o *GetSpoeTransactionParams) WithTimeout(timeout time.Duration) *GetSpoeTransactionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get spoe transaction params
func (o *GetSpoeTransactionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get spoe transaction params
func (o *GetSpoeTransactionParams) WithContext(ctx context.Context) *GetSpoeTransactionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get spoe transaction params
func (o *GetSpoeTransactionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get spoe transaction params
func (o *GetSpoeTransactionParams) WithHTTPClient(client *http.Client) *GetSpoeTransactionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get spoe transaction params
func (o *GetSpoeTransactionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get spoe transaction params
func (o *GetSpoeTransactionParams) WithID(id string) *GetSpoeTransactionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get spoe transaction params
func (o *GetSpoeTransactionParams) SetID(id string) {
	o.ID = id
}

// WithSpoe adds the spoe to the get spoe transaction params
func (o *GetSpoeTransactionParams) WithSpoe(spoe string) *GetSpoeTransactionParams {
	o.SetSpoe(spoe)
	return o
}

// SetSpoe adds the spoe to the get spoe transaction params
func (o *GetSpoeTransactionParams) SetSpoe(spoe string) {
	o.Spoe = spoe
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpoeTransactionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param spoe
	qrSpoe := o.Spoe
	qSpoe := qrSpoe
	if qSpoe != "" {

		if err := r.SetQueryParam("spoe", qSpoe); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}