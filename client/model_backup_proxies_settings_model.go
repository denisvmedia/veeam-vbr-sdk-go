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

// BackupProxiesSettingsModel Backup proxy settings.
type BackupProxiesSettingsModel struct {
	// If *true*, backup proxies are detected and assigned automatically.
	AutoSelection bool `json:"autoSelection"`
	// Array of proxy IDs.
	ProxyIds *[]string `json:"proxyIds,omitempty"`
}

// NewBackupProxiesSettingsModel instantiates a new BackupProxiesSettingsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupProxiesSettingsModel(autoSelection bool) *BackupProxiesSettingsModel {
	this := BackupProxiesSettingsModel{}
	this.AutoSelection = autoSelection
	return &this
}

// NewBackupProxiesSettingsModelWithDefaults instantiates a new BackupProxiesSettingsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupProxiesSettingsModelWithDefaults() *BackupProxiesSettingsModel {
	this := BackupProxiesSettingsModel{}
	var autoSelection bool = true
	this.AutoSelection = autoSelection
	return &this
}

// GetAutoSelection returns the AutoSelection field value
func (o *BackupProxiesSettingsModel) GetAutoSelection() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoSelection
}

// GetAutoSelectionOk returns a tuple with the AutoSelection field value
// and a boolean to check if the value has been set.
func (o *BackupProxiesSettingsModel) GetAutoSelectionOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AutoSelection, true
}

// SetAutoSelection sets field value
func (o *BackupProxiesSettingsModel) SetAutoSelection(v bool) {
	o.AutoSelection = v
}

// GetProxyIds returns the ProxyIds field value if set, zero value otherwise.
func (o *BackupProxiesSettingsModel) GetProxyIds() []string {
	if o == nil || o.ProxyIds == nil {
		var ret []string
		return ret
	}
	return *o.ProxyIds
}

// GetProxyIdsOk returns a tuple with the ProxyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupProxiesSettingsModel) GetProxyIdsOk() (*[]string, bool) {
	if o == nil || o.ProxyIds == nil {
		return nil, false
	}
	return o.ProxyIds, true
}

// HasProxyIds returns a boolean if a field has been set.
func (o *BackupProxiesSettingsModel) HasProxyIds() bool {
	if o != nil && o.ProxyIds != nil {
		return true
	}

	return false
}

// SetProxyIds gets a reference to the given []string and assigns it to the ProxyIds field.
func (o *BackupProxiesSettingsModel) SetProxyIds(v []string) {
	o.ProxyIds = &v
}

func (o BackupProxiesSettingsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["autoSelection"] = o.AutoSelection
	}
	if o.ProxyIds != nil {
		toSerialize["proxyIds"] = o.ProxyIds
	}
	return json.Marshal(toSerialize)
}

type NullableBackupProxiesSettingsModel struct {
	value *BackupProxiesSettingsModel
	isSet bool
}

func (v NullableBackupProxiesSettingsModel) Get() *BackupProxiesSettingsModel {
	return v.value
}

func (v *NullableBackupProxiesSettingsModel) Set(val *BackupProxiesSettingsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupProxiesSettingsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupProxiesSettingsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupProxiesSettingsModel(val *BackupProxiesSettingsModel) *NullableBackupProxiesSettingsModel {
	return &NullableBackupProxiesSettingsModel{value: val, isSet: true}
}

func (v NullableBackupProxiesSettingsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupProxiesSettingsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


