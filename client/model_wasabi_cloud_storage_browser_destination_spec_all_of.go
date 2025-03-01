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

// WasabiCloudStorageBrowserDestinationSpecAllOf struct for WasabiCloudStorageBrowserDestinationSpecAllOf
type WasabiCloudStorageBrowserDestinationSpecAllOf struct {
	// ID of a server you want to use to connect to the object storage. You can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used.
	HostId *string `json:"hostId,omitempty"`
	// Region where the bucket is located.
	RegionId string `json:"regionId"`
	// Name of the bucket where you want to store your backup data.
	BucketName string `json:"bucketName"`
}

// NewWasabiCloudStorageBrowserDestinationSpecAllOf instantiates a new WasabiCloudStorageBrowserDestinationSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWasabiCloudStorageBrowserDestinationSpecAllOf(regionId string, bucketName string) *WasabiCloudStorageBrowserDestinationSpecAllOf {
	this := WasabiCloudStorageBrowserDestinationSpecAllOf{}
	this.RegionId = regionId
	this.BucketName = bucketName
	return &this
}

// NewWasabiCloudStorageBrowserDestinationSpecAllOfWithDefaults instantiates a new WasabiCloudStorageBrowserDestinationSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWasabiCloudStorageBrowserDestinationSpecAllOfWithDefaults() *WasabiCloudStorageBrowserDestinationSpecAllOf {
	this := WasabiCloudStorageBrowserDestinationSpecAllOf{}
	return &this
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *WasabiCloudStorageBrowserDestinationSpecAllOf) GetHostId() string {
	if o == nil || isNil(o.HostId) {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WasabiCloudStorageBrowserDestinationSpecAllOf) GetHostIdOk() (*string, bool) {
	if o == nil || isNil(o.HostId) {
    return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *WasabiCloudStorageBrowserDestinationSpecAllOf) HasHostId() bool {
	if o != nil && !isNil(o.HostId) {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *WasabiCloudStorageBrowserDestinationSpecAllOf) SetHostId(v string) {
	o.HostId = &v
}

// GetRegionId returns the RegionId field value
func (o *WasabiCloudStorageBrowserDestinationSpecAllOf) GetRegionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value
// and a boolean to check if the value has been set.
func (o *WasabiCloudStorageBrowserDestinationSpecAllOf) GetRegionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RegionId, true
}

// SetRegionId sets field value
func (o *WasabiCloudStorageBrowserDestinationSpecAllOf) SetRegionId(v string) {
	o.RegionId = v
}

// GetBucketName returns the BucketName field value
func (o *WasabiCloudStorageBrowserDestinationSpecAllOf) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *WasabiCloudStorageBrowserDestinationSpecAllOf) GetBucketNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value
func (o *WasabiCloudStorageBrowserDestinationSpecAllOf) SetBucketName(v string) {
	o.BucketName = v
}

func (o WasabiCloudStorageBrowserDestinationSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.HostId) {
		toSerialize["hostId"] = o.HostId
	}
	if true {
		toSerialize["regionId"] = o.RegionId
	}
	if true {
		toSerialize["bucketName"] = o.BucketName
	}
	return json.Marshal(toSerialize)
}

type NullableWasabiCloudStorageBrowserDestinationSpecAllOf struct {
	value *WasabiCloudStorageBrowserDestinationSpecAllOf
	isSet bool
}

func (v NullableWasabiCloudStorageBrowserDestinationSpecAllOf) Get() *WasabiCloudStorageBrowserDestinationSpecAllOf {
	return v.value
}

func (v *NullableWasabiCloudStorageBrowserDestinationSpecAllOf) Set(val *WasabiCloudStorageBrowserDestinationSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWasabiCloudStorageBrowserDestinationSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWasabiCloudStorageBrowserDestinationSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWasabiCloudStorageBrowserDestinationSpecAllOf(val *WasabiCloudStorageBrowserDestinationSpecAllOf) *NullableWasabiCloudStorageBrowserDestinationSpecAllOf {
	return &NullableWasabiCloudStorageBrowserDestinationSpecAllOf{value: val, isSet: true}
}

func (v NullableWasabiCloudStorageBrowserDestinationSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWasabiCloudStorageBrowserDestinationSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


