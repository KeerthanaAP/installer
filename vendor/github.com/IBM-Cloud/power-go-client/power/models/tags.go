// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Tags List of user tags
//
// swagger:model Tags
type Tags []string

// Validate validates this tags
func (m Tags) Validate(formats strfmt.Registry) error {
	var res []error

	iTagsSize := int64(len(m))

	if err := validate.MaxItems("", "body", iTagsSize, 100); err != nil {
		return err
	}

	for i := 0; i < len(m); i++ {

		if err := validate.MaxLength(strconv.Itoa(i), "body", m[i], 128); err != nil {
			return err
		}

		if err := validate.Pattern(strconv.Itoa(i), "body", m[i], `^[\s]*[A-Za-z0-9:_.\-][A-Za-z0-9\s:_.\-]*$`); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this tags based on context it is used
func (m Tags) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
