// Code generated by go-swagger; DO NOT EDIT.

package resolver

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

// NewGetResolversParams creates a new GetResolversParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetResolversParams() *GetResolversParams {
	return &GetResolversParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetResolversParamsWithTimeout creates a new GetResolversParams object
// with the ability to set a timeout on a request.
func NewGetResolversParamsWithTimeout(timeout time.Duration) *GetResolversParams {
	return &GetResolversParams{
		timeout: timeout,
	}
}

// NewGetResolversParamsWithContext creates a new GetResolversParams object
// with the ability to set a context for a request.
func NewGetResolversParamsWithContext(ctx context.Context) *GetResolversParams {
	return &GetResolversParams{
		Context: ctx,
	}
}

// NewGetResolversParamsWithHTTPClient creates a new GetResolversParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetResolversParamsWithHTTPClient(client *http.Client) *GetResolversParams {
	return &GetResolversParams{
		HTTPClient: client,
	}
}

/* GetResolversParams contains all the parameters to send to the API endpoint
   for the get resolvers operation.

   Typically these are written to a http.Request.
*/
type GetResolversParams struct {

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get resolvers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResolversParams) WithDefaults() *GetResolversParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get resolvers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResolversParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get resolvers params
func (o *GetResolversParams) WithTimeout(timeout time.Duration) *GetResolversParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resolvers params
func (o *GetResolversParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resolvers params
func (o *GetResolversParams) WithContext(ctx context.Context) *GetResolversParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resolvers params
func (o *GetResolversParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resolvers params
func (o *GetResolversParams) WithHTTPClient(client *http.Client) *GetResolversParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resolvers params
func (o *GetResolversParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTransactionID adds the transactionID to the get resolvers params
func (o *GetResolversParams) WithTransactionID(transactionID *string) *GetResolversParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get resolvers params
func (o *GetResolversParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetResolversParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
