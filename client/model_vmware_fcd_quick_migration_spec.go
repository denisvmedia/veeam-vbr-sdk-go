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

// VmwareFcdQuickMigrationSpec struct for VmwareFcdQuickMigrationSpec
type VmwareFcdQuickMigrationSpec struct {
	// Array of disks that will be migrated to the `targetDatastore` associated with the `storagePolicy`.
	MountedDiskNames []string `json:"mountedDiskNames,omitempty"`
	TargetDatastore VmwareObjectModel `json:"targetDatastore"`
	StoragePolicy *VmwareObjectModel `json:"storagePolicy,omitempty"`
}

// NewVmwareFcdQuickMigrationSpec instantiates a new VmwareFcdQuickMigrationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmwareFcdQuickMigrationSpec(targetDatastore VmwareObjectModel) *VmwareFcdQuickMigrationSpec {
	this := VmwareFcdQuickMigrationSpec{}
	this.TargetDatastore = targetDatastore
	return &this
}

// NewVmwareFcdQuickMigrationSpecWithDefaults instantiates a new VmwareFcdQuickMigrationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmwareFcdQuickMigrationSpecWithDefaults() *VmwareFcdQuickMigrationSpec {
	this := VmwareFcdQuickMigrationSpec{}
	return &this
}

// GetMountedDiskNames returns the MountedDiskNames field value if set, zero value otherwise.
func (o *VmwareFcdQuickMigrationSpec) GetMountedDiskNames() []string {
	if o == nil || isNil(o.MountedDiskNames) {
		var ret []string
		return ret
	}
	return o.MountedDiskNames
}

// GetMountedDiskNamesOk returns a tuple with the MountedDiskNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareFcdQuickMigrationSpec) GetMountedDiskNamesOk() ([]string, bool) {
	if o == nil || isNil(o.MountedDiskNames) {
    return nil, false
	}
	return o.MountedDiskNames, true
}

// HasMountedDiskNames returns a boolean if a field has been set.
func (o *VmwareFcdQuickMigrationSpec) HasMountedDiskNames() bool {
	if o != nil && !isNil(o.MountedDiskNames) {
		return true
	}

	return false
}

// SetMountedDiskNames gets a reference to the given []string and assigns it to the MountedDiskNames field.
func (o *VmwareFcdQuickMigrationSpec) SetMountedDiskNames(v []string) {
	o.MountedDiskNames = v
}

// GetTargetDatastore returns the TargetDatastore field value
func (o *VmwareFcdQuickMigrationSpec) GetTargetDatastore() VmwareObjectModel {
	if o == nil {
		var ret VmwareObjectModel
		return ret
	}

	return o.TargetDatastore
}

// GetTargetDatastoreOk returns a tuple with the TargetDatastore field value
// and a boolean to check if the value has been set.
func (o *VmwareFcdQuickMigrationSpec) GetTargetDatastoreOk() (*VmwareObjectModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TargetDatastore, true
}

// SetTargetDatastore sets field value
func (o *VmwareFcdQuickMigrationSpec) SetTargetDatastore(v VmwareObjectModel) {
	o.TargetDatastore = v
}

// GetStoragePolicy returns the StoragePolicy field value if set, zero value otherwise.
func (o *VmwareFcdQuickMigrationSpec) GetStoragePolicy() VmwareObjectModel {
	if o == nil || isNil(o.StoragePolicy) {
		var ret VmwareObjectModel
		return ret
	}
	return *o.StoragePolicy
}

// GetStoragePolicyOk returns a tuple with the StoragePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmwareFcdQuickMigrationSpec) GetStoragePolicyOk() (*VmwareObjectModel, bool) {
	if o == nil || isNil(o.StoragePolicy) {
    return nil, false
	}
	return o.StoragePolicy, true
}

// HasStoragePolicy returns a boolean if a field has been set.
func (o *VmwareFcdQuickMigrationSpec) HasStoragePolicy() bool {
	if o != nil && !isNil(o.StoragePolicy) {
		return true
	}

	return false
}

// SetStoragePolicy gets a reference to the given VmwareObjectModel and assigns it to the StoragePolicy field.
func (o *VmwareFcdQuickMigrationSpec) SetStoragePolicy(v VmwareObjectModel) {
	o.StoragePolicy = &v
}

func (o VmwareFcdQuickMigrationSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MountedDiskNames) {
		toSerialize["mountedDiskNames"] = o.MountedDiskNames
	}
	if true {
		toSerialize["targetDatastore"] = o.TargetDatastore
	}
	if !isNil(o.StoragePolicy) {
		toSerialize["storagePolicy"] = o.StoragePolicy
	}
	return json.Marshal(toSerialize)
}

type NullableVmwareFcdQuickMigrationSpec struct {
	value *VmwareFcdQuickMigrationSpec
	isSet bool
}

func (v NullableVmwareFcdQuickMigrationSpec) Get() *VmwareFcdQuickMigrationSpec {
	return v.value
}

func (v *NullableVmwareFcdQuickMigrationSpec) Set(val *VmwareFcdQuickMigrationSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVmwareFcdQuickMigrationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVmwareFcdQuickMigrationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmwareFcdQuickMigrationSpec(val *VmwareFcdQuickMigrationSpec) *NullableVmwareFcdQuickMigrationSpec {
	return &NullableVmwareFcdQuickMigrationSpec{value: val, isSet: true}
}

func (v NullableVmwareFcdQuickMigrationSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmwareFcdQuickMigrationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


