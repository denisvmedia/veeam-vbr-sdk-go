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

// EScriptPeriodicityType Type of script periodicity.
type EScriptPeriodicityType string

// List of EScriptPeriodicityType
const (
	ESCRIPTPERIODICITYTYPE_DAYS EScriptPeriodicityType = "Days"
	ESCRIPTPERIODICITYTYPE_BACKUP_SESSIONS EScriptPeriodicityType = "BackupSessions"
)

// All allowed values of EScriptPeriodicityType enum
var AllowedEScriptPeriodicityTypeEnumValues = []EScriptPeriodicityType{
	"Days",
	"BackupSessions",
}

func (v *EScriptPeriodicityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EScriptPeriodicityType(value)
	for _, existing := range AllowedEScriptPeriodicityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EScriptPeriodicityType", value)
}

// NewEScriptPeriodicityTypeFromValue returns a pointer to a valid EScriptPeriodicityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEScriptPeriodicityTypeFromValue(v string) (*EScriptPeriodicityType, error) {
	ev := EScriptPeriodicityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EScriptPeriodicityType: valid values are %v", v, AllowedEScriptPeriodicityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EScriptPeriodicityType) IsValid() bool {
	for _, existing := range AllowedEScriptPeriodicityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EScriptPeriodicityType value
func (v EScriptPeriodicityType) Ptr() *EScriptPeriodicityType {
	return &v
}

type NullableEScriptPeriodicityType struct {
	value *EScriptPeriodicityType
	isSet bool
}

func (v NullableEScriptPeriodicityType) Get() *EScriptPeriodicityType {
	return v.value
}

func (v *NullableEScriptPeriodicityType) Set(val *EScriptPeriodicityType) {
	v.value = val
	v.isSet = true
}

func (v NullableEScriptPeriodicityType) IsSet() bool {
	return v.isSet
}

func (v *NullableEScriptPeriodicityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEScriptPeriodicityType(val *EScriptPeriodicityType) *NullableEScriptPeriodicityType {
	return &NullableEScriptPeriodicityType{value: val, isSet: true}
}

func (v NullableEScriptPeriodicityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEScriptPeriodicityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

