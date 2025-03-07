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

// CloudHelperApplianceSpec Helper appliance settings.
type CloudHelperApplianceSpec struct {
	Type ECloudCredentialsType `json:"type"`
}

// NewCloudHelperApplianceSpec instantiates a new CloudHelperApplianceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudHelperApplianceSpec(type_ ECloudCredentialsType) *CloudHelperApplianceSpec {
	this := CloudHelperApplianceSpec{}
	this.Type = type_
	return &this
}

// NewCloudHelperApplianceSpecWithDefaults instantiates a new CloudHelperApplianceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudHelperApplianceSpecWithDefaults() *CloudHelperApplianceSpec {
	this := CloudHelperApplianceSpec{}
	return &this
}

// GetType returns the Type field value
func (o *CloudHelperApplianceSpec) GetType() ECloudCredentialsType {
	if o == nil {
		var ret ECloudCredentialsType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CloudHelperApplianceSpec) GetTypeOk() (*ECloudCredentialsType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CloudHelperApplianceSpec) SetType(v ECloudCredentialsType) {
	o.Type = v
}

func (o CloudHelperApplianceSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCloudHelperApplianceSpec struct {
	value *CloudHelperApplianceSpec
	isSet bool
}

func (v NullableCloudHelperApplianceSpec) Get() *CloudHelperApplianceSpec {
	return v.value
}

func (v *NullableCloudHelperApplianceSpec) Set(val *CloudHelperApplianceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudHelperApplianceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudHelperApplianceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudHelperApplianceSpec(val *CloudHelperApplianceSpec) *NullableCloudHelperApplianceSpec {
	return &NullableCloudHelperApplianceSpec{value: val, isSet: true}
}

func (v NullableCloudHelperApplianceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudHelperApplianceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


