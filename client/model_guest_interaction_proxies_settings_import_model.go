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

// GuestInteractionProxiesSettingsImportModel Guest interaction proxy used to deploy the runtime process on the VM guest OS.
type GuestInteractionProxiesSettingsImportModel struct {
	// If *true*, Veeam Backup & Replication automatically selects the guest interaction proxy.
	AutomaticSelection bool `json:"automaticSelection"`
	// Array of proxies specified explicitly. The array must contain Microsoft Windows servers added to the backup infrastructure only.
	Proxies []string `json:"proxies,omitempty"`
}

// NewGuestInteractionProxiesSettingsImportModel instantiates a new GuestInteractionProxiesSettingsImportModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuestInteractionProxiesSettingsImportModel(automaticSelection bool) *GuestInteractionProxiesSettingsImportModel {
	this := GuestInteractionProxiesSettingsImportModel{}
	this.AutomaticSelection = automaticSelection
	return &this
}

// NewGuestInteractionProxiesSettingsImportModelWithDefaults instantiates a new GuestInteractionProxiesSettingsImportModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuestInteractionProxiesSettingsImportModelWithDefaults() *GuestInteractionProxiesSettingsImportModel {
	this := GuestInteractionProxiesSettingsImportModel{}
	var automaticSelection bool = true
	this.AutomaticSelection = automaticSelection
	return &this
}

// GetAutomaticSelection returns the AutomaticSelection field value
func (o *GuestInteractionProxiesSettingsImportModel) GetAutomaticSelection() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutomaticSelection
}

// GetAutomaticSelectionOk returns a tuple with the AutomaticSelection field value
// and a boolean to check if the value has been set.
func (o *GuestInteractionProxiesSettingsImportModel) GetAutomaticSelectionOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AutomaticSelection, true
}

// SetAutomaticSelection sets field value
func (o *GuestInteractionProxiesSettingsImportModel) SetAutomaticSelection(v bool) {
	o.AutomaticSelection = v
}

// GetProxies returns the Proxies field value if set, zero value otherwise.
func (o *GuestInteractionProxiesSettingsImportModel) GetProxies() []string {
	if o == nil || isNil(o.Proxies) {
		var ret []string
		return ret
	}
	return o.Proxies
}

// GetProxiesOk returns a tuple with the Proxies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuestInteractionProxiesSettingsImportModel) GetProxiesOk() ([]string, bool) {
	if o == nil || isNil(o.Proxies) {
    return nil, false
	}
	return o.Proxies, true
}

// HasProxies returns a boolean if a field has been set.
func (o *GuestInteractionProxiesSettingsImportModel) HasProxies() bool {
	if o != nil && !isNil(o.Proxies) {
		return true
	}

	return false
}

// SetProxies gets a reference to the given []string and assigns it to the Proxies field.
func (o *GuestInteractionProxiesSettingsImportModel) SetProxies(v []string) {
	o.Proxies = v
}

func (o GuestInteractionProxiesSettingsImportModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["automaticSelection"] = o.AutomaticSelection
	}
	if !isNil(o.Proxies) {
		toSerialize["proxies"] = o.Proxies
	}
	return json.Marshal(toSerialize)
}

type NullableGuestInteractionProxiesSettingsImportModel struct {
	value *GuestInteractionProxiesSettingsImportModel
	isSet bool
}

func (v NullableGuestInteractionProxiesSettingsImportModel) Get() *GuestInteractionProxiesSettingsImportModel {
	return v.value
}

func (v *NullableGuestInteractionProxiesSettingsImportModel) Set(val *GuestInteractionProxiesSettingsImportModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGuestInteractionProxiesSettingsImportModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGuestInteractionProxiesSettingsImportModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuestInteractionProxiesSettingsImportModel(val *GuestInteractionProxiesSettingsImportModel) *NullableGuestInteractionProxiesSettingsImportModel {
	return &NullableGuestInteractionProxiesSettingsImportModel{value: val, isSet: true}
}

func (v NullableGuestInteractionProxiesSettingsImportModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuestInteractionProxiesSettingsImportModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


