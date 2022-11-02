// Code generated by go-swagger; DO NOT EDIT.

package tcp_response_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// DeleteTCPResponseRuleReader is a Reader for the DeleteTCPResponseRule structure.
type DeleteTCPResponseRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTCPResponseRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteTCPResponseRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteTCPResponseRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteTCPResponseRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteTCPResponseRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTCPResponseRuleAccepted creates a DeleteTCPResponseRuleAccepted with default headers values
func NewDeleteTCPResponseRuleAccepted() *DeleteTCPResponseRuleAccepted {
	return &DeleteTCPResponseRuleAccepted{}
}

/* DeleteTCPResponseRuleAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type DeleteTCPResponseRuleAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteTCPResponseRuleAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_response_rules/{index}][%d] deleteTcpResponseRuleAccepted ", 202)
}

func (o *DeleteTCPResponseRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	return nil
}

// NewDeleteTCPResponseRuleNoContent creates a DeleteTCPResponseRuleNoContent with default headers values
func NewDeleteTCPResponseRuleNoContent() *DeleteTCPResponseRuleNoContent {
	return &DeleteTCPResponseRuleNoContent{}
}

/* DeleteTCPResponseRuleNoContent describes a response with status code 204, with default header values.

TCP Response Rule deleted
*/
type DeleteTCPResponseRuleNoContent struct {
}

func (o *DeleteTCPResponseRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_response_rules/{index}][%d] deleteTcpResponseRuleNoContent ", 204)
}

func (o *DeleteTCPResponseRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTCPResponseRuleNotFound creates a DeleteTCPResponseRuleNotFound with default headers values
func NewDeleteTCPResponseRuleNotFound() *DeleteTCPResponseRuleNotFound {
	return &DeleteTCPResponseRuleNotFound{}
}

/* DeleteTCPResponseRuleNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteTCPResponseRuleNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteTCPResponseRuleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_response_rules/{index}][%d] deleteTcpResponseRuleNotFound  %+v", 404, o.Payload)
}
func (o *DeleteTCPResponseRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteTCPResponseRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteTCPResponseRuleDefault creates a DeleteTCPResponseRuleDefault with default headers values
func NewDeleteTCPResponseRuleDefault(code int) *DeleteTCPResponseRuleDefault {
	return &DeleteTCPResponseRuleDefault{
		_statusCode: code,
	}
}

/* DeleteTCPResponseRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteTCPResponseRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete TCP response rule default response
func (o *DeleteTCPResponseRuleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTCPResponseRuleDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/tcp_response_rules/{index}][%d] deleteTCPResponseRule default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteTCPResponseRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteTCPResponseRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
