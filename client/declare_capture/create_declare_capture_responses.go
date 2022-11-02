// Code generated by go-swagger; DO NOT EDIT.

package declare_capture

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haproxytech/client-native/v4/models"
)

// CreateDeclareCaptureReader is a Reader for the CreateDeclareCapture structure.
type CreateDeclareCaptureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDeclareCaptureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDeclareCaptureCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateDeclareCaptureAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDeclareCaptureBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateDeclareCaptureConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateDeclareCaptureDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateDeclareCaptureCreated creates a CreateDeclareCaptureCreated with default headers values
func NewCreateDeclareCaptureCreated() *CreateDeclareCaptureCreated {
	return &CreateDeclareCaptureCreated{}
}

/* CreateDeclareCaptureCreated describes a response with status code 201, with default header values.

Declare capture created
*/
type CreateDeclareCaptureCreated struct {
	Payload *models.Capture
}

func (o *CreateDeclareCaptureCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/captures][%d] createDeclareCaptureCreated  %+v", 201, o.Payload)
}
func (o *CreateDeclareCaptureCreated) GetPayload() *models.Capture {
	return o.Payload
}

func (o *CreateDeclareCaptureCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Capture)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeclareCaptureAccepted creates a CreateDeclareCaptureAccepted with default headers values
func NewCreateDeclareCaptureAccepted() *CreateDeclareCaptureAccepted {
	return &CreateDeclareCaptureAccepted{}
}

/* CreateDeclareCaptureAccepted describes a response with status code 202, with default header values.

Configuration change accepted and reload requested
*/
type CreateDeclareCaptureAccepted struct {

	/* ID of the requested reload
	 */
	ReloadID string

	Payload *models.Capture
}

func (o *CreateDeclareCaptureAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/captures][%d] createDeclareCaptureAccepted  %+v", 202, o.Payload)
}
func (o *CreateDeclareCaptureAccepted) GetPayload() *models.Capture {
	return o.Payload
}

func (o *CreateDeclareCaptureAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Reload-ID
	hdrReloadID := response.GetHeader("Reload-ID")

	if hdrReloadID != "" {
		o.ReloadID = hdrReloadID
	}

	o.Payload = new(models.Capture)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeclareCaptureBadRequest creates a CreateDeclareCaptureBadRequest with default headers values
func NewCreateDeclareCaptureBadRequest() *CreateDeclareCaptureBadRequest {
	return &CreateDeclareCaptureBadRequest{}
}

/* CreateDeclareCaptureBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateDeclareCaptureBadRequest struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateDeclareCaptureBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/captures][%d] createDeclareCaptureBadRequest  %+v", 400, o.Payload)
}
func (o *CreateDeclareCaptureBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDeclareCaptureBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDeclareCaptureConflict creates a CreateDeclareCaptureConflict with default headers values
func NewCreateDeclareCaptureConflict() *CreateDeclareCaptureConflict {
	return &CreateDeclareCaptureConflict{}
}

/* CreateDeclareCaptureConflict describes a response with status code 409, with default header values.

The specified resource already exists
*/
type CreateDeclareCaptureConflict struct {

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

func (o *CreateDeclareCaptureConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/captures][%d] createDeclareCaptureConflict  %+v", 409, o.Payload)
}
func (o *CreateDeclareCaptureConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDeclareCaptureConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDeclareCaptureDefault creates a CreateDeclareCaptureDefault with default headers values
func NewCreateDeclareCaptureDefault(code int) *CreateDeclareCaptureDefault {
	return &CreateDeclareCaptureDefault{
		_statusCode: code,
	}
}

/* CreateDeclareCaptureDefault describes a response with status code -1, with default header values.

General Error
*/
type CreateDeclareCaptureDefault struct {
	_statusCode int

	/* Configuration file version
	 */
	ConfigurationVersion string

	Payload *models.Error
}

// Code gets the status code for the create declare capture default response
func (o *CreateDeclareCaptureDefault) Code() int {
	return o._statusCode
}

func (o *CreateDeclareCaptureDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/configuration/captures][%d] createDeclareCapture default  %+v", o._statusCode, o.Payload)
}
func (o *CreateDeclareCaptureDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDeclareCaptureDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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