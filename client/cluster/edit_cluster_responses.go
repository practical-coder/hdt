// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// EditClusterReader is a Reader for the EditCluster structure.
type EditClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EditClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEditClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEditClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEditClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEditClusterOK creates a EditClusterOK with default headers values
func NewEditClusterOK() *EditClusterOK {
	return &EditClusterOK{}
}

/* EditClusterOK describes a response with status code 200, with default header values.

Cluster settings changed
*/
type EditClusterOK struct {
	Payload *models.ClusterSettings
}

func (o *EditClusterOK) Error() string {
	return fmt.Sprintf("[PUT /cluster][%d] editClusterOK  %+v", 200, o.Payload)
}
func (o *EditClusterOK) GetPayload() *models.ClusterSettings {
	return o.Payload
}

func (o *EditClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditClusterBadRequest creates a EditClusterBadRequest with default headers values
func NewEditClusterBadRequest() *EditClusterBadRequest {
	return &EditClusterBadRequest{}
}

/* EditClusterBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type EditClusterBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *EditClusterBadRequest) Error() string {
	return fmt.Sprintf("[PUT /cluster][%d] editClusterBadRequest  %+v", 400, o.Payload)
}
func (o *EditClusterBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *EditClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewEditClusterDefault creates a EditClusterDefault with default headers values
func NewEditClusterDefault(code int) *EditClusterDefault {
	return &EditClusterDefault{
		_statusCode: code,
	}
}

/* EditClusterDefault describes a response with status code -1, with default header values.

General Error
*/
type EditClusterDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the edit cluster default response
func (o *EditClusterDefault) Code() int {
	return o._statusCode
}

func (o *EditClusterDefault) Error() string {
	return fmt.Sprintf("[PUT /cluster][%d] editCluster default  %+v", o._statusCode, o.Payload)
}
func (o *EditClusterDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *EditClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
