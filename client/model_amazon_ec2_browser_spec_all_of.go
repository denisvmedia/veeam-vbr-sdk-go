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

// AmazonEC2BrowserSpecAllOf struct for AmazonEC2BrowserSpecAllOf
type AmazonEC2BrowserSpecAllOf struct {
	RegionType EAmazonRegionType `json:"regionType"`
	Filters *AmazonEC2BrowserFilters `json:"filters,omitempty"`
}

// NewAmazonEC2BrowserSpecAllOf instantiates a new AmazonEC2BrowserSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonEC2BrowserSpecAllOf(regionType EAmazonRegionType) *AmazonEC2BrowserSpecAllOf {
	this := AmazonEC2BrowserSpecAllOf{}
	this.RegionType = regionType
	return &this
}

// NewAmazonEC2BrowserSpecAllOfWithDefaults instantiates a new AmazonEC2BrowserSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonEC2BrowserSpecAllOfWithDefaults() *AmazonEC2BrowserSpecAllOf {
	this := AmazonEC2BrowserSpecAllOf{}
	return &this
}

// GetRegionType returns the RegionType field value
func (o *AmazonEC2BrowserSpecAllOf) GetRegionType() EAmazonRegionType {
	if o == nil {
		var ret EAmazonRegionType
		return ret
	}

	return o.RegionType
}

// GetRegionTypeOk returns a tuple with the RegionType field value
// and a boolean to check if the value has been set.
func (o *AmazonEC2BrowserSpecAllOf) GetRegionTypeOk() (*EAmazonRegionType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RegionType, true
}

// SetRegionType sets field value
func (o *AmazonEC2BrowserSpecAllOf) SetRegionType(v EAmazonRegionType) {
	o.RegionType = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *AmazonEC2BrowserSpecAllOf) GetFilters() AmazonEC2BrowserFilters {
	if o == nil || isNil(o.Filters) {
		var ret AmazonEC2BrowserFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonEC2BrowserSpecAllOf) GetFiltersOk() (*AmazonEC2BrowserFilters, bool) {
	if o == nil || isNil(o.Filters) {
    return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *AmazonEC2BrowserSpecAllOf) HasFilters() bool {
	if o != nil && !isNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given AmazonEC2BrowserFilters and assigns it to the Filters field.
func (o *AmazonEC2BrowserSpecAllOf) SetFilters(v AmazonEC2BrowserFilters) {
	o.Filters = &v
}

func (o AmazonEC2BrowserSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["regionType"] = o.RegionType
	}
	if !isNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonEC2BrowserSpecAllOf struct {
	value *AmazonEC2BrowserSpecAllOf
	isSet bool
}

func (v NullableAmazonEC2BrowserSpecAllOf) Get() *AmazonEC2BrowserSpecAllOf {
	return v.value
}

func (v *NullableAmazonEC2BrowserSpecAllOf) Set(val *AmazonEC2BrowserSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonEC2BrowserSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonEC2BrowserSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonEC2BrowserSpecAllOf(val *AmazonEC2BrowserSpecAllOf) *NullableAmazonEC2BrowserSpecAllOf {
	return &NullableAmazonEC2BrowserSpecAllOf{value: val, isSet: true}
}

func (v NullableAmazonEC2BrowserSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonEC2BrowserSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


