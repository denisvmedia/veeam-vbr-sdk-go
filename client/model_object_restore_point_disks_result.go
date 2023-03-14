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

// ObjectRestorePointDisksResult struct for ObjectRestorePointDisksResult
type ObjectRestorePointDisksResult struct {
	// Array of disks.
	Data []ObjectRestorePointDiskModel `json:"data"`
	Pagination PaginationResult `json:"pagination"`
}

// NewObjectRestorePointDisksResult instantiates a new ObjectRestorePointDisksResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectRestorePointDisksResult(data []ObjectRestorePointDiskModel, pagination PaginationResult) *ObjectRestorePointDisksResult {
	this := ObjectRestorePointDisksResult{}
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewObjectRestorePointDisksResultWithDefaults instantiates a new ObjectRestorePointDisksResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectRestorePointDisksResultWithDefaults() *ObjectRestorePointDisksResult {
	this := ObjectRestorePointDisksResult{}
	return &this
}

// GetData returns the Data field value
func (o *ObjectRestorePointDisksResult) GetData() []ObjectRestorePointDiskModel {
	if o == nil {
		var ret []ObjectRestorePointDiskModel
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointDisksResult) GetDataOk() ([]ObjectRestorePointDiskModel, bool) {
	if o == nil {
    return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ObjectRestorePointDisksResult) SetData(v []ObjectRestorePointDiskModel) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *ObjectRestorePointDisksResult) GetPagination() PaginationResult {
	if o == nil {
		var ret PaginationResult
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointDisksResult) GetPaginationOk() (*PaginationResult, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ObjectRestorePointDisksResult) SetPagination(v PaginationResult) {
	o.Pagination = v
}

func (o ObjectRestorePointDisksResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableObjectRestorePointDisksResult struct {
	value *ObjectRestorePointDisksResult
	isSet bool
}

func (v NullableObjectRestorePointDisksResult) Get() *ObjectRestorePointDisksResult {
	return v.value
}

func (v *NullableObjectRestorePointDisksResult) Set(val *ObjectRestorePointDisksResult) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectRestorePointDisksResult) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectRestorePointDisksResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectRestorePointDisksResult(val *ObjectRestorePointDisksResult) *NullableObjectRestorePointDisksResult {
	return &NullableObjectRestorePointDisksResult{value: val, isSet: true}
}

func (v NullableObjectRestorePointDisksResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectRestorePointDisksResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


