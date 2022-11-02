// Code generated by go-swagger; DO NOT EDIT.

package spoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateSpoeMessageReader is a Reader for the CreateSpoeMessage structure.
type CreateSpoeMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSpoeMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateSpoeMessageCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSpoeMessageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateSpoeMessageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateSpoeMessageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSpoeMessageCreated creates a CreateSpoeMessageCreated with default headers values
func NewCreateSpoeMessageCreated() *CreateSpoeMessageCreated {
	return &CreateSpoeMessageCreated{}
}

/* CreateSpoeMessageCreated describes a response with status code 201, with default header values.

Spoe message created
*/
type CreateSpoeMessageCreated struct {
	Payload *models.SpoeMessage
}

func (o *CreateSpoeMessageCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_messages][%d] createSpoeMessageCreated  %+v", 201, o.Payload)
}
func (o *CreateSpoeMessageCreated) GetPayload() *models.SpoeMessage {
	return o.Payload
}

func (o *CreateSpoeMessageCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SpoeMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSpoeMessageBadRequest creates a CreateSpoeMessageBadRequest with default headers values
func NewCreateSpoeMessageBadRequest() *CreateSpoeMessageBadRequest {
	return &CreateSpoeMessageBadRequest{}
}

/* CreateSpoeMessageBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateSpoeMessageBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateSpoeMessageBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_messages][%d] createSpoeMessageBadRequest  %+v", 400, o.Payload)
}
func (o *CreateSpoeMessageBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSpoeMessageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSpoeMessageConflict creates a CreateSpoeMessageConflict with default headers values
func NewCreateSpoeMessageConflict() *CreateSpoeMessageConflict {
	return &CreateSpoeMessageConflict{}
}

/* CreateSpoeMessageConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateSpoeMessageConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateSpoeMessageConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_messages][%d] createSpoeMessageConflict  %+v", 409, o.Payload)
}
func (o *CreateSpoeMessageConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSpoeMessageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSpoeMessageDefault creates a CreateSpoeMessageDefault with default headers values
func NewCreateSpoeMessageDefault(code int) *CreateSpoeMessageDefault {
	return &CreateSpoeMessageDefault{
		_statusCode: code,
	}
}

/* CreateSpoeMessageDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateSpoeMessageDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create spoe message default response
func (o *CreateSpoeMessageDefault) Code() int {
	return o._statusCode
}

func (o *CreateSpoeMessageDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/spoe/spoe_messages][%d] createSpoeMessage default  %+v", o._statusCode, o.Payload)
}
func (o *CreateSpoeMessageDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateSpoeMessageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
