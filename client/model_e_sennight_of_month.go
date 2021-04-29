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

// ESennightOfMonth Sennight of the month.
type ESennightOfMonth string

// List of ESennightOfMonth
const (
	ESENNIGHTOFMONTH_FIRST ESennightOfMonth = "First"
	ESENNIGHTOFMONTH_SECOND ESennightOfMonth = "Second"
	ESENNIGHTOFMONTH_THIRD ESennightOfMonth = "Third"
	ESENNIGHTOFMONTH_FOURTH ESennightOfMonth = "Fourth"
	ESENNIGHTOFMONTH_FIFTH ESennightOfMonth = "Fifth"
	ESENNIGHTOFMONTH_LAST ESennightOfMonth = "Last"
)

func (v *ESennightOfMonth) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ESennightOfMonth(value)
	for _, existing := range []ESennightOfMonth{ "First", "Second", "Third", "Fourth", "Fifth", "Last",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ESennightOfMonth", value)
}

// Ptr returns reference to ESennightOfMonth value
func (v ESennightOfMonth) Ptr() *ESennightOfMonth {
	return &v
}

type NullableESennightOfMonth struct {
	value *ESennightOfMonth
	isSet bool
}

func (v NullableESennightOfMonth) Get() *ESennightOfMonth {
	return v.value
}

func (v *NullableESennightOfMonth) Set(val *ESennightOfMonth) {
	v.value = val
	v.isSet = true
}

func (v NullableESennightOfMonth) IsSet() bool {
	return v.isSet
}

func (v *NullableESennightOfMonth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableESennightOfMonth(val *ESennightOfMonth) *NullableESennightOfMonth {
	return &NullableESennightOfMonth{value: val, isSet: true}
}

func (v NullableESennightOfMonth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableESennightOfMonth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

