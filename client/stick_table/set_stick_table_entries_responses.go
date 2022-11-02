// Code generated by go-swagger; DO NOT EDIT.

package stick_table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/haproxytech/client-native/v4/models"
)

// SetStickTableEntriesReader is a Reader for the SetStickTableEntries structure.
type SetStickTableEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetStickTableEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSetStickTableEntriesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSetStickTableEntriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSetStickTableEntriesNoContent creates a SetStickTableEntriesNoContent with default headers values
func NewSetStickTableEntriesNoContent() *SetStickTableEntriesNoContent {
	return &SetStickTableEntriesNoContent{}
}

/* SetStickTableEntriesNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type SetStickTableEntriesNoContent struct {
}

func (o *SetStickTableEntriesNoContent) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/runtime/stick_table_entries][%d] setStickTableEntriesNoContent ", 204)
}

func (o *SetStickTableEntriesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetStickTableEntriesDefault creates a SetStickTableEntriesDefault with default headers values
func NewSetStickTableEntriesDefault(code int) *SetStickTableEntriesDefault {
	return &SetStickTableEntriesDefault{
		_statusCode: code,
	}
}

/* SetStickTableEntriesDefault describes a response with status code -1, with default header values.

General Error
*/
type SetStickTableEntriesDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the set stick table entries default response
func (o *SetStickTableEntriesDefault) Code() int {
	return o._statusCode
}

func (o *SetStickTableEntriesDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/runtime/stick_table_entries][%d] setStickTableEntries default  %+v", o._statusCode, o.Payload)
}
func (o *SetStickTableEntriesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetStickTableEntriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SetStickTableEntriesBody set stick table entries body
swagger:model SetStickTableEntriesBody
*/
type SetStickTableEntriesBody struct {

	// data type
	// Required: true
	DataType *models.StickTableEntry `json:"data_type"`

	// key
	// Required: true
	Key *string `json:"key"`
}

// Validate validates this set stick table entries body
func (o *SetStickTableEntriesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDataType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SetStickTableEntriesBody) validateDataType(formats strfmt.Registry) error {

	if err := validate.Required("stick_table_entry"+"."+"data_type", "body", o.DataType); err != nil {
		return err
	}

	if o.DataType != nil {
		if err := o.DataType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stick_table_entry" + "." + "data_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stick_table_entry" + "." + "data_type")
			}
			return err
		}
	}

	return nil
}

func (o *SetStickTableEntriesBody) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("stick_table_entry"+"."+"key", "body", o.Key); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this set stick table entries body based on the context it is used
func (o *SetStickTableEntriesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDataType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SetStickTableEntriesBody) contextValidateDataType(ctx context.Context, formats strfmt.Registry) error {

	if o.DataType != nil {
		if err := o.DataType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stick_table_entry" + "." + "data_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stick_table_entry" + "." + "data_type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SetStickTableEntriesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetStickTableEntriesBody) UnmarshalBinary(b []byte) error {
	var res SetStickTableEntriesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
