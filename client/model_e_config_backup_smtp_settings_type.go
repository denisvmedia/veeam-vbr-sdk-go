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

// EConfigBackupSMTPSettingsType Type of notification settings.
type EConfigBackupSMTPSettingsType string

// List of EConfigBackupSMTPSettingsType
const (
	ECONFIGBACKUPSMTPSETTINGSTYPE_GLOBAL EConfigBackupSMTPSettingsType = "Global"
	ECONFIGBACKUPSMTPSETTINGSTYPE_CUSTOM EConfigBackupSMTPSettingsType = "Custom"
)

// All allowed values of EConfigBackupSMTPSettingsType enum
var AllowedEConfigBackupSMTPSettingsTypeEnumValues = []EConfigBackupSMTPSettingsType{
	"Global",
	"Custom",
}

func (v *EConfigBackupSMTPSettingsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EConfigBackupSMTPSettingsType(value)
	for _, existing := range AllowedEConfigBackupSMTPSettingsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EConfigBackupSMTPSettingsType", value)
}

// NewEConfigBackupSMTPSettingsTypeFromValue returns a pointer to a valid EConfigBackupSMTPSettingsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEConfigBackupSMTPSettingsTypeFromValue(v string) (*EConfigBackupSMTPSettingsType, error) {
	ev := EConfigBackupSMTPSettingsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EConfigBackupSMTPSettingsType: valid values are %v", v, AllowedEConfigBackupSMTPSettingsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EConfigBackupSMTPSettingsType) IsValid() bool {
	for _, existing := range AllowedEConfigBackupSMTPSettingsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EConfigBackupSMTPSettingsType value
func (v EConfigBackupSMTPSettingsType) Ptr() *EConfigBackupSMTPSettingsType {
	return &v
}

type NullableEConfigBackupSMTPSettingsType struct {
	value *EConfigBackupSMTPSettingsType
	isSet bool
}

func (v NullableEConfigBackupSMTPSettingsType) Get() *EConfigBackupSMTPSettingsType {
	return v.value
}

func (v *NullableEConfigBackupSMTPSettingsType) Set(val *EConfigBackupSMTPSettingsType) {
	v.value = val
	v.isSet = true
}

func (v NullableEConfigBackupSMTPSettingsType) IsSet() bool {
	return v.isSet
}

func (v *NullableEConfigBackupSMTPSettingsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEConfigBackupSMTPSettingsType(val *EConfigBackupSMTPSettingsType) *NullableEConfigBackupSMTPSettingsType {
	return &NullableEConfigBackupSMTPSettingsType{value: val, isSet: true}
}

func (v NullableEConfigBackupSMTPSettingsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEConfigBackupSMTPSettingsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

