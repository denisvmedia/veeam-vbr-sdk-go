/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 
 *
 * API version: 1.0-rev1
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SmbStorageSpecAllOf struct for SmbStorageSpecAllOf
type SmbStorageSpecAllOf struct {
	Share SmbRepositoryShareSettingsModel `json:"share"`
	Repository NetworkRepositorySettingsModel `json:"repository"`
	MountServer MountServerSettingsModel `json:"mountServer"`
}

// NewSmbStorageSpecAllOf instantiates a new SmbStorageSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmbStorageSpecAllOf(share SmbRepositoryShareSettingsModel, repository NetworkRepositorySettingsModel, mountServer MountServerSettingsModel) *SmbStorageSpecAllOf {
	this := SmbStorageSpecAllOf{}
	this.Share = share
	this.Repository = repository
	this.MountServer = mountServer
	return &this
}

// NewSmbStorageSpecAllOfWithDefaults instantiates a new SmbStorageSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmbStorageSpecAllOfWithDefaults() *SmbStorageSpecAllOf {
	this := SmbStorageSpecAllOf{}
	return &this
}

// GetShare returns the Share field value
func (o *SmbStorageSpecAllOf) GetShare() SmbRepositoryShareSettingsModel {
	if o == nil {
		var ret SmbRepositoryShareSettingsModel
		return ret
	}

	return o.Share
}

// GetShareOk returns a tuple with the Share field value
// and a boolean to check if the value has been set.
func (o *SmbStorageSpecAllOf) GetShareOk() (*SmbRepositoryShareSettingsModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Share, true
}

// SetShare sets field value
func (o *SmbStorageSpecAllOf) SetShare(v SmbRepositoryShareSettingsModel) {
	o.Share = v
}

// GetRepository returns the Repository field value
func (o *SmbStorageSpecAllOf) GetRepository() NetworkRepositorySettingsModel {
	if o == nil {
		var ret NetworkRepositorySettingsModel
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *SmbStorageSpecAllOf) GetRepositoryOk() (*NetworkRepositorySettingsModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *SmbStorageSpecAllOf) SetRepository(v NetworkRepositorySettingsModel) {
	o.Repository = v
}

// GetMountServer returns the MountServer field value
func (o *SmbStorageSpecAllOf) GetMountServer() MountServerSettingsModel {
	if o == nil {
		var ret MountServerSettingsModel
		return ret
	}

	return o.MountServer
}

// GetMountServerOk returns a tuple with the MountServer field value
// and a boolean to check if the value has been set.
func (o *SmbStorageSpecAllOf) GetMountServerOk() (*MountServerSettingsModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MountServer, true
}

// SetMountServer sets field value
func (o *SmbStorageSpecAllOf) SetMountServer(v MountServerSettingsModel) {
	o.MountServer = v
}

func (o SmbStorageSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["share"] = o.Share
	}
	if true {
		toSerialize["repository"] = o.Repository
	}
	if true {
		toSerialize["mountServer"] = o.MountServer
	}
	return json.Marshal(toSerialize)
}

type NullableSmbStorageSpecAllOf struct {
	value *SmbStorageSpecAllOf
	isSet bool
}

func (v NullableSmbStorageSpecAllOf) Get() *SmbStorageSpecAllOf {
	return v.value
}

func (v *NullableSmbStorageSpecAllOf) Set(val *SmbStorageSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSmbStorageSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSmbStorageSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmbStorageSpecAllOf(val *SmbStorageSpecAllOf) *NullableSmbStorageSpecAllOf {
	return &NullableSmbStorageSpecAllOf{value: val, isSet: true}
}

func (v NullableSmbStorageSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmbStorageSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


