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

// ERepositoryExtentStatusType Performance extent status.
type ERepositoryExtentStatusType string

// List of ERepositoryExtentStatusType
const (
	EREPOSITORYEXTENTSTATUSTYPE_NORMAL ERepositoryExtentStatusType = "Normal"
	EREPOSITORYEXTENTSTATUSTYPE_EVACUATE ERepositoryExtentStatusType = "Evacuate"
	EREPOSITORYEXTENTSTATUSTYPE_PENDING ERepositoryExtentStatusType = "Pending"
	EREPOSITORYEXTENTSTATUSTYPE_SEALED ERepositoryExtentStatusType = "Sealed"
	EREPOSITORYEXTENTSTATUSTYPE_MAINTENANCE ERepositoryExtentStatusType = "Maintenance"
	EREPOSITORYEXTENTSTATUSTYPE_RESYNC_REQUIRED ERepositoryExtentStatusType = "ResyncRequired"
	EREPOSITORYEXTENTSTATUSTYPE_TENANT_EVACUATING ERepositoryExtentStatusType = "TenantEvacuating"
)

// All allowed values of ERepositoryExtentStatusType enum
var AllowedERepositoryExtentStatusTypeEnumValues = []ERepositoryExtentStatusType{
	"Normal",
	"Evacuate",
	"Pending",
	"Sealed",
	"Maintenance",
	"ResyncRequired",
	"TenantEvacuating",
}

func (v *ERepositoryExtentStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ERepositoryExtentStatusType(value)
	for _, existing := range AllowedERepositoryExtentStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ERepositoryExtentStatusType", value)
}

// NewERepositoryExtentStatusTypeFromValue returns a pointer to a valid ERepositoryExtentStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewERepositoryExtentStatusTypeFromValue(v string) (*ERepositoryExtentStatusType, error) {
	ev := ERepositoryExtentStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ERepositoryExtentStatusType: valid values are %v", v, AllowedERepositoryExtentStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ERepositoryExtentStatusType) IsValid() bool {
	for _, existing := range AllowedERepositoryExtentStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ERepositoryExtentStatusType value
func (v ERepositoryExtentStatusType) Ptr() *ERepositoryExtentStatusType {
	return &v
}

type NullableERepositoryExtentStatusType struct {
	value *ERepositoryExtentStatusType
	isSet bool
}

func (v NullableERepositoryExtentStatusType) Get() *ERepositoryExtentStatusType {
	return v.value
}

func (v *NullableERepositoryExtentStatusType) Set(val *ERepositoryExtentStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableERepositoryExtentStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableERepositoryExtentStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableERepositoryExtentStatusType(val *ERepositoryExtentStatusType) *NullableERepositoryExtentStatusType {
	return &NullableERepositoryExtentStatusType{value: val, isSet: true}
}

func (v NullableERepositoryExtentStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableERepositoryExtentStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

