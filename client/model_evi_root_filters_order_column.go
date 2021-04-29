/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 
 *
 * API version: 1.0-rev1
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// EViRootFiltersOrderColumn Sorts VMware vSphere servers by one of the VMware vSphere server parameters.
type EViRootFiltersOrderColumn string

// List of EViRootFiltersOrderColumn
const (
	EVIROOTFILTERSORDERCOLUMN_NAME EViRootFiltersOrderColumn = "Name"
	EVIROOTFILTERSORDERCOLUMN_TYPE EViRootFiltersOrderColumn = "Type"
)

func (v *EViRootFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EViRootFiltersOrderColumn(value)
	for _, existing := range []EViRootFiltersOrderColumn{ "Name", "Type",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EViRootFiltersOrderColumn", value)
}

// Ptr returns reference to EViRootFiltersOrderColumn value
func (v EViRootFiltersOrderColumn) Ptr() *EViRootFiltersOrderColumn {
	return &v
}

type NullableEViRootFiltersOrderColumn struct {
	value *EViRootFiltersOrderColumn
	isSet bool
}

func (v NullableEViRootFiltersOrderColumn) Get() *EViRootFiltersOrderColumn {
	return v.value
}

func (v *NullableEViRootFiltersOrderColumn) Set(val *EViRootFiltersOrderColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableEViRootFiltersOrderColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableEViRootFiltersOrderColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEViRootFiltersOrderColumn(val *EViRootFiltersOrderColumn) *NullableEViRootFiltersOrderColumn {
	return &NullableEViRootFiltersOrderColumn{value: val, isSet: true}
}

func (v NullableEViRootFiltersOrderColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEViRootFiltersOrderColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

