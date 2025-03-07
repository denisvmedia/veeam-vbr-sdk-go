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

// WindowsHostImportSpec struct for WindowsHostImportSpec
type WindowsHostImportSpec struct {
	// Full DNS name or IP address of the server.
	Name string `json:"name"`
	// Description of the server.
	Description string `json:"description"`
	Type EManagedServerType `json:"type"`
	Credentials *CredentialsImportModel `json:"credentials,omitempty"`
	NetworkSettings *WindowsHostPortsModel `json:"networkSettings,omitempty"`
}

// NewWindowsHostImportSpec instantiates a new WindowsHostImportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWindowsHostImportSpec(name string, description string, type_ EManagedServerType) *WindowsHostImportSpec {
	this := WindowsHostImportSpec{}
	this.Name = name
	this.Description = description
	this.Type = type_
	return &this
}

// NewWindowsHostImportSpecWithDefaults instantiates a new WindowsHostImportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWindowsHostImportSpecWithDefaults() *WindowsHostImportSpec {
	this := WindowsHostImportSpec{}
	return &this
}

// GetName returns the Name field value
func (o *WindowsHostImportSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WindowsHostImportSpec) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WindowsHostImportSpec) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *WindowsHostImportSpec) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *WindowsHostImportSpec) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *WindowsHostImportSpec) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value
func (o *WindowsHostImportSpec) GetType() EManagedServerType {
	if o == nil {
		var ret EManagedServerType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WindowsHostImportSpec) GetTypeOk() (*EManagedServerType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WindowsHostImportSpec) SetType(v EManagedServerType) {
	o.Type = v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *WindowsHostImportSpec) GetCredentials() CredentialsImportModel {
	if o == nil || isNil(o.Credentials) {
		var ret CredentialsImportModel
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsHostImportSpec) GetCredentialsOk() (*CredentialsImportModel, bool) {
	if o == nil || isNil(o.Credentials) {
    return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *WindowsHostImportSpec) HasCredentials() bool {
	if o != nil && !isNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given CredentialsImportModel and assigns it to the Credentials field.
func (o *WindowsHostImportSpec) SetCredentials(v CredentialsImportModel) {
	o.Credentials = &v
}

// GetNetworkSettings returns the NetworkSettings field value if set, zero value otherwise.
func (o *WindowsHostImportSpec) GetNetworkSettings() WindowsHostPortsModel {
	if o == nil || isNil(o.NetworkSettings) {
		var ret WindowsHostPortsModel
		return ret
	}
	return *o.NetworkSettings
}

// GetNetworkSettingsOk returns a tuple with the NetworkSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsHostImportSpec) GetNetworkSettingsOk() (*WindowsHostPortsModel, bool) {
	if o == nil || isNil(o.NetworkSettings) {
    return nil, false
	}
	return o.NetworkSettings, true
}

// HasNetworkSettings returns a boolean if a field has been set.
func (o *WindowsHostImportSpec) HasNetworkSettings() bool {
	if o != nil && !isNil(o.NetworkSettings) {
		return true
	}

	return false
}

// SetNetworkSettings gets a reference to the given WindowsHostPortsModel and assigns it to the NetworkSettings field.
func (o *WindowsHostImportSpec) SetNetworkSettings(v WindowsHostPortsModel) {
	o.NetworkSettings = &v
}

func (o WindowsHostImportSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !isNil(o.NetworkSettings) {
		toSerialize["networkSettings"] = o.NetworkSettings
	}
	return json.Marshal(toSerialize)
}

type NullableWindowsHostImportSpec struct {
	value *WindowsHostImportSpec
	isSet bool
}

func (v NullableWindowsHostImportSpec) Get() *WindowsHostImportSpec {
	return v.value
}

func (v *NullableWindowsHostImportSpec) Set(val *WindowsHostImportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableWindowsHostImportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableWindowsHostImportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWindowsHostImportSpec(val *WindowsHostImportSpec) *NullableWindowsHostImportSpec {
	return &NullableWindowsHostImportSpec{value: val, isSet: true}
}

func (v NullableWindowsHostImportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWindowsHostImportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


