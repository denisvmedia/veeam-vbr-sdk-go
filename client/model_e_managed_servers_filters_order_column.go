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

// EManagedServersFiltersOrderColumn the model 'EManagedServersFiltersOrderColumn'
type EManagedServersFiltersOrderColumn string

// List of EManagedServersFiltersOrderColumn
const (
	EMANAGEDSERVERSFILTERSORDERCOLUMN_NAME EManagedServersFiltersOrderColumn = "Name"
	EMANAGEDSERVERSFILTERSORDERCOLUMN_TYPE EManagedServersFiltersOrderColumn = "Type"
	EMANAGEDSERVERSFILTERSORDERCOLUMN_DESCRIPTION EManagedServersFiltersOrderColumn = "Description"
)

// All allowed values of EManagedServersFiltersOrderColumn enum
var AllowedEManagedServersFiltersOrderColumnEnumValues = []EManagedServersFiltersOrderColumn{
	"Name",
	"Type",
	"Description",
}

func (v *EManagedServersFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EManagedServersFiltersOrderColumn(value)
	for _, existing := range AllowedEManagedServersFiltersOrderColumnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EManagedServersFiltersOrderColumn", value)
}

// NewEManagedServersFiltersOrderColumnFromValue returns a pointer to a valid EManagedServersFiltersOrderColumn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEManagedServersFiltersOrderColumnFromValue(v string) (*EManagedServersFiltersOrderColumn, error) {
	ev := EManagedServersFiltersOrderColumn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EManagedServersFiltersOrderColumn: valid values are %v", v, AllowedEManagedServersFiltersOrderColumnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EManagedServersFiltersOrderColumn) IsValid() bool {
	for _, existing := range AllowedEManagedServersFiltersOrderColumnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EManagedServersFiltersOrderColumn value
func (v EManagedServersFiltersOrderColumn) Ptr() *EManagedServersFiltersOrderColumn {
	return &v
}

type NullableEManagedServersFiltersOrderColumn struct {
	value *EManagedServersFiltersOrderColumn
	isSet bool
}

func (v NullableEManagedServersFiltersOrderColumn) Get() *EManagedServersFiltersOrderColumn {
	return v.value
}

func (v *NullableEManagedServersFiltersOrderColumn) Set(val *EManagedServersFiltersOrderColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableEManagedServersFiltersOrderColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableEManagedServersFiltersOrderColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEManagedServersFiltersOrderColumn(val *EManagedServersFiltersOrderColumn) *NullableEManagedServersFiltersOrderColumn {
	return &NullableEManagedServersFiltersOrderColumn{value: val, isSet: true}
}

func (v NullableEManagedServersFiltersOrderColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEManagedServersFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

