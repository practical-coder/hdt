// Code generated by go-swagger; DO NOT EDIT.

package server_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// ReplaceServerSwitchingRuleReader is a Reader for the ReplaceServerSwitchingRule structure.
type ReplaceServerSwitchingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceServerSwitchingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceServerSwitchingRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceServerSwitchingRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceServerSwitchingRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceServerSwitchingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceServerSwitchingRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceServerSwitchingRuleOK creates a ReplaceServerSwitchingRuleOK with default headers values
func NewReplaceServerSwitchingRuleOK() *ReplaceServerSwitchingRuleOK {
	return &ReplaceServerSwitchingRuleOK{}
}

/* ReplaceServerSwitchingRuleOK describes a response with status code 200, with default header values.

Server Switching Rule replaced
*/
type ReplaceServerSwitchingRuleOK struct {
	Payload *models.ServerSwitchingRule
}

func (o *ReplaceServerSwitchingRuleOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/server_switching_rules/{index}][%d] replaceServerSwitchingRuleOK  %+v", 200, o.Payload)
}
func (o *ReplaceServerSwitchingRuleOK) GetPayload() *models.ServerSwitchingRule {
	return o.Payload
}

func (o *ReplaceServerSwitchingRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerSwitchingRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceServerSwitchingRuleAccepted creates a ReplaceServerSwitchingRuleAccepted with default headers values
func NewReplaceServerSwitchingRuleAccepted() *ReplaceServerSwitchingRuleAccepted {
	return &ReplaceServerSwitchingRuleAccepted{}
}

/* ReplaceServerSwitchingRuleAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type ReplaceServerSwitchingRuleAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.ServerSwitchingRule
}

func (o *ReplaceServerSwitchingRuleAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/server_switching_rules/{index}][%d] replaceServerSwitchingRuleAccepted  %+v", 202, o.Payload)
}
func (o *ReplaceServerSwitchingRuleAccepted) GetPayload() *models.ServerSwitchingRule {
	return o.Payload
}

func (o *ReplaceServerSwitchingRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.ServerSwitchingRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceServerSwitchingRuleBadRequest creates a ReplaceServerSwitchingRuleBadRequest with default headers values
func NewReplaceServerSwitchingRuleBadRequest() *ReplaceServerSwitchingRuleBadRequest {
	return &ReplaceServerSwitchingRuleBadRequest{}
}

/* ReplaceServerSwitchingRuleBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ReplaceServerSwitchingRuleBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceServerSwitchingRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/server_switching_rules/{index}][%d] replaceServerSwitchingRuleBadRequest  %+v", 400, o.Payload)
}
func (o *ReplaceServerSwitchingRuleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceServerSwitchingRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceServerSwitchingRuleNotFound creates a ReplaceServerSwitchingRuleNotFound with default headers values
func NewReplaceServerSwitchingRuleNotFound() *ReplaceServerSwitchingRuleNotFound {
	return &ReplaceServerSwitchingRuleNotFound{}
}

/* ReplaceServerSwitchingRuleNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type ReplaceServerSwitchingRuleNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *ReplaceServerSwitchingRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/server_switching_rules/{index}][%d] replaceServerSwitchingRuleNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceServerSwitchingRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceServerSwitchingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceServerSwitchingRuleDefault creates a ReplaceServerSwitchingRuleDefault with default headers values
func NewReplaceServerSwitchingRuleDefault(code int) *ReplaceServerSwitchingRuleDefault {
	return &ReplaceServerSwitchingRuleDefault{
		_statusCode: code,
	}
}

/* ReplaceServerSwitchingRuleDefault describes a response with status code -1, with default header values.

General Error
*/
type ReplaceServerSwitchingRuleDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the replace server switching rule default response
func (o *ReplaceServerSwitchingRuleDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceServerSwitchingRuleDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/server_switching_rules/{index}][%d] replaceServerSwitchingRule default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceServerSwitchingRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceServerSwitchingRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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