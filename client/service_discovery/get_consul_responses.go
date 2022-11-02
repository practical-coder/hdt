// Code generated by go-swagger; DO NOT EDIT.

package service_discovery

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

	"github.com/haproxytech/client-native/v4/models"
)

// GetConsulReader is a Reader for the GetConsul structure.
type GetConsulReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConsulReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConsulOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetConsulNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetConsulDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetConsulOK creates a GetConsulOK with default headers values
func NewGetConsulOK() *GetConsulOK {
	return &GetConsulOK{}
}

/* GetConsulOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetConsulOK struct {
	Payload *GetConsulOKBody
}

func (o *GetConsulOK) Error() string {
	return fmt.Sprintf("[GET /service_discovery/consul/{id}][%d] getConsulOK  %+v", 200, o.Payload)
}
func (o *GetConsulOK) GetPayload() *GetConsulOKBody {
	return o.Payload
}

func (o *GetConsulOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetConsulOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConsulNotFound creates a GetConsulNotFound with default headers values
func NewGetConsulNotFound() *GetConsulNotFound {
	return &GetConsulNotFound{}
}

/* GetConsulNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetConsulNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *GetConsulNotFound) Error() string {
	return fmt.Sprintf("[GET /service_discovery/consul/{id}][%d] getConsulNotFound  %+v", 404, o.Payload)
}
func (o *GetConsulNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetConsulNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetConsulDefault creates a GetConsulDefault with default headers values
func NewGetConsulDefault(code int) *GetConsulDefault {
	return &GetConsulDefault{
		_statusCode: code,
	}
}

/* GetConsulDefault describes a response with status code -1, with default header values.

General Error
*/
type GetConsulDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get consul default response
func (o *GetConsulDefault) Code() int {
	return o._statusCode
}

func (o *GetConsulDefault) Error() string {
	return fmt.Sprintf("[GET /service_discovery/consul/{id}][%d] getConsul default  %+v", o._statusCode, o.Payload)
}
func (o *GetConsulDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetConsulDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetConsulOKBody get consul o k body
swagger:model GetConsulOKBody
*/
type GetConsulOKBody struct {

	// data
	Data *models.Consul `json:"data,omitempty"`
}

// Validate validates this get consul o k body
func (o *GetConsulOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConsulOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getConsulOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getConsulOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get consul o k body based on the context it is used
func (o *GetConsulOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConsulOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getConsulOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getConsulOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetConsulOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetConsulOKBody) UnmarshalBinary(b []byte) error {
	var res GetConsulOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}