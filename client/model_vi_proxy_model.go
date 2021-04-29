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
)

// ViProxyModel struct for ViProxyModel
type ViProxyModel struct {
	ProxyModel
	Server ProxyServerSettingsModel `json:"server"`
}

// NewViProxyModel instantiates a new ViProxyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViProxyModel(server ProxyServerSettingsModel, id string, name string, description string, type_ EProxyType) *ViProxyModel {
	this := ViProxyModel{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Type = type_
	this.Server = server
	return &this
}

// NewViProxyModelWithDefaults instantiates a new ViProxyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViProxyModelWithDefaults() *ViProxyModel {
	this := ViProxyModel{}
	return &this
}

// GetServer returns the Server field value
func (o *ViProxyModel) GetServer() ProxyServerSettingsModel {
	if o == nil {
		var ret ProxyServerSettingsModel
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *ViProxyModel) GetServerOk() (*ProxyServerSettingsModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *ViProxyModel) SetServer(v ProxyServerSettingsModel) {
	o.Server = v
}

func (o ViProxyModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedProxyModel, errProxyModel := json.Marshal(o.ProxyModel)
	if errProxyModel != nil {
		return []byte{}, errProxyModel
	}
	errProxyModel = json.Unmarshal([]byte(serializedProxyModel), &toSerialize)
	if errProxyModel != nil {
		return []byte{}, errProxyModel
	}
	if true {
		toSerialize["server"] = o.Server
	}
	return json.Marshal(toSerialize)
}

type NullableViProxyModel struct {
	value *ViProxyModel
	isSet bool
}

func (v NullableViProxyModel) Get() *ViProxyModel {
	return v.value
}

func (v *NullableViProxyModel) Set(val *ViProxyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableViProxyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableViProxyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViProxyModel(val *ViProxyModel) *NullableViProxyModel {
	return &NullableViProxyModel{value: val, isSet: true}
}

func (v NullableViProxyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViProxyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


