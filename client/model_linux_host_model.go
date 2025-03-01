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

// LinuxHostModel struct for LinuxHostModel
type LinuxHostModel struct {
	ManagedServerModel
	SshSettings *LinuxHostSSHSettingsModel `json:"sshSettings,omitempty"`
	CredentialsStorageType *ECredentialsStorageType `json:"credentialsStorageType,omitempty"`
}

// NewLinuxHostModel instantiates a new LinuxHostModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinuxHostModel(id string, name string, description string, type_ EManagedServerType, credentialsId string) *LinuxHostModel {
	this := LinuxHostModel{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Type = type_
	this.CredentialsId = credentialsId
	return &this
}

// NewLinuxHostModelWithDefaults instantiates a new LinuxHostModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinuxHostModelWithDefaults() *LinuxHostModel {
	this := LinuxHostModel{}
	return &this
}

// GetSshSettings returns the SshSettings field value if set, zero value otherwise.
func (o *LinuxHostModel) GetSshSettings() LinuxHostSSHSettingsModel {
	if o == nil || isNil(o.SshSettings) {
		var ret LinuxHostSSHSettingsModel
		return ret
	}
	return *o.SshSettings
}

// GetSshSettingsOk returns a tuple with the SshSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxHostModel) GetSshSettingsOk() (*LinuxHostSSHSettingsModel, bool) {
	if o == nil || isNil(o.SshSettings) {
    return nil, false
	}
	return o.SshSettings, true
}

// HasSshSettings returns a boolean if a field has been set.
func (o *LinuxHostModel) HasSshSettings() bool {
	if o != nil && !isNil(o.SshSettings) {
		return true
	}

	return false
}

// SetSshSettings gets a reference to the given LinuxHostSSHSettingsModel and assigns it to the SshSettings field.
func (o *LinuxHostModel) SetSshSettings(v LinuxHostSSHSettingsModel) {
	o.SshSettings = &v
}

// GetCredentialsStorageType returns the CredentialsStorageType field value if set, zero value otherwise.
func (o *LinuxHostModel) GetCredentialsStorageType() ECredentialsStorageType {
	if o == nil || isNil(o.CredentialsStorageType) {
		var ret ECredentialsStorageType
		return ret
	}
	return *o.CredentialsStorageType
}

// GetCredentialsStorageTypeOk returns a tuple with the CredentialsStorageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxHostModel) GetCredentialsStorageTypeOk() (*ECredentialsStorageType, bool) {
	if o == nil || isNil(o.CredentialsStorageType) {
    return nil, false
	}
	return o.CredentialsStorageType, true
}

// HasCredentialsStorageType returns a boolean if a field has been set.
func (o *LinuxHostModel) HasCredentialsStorageType() bool {
	if o != nil && !isNil(o.CredentialsStorageType) {
		return true
	}

	return false
}

// SetCredentialsStorageType gets a reference to the given ECredentialsStorageType and assigns it to the CredentialsStorageType field.
func (o *LinuxHostModel) SetCredentialsStorageType(v ECredentialsStorageType) {
	o.CredentialsStorageType = &v
}

func (o LinuxHostModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedManagedServerModel, errManagedServerModel := json.Marshal(o.ManagedServerModel)
	if errManagedServerModel != nil {
		return []byte{}, errManagedServerModel
	}
	errManagedServerModel = json.Unmarshal([]byte(serializedManagedServerModel), &toSerialize)
	if errManagedServerModel != nil {
		return []byte{}, errManagedServerModel
	}
	if !isNil(o.SshSettings) {
		toSerialize["sshSettings"] = o.SshSettings
	}
	if !isNil(o.CredentialsStorageType) {
		toSerialize["credentialsStorageType"] = o.CredentialsStorageType
	}
	return json.Marshal(toSerialize)
}

type NullableLinuxHostModel struct {
	value *LinuxHostModel
	isSet bool
}

func (v NullableLinuxHostModel) Get() *LinuxHostModel {
	return v.value
}

func (v *NullableLinuxHostModel) Set(val *LinuxHostModel) {
	v.value = val
	v.isSet = true
}

func (v NullableLinuxHostModel) IsSet() bool {
	return v.isSet
}

func (v *NullableLinuxHostModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinuxHostModel(val *LinuxHostModel) *NullableLinuxHostModel {
	return &NullableLinuxHostModel{value: val, isSet: true}
}

func (v NullableLinuxHostModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinuxHostModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


