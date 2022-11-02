// Code generated by go-swagger; DO NOT EDIT.

package log_forward

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// DeleteLogForwardReader is a Reader for the DeleteLogForward structure.
type DeleteLogForwardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLogForwardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteLogForwardAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteLogForwardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteLogForwardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteLogForwardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteLogForwardAccepted creates a DeleteLogForwardAccepted with default headers values
func NewDeleteLogForwardAccepted() *DeleteLogForwardAccepted {
	return &DeleteLogForwardAccepted{}
}

/* DeleteLogForwardAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteLogForwardAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteLogForwardAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/log_forwards/{name}][%d] deleteLogForwardAccepted ", 202)
}

func (o *DeleteLogForwardAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteLogForwardNoContent creates a DeleteLogForwardNoContent with default headers values
func NewDeleteLogForwardNoContent() *DeleteLogForwardNoContent {
	return &DeleteLogForwardNoContent{}
}

/* DeleteLogForwardNoContent describes a response with status code 204, with default header values.

Log Forward deleted
*/
type DeleteLogForwardNoContent struct {
}

func (o *DeleteLogForwardNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/log_forwards/{name}][%d] deleteLogForwardNoContent ", 204)
}

func (o *DeleteLogForwardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLogForwardNotFound creates a DeleteLogForwardNotFound with default headers values
func NewDeleteLogForwardNotFound() *DeleteLogForwardNotFound {
	return &DeleteLogForwardNotFound{}
}

/* DeleteLogForwardNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteLogForwardNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteLogForwardNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/log_forwards/{name}][%d] deleteLogForwardNotFound  %+v", 404, o.Payload)
}
func (o *DeleteLogForwardNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteLogForwardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteLogForwardDefault creates a DeleteLogForwardDefault with default headers values
func NewDeleteLogForwardDefault(code int) *DeleteLogForwardDefault {
	return &DeleteLogForwardDefault{
		_statusCode: code,
	}
}

/* DeleteLogForwardDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteLogForwardDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete log forward default response
func (o *DeleteLogForwardDefault) Code() int {
	return o._statusCode
}

func (o *DeleteLogForwardDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/log_forwards/{name}][%d] deleteLogForward default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteLogForwardDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteLogForwardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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