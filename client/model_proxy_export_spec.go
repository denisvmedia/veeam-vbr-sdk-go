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

// ProxyExportSpec struct for ProxyExportSpec
type ProxyExportSpec struct {
	// Array of backup proxy IDs.
	Ids *[]string `json:"ids,omitempty"`
	// Array of backup proxy types.
	Types *[]EProxyType `json:"types,omitempty"`
	// Array of backup proxy names. Wildcard characters are supported.
	Names *[]string `json:"names,omitempty"`
}

// NewProxyExportSpec instantiates a new ProxyExportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyExportSpec() *ProxyExportSpec {
	this := ProxyExportSpec{}
	return &this
}

// NewProxyExportSpecWithDefaults instantiates a new ProxyExportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyExportSpecWithDefaults() *ProxyExportSpec {
	this := ProxyExportSpec{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *ProxyExportSpec) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyExportSpec) GetIdsOk() (*[]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *ProxyExportSpec) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *ProxyExportSpec) SetIds(v []string) {
	o.Ids = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *ProxyExportSpec) GetTypes() []EProxyType {
	if o == nil || o.Types == nil {
		var ret []EProxyType
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyExportSpec) GetTypesOk() (*[]EProxyType, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *ProxyExportSpec) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []EProxyType and assigns it to the Types field.
func (o *ProxyExportSpec) SetTypes(v []EProxyType) {
	o.Types = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *ProxyExportSpec) GetNames() []string {
	if o == nil || o.Names == nil {
		var ret []string
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyExportSpec) GetNamesOk() (*[]string, bool) {
	if o == nil || o.Names == nil {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *ProxyExportSpec) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given []string and assigns it to the Names field.
func (o *ProxyExportSpec) SetNames(v []string) {
	o.Names = &v
}

func (o ProxyExportSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	if o.Names != nil {
		toSerialize["names"] = o.Names
	}
	return json.Marshal(toSerialize)
}

type NullableProxyExportSpec struct {
	value *ProxyExportSpec
	isSet bool
}

func (v NullableProxyExportSpec) Get() *ProxyExportSpec {
	return v.value
}

func (v *NullableProxyExportSpec) Set(val *ProxyExportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyExportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyExportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyExportSpec(val *ProxyExportSpec) *NullableProxyExportSpec {
	return &NullableProxyExportSpec{value: val, isSet: true}
}

func (v NullableProxyExportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyExportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


