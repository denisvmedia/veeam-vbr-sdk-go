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

// SessionsResult struct for SessionsResult
type SessionsResult struct {
	// Array of sessions.
	Data []SessionModel `json:"data"`
	Pagination PaginationResult `json:"pagination"`
}

// NewSessionsResult instantiates a new SessionsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionsResult(data []SessionModel, pagination PaginationResult) *SessionsResult {
	this := SessionsResult{}
	this.Data = data
	this.Pagination = pagination
	return &this
}

// NewSessionsResultWithDefaults instantiates a new SessionsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionsResultWithDefaults() *SessionsResult {
	this := SessionsResult{}
	return &this
}

// GetData returns the Data field value
func (o *SessionsResult) GetData() []SessionModel {
	if o == nil {
		var ret []SessionModel
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SessionsResult) GetDataOk() (*[]SessionModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SessionsResult) SetData(v []SessionModel) {
	o.Data = v
}

// GetPagination returns the Pagination field value
func (o *SessionsResult) GetPagination() PaginationResult {
	if o == nil {
		var ret PaginationResult
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *SessionsResult) GetPaginationOk() (*PaginationResult, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *SessionsResult) SetPagination(v PaginationResult) {
	o.Pagination = v
}

func (o SessionsResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableSessionsResult struct {
	value *SessionsResult
	isSet bool
}

func (v NullableSessionsResult) Get() *SessionsResult {
	return v.value
}

func (v *NullableSessionsResult) Set(val *SessionsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionsResult(val *SessionsResult) *NullableSessionsResult {
	return &NullableSessionsResult{value: val, isSet: true}
}

func (v NullableSessionsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


