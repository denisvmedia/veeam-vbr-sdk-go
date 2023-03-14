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

// IBMCloudStorageBucketModel IBM Cloud bucket where backup data is stored.
type IBMCloudStorageBucketModel struct {
	// Name of an IBM Cloud bucket.
	BucketName string `json:"bucketName"`
	// Name of the folder that the object storage repository is mapped to.
	FolderName string `json:"folderName"`
	StorageConsumptionLimit *ObjectStorageConsumptionLimitModel `json:"storageConsumptionLimit,omitempty"`
	Immutability *ObjectStorageImmutabilityModel `json:"immutability,omitempty"`
}

// NewIBMCloudStorageBucketModel instantiates a new IBMCloudStorageBucketModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIBMCloudStorageBucketModel(bucketName string, folderName string) *IBMCloudStorageBucketModel {
	this := IBMCloudStorageBucketModel{}
	this.BucketName = bucketName
	this.FolderName = folderName
	return &this
}

// NewIBMCloudStorageBucketModelWithDefaults instantiates a new IBMCloudStorageBucketModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIBMCloudStorageBucketModelWithDefaults() *IBMCloudStorageBucketModel {
	this := IBMCloudStorageBucketModel{}
	return &this
}

// GetBucketName returns the BucketName field value
func (o *IBMCloudStorageBucketModel) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageBucketModel) GetBucketNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value
func (o *IBMCloudStorageBucketModel) SetBucketName(v string) {
	o.BucketName = v
}

// GetFolderName returns the FolderName field value
func (o *IBMCloudStorageBucketModel) GetFolderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FolderName
}

// GetFolderNameOk returns a tuple with the FolderName field value
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageBucketModel) GetFolderNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FolderName, true
}

// SetFolderName sets field value
func (o *IBMCloudStorageBucketModel) SetFolderName(v string) {
	o.FolderName = v
}

// GetStorageConsumptionLimit returns the StorageConsumptionLimit field value if set, zero value otherwise.
func (o *IBMCloudStorageBucketModel) GetStorageConsumptionLimit() ObjectStorageConsumptionLimitModel {
	if o == nil || isNil(o.StorageConsumptionLimit) {
		var ret ObjectStorageConsumptionLimitModel
		return ret
	}
	return *o.StorageConsumptionLimit
}

// GetStorageConsumptionLimitOk returns a tuple with the StorageConsumptionLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageBucketModel) GetStorageConsumptionLimitOk() (*ObjectStorageConsumptionLimitModel, bool) {
	if o == nil || isNil(o.StorageConsumptionLimit) {
    return nil, false
	}
	return o.StorageConsumptionLimit, true
}

// HasStorageConsumptionLimit returns a boolean if a field has been set.
func (o *IBMCloudStorageBucketModel) HasStorageConsumptionLimit() bool {
	if o != nil && !isNil(o.StorageConsumptionLimit) {
		return true
	}

	return false
}

// SetStorageConsumptionLimit gets a reference to the given ObjectStorageConsumptionLimitModel and assigns it to the StorageConsumptionLimit field.
func (o *IBMCloudStorageBucketModel) SetStorageConsumptionLimit(v ObjectStorageConsumptionLimitModel) {
	o.StorageConsumptionLimit = &v
}

// GetImmutability returns the Immutability field value if set, zero value otherwise.
func (o *IBMCloudStorageBucketModel) GetImmutability() ObjectStorageImmutabilityModel {
	if o == nil || isNil(o.Immutability) {
		var ret ObjectStorageImmutabilityModel
		return ret
	}
	return *o.Immutability
}

// GetImmutabilityOk returns a tuple with the Immutability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IBMCloudStorageBucketModel) GetImmutabilityOk() (*ObjectStorageImmutabilityModel, bool) {
	if o == nil || isNil(o.Immutability) {
    return nil, false
	}
	return o.Immutability, true
}

// HasImmutability returns a boolean if a field has been set.
func (o *IBMCloudStorageBucketModel) HasImmutability() bool {
	if o != nil && !isNil(o.Immutability) {
		return true
	}

	return false
}

// SetImmutability gets a reference to the given ObjectStorageImmutabilityModel and assigns it to the Immutability field.
func (o *IBMCloudStorageBucketModel) SetImmutability(v ObjectStorageImmutabilityModel) {
	o.Immutability = &v
}

func (o IBMCloudStorageBucketModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bucketName"] = o.BucketName
	}
	if true {
		toSerialize["folderName"] = o.FolderName
	}
	if !isNil(o.StorageConsumptionLimit) {
		toSerialize["storageConsumptionLimit"] = o.StorageConsumptionLimit
	}
	if !isNil(o.Immutability) {
		toSerialize["immutability"] = o.Immutability
	}
	return json.Marshal(toSerialize)
}

type NullableIBMCloudStorageBucketModel struct {
	value *IBMCloudStorageBucketModel
	isSet bool
}

func (v NullableIBMCloudStorageBucketModel) Get() *IBMCloudStorageBucketModel {
	return v.value
}

func (v *NullableIBMCloudStorageBucketModel) Set(val *IBMCloudStorageBucketModel) {
	v.value = val
	v.isSet = true
}

func (v NullableIBMCloudStorageBucketModel) IsSet() bool {
	return v.isSet
}

func (v *NullableIBMCloudStorageBucketModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIBMCloudStorageBucketModel(val *IBMCloudStorageBucketModel) *NullableIBMCloudStorageBucketModel {
	return &NullableIBMCloudStorageBucketModel{value: val, isSet: true}
}

func (v NullableIBMCloudStorageBucketModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIBMCloudStorageBucketModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


