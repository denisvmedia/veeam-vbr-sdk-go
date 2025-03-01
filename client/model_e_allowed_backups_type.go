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

// EAllowedBackupsType Type of backup files that can be stored on the extent.
type EAllowedBackupsType string

// List of EAllowedBackupsType
const (
	EALLOWEDBACKUPSTYPE_ALL EAllowedBackupsType = "All"
	EALLOWEDBACKUPSTYPE_FULLS_ONLY EAllowedBackupsType = "FullsOnly"
	EALLOWEDBACKUPSTYPE_INCREMENTS_ONLY EAllowedBackupsType = "IncrementsOnly"
	EALLOWEDBACKUPSTYPE_NONE EAllowedBackupsType = "None"
)

// All allowed values of EAllowedBackupsType enum
var AllowedEAllowedBackupsTypeEnumValues = []EAllowedBackupsType{
	"All",
	"FullsOnly",
	"IncrementsOnly",
	"None",
}

func (v *EAllowedBackupsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EAllowedBackupsType(value)
	for _, existing := range AllowedEAllowedBackupsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EAllowedBackupsType", value)
}

// NewEAllowedBackupsTypeFromValue returns a pointer to a valid EAllowedBackupsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEAllowedBackupsTypeFromValue(v string) (*EAllowedBackupsType, error) {
	ev := EAllowedBackupsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EAllowedBackupsType: valid values are %v", v, AllowedEAllowedBackupsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EAllowedBackupsType) IsValid() bool {
	for _, existing := range AllowedEAllowedBackupsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EAllowedBackupsType value
func (v EAllowedBackupsType) Ptr() *EAllowedBackupsType {
	return &v
}

type NullableEAllowedBackupsType struct {
	value *EAllowedBackupsType
	isSet bool
}

func (v NullableEAllowedBackupsType) Get() *EAllowedBackupsType {
	return v.value
}

func (v *NullableEAllowedBackupsType) Set(val *EAllowedBackupsType) {
	v.value = val
	v.isSet = true
}

func (v NullableEAllowedBackupsType) IsSet() bool {
	return v.isSet
}

func (v *NullableEAllowedBackupsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEAllowedBackupsType(val *EAllowedBackupsType) *NullableEAllowedBackupsType {
	return &NullableEAllowedBackupsType{value: val, isSet: true}
}

func (v NullableEAllowedBackupsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEAllowedBackupsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

