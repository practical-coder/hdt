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

// DeleteSpoeGroupReader is a Reader for the DeleteSpoeGroup structure.
type DeleteSpoeGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSpoeGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSpoeGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteSpoeGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteSpoeGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSpoeGroupNoContent creates a DeleteSpoeGroupNoContent with default headers values
func NewDeleteSpoeGroupNoContent() *DeleteSpoeGroupNoContent {
	return &DeleteSpoeGroupNoContent{}
}

/* DeleteSpoeGroupNoContent describes a response with status code 204, with default header values.

Spoe group deleted
*/
type DeleteSpoeGroupNoContent struct {
}

func (o *DeleteSpoeGroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_groups/{name}][%d] deleteSpoeGroupNoContent ", 204)
}

func (o *DeleteSpoeGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSpoeGroupNotFound creates a DeleteSpoeGroupNotFound with default headers values
func NewDeleteSpoeGroupNotFound() *DeleteSpoeGroupNotFound {
	return &DeleteSpoeGroupNotFound{}
}

/* DeleteSpoeGroupNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type DeleteSpoeGroupNotFound struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *DeleteSpoeGroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_groups/{name}][%d] deleteSpoeGroupNotFound  %+v", 404, o.Payload)
}
func (o *DeleteSpoeGroupNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSpoeGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSpoeGroupDefault creates a DeleteSpoeGroupDefault with default headers values
func NewDeleteSpoeGroupDefault(code int) *DeleteSpoeGroupDefault {
	return &DeleteSpoeGroupDefault{
		_statusCode: code,
	}
}

/* DeleteSpoeGroupDefault describes a response with status code -1, with default header values.

General Error
*/
type DeleteSpoeGroupDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the delete spoe group default response
func (o *DeleteSpoeGroupDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSpoeGroupDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/spoe/spoe_groups/{name}][%d] deleteSpoeGroup default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteSpoeGroupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteSpoeGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
