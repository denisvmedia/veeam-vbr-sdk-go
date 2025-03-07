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

// ViProxySpec struct for ViProxySpec
type ViProxySpec struct {
	ProxySpec
	Server ProxyServerSettingsModel `json:"server"`
}

// NewViProxySpec instantiates a new ViProxySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViProxySpec(server ProxyServerSettingsModel, description string, type_ EProxyType) *ViProxySpec {
	this := ViProxySpec{}
	this.Description = description
	this.Type = type_
	this.Server = server
	return &this
}

// NewViProxySpecWithDefaults instantiates a new ViProxySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViProxySpecWithDefaults() *ViProxySpec {
	this := ViProxySpec{}
	return &this
}

// GetServer returns the Server field value
func (o *ViProxySpec) GetServer() ProxyServerSettingsModel {
	if o == nil {
		var ret ProxyServerSettingsModel
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *ViProxySpec) GetServerOk() (*ProxyServerSettingsModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *ViProxySpec) SetServer(v ProxyServerSettingsModel) {
	o.Server = v
}

func (o ViProxySpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedProxySpec, errProxySpec := json.Marshal(o.ProxySpec)
	if errProxySpec != nil {
		return []byte{}, errProxySpec
	}
	errProxySpec = json.Unmarshal([]byte(serializedProxySpec), &toSerialize)
	if errProxySpec != nil {
		return []byte{}, errProxySpec
	}
	if true {
		toSerialize["server"] = o.Server
	}
	return json.Marshal(toSerialize)
}

type NullableViProxySpec struct {
	value *ViProxySpec
	isSet bool
}

func (v NullableViProxySpec) Get() *ViProxySpec {
	return v.value
}

func (v *NullableViProxySpec) Set(val *ViProxySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableViProxySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableViProxySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViProxySpec(val *ViProxySpec) *NullableViProxySpec {
	return &NullableViProxySpec{value: val, isSet: true}
}

func (v NullableViProxySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViProxySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


