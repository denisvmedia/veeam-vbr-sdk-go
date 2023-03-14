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

// AmazonS3StorageAccountModel Account used to access the Amazon S3 storage.
type AmazonS3StorageAccountModel struct {
	// ID of the cloud credentials record.
	CredentialsId string `json:"credentialsId"`
	RegionType EAmazonRegionType `json:"regionType"`
	ConnectionSettings *ObjectStorageConnectionModel `json:"connectionSettings,omitempty"`
}

// NewAmazonS3StorageAccountModel instantiates a new AmazonS3StorageAccountModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonS3StorageAccountModel(credentialsId string, regionType EAmazonRegionType) *AmazonS3StorageAccountModel {
	this := AmazonS3StorageAccountModel{}
	this.CredentialsId = credentialsId
	this.RegionType = regionType
	return &this
}

// NewAmazonS3StorageAccountModelWithDefaults instantiates a new AmazonS3StorageAccountModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonS3StorageAccountModelWithDefaults() *AmazonS3StorageAccountModel {
	this := AmazonS3StorageAccountModel{}
	return &this
}

// GetCredentialsId returns the CredentialsId field value
func (o *AmazonS3StorageAccountModel) GetCredentialsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CredentialsId
}

// GetCredentialsIdOk returns a tuple with the CredentialsId field value
// and a boolean to check if the value has been set.
func (o *AmazonS3StorageAccountModel) GetCredentialsIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CredentialsId, true
}

// SetCredentialsId sets field value
func (o *AmazonS3StorageAccountModel) SetCredentialsId(v string) {
	o.CredentialsId = v
}

// GetRegionType returns the RegionType field value
func (o *AmazonS3StorageAccountModel) GetRegionType() EAmazonRegionType {
	if o == nil {
		var ret EAmazonRegionType
		return ret
	}

	return o.RegionType
}

// GetRegionTypeOk returns a tuple with the RegionType field value
// and a boolean to check if the value has been set.
func (o *AmazonS3StorageAccountModel) GetRegionTypeOk() (*EAmazonRegionType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RegionType, true
}

// SetRegionType sets field value
func (o *AmazonS3StorageAccountModel) SetRegionType(v EAmazonRegionType) {
	o.RegionType = v
}

// GetConnectionSettings returns the ConnectionSettings field value if set, zero value otherwise.
func (o *AmazonS3StorageAccountModel) GetConnectionSettings() ObjectStorageConnectionModel {
	if o == nil || isNil(o.ConnectionSettings) {
		var ret ObjectStorageConnectionModel
		return ret
	}
	return *o.ConnectionSettings
}

// GetConnectionSettingsOk returns a tuple with the ConnectionSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonS3StorageAccountModel) GetConnectionSettingsOk() (*ObjectStorageConnectionModel, bool) {
	if o == nil || isNil(o.ConnectionSettings) {
    return nil, false
	}
	return o.ConnectionSettings, true
}

// HasConnectionSettings returns a boolean if a field has been set.
func (o *AmazonS3StorageAccountModel) HasConnectionSettings() bool {
	if o != nil && !isNil(o.ConnectionSettings) {
		return true
	}

	return false
}

// SetConnectionSettings gets a reference to the given ObjectStorageConnectionModel and assigns it to the ConnectionSettings field.
func (o *AmazonS3StorageAccountModel) SetConnectionSettings(v ObjectStorageConnectionModel) {
	o.ConnectionSettings = &v
}

func (o AmazonS3StorageAccountModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["credentialsId"] = o.CredentialsId
	}
	if true {
		toSerialize["regionType"] = o.RegionType
	}
	if !isNil(o.ConnectionSettings) {
		toSerialize["connectionSettings"] = o.ConnectionSettings
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonS3StorageAccountModel struct {
	value *AmazonS3StorageAccountModel
	isSet bool
}

func (v NullableAmazonS3StorageAccountModel) Get() *AmazonS3StorageAccountModel {
	return v.value
}

func (v *NullableAmazonS3StorageAccountModel) Set(val *AmazonS3StorageAccountModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonS3StorageAccountModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonS3StorageAccountModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonS3StorageAccountModel(val *AmazonS3StorageAccountModel) *NullableAmazonS3StorageAccountModel {
	return &NullableAmazonS3StorageAccountModel{value: val, isSet: true}
}

func (v NullableAmazonS3StorageAccountModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonS3StorageAccountModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


