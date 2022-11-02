// Code generated by go-swagger; DO NOT EDIT.

package nameserver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// DeleteNameserverReader is a Reader for the DeleteNameserver structure.
type DeleteNameserverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNameserverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteNameserverAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteNameserverNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteNameserverNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteNameserverDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNameserverAccepted creates a DeleteNameserverAccepted with default headers values
func NewDeleteNameserverAccepted() *DeleteNameserverAccepted {
	return &DeleteNameserverAccepted{}
}

/* DeleteNameserverAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteNameserverAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteNameserverAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/nameservers/{name}][%d] deleteNameserverAccepted ", 202)
}

func (o *DeleteNameserverAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteNameserverNoContent creates a DeleteNameserverNoContent with default headers values
func NewDeleteNameserverNoContent() *DeleteNameserverNoContent {
	return &DeleteNameserverNoContent{}
}

/* DeleteNameserverNoContent describes a response with status code 204, with default header values.

Nameserver deleted
*/
type DeleteNameserverNoContent struct {
}

func (o *DeleteNameserverNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/nameservers/{name}][%d] deleteNameserverNoContent ", 204)
}

func (o *DeleteNameserverNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNameserverNotFound creates a DeleteNameserverNotFound with default headers values
func NewDeleteNameserverNotFound() *DeleteNameserverNotFound {
	return &DeleteNameserverNotFound{}
}

/* DeleteNameserverNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteNameserverNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteNameserverNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/nameservers/{name}][%d] deleteNameserverNotFound  %+v", 404, o.Payload)
}
func (o *DeleteNameserverNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteNameserverNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteNameserverDefault creates a DeleteNameserverDefault with default headers values
func NewDeleteNameserverDefault(code int) *DeleteNameserverDefault {
	return &DeleteNameserverDefault{
		_statusCode: code,
	}
}

/* DeleteNameserverDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteNameserverDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete nameserver default response
func (o *DeleteNameserverDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNameserverDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/nameservers/{name}][%d] deleteNameserver default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteNameserverDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteNameserverDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
