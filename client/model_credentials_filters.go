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

// CredentialsFilters struct for CredentialsFilters
type CredentialsFilters struct {
	// Number of credentials records to skip.
	Skip *int32 `json:"skip,omitempty"`
	// Maximum number of credentials records to return.
	Limit *int32 `json:"limit,omitempty"`
	OrderColumn *ECredentialsFiltersOrderColumn `json:"orderColumn,omitempty"`
	// Sorts credentials in the ascending order by the `orderColumn` parameter.
	OrderAsc *bool `json:"orderAsc,omitempty"`
	// Filters credentials by the `nameFilter` pattern. The pattern can match any credentials parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end.
	NameFilter *string `json:"nameFilter,omitempty"`
}

// NewCredentialsFilters instantiates a new CredentialsFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsFilters() *CredentialsFilters {
	this := CredentialsFilters{}
	return &this
}

// NewCredentialsFiltersWithDefaults instantiates a new CredentialsFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsFiltersWithDefaults() *CredentialsFilters {
	this := CredentialsFilters{}
	return &this
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *CredentialsFilters) GetSkip() int32 {
	if o == nil || isNil(o.Skip) {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsFilters) GetSkipOk() (*int32, bool) {
	if o == nil || isNil(o.Skip) {
    return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *CredentialsFilters) HasSkip() bool {
	if o != nil && !isNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *CredentialsFilters) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *CredentialsFilters) GetLimit() int32 {
	if o == nil || isNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsFilters) GetLimitOk() (*int32, bool) {
	if o == nil || isNil(o.Limit) {
    return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *CredentialsFilters) HasLimit() bool {
	if o != nil && !isNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *CredentialsFilters) SetLimit(v int32) {
	o.Limit = &v
}

// GetOrderColumn returns the OrderColumn field value if set, zero value otherwise.
func (o *CredentialsFilters) GetOrderColumn() ECredentialsFiltersOrderColumn {
	if o == nil || isNil(o.OrderColumn) {
		var ret ECredentialsFiltersOrderColumn
		return ret
	}
	return *o.OrderColumn
}

// GetOrderColumnOk returns a tuple with the OrderColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsFilters) GetOrderColumnOk() (*ECredentialsFiltersOrderColumn, bool) {
	if o == nil || isNil(o.OrderColumn) {
    return nil, false
	}
	return o.OrderColumn, true
}

// HasOrderColumn returns a boolean if a field has been set.
func (o *CredentialsFilters) HasOrderColumn() bool {
	if o != nil && !isNil(o.OrderColumn) {
		return true
	}

	return false
}

// SetOrderColumn gets a reference to the given ECredentialsFiltersOrderColumn and assigns it to the OrderColumn field.
func (o *CredentialsFilters) SetOrderColumn(v ECredentialsFiltersOrderColumn) {
	o.OrderColumn = &v
}

// GetOrderAsc returns the OrderAsc field value if set, zero value otherwise.
func (o *CredentialsFilters) GetOrderAsc() bool {
	if o == nil || isNil(o.OrderAsc) {
		var ret bool
		return ret
	}
	return *o.OrderAsc
}

// GetOrderAscOk returns a tuple with the OrderAsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsFilters) GetOrderAscOk() (*bool, bool) {
	if o == nil || isNil(o.OrderAsc) {
    return nil, false
	}
	return o.OrderAsc, true
}

// HasOrderAsc returns a boolean if a field has been set.
func (o *CredentialsFilters) HasOrderAsc() bool {
	if o != nil && !isNil(o.OrderAsc) {
		return true
	}

	return false
}

// SetOrderAsc gets a reference to the given bool and assigns it to the OrderAsc field.
func (o *CredentialsFilters) SetOrderAsc(v bool) {
	o.OrderAsc = &v
}

// GetNameFilter returns the NameFilter field value if set, zero value otherwise.
func (o *CredentialsFilters) GetNameFilter() string {
	if o == nil || isNil(o.NameFilter) {
		var ret string
		return ret
	}
	return *o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsFilters) GetNameFilterOk() (*string, bool) {
	if o == nil || isNil(o.NameFilter) {
    return nil, false
	}
	return o.NameFilter, true
}

// HasNameFilter returns a boolean if a field has been set.
func (o *CredentialsFilters) HasNameFilter() bool {
	if o != nil && !isNil(o.NameFilter) {
		return true
	}

	return false
}

// SetNameFilter gets a reference to the given string and assigns it to the NameFilter field.
func (o *CredentialsFilters) SetNameFilter(v string) {
	o.NameFilter = &v
}

func (o CredentialsFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Skip) {
		toSerialize["skip"] = o.Skip
	}
	if !isNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !isNil(o.OrderColumn) {
		toSerialize["orderColumn"] = o.OrderColumn
	}
	if !isNil(o.OrderAsc) {
		toSerialize["orderAsc"] = o.OrderAsc
	}
	if !isNil(o.NameFilter) {
		toSerialize["nameFilter"] = o.NameFilter
	}
	return json.Marshal(toSerialize)
}

type NullableCredentialsFilters struct {
	value *CredentialsFilters
	isSet bool
}

func (v NullableCredentialsFilters) Get() *CredentialsFilters {
	return v.value
}

func (v *NullableCredentialsFilters) Set(val *CredentialsFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsFilters(val *CredentialsFilters) *NullableCredentialsFilters {
	return &NullableCredentialsFilters{value: val, isSet: true}
}

func (v NullableCredentialsFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


