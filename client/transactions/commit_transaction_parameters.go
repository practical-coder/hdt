// Code generated by go-swagger; DO NOT EDIT.

package transactions

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
	"github.com/go-openapi/swag"
)

// NewCommitTransactionParams creates a new CommitTransactionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCommitTransactionParams() *CommitTransactionParams {
	return &CommitTransactionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCommitTransactionParamsWithTimeout creates a new CommitTransactionParams object
// with the ability to set a timeout on a request.
func NewCommitTransactionParamsWithTimeout(timeout time.Duration) *CommitTransactionParams {
	return &CommitTransactionParams{
		timeout: timeout,
	}
}

// NewCommitTransactionParamsWithContext creates a new CommitTransactionParams object
// with the ability to set a context for a request.
func NewCommitTransactionParamsWithContext(ctx context.Context) *CommitTransactionParams {
	return &CommitTransactionParams{
		Context: ctx,
	}
}

// NewCommitTransactionParamsWithHTTPClient creates a new CommitTransactionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCommitTransactionParamsWithHTTPClient(client *http.Client) *CommitTransactionParams {
	return &CommitTransactionParams{
		HTTPClient: client,
	}
}

/* CommitTransactionParams contains all the parameters to send to the API endpoint
   for the commit transaction operation.

   Typically these are written to a http.Request.
*/
type CommitTransactionParams struct {

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* ID.

	   Transaction id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the commit transaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CommitTransactionParams) WithDefaults() *CommitTransactionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the commit transaction params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CommitTransactionParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := CommitTransactionParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the commit transaction params
func (o *CommitTransactionParams) WithTimeout(timeout time.Duration) *CommitTransactionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the commit transaction params
func (o *CommitTransactionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the commit transaction params
func (o *CommitTransactionParams) WithContext(ctx context.Context) *CommitTransactionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the commit transaction params
func (o *CommitTransactionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the commit transaction params
func (o *CommitTransactionParams) WithHTTPClient(client *http.Client) *CommitTransactionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the commit transaction params
func (o *CommitTransactionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceReload adds the forceReload to the commit transaction params
func (o *CommitTransactionParams) WithForceReload(forceReload *bool) *CommitTransactionParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the commit transaction params
func (o *CommitTransactionParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithID adds the id to the commit transaction params
func (o *CommitTransactionParams) WithID(id string) *CommitTransactionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the commit transaction params
func (o *CommitTransactionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CommitTransactionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ForceReload != nil {

		// query param force_reload
		var qrForceReload bool

		if o.ForceReload != nil {
			qrForceReload = *o.ForceReload
		}
		qForceReload := swag.FormatBool(qrForceReload)
		if qForceReload != "" {

			if err := r.SetQueryParam("force_reload", qForceReload); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
