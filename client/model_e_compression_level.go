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

// ECompressionLevel Compression level.
type ECompressionLevel string

// List of ECompressionLevel
const (
	ECOMPRESSIONLEVEL_NONE ECompressionLevel = "None"
	ECOMPRESSIONLEVEL_DEDUP_FRIENDLY ECompressionLevel = "DedupFriendly"
	ECOMPRESSIONLEVEL_OPTIMAL ECompressionLevel = "Optimal"
	ECOMPRESSIONLEVEL_HIGH ECompressionLevel = "High"
	ECOMPRESSIONLEVEL_EXTREME ECompressionLevel = "Extreme"
)

// All allowed values of ECompressionLevel enum
var AllowedECompressionLevelEnumValues = []ECompressionLevel{
	"None",
	"DedupFriendly",
	"Optimal",
	"High",
	"Extreme",
}

func (v *ECompressionLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ECompressionLevel(value)
	for _, existing := range AllowedECompressionLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ECompressionLevel", value)
}

// NewECompressionLevelFromValue returns a pointer to a valid ECompressionLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewECompressionLevelFromValue(v string) (*ECompressionLevel, error) {
	ev := ECompressionLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ECompressionLevel: valid values are %v", v, AllowedECompressionLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ECompressionLevel) IsValid() bool {
	for _, existing := range AllowedECompressionLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ECompressionLevel value
func (v ECompressionLevel) Ptr() *ECompressionLevel {
	return &v
}

type NullableECompressionLevel struct {
	value *ECompressionLevel
	isSet bool
}

func (v NullableECompressionLevel) Get() *ECompressionLevel {
	return v.value
}

func (v *NullableECompressionLevel) Set(val *ECompressionLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableECompressionLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableECompressionLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECompressionLevel(val *ECompressionLevel) *NullableECompressionLevel {
	return &NullableECompressionLevel{value: val, isSet: true}
}

func (v NullableECompressionLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECompressionLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

