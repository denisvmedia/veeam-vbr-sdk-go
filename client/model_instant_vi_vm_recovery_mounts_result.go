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

// InstantViVMRecoveryMountsResult struct for InstantViVMRecoveryMountsResult
type InstantViVMRecoveryMountsResult struct {
	// Array of VM mounts.
	Data []InstantViVMRecoveryMount `json:"data"`
	Pagination PaginationResult `json:"pagination"`
}

// NewInstantViVMRecoveryMountsResult instantiates a new InstantViVMRecoveryMountsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstantViVMRecoveryMountsResult(data []InstantViVMRecoveryMount, pagination PaginationResult) *InstantViVMRecoveryMountsResult {
	this := InstantViVMRecoveryMountsResult{}
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewInstantViVMRecoveryMountsResultWithDefaults instantiates a new InstantViVMRecoveryMountsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstantViVMRecoveryMountsResultWithDefaults() *InstantViVMRecoveryMountsResult {
	this := InstantViVMRecoveryMountsResult{}
	return &this
}

// GetData returns the Data field value
func (o *InstantViVMRecoveryMountsResult) GetData() []InstantViVMRecoveryMount {
	if o == nil {
		var ret []InstantViVMRecoveryMount
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InstantViVMRecoveryMountsResult) GetDataOk() ([]InstantViVMRecoveryMount, bool) {
	if o == nil {
    return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *InstantViVMRecoveryMountsResult) SetData(v []InstantViVMRecoveryMount) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *InstantViVMRecoveryMountsResult) GetPagination() PaginationResult {
	if o == nil {
		var ret PaginationResult
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *InstantViVMRecoveryMountsResult) GetPaginationOk() (*PaginationResult, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *InstantViVMRecoveryMountsResult) SetPagination(v PaginationResult) {
	o.Pagination = v
}

func (o InstantViVMRecoveryMountsResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableInstantViVMRecoveryMountsResult struct {
	value *InstantViVMRecoveryMountsResult
	isSet bool
}

func (v NullableInstantViVMRecoveryMountsResult) Get() *InstantViVMRecoveryMountsResult {
	return v.value
}

func (v *NullableInstantViVMRecoveryMountsResult) Set(val *InstantViVMRecoveryMountsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableInstantViVMRecoveryMountsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableInstantViVMRecoveryMountsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstantViVMRecoveryMountsResult(val *InstantViVMRecoveryMountsResult) *NullableInstantViVMRecoveryMountsResult {
	return &NullableInstantViVMRecoveryMountsResult{value: val, isSet: true}
}

func (v NullableInstantViVMRecoveryMountsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstantViVMRecoveryMountsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


