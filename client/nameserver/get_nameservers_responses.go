// Code generated by go-swagger; DO NOT EDIT.

package nameserver

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

// GetNameserversReader is a Reader for the GetNameservers structure.
type GetNameserversReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNameserversReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNameserversOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNameserversDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNameserversOK creates a GetNameserversOK with default headers values
func NewGetNameserversOK() *GetNameserversOK {
	return &GetNameserversOK{}
}

/* GetNameserversOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNameserversOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetNameserversOKBody
}

func (o *GetNameserversOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/nameservers][%d] getNameserversOK  %+v", 200, o.Payload)
}
func (o *GetNameserversOK) GetPayload() *GetNameserversOKBody {
	return o.Payload
}

func (o *GetNameserversOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetNameserversOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNameserversDefault creates a GetNameserversDefault with default headers values
func NewGetNameserversDefault(code int) *GetNameserversDefault {
	return &GetNameserversDefault{
		_statusCode: code,
	}
}

/* GetNameserversDefault describes a response with status code -1, with default header values.

General Error
*/
type GetNameserversDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get nameservers default response
func (o *GetNameserversDefault) Code() int {
	return o._statusCode
}

func (o *GetNameserversDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/nameservers][%d] getNameservers default  %+v", o._statusCode, o.Payload)
}
func (o *GetNameserversDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNameserversDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetNameserversOKBody get nameservers o k body
swagger:model GetNameserversOKBody
*/
type GetNameserversOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.Nameservers `json:"data"`
}

// Validate validates this get nameservers o k body
func (o *GetNameserversOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNameserversOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getNameserversOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getNameserversOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getNameserversOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get nameservers o k body based on the context it is used
func (o *GetNameserversOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNameserversOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getNameserversOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getNameserversOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNameserversOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNameserversOKBody) UnmarshalBinary(b []byte) error {
	var res GetNameserversOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
