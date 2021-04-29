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

// EManagedServerType Type of the server.
type EManagedServerType string

// List of EManagedServerType
const (
	EMANAGEDSERVERTYPE_WINDOWS_HOST EManagedServerType = "WindowsHost"
	EMANAGEDSERVERTYPE_LINUX_HOST EManagedServerType = "LinuxHost"
	EMANAGEDSERVERTYPE_VI_HOST EManagedServerType = "ViHost"
)

func (v *EManagedServerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EManagedServerType(value)
	for _, existing := range []EManagedServerType{ "WindowsHost", "LinuxHost", "ViHost",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EManagedServerType", value)
}

// Ptr returns reference to EManagedServerType value
func (v EManagedServerType) Ptr() *EManagedServerType {
	return &v
}

type NullableEManagedServerType struct {
	value *EManagedServerType
	isSet bool
}

func (v NullableEManagedServerType) Get() *EManagedServerType {
	return v.value
}

func (v *NullableEManagedServerType) Set(val *EManagedServerType) {
	v.value = val
	v.isSet = true
}

func (v NullableEManagedServerType) IsSet() bool {
	return v.isSet
}

func (v *NullableEManagedServerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEManagedServerType(val *EManagedServerType) *NullableEManagedServerType {
	return &NullableEManagedServerType{value: val, isSet: true}
}

func (v NullableEManagedServerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEManagedServerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

