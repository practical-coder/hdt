// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// SpoeFiles SPOE files
//
// SPOE files
//
// swagger:model spoe_files
type SpoeFiles []string

// Validate validates this spoe files
func (m SpoeFiles) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this spoe files based on context it is used
func (m SpoeFiles) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
