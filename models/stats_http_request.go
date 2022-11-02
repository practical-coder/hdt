// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StatsHTTPRequest stats http request
//
// swagger:model stats_http_request
type StatsHTTPRequest struct {

	// cond
	Cond string `json:"cond,omitempty"`

	// cond test
	CondTest string `json:"cond_test,omitempty"`

	// realm
	Realm string `json:"realm,omitempty"`

	// type
	// Required: true
	// Enum: [allow deny auth]
	Type *string `json:"type"`
}

// Validate validates this stats http request
func (m *StatsHTTPRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var statsHttpRequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["allow","deny","auth"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statsHttpRequestTypeTypePropEnum = append(statsHttpRequestTypeTypePropEnum, v)
	}
}

const (

	// StatsHTTPRequestTypeAllow captures enum value "allow"
	StatsHTTPRequestTypeAllow string = "allow"

	// StatsHTTPRequestTypeDeny captures enum value "deny"
	StatsHTTPRequestTypeDeny string = "deny"

	// StatsHTTPRequestTypeAuth captures enum value "auth"
	StatsHTTPRequestTypeAuth string = "auth"
)

// prop value enum
func (m *StatsHTTPRequest) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, statsHttpRequestTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StatsHTTPRequest) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this stats http request based on context it is used
func (m *StatsHTTPRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StatsHTTPRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatsHTTPRequest) UnmarshalBinary(b []byte) error {
	var res StatsHTTPRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
