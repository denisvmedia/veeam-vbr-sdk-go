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
	"time"
)

// ConfigBackupLastSuccessfulModel Last successful backup.
type ConfigBackupLastSuccessfulModel struct {
	// Date and time when the last successful backup was created.
	LastSuccessfulTime *time.Time `json:"lastSuccessfulTime,omitempty"`
	// ID of the job session.
	SessionId *string `json:"sessionId,omitempty"`
}

// NewConfigBackupLastSuccessfulModel instantiates a new ConfigBackupLastSuccessfulModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigBackupLastSuccessfulModel() *ConfigBackupLastSuccessfulModel {
	this := ConfigBackupLastSuccessfulModel{}
	return &this
}

// NewConfigBackupLastSuccessfulModelWithDefaults instantiates a new ConfigBackupLastSuccessfulModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigBackupLastSuccessfulModelWithDefaults() *ConfigBackupLastSuccessfulModel {
	this := ConfigBackupLastSuccessfulModel{}
	return &this
}

// GetLastSuccessfulTime returns the LastSuccessfulTime field value if set, zero value otherwise.
func (o *ConfigBackupLastSuccessfulModel) GetLastSuccessfulTime() time.Time {
	if o == nil || o.LastSuccessfulTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSuccessfulTime
}

// GetLastSuccessfulTimeOk returns a tuple with the LastSuccessfulTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigBackupLastSuccessfulModel) GetLastSuccessfulTimeOk() (*time.Time, bool) {
	if o == nil || o.LastSuccessfulTime == nil {
		return nil, false
	}
	return o.LastSuccessfulTime, true
}

// HasLastSuccessfulTime returns a boolean if a field has been set.
func (o *ConfigBackupLastSuccessfulModel) HasLastSuccessfulTime() bool {
	if o != nil && o.LastSuccessfulTime != nil {
		return true
	}

	return false
}

// SetLastSuccessfulTime gets a reference to the given time.Time and assigns it to the LastSuccessfulTime field.
func (o *ConfigBackupLastSuccessfulModel) SetLastSuccessfulTime(v time.Time) {
	o.LastSuccessfulTime = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *ConfigBackupLastSuccessfulModel) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigBackupLastSuccessfulModel) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *ConfigBackupLastSuccessfulModel) HasSessionId() bool {
	if o != nil && o.SessionId != nil {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *ConfigBackupLastSuccessfulModel) SetSessionId(v string) {
	o.SessionId = &v
}

func (o ConfigBackupLastSuccessfulModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastSuccessfulTime != nil {
		toSerialize["lastSuccessfulTime"] = o.LastSuccessfulTime
	}
	if o.SessionId != nil {
		toSerialize["sessionId"] = o.SessionId
	}
	return json.Marshal(toSerialize)
}

type NullableConfigBackupLastSuccessfulModel struct {
	value *ConfigBackupLastSuccessfulModel
	isSet bool
}

func (v NullableConfigBackupLastSuccessfulModel) Get() *ConfigBackupLastSuccessfulModel {
	return v.value
}

func (v *NullableConfigBackupLastSuccessfulModel) Set(val *ConfigBackupLastSuccessfulModel) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigBackupLastSuccessfulModel) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigBackupLastSuccessfulModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigBackupLastSuccessfulModel(val *ConfigBackupLastSuccessfulModel) *NullableConfigBackupLastSuccessfulModel {
	return &NullableConfigBackupLastSuccessfulModel{value: val, isSet: true}
}

func (v NullableConfigBackupLastSuccessfulModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigBackupLastSuccessfulModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


