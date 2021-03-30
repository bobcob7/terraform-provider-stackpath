// Code generated by go-swagger; DO NOT EDIT.

package workload_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Workloadv1InstanceInstancePhase An instance's startup state
//
// - INSTANCE_PHASE_UNSPECIFIED: StackPath is unable to determine the instance's startup state
//  - STARTING: The instance is still initializing
//  - RUNNING: The instance is running
//  - FAILED: The instance failed to start
//  - COMPLETED: The instance finished running
//  - SCHEDULING: The instance is being scheduled
//  - STOPPED: The instance is stopped
//
// swagger:model workloadv1InstanceInstancePhase
type Workloadv1InstanceInstancePhase string

const (

	// Workloadv1InstanceInstancePhaseINSTANCEPHASEUNSPECIFIED captures enum value "INSTANCE_PHASE_UNSPECIFIED"
	Workloadv1InstanceInstancePhaseINSTANCEPHASEUNSPECIFIED Workloadv1InstanceInstancePhase = "INSTANCE_PHASE_UNSPECIFIED"

	// Workloadv1InstanceInstancePhaseSTARTING captures enum value "STARTING"
	Workloadv1InstanceInstancePhaseSTARTING Workloadv1InstanceInstancePhase = "STARTING"

	// Workloadv1InstanceInstancePhaseRUNNING captures enum value "RUNNING"
	Workloadv1InstanceInstancePhaseRUNNING Workloadv1InstanceInstancePhase = "RUNNING"

	// Workloadv1InstanceInstancePhaseFAILED captures enum value "FAILED"
	Workloadv1InstanceInstancePhaseFAILED Workloadv1InstanceInstancePhase = "FAILED"

	// Workloadv1InstanceInstancePhaseCOMPLETED captures enum value "COMPLETED"
	Workloadv1InstanceInstancePhaseCOMPLETED Workloadv1InstanceInstancePhase = "COMPLETED"

	// Workloadv1InstanceInstancePhaseSCHEDULING captures enum value "SCHEDULING"
	Workloadv1InstanceInstancePhaseSCHEDULING Workloadv1InstanceInstancePhase = "SCHEDULING"

	// Workloadv1InstanceInstancePhaseSTOPPED captures enum value "STOPPED"
	Workloadv1InstanceInstancePhaseSTOPPED Workloadv1InstanceInstancePhase = "STOPPED"
)

// for schema
var workloadv1InstanceInstancePhaseEnum []interface{}

func init() {
	var res []Workloadv1InstanceInstancePhase
	if err := json.Unmarshal([]byte(`["INSTANCE_PHASE_UNSPECIFIED","STARTING","RUNNING","FAILED","COMPLETED","SCHEDULING","STOPPED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workloadv1InstanceInstancePhaseEnum = append(workloadv1InstanceInstancePhaseEnum, v)
	}
}

func (m Workloadv1InstanceInstancePhase) validateWorkloadv1InstanceInstancePhaseEnum(path, location string, value Workloadv1InstanceInstancePhase) error {
	if err := validate.EnumCase(path, location, value, workloadv1InstanceInstancePhaseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this workloadv1 instance instance phase
func (m Workloadv1InstanceInstancePhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateWorkloadv1InstanceInstancePhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this workloadv1 instance instance phase based on context it is used
func (m Workloadv1InstanceInstancePhase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
