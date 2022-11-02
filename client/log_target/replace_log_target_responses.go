// Code generated by go-swagger; DO NOT EDIT.

package log_target

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// ReplaceLogTargetReader is a Reader for the ReplaceLogTarget structure.
type ReplaceLogTargetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceLogTargetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceLogTargetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceLogTargetAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceLogTargetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceLogTargetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceLogTargetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceLogTargetOK creates a ReplaceLogTargetOK with default headers values
func NewReplaceLogTargetOK() *ReplaceLogTargetOK {
	return &ReplaceLogTargetOK{}
}

/* ReplaceLogTargetOK describes a response with status code 200, with default header values.

Log Target replaced
*/
type ReplaceLogTargetOK struct {
	Payload *models.LogTarget
}

func (o *ReplaceLogTargetOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/log_targets/{index}][%d] replaceLogTargetOK  %+v", 200, o.Payload)
}
func (o *ReplaceLogTargetOK) GetPayload() *models.LogTarget {
	return o.Payload
}

func (o *ReplaceLogTargetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogTarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceLogTargetAccepted creates a ReplaceLogTargetAccepted with default headers values
func NewReplaceLogTargetAccepted() *ReplaceLogTargetAccepted {
	return &ReplaceLogTargetAccepted{}
}

/* ReplaceLogTargetAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceLogTargetAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.LogTarget
}

func (o *ReplaceLogTargetAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/log_targets/{index}][%d] replaceLogTargetAccepted  %+v", 202, o.Payload)
}
func (o *ReplaceLogTargetAccepted) GetPayload() *models.LogTarget {
	return o.Payload
}

func (o *ReplaceLogTargetAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.LogTarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceLogTargetBadRequest creates a ReplaceLogTargetBadRequest with default headers values
func NewReplaceLogTargetBadRequest() *ReplaceLogTargetBadRequest {
	return &ReplaceLogTargetBadRequest{}
}

/* ReplaceLogTargetBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceLogTargetBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceLogTargetBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/log_targets/{index}][%d] replaceLogTargetBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceLogTargetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceLogTargetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceLogTargetNotFound creates a ReplaceLogTargetNotFound with default headers values
func NewReplaceLogTargetNotFound() *ReplaceLogTargetNotFound {
	return &ReplaceLogTargetNotFound{}
}

/* ReplaceLogTargetNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceLogTargetNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceLogTargetNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/log_targets/{index}][%d] replaceLogTargetNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceLogTargetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceLogTargetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceLogTargetDefault creates a ReplaceLogTargetDefault with default headers values
func NewReplaceLogTargetDefault(code int) *ReplaceLogTargetDefault {
	return &ReplaceLogTargetDefault{
		_statusCode: code,
	}
}

/* ReplaceLogTargetDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceLogTargetDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace log target default response
func (o *ReplaceLogTargetDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceLogTargetDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/log_targets/{index}][%d] replaceLogTarget default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceLogTargetDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceLogTargetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
