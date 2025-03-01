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

// BackupApplicationAwareProcessingImportModel Application-aware processing settings.
type BackupApplicationAwareProcessingImportModel struct {
	// If *true*, application-aware processing is enabled.
	IsEnabled bool `json:"isEnabled"`
	// Array of VMware vSphere objects and their application settings.
	AppSettings []BackupApplicationSettingsImportModel `json:"appSettings,omitempty"`
}

// NewBackupApplicationAwareProcessingImportModel instantiates a new BackupApplicationAwareProcessingImportModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupApplicationAwareProcessingImportModel(isEnabled bool) *BackupApplicationAwareProcessingImportModel {
	this := BackupApplicationAwareProcessingImportModel{}
	this.IsEnabled = isEnabled
	return &this
}

// NewBackupApplicationAwareProcessingImportModelWithDefaults instantiates a new BackupApplicationAwareProcessingImportModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupApplicationAwareProcessingImportModelWithDefaults() *BackupApplicationAwareProcessingImportModel {
	this := BackupApplicationAwareProcessingImportModel{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *BackupApplicationAwareProcessingImportModel) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *BackupApplicationAwareProcessingImportModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *BackupApplicationAwareProcessingImportModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetAppSettings returns the AppSettings field value if set, zero value otherwise.
func (o *BackupApplicationAwareProcessingImportModel) GetAppSettings() []BackupApplicationSettingsImportModel {
	if o == nil || isNil(o.AppSettings) {
		var ret []BackupApplicationSettingsImportModel
		return ret
	}
	return o.AppSettings
}

// GetAppSettingsOk returns a tuple with the AppSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupApplicationAwareProcessingImportModel) GetAppSettingsOk() ([]BackupApplicationSettingsImportModel, bool) {
	if o == nil || isNil(o.AppSettings) {
    return nil, false
	}
	return o.AppSettings, true
}

// HasAppSettings returns a boolean if a field has been set.
func (o *BackupApplicationAwareProcessingImportModel) HasAppSettings() bool {
	if o != nil && !isNil(o.AppSettings) {
		return true
	}

	return false
}

// SetAppSettings gets a reference to the given []BackupApplicationSettingsImportModel and assigns it to the AppSettings field.
func (o *BackupApplicationAwareProcessingImportModel) SetAppSettings(v []BackupApplicationSettingsImportModel) {
	o.AppSettings = v
}

func (o BackupApplicationAwareProcessingImportModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !isNil(o.AppSettings) {
		toSerialize["appSettings"] = o.AppSettings
	}
	return json.Marshal(toSerialize)
}

type NullableBackupApplicationAwareProcessingImportModel struct {
	value *BackupApplicationAwareProcessingImportModel
	isSet bool
}

func (v NullableBackupApplicationAwareProcessingImportModel) Get() *BackupApplicationAwareProcessingImportModel {
	return v.value
}

func (v *NullableBackupApplicationAwareProcessingImportModel) Set(val *BackupApplicationAwareProcessingImportModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupApplicationAwareProcessingImportModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupApplicationAwareProcessingImportModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupApplicationAwareProcessingImportModel(val *BackupApplicationAwareProcessingImportModel) *NullableBackupApplicationAwareProcessingImportModel {
	return &NullableBackupApplicationAwareProcessingImportModel{value: val, isSet: true}
}

func (v NullableBackupApplicationAwareProcessingImportModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupApplicationAwareProcessingImportModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


