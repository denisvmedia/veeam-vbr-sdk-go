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

// AzureComputeBrowserSpec struct for AzureComputeBrowserSpec
type AzureComputeBrowserSpec struct {
	CloudBrowserSpec
	Filters *AzureComputeBrowserFilters `json:"filters,omitempty"`
}

// NewAzureComputeBrowserSpec instantiates a new AzureComputeBrowserSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureComputeBrowserSpec(credentialsId string, serviceType ECloudServiceType) *AzureComputeBrowserSpec {
	this := AzureComputeBrowserSpec{}
	this.CredentialsId = credentialsId
	this.ServiceType = serviceType
	return &this
}

// NewAzureComputeBrowserSpecWithDefaults instantiates a new AzureComputeBrowserSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureComputeBrowserSpecWithDefaults() *AzureComputeBrowserSpec {
	this := AzureComputeBrowserSpec{}
	return &this
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *AzureComputeBrowserSpec) GetFilters() AzureComputeBrowserFilters {
	if o == nil || isNil(o.Filters) {
		var ret AzureComputeBrowserFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureComputeBrowserSpec) GetFiltersOk() (*AzureComputeBrowserFilters, bool) {
	if o == nil || isNil(o.Filters) {
    return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *AzureComputeBrowserSpec) HasFilters() bool {
	if o != nil && !isNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given AzureComputeBrowserFilters and assigns it to the Filters field.
func (o *AzureComputeBrowserSpec) SetFilters(v AzureComputeBrowserFilters) {
	o.Filters = &v
}

func (o AzureComputeBrowserSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBrowserSpec, errCloudBrowserSpec := json.Marshal(o.CloudBrowserSpec)
	if errCloudBrowserSpec != nil {
		return []byte{}, errCloudBrowserSpec
	}
	errCloudBrowserSpec = json.Unmarshal([]byte(serializedCloudBrowserSpec), &toSerialize)
	if errCloudBrowserSpec != nil {
		return []byte{}, errCloudBrowserSpec
	}
	if !isNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

type NullableAzureComputeBrowserSpec struct {
	value *AzureComputeBrowserSpec
	isSet bool
}

func (v NullableAzureComputeBrowserSpec) Get() *AzureComputeBrowserSpec {
	return v.value
}

func (v *NullableAzureComputeBrowserSpec) Set(val *AzureComputeBrowserSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureComputeBrowserSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureComputeBrowserSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureComputeBrowserSpec(val *AzureComputeBrowserSpec) *NullableAzureComputeBrowserSpec {
	return &NullableAzureComputeBrowserSpec{value: val, isSet: true}
}

func (v NullableAzureComputeBrowserSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureComputeBrowserSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


