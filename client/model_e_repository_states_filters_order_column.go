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

// ERepositoryStatesFiltersOrderColumn Orders repositories by the specified column.
type ERepositoryStatesFiltersOrderColumn string

// List of ERepositoryStatesFiltersOrderColumn
const (
	EREPOSITORYSTATESFILTERSORDERCOLUMN_NAME ERepositoryStatesFiltersOrderColumn = "Name"
	EREPOSITORYSTATESFILTERSORDERCOLUMN_TYPE ERepositoryStatesFiltersOrderColumn = "Type"
	EREPOSITORYSTATESFILTERSORDERCOLUMN_HOST ERepositoryStatesFiltersOrderColumn = "Host"
	EREPOSITORYSTATESFILTERSORDERCOLUMN_PATH ERepositoryStatesFiltersOrderColumn = "Path"
	EREPOSITORYSTATESFILTERSORDERCOLUMN_CAPACITY_GB ERepositoryStatesFiltersOrderColumn = "CapacityGB"
	EREPOSITORYSTATESFILTERSORDERCOLUMN_FREE_GB ERepositoryStatesFiltersOrderColumn = "FreeGB"
	EREPOSITORYSTATESFILTERSORDERCOLUMN_USED_SPACE_GB ERepositoryStatesFiltersOrderColumn = "UsedSpaceGB"
	EREPOSITORYSTATESFILTERSORDERCOLUMN_DESCRIPTION ERepositoryStatesFiltersOrderColumn = "Description"
)

// All allowed values of ERepositoryStatesFiltersOrderColumn enum
var AllowedERepositoryStatesFiltersOrderColumnEnumValues = []ERepositoryStatesFiltersOrderColumn{
	"Name",
	"Type",
	"Host",
	"Path",
	"CapacityGB",
	"FreeGB",
	"UsedSpaceGB",
	"Description",
}

func (v *ERepositoryStatesFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ERepositoryStatesFiltersOrderColumn(value)
	for _, existing := range AllowedERepositoryStatesFiltersOrderColumnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ERepositoryStatesFiltersOrderColumn", value)
}

// NewERepositoryStatesFiltersOrderColumnFromValue returns a pointer to a valid ERepositoryStatesFiltersOrderColumn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewERepositoryStatesFiltersOrderColumnFromValue(v string) (*ERepositoryStatesFiltersOrderColumn, error) {
	ev := ERepositoryStatesFiltersOrderColumn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ERepositoryStatesFiltersOrderColumn: valid values are %v", v, AllowedERepositoryStatesFiltersOrderColumnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ERepositoryStatesFiltersOrderColumn) IsValid() bool {
	for _, existing := range AllowedERepositoryStatesFiltersOrderColumnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ERepositoryStatesFiltersOrderColumn value
func (v ERepositoryStatesFiltersOrderColumn) Ptr() *ERepositoryStatesFiltersOrderColumn {
	return &v
}

type NullableERepositoryStatesFiltersOrderColumn struct {
	value *ERepositoryStatesFiltersOrderColumn
	isSet bool
}

func (v NullableERepositoryStatesFiltersOrderColumn) Get() *ERepositoryStatesFiltersOrderColumn {
	return v.value
}

func (v *NullableERepositoryStatesFiltersOrderColumn) Set(val *ERepositoryStatesFiltersOrderColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableERepositoryStatesFiltersOrderColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableERepositoryStatesFiltersOrderColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableERepositoryStatesFiltersOrderColumn(val *ERepositoryStatesFiltersOrderColumn) *NullableERepositoryStatesFiltersOrderColumn {
	return &NullableERepositoryStatesFiltersOrderColumn{value: val, isSet: true}
}

func (v NullableERepositoryStatesFiltersOrderColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableERepositoryStatesFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

