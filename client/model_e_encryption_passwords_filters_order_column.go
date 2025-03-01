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

// EEncryptionPasswordsFiltersOrderColumn the model 'EEncryptionPasswordsFiltersOrderColumn'
type EEncryptionPasswordsFiltersOrderColumn string

// List of EEncryptionPasswordsFiltersOrderColumn
const (
	EENCRYPTIONPASSWORDSFILTERSORDERCOLUMN_HINT EEncryptionPasswordsFiltersOrderColumn = "Hint"
	EENCRYPTIONPASSWORDSFILTERSORDERCOLUMN_MODIFICATION_TIME EEncryptionPasswordsFiltersOrderColumn = "ModificationTime"
)

// All allowed values of EEncryptionPasswordsFiltersOrderColumn enum
var AllowedEEncryptionPasswordsFiltersOrderColumnEnumValues = []EEncryptionPasswordsFiltersOrderColumn{
	"Hint",
	"ModificationTime",
}

func (v *EEncryptionPasswordsFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EEncryptionPasswordsFiltersOrderColumn(value)
	for _, existing := range AllowedEEncryptionPasswordsFiltersOrderColumnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EEncryptionPasswordsFiltersOrderColumn", value)
}

// NewEEncryptionPasswordsFiltersOrderColumnFromValue returns a pointer to a valid EEncryptionPasswordsFiltersOrderColumn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEEncryptionPasswordsFiltersOrderColumnFromValue(v string) (*EEncryptionPasswordsFiltersOrderColumn, error) {
	ev := EEncryptionPasswordsFiltersOrderColumn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EEncryptionPasswordsFiltersOrderColumn: valid values are %v", v, AllowedEEncryptionPasswordsFiltersOrderColumnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EEncryptionPasswordsFiltersOrderColumn) IsValid() bool {
	for _, existing := range AllowedEEncryptionPasswordsFiltersOrderColumnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EEncryptionPasswordsFiltersOrderColumn value
func (v EEncryptionPasswordsFiltersOrderColumn) Ptr() *EEncryptionPasswordsFiltersOrderColumn {
	return &v
}

type NullableEEncryptionPasswordsFiltersOrderColumn struct {
	value *EEncryptionPasswordsFiltersOrderColumn
	isSet bool
}

func (v NullableEEncryptionPasswordsFiltersOrderColumn) Get() *EEncryptionPasswordsFiltersOrderColumn {
	return v.value
}

func (v *NullableEEncryptionPasswordsFiltersOrderColumn) Set(val *EEncryptionPasswordsFiltersOrderColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableEEncryptionPasswordsFiltersOrderColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableEEncryptionPasswordsFiltersOrderColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEEncryptionPasswordsFiltersOrderColumn(val *EEncryptionPasswordsFiltersOrderColumn) *NullableEEncryptionPasswordsFiltersOrderColumn {
	return &NullableEEncryptionPasswordsFiltersOrderColumn{value: val, isSet: true}
}

func (v NullableEEncryptionPasswordsFiltersOrderColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEEncryptionPasswordsFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

