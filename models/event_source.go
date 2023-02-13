// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// EventSource Source of events
//
// - EVENT_SOURCE_UNSPECIFIED: Default
//   - EVENT_SOURCE_SYSTEM: Event is generated as a result of status change reported by edge node / application
//   - EVENT_SOURCE_USER: Event is generated as a result of an user's configuration action
//
// swagger:model EventSource
type EventSource string

func NewEventSource(value EventSource) *EventSource {
	return &value
}

// Pointer returns a pointer to a freshly-allocated EventSource.
func (m EventSource) Pointer() *EventSource {
	return &m
}

const (

	// EventSourceEVENTSOURCEUNSPECIFIED captures enum value "EVENT_SOURCE_UNSPECIFIED"
	EventSourceEVENTSOURCEUNSPECIFIED EventSource = "EVENT_SOURCE_UNSPECIFIED"

	// EventSourceEVENTSOURCESYSTEM captures enum value "EVENT_SOURCE_SYSTEM"
	EventSourceEVENTSOURCESYSTEM EventSource = "EVENT_SOURCE_SYSTEM"

	// EventSourceEVENTSOURCEUSER captures enum value "EVENT_SOURCE_USER"
	EventSourceEVENTSOURCEUSER EventSource = "EVENT_SOURCE_USER"
)

// for schema
var eventSourceEnum []interface{}

func init() {
	var res []EventSource
	if err := json.Unmarshal([]byte(`["EVENT_SOURCE_UNSPECIFIED","EVENT_SOURCE_SYSTEM","EVENT_SOURCE_USER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eventSourceEnum = append(eventSourceEnum, v)
	}
}

func (m EventSource) validateEventSourceEnum(path, location string, value EventSource) error {
	if err := validate.EnumCase(path, location, value, eventSourceEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this event source
func (m EventSource) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEventSourceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this event source based on context it is used
func (m EventSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}