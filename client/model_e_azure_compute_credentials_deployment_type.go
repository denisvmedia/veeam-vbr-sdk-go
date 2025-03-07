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

// EAzureComputeCredentialsDeploymentType Deployment type (global Microsoft Azure or Microsoft Azure Stack Hub).
type EAzureComputeCredentialsDeploymentType string

// List of EAzureComputeCredentialsDeploymentType
const (
	EAZURECOMPUTECREDENTIALSDEPLOYMENTTYPE_MICROSOFT_AZURE EAzureComputeCredentialsDeploymentType = "MicrosoftAzure"
	EAZURECOMPUTECREDENTIALSDEPLOYMENTTYPE_MICROSOFT_AZURE_STACK EAzureComputeCredentialsDeploymentType = "MicrosoftAzureStack"
)

// All allowed values of EAzureComputeCredentialsDeploymentType enum
var AllowedEAzureComputeCredentialsDeploymentTypeEnumValues = []EAzureComputeCredentialsDeploymentType{
	"MicrosoftAzure",
	"MicrosoftAzureStack",
}

func (v *EAzureComputeCredentialsDeploymentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EAzureComputeCredentialsDeploymentType(value)
	for _, existing := range AllowedEAzureComputeCredentialsDeploymentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EAzureComputeCredentialsDeploymentType", value)
}

// NewEAzureComputeCredentialsDeploymentTypeFromValue returns a pointer to a valid EAzureComputeCredentialsDeploymentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEAzureComputeCredentialsDeploymentTypeFromValue(v string) (*EAzureComputeCredentialsDeploymentType, error) {
	ev := EAzureComputeCredentialsDeploymentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EAzureComputeCredentialsDeploymentType: valid values are %v", v, AllowedEAzureComputeCredentialsDeploymentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EAzureComputeCredentialsDeploymentType) IsValid() bool {
	for _, existing := range AllowedEAzureComputeCredentialsDeploymentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EAzureComputeCredentialsDeploymentType value
func (v EAzureComputeCredentialsDeploymentType) Ptr() *EAzureComputeCredentialsDeploymentType {
	return &v
}

type NullableEAzureComputeCredentialsDeploymentType struct {
	value *EAzureComputeCredentialsDeploymentType
	isSet bool
}

func (v NullableEAzureComputeCredentialsDeploymentType) Get() *EAzureComputeCredentialsDeploymentType {
	return v.value
}

func (v *NullableEAzureComputeCredentialsDeploymentType) Set(val *EAzureComputeCredentialsDeploymentType) {
	v.value = val
	v.isSet = true
}

func (v NullableEAzureComputeCredentialsDeploymentType) IsSet() bool {
	return v.isSet
}

func (v *NullableEAzureComputeCredentialsDeploymentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEAzureComputeCredentialsDeploymentType(val *EAzureComputeCredentialsDeploymentType) *NullableEAzureComputeCredentialsDeploymentType {
	return &NullableEAzureComputeCredentialsDeploymentType{value: val, isSet: true}
}

func (v NullableEAzureComputeCredentialsDeploymentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEAzureComputeCredentialsDeploymentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

