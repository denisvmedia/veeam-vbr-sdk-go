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

// ProxySpec struct for ProxySpec
type ProxySpec struct {
	// Description of the backup proxy.
	Description string `json:"description"`
	Type EProxyType `json:"type"`
}

// NewProxySpec instantiates a new ProxySpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxySpec(description string, type_ EProxyType) *ProxySpec {
	this := ProxySpec{}
	this.Description = description
	this.Type = type_
	return &this
}

// NewProxySpecWithDefaults instantiates a new ProxySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxySpecWithDefaults() *ProxySpec {
	this := ProxySpec{}
	return &this
}

// GetDescription returns the Description field value
func (o *ProxySpec) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ProxySpec) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ProxySpec) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value
func (o *ProxySpec) GetType() EProxyType {
	if o == nil {
		var ret EProxyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProxySpec) GetTypeOk() (*EProxyType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ProxySpec) SetType(v EProxyType) {
	o.Type = v
}

func (o ProxySpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableProxySpec struct {
	value *ProxySpec
	isSet bool
}

func (v NullableProxySpec) Get() *ProxySpec {
	return v.value
}

func (v *NullableProxySpec) Set(val *ProxySpec) {
	v.value = val
	v.isSet = true
}

func (v NullableProxySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableProxySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxySpec(val *ProxySpec) *NullableProxySpec {
	return &NullableProxySpec{value: val, isSet: true}
}

func (v NullableProxySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


