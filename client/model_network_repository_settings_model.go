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

// NetworkRepositorySettingsModel Repository settings.
type NetworkRepositorySettingsModel struct {
	// If *true*, the maximum number of concurrent tasks is limited.
	EnableTaskLimit *bool `json:"enableTaskLimit,omitempty"`
	// Maximum number of concurrent tasks.
	MaxTaskCount *int32 `json:"maxTaskCount,omitempty"`
	// If *true*, reading and writing speed is limited.
	EnableReadWriteLimit *bool `json:"enableReadWriteLimit,omitempty"`
	// Maximum rate that restricts the total speed of reading and writing data to the backup repository disk.
	ReadWriteRate *int32 `json:"readWriteRate,omitempty"`
	AdvancedSettings *RepositoryAdvancedSettingsModel `json:"advancedSettings,omitempty"`
}

// NewNetworkRepositorySettingsModel instantiates a new NetworkRepositorySettingsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkRepositorySettingsModel() *NetworkRepositorySettingsModel {
	this := NetworkRepositorySettingsModel{}
	return &this
}

// NewNetworkRepositorySettingsModelWithDefaults instantiates a new NetworkRepositorySettingsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkRepositorySettingsModelWithDefaults() *NetworkRepositorySettingsModel {
	this := NetworkRepositorySettingsModel{}
	return &this
}

// GetEnableTaskLimit returns the EnableTaskLimit field value if set, zero value otherwise.
func (o *NetworkRepositorySettingsModel) GetEnableTaskLimit() bool {
	if o == nil || isNil(o.EnableTaskLimit) {
		var ret bool
		return ret
	}
	return *o.EnableTaskLimit
}

// GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRepositorySettingsModel) GetEnableTaskLimitOk() (*bool, bool) {
	if o == nil || isNil(o.EnableTaskLimit) {
    return nil, false
	}
	return o.EnableTaskLimit, true
}

// HasEnableTaskLimit returns a boolean if a field has been set.
func (o *NetworkRepositorySettingsModel) HasEnableTaskLimit() bool {
	if o != nil && !isNil(o.EnableTaskLimit) {
		return true
	}

	return false
}

// SetEnableTaskLimit gets a reference to the given bool and assigns it to the EnableTaskLimit field.
func (o *NetworkRepositorySettingsModel) SetEnableTaskLimit(v bool) {
	o.EnableTaskLimit = &v
}

// GetMaxTaskCount returns the MaxTaskCount field value if set, zero value otherwise.
func (o *NetworkRepositorySettingsModel) GetMaxTaskCount() int32 {
	if o == nil || isNil(o.MaxTaskCount) {
		var ret int32
		return ret
	}
	return *o.MaxTaskCount
}

// GetMaxTaskCountOk returns a tuple with the MaxTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRepositorySettingsModel) GetMaxTaskCountOk() (*int32, bool) {
	if o == nil || isNil(o.MaxTaskCount) {
    return nil, false
	}
	return o.MaxTaskCount, true
}

// HasMaxTaskCount returns a boolean if a field has been set.
func (o *NetworkRepositorySettingsModel) HasMaxTaskCount() bool {
	if o != nil && !isNil(o.MaxTaskCount) {
		return true
	}

	return false
}

// SetMaxTaskCount gets a reference to the given int32 and assigns it to the MaxTaskCount field.
func (o *NetworkRepositorySettingsModel) SetMaxTaskCount(v int32) {
	o.MaxTaskCount = &v
}

// GetEnableReadWriteLimit returns the EnableReadWriteLimit field value if set, zero value otherwise.
func (o *NetworkRepositorySettingsModel) GetEnableReadWriteLimit() bool {
	if o == nil || isNil(o.EnableReadWriteLimit) {
		var ret bool
		return ret
	}
	return *o.EnableReadWriteLimit
}

// GetEnableReadWriteLimitOk returns a tuple with the EnableReadWriteLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRepositorySettingsModel) GetEnableReadWriteLimitOk() (*bool, bool) {
	if o == nil || isNil(o.EnableReadWriteLimit) {
    return nil, false
	}
	return o.EnableReadWriteLimit, true
}

// HasEnableReadWriteLimit returns a boolean if a field has been set.
func (o *NetworkRepositorySettingsModel) HasEnableReadWriteLimit() bool {
	if o != nil && !isNil(o.EnableReadWriteLimit) {
		return true
	}

	return false
}

// SetEnableReadWriteLimit gets a reference to the given bool and assigns it to the EnableReadWriteLimit field.
func (o *NetworkRepositorySettingsModel) SetEnableReadWriteLimit(v bool) {
	o.EnableReadWriteLimit = &v
}

// GetReadWriteRate returns the ReadWriteRate field value if set, zero value otherwise.
func (o *NetworkRepositorySettingsModel) GetReadWriteRate() int32 {
	if o == nil || isNil(o.ReadWriteRate) {
		var ret int32
		return ret
	}
	return *o.ReadWriteRate
}

// GetReadWriteRateOk returns a tuple with the ReadWriteRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRepositorySettingsModel) GetReadWriteRateOk() (*int32, bool) {
	if o == nil || isNil(o.ReadWriteRate) {
    return nil, false
	}
	return o.ReadWriteRate, true
}

// HasReadWriteRate returns a boolean if a field has been set.
func (o *NetworkRepositorySettingsModel) HasReadWriteRate() bool {
	if o != nil && !isNil(o.ReadWriteRate) {
		return true
	}

	return false
}

// SetReadWriteRate gets a reference to the given int32 and assigns it to the ReadWriteRate field.
func (o *NetworkRepositorySettingsModel) SetReadWriteRate(v int32) {
	o.ReadWriteRate = &v
}

// GetAdvancedSettings returns the AdvancedSettings field value if set, zero value otherwise.
func (o *NetworkRepositorySettingsModel) GetAdvancedSettings() RepositoryAdvancedSettingsModel {
	if o == nil || isNil(o.AdvancedSettings) {
		var ret RepositoryAdvancedSettingsModel
		return ret
	}
	return *o.AdvancedSettings
}

// GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRepositorySettingsModel) GetAdvancedSettingsOk() (*RepositoryAdvancedSettingsModel, bool) {
	if o == nil || isNil(o.AdvancedSettings) {
    return nil, false
	}
	return o.AdvancedSettings, true
}

// HasAdvancedSettings returns a boolean if a field has been set.
func (o *NetworkRepositorySettingsModel) HasAdvancedSettings() bool {
	if o != nil && !isNil(o.AdvancedSettings) {
		return true
	}

	return false
}

// SetAdvancedSettings gets a reference to the given RepositoryAdvancedSettingsModel and assigns it to the AdvancedSettings field.
func (o *NetworkRepositorySettingsModel) SetAdvancedSettings(v RepositoryAdvancedSettingsModel) {
	o.AdvancedSettings = &v
}

func (o NetworkRepositorySettingsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EnableTaskLimit) {
		toSerialize["enableTaskLimit"] = o.EnableTaskLimit
	}
	if !isNil(o.MaxTaskCount) {
		toSerialize["maxTaskCount"] = o.MaxTaskCount
	}
	if !isNil(o.EnableReadWriteLimit) {
		toSerialize["enableReadWriteLimit"] = o.EnableReadWriteLimit
	}
	if !isNil(o.ReadWriteRate) {
		toSerialize["readWriteRate"] = o.ReadWriteRate
	}
	if !isNil(o.AdvancedSettings) {
		toSerialize["advancedSettings"] = o.AdvancedSettings
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkRepositorySettingsModel struct {
	value *NetworkRepositorySettingsModel
	isSet bool
}

func (v NullableNetworkRepositorySettingsModel) Get() *NetworkRepositorySettingsModel {
	return v.value
}

func (v *NullableNetworkRepositorySettingsModel) Set(val *NetworkRepositorySettingsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkRepositorySettingsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkRepositorySettingsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkRepositorySettingsModel(val *NetworkRepositorySettingsModel) *NullableNetworkRepositorySettingsModel {
	return &NullableNetworkRepositorySettingsModel{value: val, isSet: true}
}

func (v NullableNetworkRepositorySettingsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkRepositorySettingsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


