// Code generated by go-swagger; DO NOT EDIT.

package peer_entry

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

// GetPeerEntriesReader is a Reader for the GetPeerEntries structure.
type GetPeerEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPeerEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPeerEntriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPeerEntriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPeerEntriesOK creates a GetPeerEntriesOK with default headers values
func NewGetPeerEntriesOK() *GetPeerEntriesOK {
	return &GetPeerEntriesOK{}
}

/* GetPeerEntriesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetPeerEntriesOK struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *GetPeerEntriesOKBody
}

func (o *GetPeerEntriesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/peer_entries][%d] getPeerEntriesOK  %+v", 200, o.Payload)
}
func (o *GetPeerEntriesOK) GetPayload() *GetPeerEntriesOKBody {
	return o.Payload
}

func (o *GetPeerEntriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Configuration-Version
	hdrConfigurationVersion := response.GetHeader("Configuration-Version")

	if hdrConfigurationVersion != "" {
		o.ConfigurationVersion = hdrConfigurationVersion
	}

	o.Payload = new(GetPeerEntriesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPeerEntriesDefault creates a GetPeerEntriesDefault with default headers values
func NewGetPeerEntriesDefault(code int) *GetPeerEntriesDefault {
	return &GetPeerEntriesDefault{
		_statusCode: code,
	}
}

/* GetPeerEntriesDefault describes a response with status code -1, with default header values.

General Error
*/
type GetPeerEntriesDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get peer entries default response
func (o *GetPeerEntriesDefault) Code() int {
	return o._statusCode
}

func (o *GetPeerEntriesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/peer_entries][%d] getPeerEntries default  %+v", o._statusCode, o.Payload)
}
func (o *GetPeerEntriesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPeerEntriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetPeerEntriesOKBody get peer entries o k body
swagger:model GetPeerEntriesOKBody
*/
type GetPeerEntriesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.PeerEntries `json:"data"`
}

// Validate validates this get peer entries o k body
func (o *GetPeerEntriesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPeerEntriesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getPeerEntriesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getPeerEntriesOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getPeerEntriesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this get peer entries o k body based on the context it is used
func (o *GetPeerEntriesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPeerEntriesOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Data.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getPeerEntriesOK" + "." + "data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("getPeerEntriesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPeerEntriesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPeerEntriesOKBody) UnmarshalBinary(b []byte) error {
	var res GetPeerEntriesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
