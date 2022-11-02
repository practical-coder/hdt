// Code generated by go-swagger; DO NOT EDIT.

package http_response_rule

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

// NewDeleteHTTPResponseRuleParams creates a new DeleteHTTPResponseRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteHTTPResponseRuleParams() *DeleteHTTPResponseRuleParams {
	return &DeleteHTTPResponseRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteHTTPResponseRuleParamsWithTimeout creates a new DeleteHTTPResponseRuleParams object
// with the ability to set a timeout on a request.
func NewDeleteHTTPResponseRuleParamsWithTimeout(timeout time.Duration) *DeleteHTTPResponseRuleParams {
	return &DeleteHTTPResponseRuleParams{
		timeout: timeout,
	}
}

// NewDeleteHTTPResponseRuleParamsWithContext creates a new DeleteHTTPResponseRuleParams object
// with the ability to set a context for a request.
func NewDeleteHTTPResponseRuleParamsWithContext(ctx context.Context) *DeleteHTTPResponseRuleParams {
	return &DeleteHTTPResponseRuleParams{
		Context: ctx,
	}
}

// NewDeleteHTTPResponseRuleParamsWithHTTPClient creates a new DeleteHTTPResponseRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteHTTPResponseRuleParamsWithHTTPClient(client *http.Client) *DeleteHTTPResponseRuleParams {
	return &DeleteHTTPResponseRuleParams{
		HTTPClient: client,
	}
}

/* DeleteHTTPResponseRuleParams contains all the parameters to send to the API endpoint
   for the delete HTTP response rule operation.

   Typically these are written to a http.Request.
*/
type DeleteHTTPResponseRuleParams struct {

	/* ForceReload.

	   If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.
	*/
	ForceReload *bool

	/* Index.

	   HTTP Response Rule Index
	*/
	Index int64

	/* ParentName.

	   Parent name
	*/
	ParentName string

	/* ParentType.

	   Parent type
	*/
	ParentType string

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

// WithDefaults hydrates default values in the delete HTTP response rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteHTTPResponseRuleParams) WithDefaults() *DeleteHTTPResponseRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete HTTP response rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteHTTPResponseRuleParams) SetDefaults() {
	var (
		forceReloadDefault = bool(false)
	)

	val := DeleteHTTPResponseRuleParams{
		ForceReload: &forceReloadDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete HTTP response rule params
func (o *DeleteHTTPResponseRuleParams) WithTimeout(timeout time.Duration) *DeleteHTTPResponseRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete HTTP response rule params
func (o *DeleteHTTPResponseRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete HTTP response rule params
func (o *DeleteHTTPResponseRuleParams) WithContext(ctx context.Context) *DeleteHTTPResponseRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete HTTP response rule params
func (o *DeleteHTTPResponseRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete HTTP response rule params
func (o *DeleteHTTPResponseRuleParams) WithHTTPClient(client *http.Client) *DeleteHTTPResponseRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete HTTP response rule params
func (o *DeleteHTTPResponseRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceReload adds the forceReload to the delete HTTP response rule params
func (o *DeleteHTTPResponseRuleParams) WithForceReload(forceReload *bool) *DeleteHTTPResponseRuleParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the delete HTTP response rule params
func (o *DeleteHTTPResponseRuleParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithIndex adds the index to the delete HTTP response rule params
func (o *DeleteHTTPResponseRuleParams) WithIndex(index int64) *DeleteHTTPResponseRuleParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the delete HTTP response rule params
func (o *DeleteHTTPResponseRuleParams) SetIndex(index int64) {
	o.Index = index
}

// WithParentName adds the parentName to the delete HTTP response rule params
func (o *DeleteHTTPResponseRuleParams) WithParentName(parentName string) *DeleteHTTPResponseRuleParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the delete HTTP response rule params
func (o *DeleteHTTPResponseRuleParams) SetParentName(parentName string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the delete HTTP response rule params
func (o *DeleteHTTPResponseRuleParams) WithParentType(parentType string) *DeleteHTTPResponseRuleParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the delete HTTP response rule params
func (o *DeleteHTTPResponseRuleParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the delete HTTP response rule params
func (o *DeleteHTTPResponseRuleParams) WithTransactionID(transactionID *string) *DeleteHTTPResponseRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the delete HTTP response rule params
func (o *DeleteHTTPResponseRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the delete HTTP response rule params
func (o *DeleteHTTPResponseRuleParams) WithVersion(version *int64) *DeleteHTTPResponseRuleParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete HTTP response rule params
func (o *DeleteHTTPResponseRuleParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteHTTPResponseRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}

	// query param parent_name
	qrParentName := o.ParentName
	qParentName := qrParentName
	if qParentName != "" {

		if err := r.SetQueryParam("parent_name", qParentName); err != nil {
			return err
		}
	}

	// query param parent_type
	qrParentType := o.ParentType
	qParentType := qrParentType
	if qParentType != "" {

		if err := r.SetQueryParam("parent_type", qParentType); err != nil {
			return err
		}
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
