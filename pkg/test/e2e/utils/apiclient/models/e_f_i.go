// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EFI If set, EFI will be used instead of BIOS.
//
// swagger:model EFI
type EFI struct {

	// If set, SecureBoot will be enabled and the OVMF roms will be swapped for
	// SecureBoot-enabled ones.
	// Requires SMM to be enabled.
	// Defaults to true
	// +optional
	SecureBoot bool `json:"secureBoot,omitempty"`
}

// Validate validates this e f i
func (m *EFI) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this e f i based on context it is used
func (m *EFI) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EFI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EFI) UnmarshalBinary(b []byte) error {
	var res EFI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}