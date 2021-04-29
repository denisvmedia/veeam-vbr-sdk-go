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

// ObjectRestorePointsResult struct for ObjectRestorePointsResult
type ObjectRestorePointsResult struct {
	// Array of restore points.
	Data []ObjectRestorePointModel `json:"data"`
	Pagination PaginationResult `json:"pagination"`
}

// NewObjectRestorePointsResult instantiates a new ObjectRestorePointsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectRestorePointsResult(data []ObjectRestorePointModel, pagination PaginationResult) *ObjectRestorePointsResult {
	this := ObjectRestorePointsResult{}
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewObjectRestorePointsResultWithDefaults instantiates a new ObjectRestorePointsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectRestorePointsResultWithDefaults() *ObjectRestorePointsResult {
	this := ObjectRestorePointsResult{}
	return &this
}

// GetData returns the Data field value
func (o *ObjectRestorePointsResult) GetData() []ObjectRestorePointModel {
	if o == nil {
		var ret []ObjectRestorePointModel
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointsResult) GetDataOk() (*[]ObjectRestorePointModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ObjectRestorePointsResult) SetData(v []ObjectRestorePointModel) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *ObjectRestorePointsResult) GetPagination() PaginationResult {
	if o == nil {
		var ret PaginationResult
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ObjectRestorePointsResult) GetPaginationOk() (*PaginationResult, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ObjectRestorePointsResult) SetPagination(v PaginationResult) {
	o.Pagination = v
}

func (o ObjectRestorePointsResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableObjectRestorePointsResult struct {
	value *ObjectRestorePointsResult
	isSet bool
}

func (v NullableObjectRestorePointsResult) Get() *ObjectRestorePointsResult {
	return v.value
}

func (v *NullableObjectRestorePointsResult) Set(val *ObjectRestorePointsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectRestorePointsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectRestorePointsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectRestorePointsResult(val *ObjectRestorePointsResult) *NullableObjectRestorePointsResult {
	return &NullableObjectRestorePointsResult{value: val, isSet: true}
}

func (v NullableObjectRestorePointsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectRestorePointsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


