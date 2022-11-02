// Code generated by go-swagger; DO NOT EDIT.

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// GetServicesEndpointsReader is a Reader for the GetServicesEndpoints structure.
type GetServicesEndpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServicesEndpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServicesEndpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetServicesEndpointsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServicesEndpointsOK creates a GetServicesEndpointsOK with default headers values
func NewGetServicesEndpointsOK() *GetServicesEndpointsOK {
	return &GetServicesEndpointsOK{}
}

/* GetServicesEndpointsOK describes a response with status code 200, with default header values.

Success
*/
type GetServicesEndpointsOK struct {
	Payload models.Endpoints
}

func (o *GetServicesEndpointsOK) Error() string {
	return fmt.Sprintf("[GET /services][%d] getServicesEndpointsOK  %+v", 200, o.Payload)
}
func (o *GetServicesEndpointsOK) GetPayload() models.Endpoints {
	return o.Payload
}

func (o *GetServicesEndpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServicesEndpointsDefault creates a GetServicesEndpointsDefault with default headers values
func NewGetServicesEndpointsDefault(code int) *GetServicesEndpointsDefault {
	return &GetServicesEndpointsDefault{
		_statusCode: code,
	}
}

/* GetServicesEndpointsDefault describes a response with status code -1, with default header values.

General Error
*/
type GetServicesEndpointsDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the get services endpoints default response
func (o *GetServicesEndpointsDefault) Code() int {
	return o._statusCode
}

func (o *GetServicesEndpointsDefault) Error() string {
	return fmt.Sprintf("[GET /services][%d] getServicesEndpoints default  %+v", o._statusCode, o.Payload)
}
func (o *GetServicesEndpointsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServicesEndpointsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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