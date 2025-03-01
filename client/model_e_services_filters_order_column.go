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

// EServicesFiltersOrderColumn the model 'EServicesFiltersOrderColumn'
type EServicesFiltersOrderColumn string

// List of EServicesFiltersOrderColumn
const (
	ESERVICESFILTERSORDERCOLUMN_NAME EServicesFiltersOrderColumn = "Name"
	ESERVICESFILTERSORDERCOLUMN_PORT EServicesFiltersOrderColumn = "Port"
)

// All allowed values of EServicesFiltersOrderColumn enum
var AllowedEServicesFiltersOrderColumnEnumValues = []EServicesFiltersOrderColumn{
	"Name",
	"Port",
}

func (v *EServicesFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EServicesFiltersOrderColumn(value)
	for _, existing := range AllowedEServicesFiltersOrderColumnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EServicesFiltersOrderColumn", value)
}

// NewEServicesFiltersOrderColumnFromValue returns a pointer to a valid EServicesFiltersOrderColumn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEServicesFiltersOrderColumnFromValue(v string) (*EServicesFiltersOrderColumn, error) {
	ev := EServicesFiltersOrderColumn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EServicesFiltersOrderColumn: valid values are %v", v, AllowedEServicesFiltersOrderColumnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EServicesFiltersOrderColumn) IsValid() bool {
	for _, existing := range AllowedEServicesFiltersOrderColumnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EServicesFiltersOrderColumn value
func (v EServicesFiltersOrderColumn) Ptr() *EServicesFiltersOrderColumn {
	return &v
}

type NullableEServicesFiltersOrderColumn struct {
	value *EServicesFiltersOrderColumn
	isSet bool
}

func (v NullableEServicesFiltersOrderColumn) Get() *EServicesFiltersOrderColumn {
	return v.value
}

func (v *NullableEServicesFiltersOrderColumn) Set(val *EServicesFiltersOrderColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableEServicesFiltersOrderColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableEServicesFiltersOrderColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEServicesFiltersOrderColumn(val *EServicesFiltersOrderColumn) *NullableEServicesFiltersOrderColumn {
	return &NullableEServicesFiltersOrderColumn{value: val, isSet: true}
}

func (v NullableEServicesFiltersOrderColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEServicesFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

