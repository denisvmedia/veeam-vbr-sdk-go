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

// ArchiveTierModel Archive tier.
type ArchiveTierModel struct {
	// If *true*, the archive tier is enabled.
	IsEnabled bool `json:"isEnabled"`
	// ID of an object storage repository added as an archive extent.
	ExtentId *string `json:"extentId,omitempty"`
	// Number of days after which backup chains on the capacity extent are moved to the archive extent. Specify *0* to offload inactive backup chains on the same day they are created.
	ArchivePeriodDays *int32 `json:"archivePeriodDays,omitempty"`
	AdvancedSettings *ArchiveTierAdvancedSettingsModel `json:"advancedSettings,omitempty"`
}

// NewArchiveTierModel instantiates a new ArchiveTierModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchiveTierModel(isEnabled bool) *ArchiveTierModel {
	this := ArchiveTierModel{}
	this.IsEnabled = isEnabled
	return &this
}

// NewArchiveTierModelWithDefaults instantiates a new ArchiveTierModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchiveTierModelWithDefaults() *ArchiveTierModel {
	this := ArchiveTierModel{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *ArchiveTierModel) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *ArchiveTierModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *ArchiveTierModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetExtentId returns the ExtentId field value if set, zero value otherwise.
func (o *ArchiveTierModel) GetExtentId() string {
	if o == nil || isNil(o.ExtentId) {
		var ret string
		return ret
	}
	return *o.ExtentId
}

// GetExtentIdOk returns a tuple with the ExtentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveTierModel) GetExtentIdOk() (*string, bool) {
	if o == nil || isNil(o.ExtentId) {
    return nil, false
	}
	return o.ExtentId, true
}

// HasExtentId returns a boolean if a field has been set.
func (o *ArchiveTierModel) HasExtentId() bool {
	if o != nil && !isNil(o.ExtentId) {
		return true
	}

	return false
}

// SetExtentId gets a reference to the given string and assigns it to the ExtentId field.
func (o *ArchiveTierModel) SetExtentId(v string) {
	o.ExtentId = &v
}

// GetArchivePeriodDays returns the ArchivePeriodDays field value if set, zero value otherwise.
func (o *ArchiveTierModel) GetArchivePeriodDays() int32 {
	if o == nil || isNil(o.ArchivePeriodDays) {
		var ret int32
		return ret
	}
	return *o.ArchivePeriodDays
}

// GetArchivePeriodDaysOk returns a tuple with the ArchivePeriodDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveTierModel) GetArchivePeriodDaysOk() (*int32, bool) {
	if o == nil || isNil(o.ArchivePeriodDays) {
    return nil, false
	}
	return o.ArchivePeriodDays, true
}

// HasArchivePeriodDays returns a boolean if a field has been set.
func (o *ArchiveTierModel) HasArchivePeriodDays() bool {
	if o != nil && !isNil(o.ArchivePeriodDays) {
		return true
	}

	return false
}

// SetArchivePeriodDays gets a reference to the given int32 and assigns it to the ArchivePeriodDays field.
func (o *ArchiveTierModel) SetArchivePeriodDays(v int32) {
	o.ArchivePeriodDays = &v
}

// GetAdvancedSettings returns the AdvancedSettings field value if set, zero value otherwise.
func (o *ArchiveTierModel) GetAdvancedSettings() ArchiveTierAdvancedSettingsModel {
	if o == nil || isNil(o.AdvancedSettings) {
		var ret ArchiveTierAdvancedSettingsModel
		return ret
	}
	return *o.AdvancedSettings
}

// GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveTierModel) GetAdvancedSettingsOk() (*ArchiveTierAdvancedSettingsModel, bool) {
	if o == nil || isNil(o.AdvancedSettings) {
    return nil, false
	}
	return o.AdvancedSettings, true
}

// HasAdvancedSettings returns a boolean if a field has been set.
func (o *ArchiveTierModel) HasAdvancedSettings() bool {
	if o != nil && !isNil(o.AdvancedSettings) {
		return true
	}

	return false
}

// SetAdvancedSettings gets a reference to the given ArchiveTierAdvancedSettingsModel and assigns it to the AdvancedSettings field.
func (o *ArchiveTierModel) SetAdvancedSettings(v ArchiveTierAdvancedSettingsModel) {
	o.AdvancedSettings = &v
}

func (o ArchiveTierModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !isNil(o.ExtentId) {
		toSerialize["extentId"] = o.ExtentId
	}
	if !isNil(o.ArchivePeriodDays) {
		toSerialize["archivePeriodDays"] = o.ArchivePeriodDays
	}
	if !isNil(o.AdvancedSettings) {
		toSerialize["advancedSettings"] = o.AdvancedSettings
	}
	return json.Marshal(toSerialize)
}

type NullableArchiveTierModel struct {
	value *ArchiveTierModel
	isSet bool
}

func (v NullableArchiveTierModel) Get() *ArchiveTierModel {
	return v.value
}

func (v *NullableArchiveTierModel) Set(val *ArchiveTierModel) {
	v.value = val
	v.isSet = true
}

func (v NullableArchiveTierModel) IsSet() bool {
	return v.isSet
}

func (v *NullableArchiveTierModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchiveTierModel(val *ArchiveTierModel) *NullableArchiveTierModel {
	return &NullableArchiveTierModel{value: val, isSet: true}
}

func (v NullableArchiveTierModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchiveTierModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


