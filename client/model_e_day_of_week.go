/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br>Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Some request and response bodies refer to reusable schema objects that are defined in the `schemas` section of the REST API specification. Schemas, in turn, may inherit a part of their properties by referring to other schemas, and be polymorphic by referring to multiple alternate schemas.

API version: 1.1-rev0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// EDayOfWeek Day of the week.
type EDayOfWeek string

// List of EDayOfWeek
const (
	EDAYOFWEEK_MONDAY EDayOfWeek = "monday"
	EDAYOFWEEK_TUESDAY EDayOfWeek = "tuesday"
	EDAYOFWEEK_WEDNESDAY EDayOfWeek = "wednesday"
	EDAYOFWEEK_THURSDAY EDayOfWeek = "thursday"
	EDAYOFWEEK_FRIDAY EDayOfWeek = "friday"
	EDAYOFWEEK_SATURDAY EDayOfWeek = "saturday"
	EDAYOFWEEK_SUNDAY EDayOfWeek = "sunday"
)

// All allowed values of EDayOfWeek enum
var AllowedEDayOfWeekEnumValues = []EDayOfWeek{
	"monday",
	"tuesday",
	"wednesday",
	"thursday",
	"friday",
	"saturday",
	"sunday",
}

func (v *EDayOfWeek) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EDayOfWeek(value)
	for _, existing := range AllowedEDayOfWeekEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EDayOfWeek", value)
}

// NewEDayOfWeekFromValue returns a pointer to a valid EDayOfWeek
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEDayOfWeekFromValue(v string) (*EDayOfWeek, error) {
	ev := EDayOfWeek(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EDayOfWeek: valid values are %v", v, AllowedEDayOfWeekEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EDayOfWeek) IsValid() bool {
	for _, existing := range AllowedEDayOfWeekEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EDayOfWeek value
func (v EDayOfWeek) Ptr() *EDayOfWeek {
	return &v
}

type NullableEDayOfWeek struct {
	value *EDayOfWeek
	isSet bool
}

func (v NullableEDayOfWeek) Get() *EDayOfWeek {
	return v.value
}

func (v *NullableEDayOfWeek) Set(val *EDayOfWeek) {
	v.value = val
	v.isSet = true
}

func (v NullableEDayOfWeek) IsSet() bool {
	return v.isSet
}

func (v *NullableEDayOfWeek) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEDayOfWeek(val *EDayOfWeek) *NullableEDayOfWeek {
	return &NullableEDayOfWeek{value: val, isSet: true}
}

func (v NullableEDayOfWeek) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEDayOfWeek) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

