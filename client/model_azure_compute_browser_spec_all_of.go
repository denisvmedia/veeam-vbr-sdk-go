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

// AzureComputeBrowserSpecAllOf struct for AzureComputeBrowserSpecAllOf
type AzureComputeBrowserSpecAllOf struct {
	Filters *AzureComputeBrowserFilters `json:"filters,omitempty"`
}

// NewAzureComputeBrowserSpecAllOf instantiates a new AzureComputeBrowserSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureComputeBrowserSpecAllOf() *AzureComputeBrowserSpecAllOf {
	this := AzureComputeBrowserSpecAllOf{}
	return &this
}

// NewAzureComputeBrowserSpecAllOfWithDefaults instantiates a new AzureComputeBrowserSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureComputeBrowserSpecAllOfWithDefaults() *AzureComputeBrowserSpecAllOf {
	this := AzureComputeBrowserSpecAllOf{}
	return &this
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *AzureComputeBrowserSpecAllOf) GetFilters() AzureComputeBrowserFilters {
	if o == nil || isNil(o.Filters) {
		var ret AzureComputeBrowserFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureComputeBrowserSpecAllOf) GetFiltersOk() (*AzureComputeBrowserFilters, bool) {
	if o == nil || isNil(o.Filters) {
    return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *AzureComputeBrowserSpecAllOf) HasFilters() bool {
	if o != nil && !isNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given AzureComputeBrowserFilters and assigns it to the Filters field.
func (o *AzureComputeBrowserSpecAllOf) SetFilters(v AzureComputeBrowserFilters) {
	o.Filters = &v
}

func (o AzureComputeBrowserSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

type NullableAzureComputeBrowserSpecAllOf struct {
	value *AzureComputeBrowserSpecAllOf
	isSet bool
}

func (v NullableAzureComputeBrowserSpecAllOf) Get() *AzureComputeBrowserSpecAllOf {
	return v.value
}

func (v *NullableAzureComputeBrowserSpecAllOf) Set(val *AzureComputeBrowserSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureComputeBrowserSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureComputeBrowserSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureComputeBrowserSpecAllOf(val *AzureComputeBrowserSpecAllOf) *NullableAzureComputeBrowserSpecAllOf {
	return &NullableAzureComputeBrowserSpecAllOf{value: val, isSet: true}
}

func (v NullableAzureComputeBrowserSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureComputeBrowserSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


