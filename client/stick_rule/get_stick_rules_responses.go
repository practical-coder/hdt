// Code generated by go-swagger; DO NOT EDIT.

package stick_rule

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

// GetStickRulesReader is a Reader for the GetStickRules structure.
type GetStickRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStickRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStickRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetStickRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetStickRulesOK creates a GetStickRulesOK with default headers values
func NewGetStickRulesOK() *GetStickRulesOK {
	return &GetStickRulesOK{}
}

/* GetStickRulesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetStickRulesOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetStickRulesOKBody
}

func (o *GetStickRulesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/stick_rules][%d] getStickRulesOK  %+v", 200, o.Payload)
}
func (o *GetStickRulesOK) GetPayload() *GetStickRulesOKBody {
	return o.Payload
}

func (o *GetStickRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetStickRulesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStickRulesDefault creates a GetStickRulesDefault with default headers values
func NewGetStickRulesDefault(code int) *GetStickRulesDefault {
	return &GetStickRulesDefault{
		_statusCode: code,
	}
}

/* GetStickRulesDefault describes a response with status code -1, with default header values.

General Error
*/
type GetStickRulesDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get stick rules default response
func (o *GetStickRulesDefault) Code() int {
	return o._statusCode
}

func (o *GetStickRulesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/stick_rules][%d] getStickRules default  %+v", o._statusCode, o.Payload)
}
func (o *GetStickRulesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStickRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetStickRulesOKBody get stick rules o k body
swagger:model GetStickRulesOKBody
*/
type GetStickRulesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.StickRules `json:"data"`
}

// Validate validates this get stick rules o k body
func (o *GetStickRulesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStickRulesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getStickRulesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getStickRulesOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getStickRulesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get stick rules o k body based on the context it is used
func (o *GetStickRulesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStickRulesOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getStickRulesOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getStickRulesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetStickRulesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStickRulesOKBody) UnmarshalBinary(b []byte) error {
	var res GetStickRulesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}