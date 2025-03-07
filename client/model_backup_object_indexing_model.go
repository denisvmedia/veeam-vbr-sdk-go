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

// BackupObjectIndexingModel Guest OS indexing options for the VM.
type BackupObjectIndexingModel struct {
	GuestFSIndexingMode EGuestFSIndexingMode `json:"guestFSIndexingMode"`
	// Array of folders. Environmental variables and full paths to folders can be used.
	IndexingList []string `json:"indexingList,omitempty"`
}

// NewBackupObjectIndexingModel instantiates a new BackupObjectIndexingModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupObjectIndexingModel(guestFSIndexingMode EGuestFSIndexingMode) *BackupObjectIndexingModel {
	this := BackupObjectIndexingModel{}
	this.GuestFSIndexingMode = guestFSIndexingMode
	return &this
}

// NewBackupObjectIndexingModelWithDefaults instantiates a new BackupObjectIndexingModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupObjectIndexingModelWithDefaults() *BackupObjectIndexingModel {
	this := BackupObjectIndexingModel{}
	return &this
}

// GetGuestFSIndexingMode returns the GuestFSIndexingMode field value
func (o *BackupObjectIndexingModel) GetGuestFSIndexingMode() EGuestFSIndexingMode {
	if o == nil {
		var ret EGuestFSIndexingMode
		return ret
	}

	return o.GuestFSIndexingMode
}

// GetGuestFSIndexingModeOk returns a tuple with the GuestFSIndexingMode field value
// and a boolean to check if the value has been set.
func (o *BackupObjectIndexingModel) GetGuestFSIndexingModeOk() (*EGuestFSIndexingMode, bool) {
	if o == nil {
    return nil, false
	}
	return &o.GuestFSIndexingMode, true
}

// SetGuestFSIndexingMode sets field value
func (o *BackupObjectIndexingModel) SetGuestFSIndexingMode(v EGuestFSIndexingMode) {
	o.GuestFSIndexingMode = v
}

// GetIndexingList returns the IndexingList field value if set, zero value otherwise.
func (o *BackupObjectIndexingModel) GetIndexingList() []string {
	if o == nil || isNil(o.IndexingList) {
		var ret []string
		return ret
	}
	return o.IndexingList
}

// GetIndexingListOk returns a tuple with the IndexingList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupObjectIndexingModel) GetIndexingListOk() ([]string, bool) {
	if o == nil || isNil(o.IndexingList) {
    return nil, false
	}
	return o.IndexingList, true
}

// HasIndexingList returns a boolean if a field has been set.
func (o *BackupObjectIndexingModel) HasIndexingList() bool {
	if o != nil && !isNil(o.IndexingList) {
		return true
	}

	return false
}

// SetIndexingList gets a reference to the given []string and assigns it to the IndexingList field.
func (o *BackupObjectIndexingModel) SetIndexingList(v []string) {
	o.IndexingList = v
}

func (o BackupObjectIndexingModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["guestFSIndexingMode"] = o.GuestFSIndexingMode
	}
	if !isNil(o.IndexingList) {
		toSerialize["indexingList"] = o.IndexingList
	}
	return json.Marshal(toSerialize)
}

type NullableBackupObjectIndexingModel struct {
	value *BackupObjectIndexingModel
	isSet bool
}

func (v NullableBackupObjectIndexingModel) Get() *BackupObjectIndexingModel {
	return v.value
}

func (v *NullableBackupObjectIndexingModel) Set(val *BackupObjectIndexingModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupObjectIndexingModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupObjectIndexingModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupObjectIndexingModel(val *BackupObjectIndexingModel) *NullableBackupObjectIndexingModel {
	return &NullableBackupObjectIndexingModel{value: val, isSet: true}
}

func (v NullableBackupObjectIndexingModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupObjectIndexingModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


