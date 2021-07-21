// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
)

// EtcdBackupConfigConditionType EtcdBackupConfigConditionType is used to indicate the type of a EtcdBackupConfig condition. For all condition
// types, the `true` value must indicate success. All condition types must be registered within
// the `AllClusterConditionTypes` variable.
//
// swagger:model EtcdBackupConfigConditionType
type EtcdBackupConfigConditionType string

// Validate validates this etcd backup config condition type
func (m EtcdBackupConfigConditionType) Validate(formats strfmt.Registry) error {
	return nil
}