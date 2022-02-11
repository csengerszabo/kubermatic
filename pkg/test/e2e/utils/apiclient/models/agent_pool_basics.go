// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AgentPoolBasics agent pool basics
//
// swagger:model AgentPoolBasics
type AgentPoolBasics struct {

	// AvailabilityZones - The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType property is 'VirtualMachineScaleSets'.
	AvailabilityZones []string `json:"availabilityZones"`

	// Required: Count - Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive) for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.
	Count int32 `json:"count,omitempty"`

	// EnableAutoScaling - Whether to enable auto-scaler
	EnableAutoScaling bool `json:"enableAutoScaling,omitempty"`

	// Mode - Possible values include: 'System', 'User'.
	Mode string `json:"mode,omitempty"`

	// OrchestratorVersion - As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor version must be within two minor versions of the control plane version. The node pool version cannot be greater than the control plane version. For more information see [upgrading a node pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool).
	OrchestratorVersion string `json:"orchestratorVersion,omitempty"`

	// The OSDiskSize for Agent agentpool cannot be less than 30GB or larger than 2048GB.
	OsDiskSizeGB int32 `json:"osDiskSizeGB,omitempty"`

	// Required: VMSize - VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods might fail to run correctly. For more details on restricted VM sizes, see: https://docs.microsoft.com/azure/aks/quotas-skus-regions
	VMSize string `json:"vmSize,omitempty"`

	// scaling config
	ScalingConfig *NodegroupScalingConfig `json:"scalingConfig,omitempty"`
}

// Validate validates this agent pool basics
func (m *AgentPoolBasics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScalingConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgentPoolBasics) validateScalingConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.ScalingConfig) { // not required
		return nil
	}

	if m.ScalingConfig != nil {
		if err := m.ScalingConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scalingConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this agent pool basics based on the context it is used
func (m *AgentPoolBasics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScalingConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgentPoolBasics) contextValidateScalingConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.ScalingConfig != nil {
		if err := m.ScalingConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scalingConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AgentPoolBasics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgentPoolBasics) UnmarshalBinary(b []byte) error {
	var res AgentPoolBasics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
