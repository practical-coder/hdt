// Code generated by go-swagger; DO NOT EDIT.

package frontend

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

// NewGetFrontendsParams creates a new GetFrontendsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFrontendsParams() *GetFrontendsParams {
	return &GetFrontendsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFrontendsParamsWithTimeout creates a new GetFrontendsParams object
// with the ability to set a timeout on a request.
func NewGetFrontendsParamsWithTimeout(timeout time.Duration) *GetFrontendsParams {
	return &GetFrontendsParams{
		timeout: timeout,
	}
}

// NewGetFrontendsParamsWithContext creates a new GetFrontendsParams object
// with the ability to set a context for a request.
func NewGetFrontendsParamsWithContext(ctx context.Context) *GetFrontendsParams {
	return &GetFrontendsParams{
		Context: ctx,
	}
}

// NewGetFrontendsParamsWithHTTPClient creates a new GetFrontendsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFrontendsParamsWithHTTPClient(client *http.Client) *GetFrontendsParams {
	return &GetFrontendsParams{
		HTTPClient: client,
	}
}

/* GetFrontendsParams contains all the parameters to send to the API endpoint
   for the get frontends operation.

   Typically these are written to a http.Request.
*/
type GetFrontendsParams struct {

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get frontends params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFrontendsParams) WithDefaults() *GetFrontendsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get frontends params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFrontendsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get frontends params
func (o *GetFrontendsParams) WithTimeout(timeout time.Duration) *GetFrontendsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get frontends params
func (o *GetFrontendsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get frontends params
func (o *GetFrontendsParams) WithContext(ctx context.Context) *GetFrontendsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get frontends params
func (o *GetFrontendsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get frontends params
func (o *GetFrontendsParams) WithHTTPClient(client *http.Client) *GetFrontendsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get frontends params
func (o *GetFrontendsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTransactionID adds the transactionID to the get frontends params
func (o *GetFrontendsParams) WithTransactionID(transactionID *string) *GetFrontendsParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get frontends params
func (o *GetFrontendsParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFrontendsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TransactionID != nil {

		// query param transaction_id
		var qrTransactionID string

		if o.TransactionID != nil {
			qrTransactionID = *o.TransactionID
		}
		qTransactionID := qrTransactionID
		if qTransactionID != "" {

			if err := r.SetQueryParam("transaction_id", qTransactionID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}