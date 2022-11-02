// Code generated by go-swagger; DO NOT EDIT.

package declare_capture

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

// NewDeleteDeclareCaptureParams creates a new DeleteDeclareCaptureParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDeclareCaptureParams() *DeleteDeclareCaptureParams {
	return &DeleteDeclareCaptureParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeclareCaptureParamsWithTimeout creates a new DeleteDeclareCaptureParams object
// with the ability to set a timeout on a request.
func NewDeleteDeclareCaptureParamsWithTimeout(timeout time.Duration) *DeleteDeclareCaptureParams {
	return &DeleteDeclareCaptureParams{
		timeout: timeout,
	}
}

// NewDeleteDeclareCaptureParamsWithContext creates a new DeleteDeclareCaptureParams object
// with the ability to set a context for a request.
func NewDeleteDeclareCaptureParamsWithContext(ctx context.Context) *DeleteDeclareCaptureParams {
	return &DeleteDeclareCaptureParams{
		Context: ctx,
	}
}

// NewDeleteDeclareCaptureParamsWithHTTPClient creates a new DeleteDeclareCaptureParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDeclareCaptureParamsWithHTTPClient(client *http.Client) *DeleteDeclareCaptureParams {
	return &DeleteDeclareCaptureParams{
		HTTPClient: client,
	}
}

/* DeleteDeclareCaptureParams contains all the parameters to send to the API endpoint
   for the delete declare capture operation.

   Typically these are written to a http.Request.
*/
type DeleteDeclareCaptureParams struct {

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Frontend.

	   Parent frontend name
	*/
	Frontend string

	/* Index.

	   Declare Capture Index
	*/
	Index int64

	/* TransactionID.

	   ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	*/
	TransactionID *string

	/* Version.

	   Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.
	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete declare capture params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeclareCaptureParams) WithDefaults() *DeleteDeclareCaptureParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete declare capture params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeclareCaptureParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := DeleteDeclareCaptureParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete declare capture params
func (o *DeleteDeclareCaptureParams) WithTimeout(timeout time.Duration) *DeleteDeclareCaptureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete declare capture params
func (o *DeleteDeclareCaptureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete declare capture params
func (o *DeleteDeclareCaptureParams) WithContext(ctx context.Context) *DeleteDeclareCaptureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete declare capture params
func (o *DeleteDeclareCaptureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete declare capture params
func (o *DeleteDeclareCaptureParams) WithHTTPClient(client *http.Client) *DeleteDeclareCaptureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete declare capture params
func (o *DeleteDeclareCaptureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceReload adds the forceReload to the delete declare capture params
func (o *DeleteDeclareCaptureParams) WithForceReload(forceReload *bool) *DeleteDeclareCaptureParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the delete declare capture params
func (o *DeleteDeclareCaptureParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithFrontend adds the frontend to the delete declare capture params
func (o *DeleteDeclareCaptureParams) WithFrontend(frontend string) *DeleteDeclareCaptureParams {
	o.SetFrontend(frontend)
	return o
}

// SetFrontend adds the frontend to the delete declare capture params
func (o *DeleteDeclareCaptureParams) SetFrontend(frontend string) {
	o.Frontend = frontend
}

// WithIndex adds the index to the delete declare capture params
func (o *DeleteDeclareCaptureParams) WithIndex(index int64) *DeleteDeclareCaptureParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the delete declare capture params
func (o *DeleteDeclareCaptureParams) SetIndex(index int64) {
	o.Index = index
}

// WithTransactionID adds the transactionID to the delete declare capture params
func (o *DeleteDeclareCaptureParams) WithTransactionID(transactionID *string) *DeleteDeclareCaptureParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the delete declare capture params
func (o *DeleteDeclareCaptureParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the delete declare capture params
func (o *DeleteDeclareCaptureParams) WithVersion(version *int64) *DeleteDeclareCaptureParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete declare capture params
func (o *DeleteDeclareCaptureParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeclareCaptureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param frontend
	qrFrontend := o.Frontend
	qFrontend := qrFrontend
	if qFrontend != "" {

		if err := r.SetQueryParam("frontend", qFrontend); err != nil {
			return err
		}
	}

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}

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

	if o.Version != nil {

		// query param version
		var qrVersion int64

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
