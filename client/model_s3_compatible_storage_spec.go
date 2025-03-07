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

// S3CompatibleStorageSpec S3 compatible storage.
type S3CompatibleStorageSpec struct {
	RepositorySpec
	// If *true*, the maximum number of concurrent tasks is limited.
	EnableTaskLimit *bool `json:"enableTaskLimit,omitempty"`
	// Maximum number of concurrent tasks.
	MaxTaskCount *int32 `json:"maxTaskCount,omitempty"`
	Account S3CompatibleStorageAccountModel `json:"account"`
	Bucket S3CompatibleStorageBucketModel `json:"bucket"`
	MountServer MountServerSettingsModel `json:"mountServer"`
	ProxyAppliance *S3CompatibleProxyModel `json:"proxyAppliance,omitempty"`
}

// NewS3CompatibleStorageSpec instantiates a new S3CompatibleStorageSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3CompatibleStorageSpec(account S3CompatibleStorageAccountModel, bucket S3CompatibleStorageBucketModel, mountServer MountServerSettingsModel, name string, description string, type_ ERepositoryType) *S3CompatibleStorageSpec {
	this := S3CompatibleStorageSpec{}
	this.Name = name
	this.Description = description
	this.Type = type_
	this.Account = account
	this.Bucket = bucket
	this.MountServer = mountServer
	return &this
}

// NewS3CompatibleStorageSpecWithDefaults instantiates a new S3CompatibleStorageSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3CompatibleStorageSpecWithDefaults() *S3CompatibleStorageSpec {
	this := S3CompatibleStorageSpec{}
	return &this
}

// GetEnableTaskLimit returns the EnableTaskLimit field value if set, zero value otherwise.
func (o *S3CompatibleStorageSpec) GetEnableTaskLimit() bool {
	if o == nil || isNil(o.EnableTaskLimit) {
		var ret bool
		return ret
	}
	return *o.EnableTaskLimit
}

// GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageSpec) GetEnableTaskLimitOk() (*bool, bool) {
	if o == nil || isNil(o.EnableTaskLimit) {
    return nil, false
	}
	return o.EnableTaskLimit, true
}

// HasEnableTaskLimit returns a boolean if a field has been set.
func (o *S3CompatibleStorageSpec) HasEnableTaskLimit() bool {
	if o != nil && !isNil(o.EnableTaskLimit) {
		return true
	}

	return false
}

// SetEnableTaskLimit gets a reference to the given bool and assigns it to the EnableTaskLimit field.
func (o *S3CompatibleStorageSpec) SetEnableTaskLimit(v bool) {
	o.EnableTaskLimit = &v
}

// GetMaxTaskCount returns the MaxTaskCount field value if set, zero value otherwise.
func (o *S3CompatibleStorageSpec) GetMaxTaskCount() int32 {
	if o == nil || isNil(o.MaxTaskCount) {
		var ret int32
		return ret
	}
	return *o.MaxTaskCount
}

// GetMaxTaskCountOk returns a tuple with the MaxTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageSpec) GetMaxTaskCountOk() (*int32, bool) {
	if o == nil || isNil(o.MaxTaskCount) {
    return nil, false
	}
	return o.MaxTaskCount, true
}

// HasMaxTaskCount returns a boolean if a field has been set.
func (o *S3CompatibleStorageSpec) HasMaxTaskCount() bool {
	if o != nil && !isNil(o.MaxTaskCount) {
		return true
	}

	return false
}

// SetMaxTaskCount gets a reference to the given int32 and assigns it to the MaxTaskCount field.
func (o *S3CompatibleStorageSpec) SetMaxTaskCount(v int32) {
	o.MaxTaskCount = &v
}

// GetAccount returns the Account field value
func (o *S3CompatibleStorageSpec) GetAccount() S3CompatibleStorageAccountModel {
	if o == nil {
		var ret S3CompatibleStorageAccountModel
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageSpec) GetAccountOk() (*S3CompatibleStorageAccountModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *S3CompatibleStorageSpec) SetAccount(v S3CompatibleStorageAccountModel) {
	o.Account = v
}

// GetBucket returns the Bucket field value
func (o *S3CompatibleStorageSpec) GetBucket() S3CompatibleStorageBucketModel {
	if o == nil {
		var ret S3CompatibleStorageBucketModel
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageSpec) GetBucketOk() (*S3CompatibleStorageBucketModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *S3CompatibleStorageSpec) SetBucket(v S3CompatibleStorageBucketModel) {
	o.Bucket = v
}

// GetMountServer returns the MountServer field value
func (o *S3CompatibleStorageSpec) GetMountServer() MountServerSettingsModel {
	if o == nil {
		var ret MountServerSettingsModel
		return ret
	}

	return o.MountServer
}

// GetMountServerOk returns a tuple with the MountServer field value
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageSpec) GetMountServerOk() (*MountServerSettingsModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MountServer, true
}

// SetMountServer sets field value
func (o *S3CompatibleStorageSpec) SetMountServer(v MountServerSettingsModel) {
	o.MountServer = v
}

// GetProxyAppliance returns the ProxyAppliance field value if set, zero value otherwise.
func (o *S3CompatibleStorageSpec) GetProxyAppliance() S3CompatibleProxyModel {
	if o == nil || isNil(o.ProxyAppliance) {
		var ret S3CompatibleProxyModel
		return ret
	}
	return *o.ProxyAppliance
}

// GetProxyApplianceOk returns a tuple with the ProxyAppliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageSpec) GetProxyApplianceOk() (*S3CompatibleProxyModel, bool) {
	if o == nil || isNil(o.ProxyAppliance) {
    return nil, false
	}
	return o.ProxyAppliance, true
}

// HasProxyAppliance returns a boolean if a field has been set.
func (o *S3CompatibleStorageSpec) HasProxyAppliance() bool {
	if o != nil && !isNil(o.ProxyAppliance) {
		return true
	}

	return false
}

// SetProxyAppliance gets a reference to the given S3CompatibleProxyModel and assigns it to the ProxyAppliance field.
func (o *S3CompatibleStorageSpec) SetProxyAppliance(v S3CompatibleProxyModel) {
	o.ProxyAppliance = &v
}

func (o S3CompatibleStorageSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRepositorySpec, errRepositorySpec := json.Marshal(o.RepositorySpec)
	if errRepositorySpec != nil {
		return []byte{}, errRepositorySpec
	}
	errRepositorySpec = json.Unmarshal([]byte(serializedRepositorySpec), &toSerialize)
	if errRepositorySpec != nil {
		return []byte{}, errRepositorySpec
	}
	if !isNil(o.EnableTaskLimit) {
		toSerialize["enableTaskLimit"] = o.EnableTaskLimit
	}
	if !isNil(o.MaxTaskCount) {
		toSerialize["maxTaskCount"] = o.MaxTaskCount
	}
	if true {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["bucket"] = o.Bucket
	}
	if true {
		toSerialize["mountServer"] = o.MountServer
	}
	if !isNil(o.ProxyAppliance) {
		toSerialize["proxyAppliance"] = o.ProxyAppliance
	}
	return json.Marshal(toSerialize)
}

type NullableS3CompatibleStorageSpec struct {
	value *S3CompatibleStorageSpec
	isSet bool
}

func (v NullableS3CompatibleStorageSpec) Get() *S3CompatibleStorageSpec {
	return v.value
}

func (v *NullableS3CompatibleStorageSpec) Set(val *S3CompatibleStorageSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableS3CompatibleStorageSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableS3CompatibleStorageSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3CompatibleStorageSpec(val *S3CompatibleStorageSpec) *NullableS3CompatibleStorageSpec {
	return &NullableS3CompatibleStorageSpec{value: val, isSet: true}
}

func (v NullableS3CompatibleStorageSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3CompatibleStorageSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


