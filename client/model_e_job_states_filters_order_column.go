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

// EJobStatesFiltersOrderColumn Orders job states by the specified column.
type EJobStatesFiltersOrderColumn string

// List of EJobStatesFiltersOrderColumn
const (
	EJOBSTATESFILTERSORDERCOLUMN_NAME EJobStatesFiltersOrderColumn = "Name"
	EJOBSTATESFILTERSORDERCOLUMN_TYPE EJobStatesFiltersOrderColumn = "Type"
	EJOBSTATESFILTERSORDERCOLUMN_STATUS EJobStatesFiltersOrderColumn = "Status"
	EJOBSTATESFILTERSORDERCOLUMN_LAST_RUN EJobStatesFiltersOrderColumn = "LastRun"
	EJOBSTATESFILTERSORDERCOLUMN_LAST_RESULT EJobStatesFiltersOrderColumn = "LastResult"
	EJOBSTATESFILTERSORDERCOLUMN_NEXT_RUN EJobStatesFiltersOrderColumn = "NextRun"
	EJOBSTATESFILTERSORDERCOLUMN_DESCRIPTION EJobStatesFiltersOrderColumn = "Description"
	EJOBSTATESFILTERSORDERCOLUMN_REPOSITORY_ID EJobStatesFiltersOrderColumn = "RepositoryId"
	EJOBSTATESFILTERSORDERCOLUMN_OBJECTS_COUNT EJobStatesFiltersOrderColumn = "ObjectsCount"
)

// All allowed values of EJobStatesFiltersOrderColumn enum
var AllowedEJobStatesFiltersOrderColumnEnumValues = []EJobStatesFiltersOrderColumn{
	"Name",
	"Type",
	"Status",
	"LastRun",
	"LastResult",
	"NextRun",
	"Description",
	"RepositoryId",
	"ObjectsCount",
}

func (v *EJobStatesFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EJobStatesFiltersOrderColumn(value)
	for _, existing := range AllowedEJobStatesFiltersOrderColumnEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EJobStatesFiltersOrderColumn", value)
}

// NewEJobStatesFiltersOrderColumnFromValue returns a pointer to a valid EJobStatesFiltersOrderColumn
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEJobStatesFiltersOrderColumnFromValue(v string) (*EJobStatesFiltersOrderColumn, error) {
	ev := EJobStatesFiltersOrderColumn(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EJobStatesFiltersOrderColumn: valid values are %v", v, AllowedEJobStatesFiltersOrderColumnEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EJobStatesFiltersOrderColumn) IsValid() bool {
	for _, existing := range AllowedEJobStatesFiltersOrderColumnEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EJobStatesFiltersOrderColumn value
func (v EJobStatesFiltersOrderColumn) Ptr() *EJobStatesFiltersOrderColumn {
	return &v
}

type NullableEJobStatesFiltersOrderColumn struct {
	value *EJobStatesFiltersOrderColumn
	isSet bool
}

func (v NullableEJobStatesFiltersOrderColumn) Get() *EJobStatesFiltersOrderColumn {
	return v.value
}

func (v *NullableEJobStatesFiltersOrderColumn) Set(val *EJobStatesFiltersOrderColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableEJobStatesFiltersOrderColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableEJobStatesFiltersOrderColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEJobStatesFiltersOrderColumn(val *EJobStatesFiltersOrderColumn) *NullableEJobStatesFiltersOrderColumn {
	return &NullableEJobStatesFiltersOrderColumn{value: val, isSet: true}
}

func (v NullableEJobStatesFiltersOrderColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEJobStatesFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

