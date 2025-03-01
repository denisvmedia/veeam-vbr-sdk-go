/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br>Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Some request and response bodies refer to reusable schema objects that are defined in the `schemas` section of the REST API specification. Schemas, in turn, may inherit a part of their properties by referring to other schemas, and be polymorphic by referring to multiple alternate schemas.

API version: 1.1-rev0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CloudDeviceCodeSpec struct for CloudDeviceCodeSpec
type CloudDeviceCodeSpec struct {
	Type ECloudCredentialsType `json:"type"`
}

// NewCloudDeviceCodeSpec instantiates a new CloudDeviceCodeSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudDeviceCodeSpec(type_ ECloudCredentialsType) *CloudDeviceCodeSpec {
	this := CloudDeviceCodeSpec{}
	this.Type = type_
	return &this
}

// NewCloudDeviceCodeSpecWithDefaults instantiates a new CloudDeviceCodeSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudDeviceCodeSpecWithDefaults() *CloudDeviceCodeSpec {
	this := CloudDeviceCodeSpec{}
	return &this
}

// GetType returns the Type field value
func (o *CloudDeviceCodeSpec) GetType() ECloudCredentialsType {
	if o == nil {
		var ret ECloudCredentialsType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CloudDeviceCodeSpec) GetTypeOk() (*ECloudCredentialsType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CloudDeviceCodeSpec) SetType(v ECloudCredentialsType) {
	o.Type = v
}

func (o CloudDeviceCodeSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCloudDeviceCodeSpec struct {
	value *CloudDeviceCodeSpec
	isSet bool
}

func (v NullableCloudDeviceCodeSpec) Get() *CloudDeviceCodeSpec {
	return v.value
}

func (v *NullableCloudDeviceCodeSpec) Set(val *CloudDeviceCodeSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudDeviceCodeSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudDeviceCodeSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudDeviceCodeSpec(val *CloudDeviceCodeSpec) *NullableCloudDeviceCodeSpec {
	return &NullableCloudDeviceCodeSpec{value: val, isSet: true}
}

func (v NullableCloudDeviceCodeSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudDeviceCodeSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


