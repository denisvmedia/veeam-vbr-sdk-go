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

// AmazonSnowballEdgeStorageAccountModel AWS account used to access the AWS Snowball Edge storage.
type AmazonSnowballEdgeStorageAccountModel struct {
	// Service point address and port number of the AWS Snowball Edge device.
	ServicePoint string `json:"servicePoint"`
	// For AWS Snowball Edge, the region is `snow`.
	RegionId string `json:"regionId"`
	// ID of the cloud credentials record.
	CredentialsId string `json:"credentialsId"`
	ConnectionSettings *ObjectStorageConnectionModel `json:"connectionSettings,omitempty"`
}

// NewAmazonSnowballEdgeStorageAccountModel instantiates a new AmazonSnowballEdgeStorageAccountModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonSnowballEdgeStorageAccountModel(servicePoint string, regionId string, credentialsId string) *AmazonSnowballEdgeStorageAccountModel {
	this := AmazonSnowballEdgeStorageAccountModel{}
	this.ServicePoint = servicePoint
	this.RegionId = regionId
	this.CredentialsId = credentialsId
	return &this
}

// NewAmazonSnowballEdgeStorageAccountModelWithDefaults instantiates a new AmazonSnowballEdgeStorageAccountModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonSnowballEdgeStorageAccountModelWithDefaults() *AmazonSnowballEdgeStorageAccountModel {
	this := AmazonSnowballEdgeStorageAccountModel{}
	return &this
}

// GetServicePoint returns the ServicePoint field value
func (o *AmazonSnowballEdgeStorageAccountModel) GetServicePoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServicePoint
}

// GetServicePointOk returns a tuple with the ServicePoint field value
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeStorageAccountModel) GetServicePointOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServicePoint, true
}

// SetServicePoint sets field value
func (o *AmazonSnowballEdgeStorageAccountModel) SetServicePoint(v string) {
	o.ServicePoint = v
}

// GetRegionId returns the RegionId field value
func (o *AmazonSnowballEdgeStorageAccountModel) GetRegionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeStorageAccountModel) GetRegionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RegionId, true
}

// SetRegionId sets field value
func (o *AmazonSnowballEdgeStorageAccountModel) SetRegionId(v string) {
	o.RegionId = v
}

// GetCredentialsId returns the CredentialsId field value
func (o *AmazonSnowballEdgeStorageAccountModel) GetCredentialsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CredentialsId
}

// GetCredentialsIdOk returns a tuple with the CredentialsId field value
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeStorageAccountModel) GetCredentialsIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CredentialsId, true
}

// SetCredentialsId sets field value
func (o *AmazonSnowballEdgeStorageAccountModel) SetCredentialsId(v string) {
	o.CredentialsId = v
}

// GetConnectionSettings returns the ConnectionSettings field value if set, zero value otherwise.
func (o *AmazonSnowballEdgeStorageAccountModel) GetConnectionSettings() ObjectStorageConnectionModel {
	if o == nil || isNil(o.ConnectionSettings) {
		var ret ObjectStorageConnectionModel
		return ret
	}
	return *o.ConnectionSettings
}

// GetConnectionSettingsOk returns a tuple with the ConnectionSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeStorageAccountModel) GetConnectionSettingsOk() (*ObjectStorageConnectionModel, bool) {
	if o == nil || isNil(o.ConnectionSettings) {
    return nil, false
	}
	return o.ConnectionSettings, true
}

// HasConnectionSettings returns a boolean if a field has been set.
func (o *AmazonSnowballEdgeStorageAccountModel) HasConnectionSettings() bool {
	if o != nil && !isNil(o.ConnectionSettings) {
		return true
	}

	return false
}

// SetConnectionSettings gets a reference to the given ObjectStorageConnectionModel and assigns it to the ConnectionSettings field.
func (o *AmazonSnowballEdgeStorageAccountModel) SetConnectionSettings(v ObjectStorageConnectionModel) {
	o.ConnectionSettings = &v
}

func (o AmazonSnowballEdgeStorageAccountModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["servicePoint"] = o.ServicePoint
	}
	if true {
		toSerialize["regionId"] = o.RegionId
	}
	if true {
		toSerialize["credentialsId"] = o.CredentialsId
	}
	if !isNil(o.ConnectionSettings) {
		toSerialize["connectionSettings"] = o.ConnectionSettings
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonSnowballEdgeStorageAccountModel struct {
	value *AmazonSnowballEdgeStorageAccountModel
	isSet bool
}

func (v NullableAmazonSnowballEdgeStorageAccountModel) Get() *AmazonSnowballEdgeStorageAccountModel {
	return v.value
}

func (v *NullableAmazonSnowballEdgeStorageAccountModel) Set(val *AmazonSnowballEdgeStorageAccountModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonSnowballEdgeStorageAccountModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonSnowballEdgeStorageAccountModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonSnowballEdgeStorageAccountModel(val *AmazonSnowballEdgeStorageAccountModel) *NullableAmazonSnowballEdgeStorageAccountModel {
	return &NullableAmazonSnowballEdgeStorageAccountModel{value: val, isSet: true}
}

func (v NullableAmazonSnowballEdgeStorageAccountModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonSnowballEdgeStorageAccountModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


