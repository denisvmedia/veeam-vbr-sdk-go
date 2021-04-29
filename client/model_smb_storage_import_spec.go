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

// SmbStorageImportSpec struct for SmbStorageImportSpec
type SmbStorageImportSpec struct {
	// Name of the backup repository.
	Name string `json:"name"`
	// Description of the backup repository.
	Description string `json:"description"`
	// VMware vSphere tag assigned to the backup repository.
	Tag string `json:"tag"`
	Type ERepositoryType `json:"type"`
	Share SmbRepositoryShareSettingsSpec `json:"share"`
	Repository NetworkRepositorySettingsModel `json:"repository"`
	MountServer MountServerSettingsImportSpec `json:"mountServer"`
}

// NewSmbStorageImportSpec instantiates a new SmbStorageImportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmbStorageImportSpec(name string, description string, tag string, type_ ERepositoryType, share SmbRepositoryShareSettingsSpec, repository NetworkRepositorySettingsModel, mountServer MountServerSettingsImportSpec) *SmbStorageImportSpec {
	this := SmbStorageImportSpec{}
	this.Name = name
	this.Description = description
	this.Tag = tag
	this.Type = type_
	this.Share = share
	this.Repository = repository
	this.MountServer = mountServer
	return &this
}

// NewSmbStorageImportSpecWithDefaults instantiates a new SmbStorageImportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmbStorageImportSpecWithDefaults() *SmbStorageImportSpec {
	this := SmbStorageImportSpec{}
	return &this
}

// GetName returns the Name field value
func (o *SmbStorageImportSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SmbStorageImportSpec) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SmbStorageImportSpec) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *SmbStorageImportSpec) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SmbStorageImportSpec) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SmbStorageImportSpec) SetDescription(v string) {
	o.Description = v
}

// GetTag returns the Tag field value
func (o *SmbStorageImportSpec) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *SmbStorageImportSpec) GetTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *SmbStorageImportSpec) SetTag(v string) {
	o.Tag = v
}

// GetType returns the Type field value
func (o *SmbStorageImportSpec) GetType() ERepositoryType {
	if o == nil {
		var ret ERepositoryType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SmbStorageImportSpec) GetTypeOk() (*ERepositoryType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SmbStorageImportSpec) SetType(v ERepositoryType) {
	o.Type = v
}

// GetShare returns the Share field value
func (o *SmbStorageImportSpec) GetShare() SmbRepositoryShareSettingsSpec {
	if o == nil {
		var ret SmbRepositoryShareSettingsSpec
		return ret
	}

	return o.Share
}

// GetShareOk returns a tuple with the Share field value
// and a boolean to check if the value has been set.
func (o *SmbStorageImportSpec) GetShareOk() (*SmbRepositoryShareSettingsSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Share, true
}

// SetShare sets field value
func (o *SmbStorageImportSpec) SetShare(v SmbRepositoryShareSettingsSpec) {
	o.Share = v
}

// GetRepository returns the Repository field value
func (o *SmbStorageImportSpec) GetRepository() NetworkRepositorySettingsModel {
	if o == nil {
		var ret NetworkRepositorySettingsModel
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *SmbStorageImportSpec) GetRepositoryOk() (*NetworkRepositorySettingsModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *SmbStorageImportSpec) SetRepository(v NetworkRepositorySettingsModel) {
	o.Repository = v
}

// GetMountServer returns the MountServer field value
func (o *SmbStorageImportSpec) GetMountServer() MountServerSettingsImportSpec {
	if o == nil {
		var ret MountServerSettingsImportSpec
		return ret
	}

	return o.MountServer
}

// GetMountServerOk returns a tuple with the MountServer field value
// and a boolean to check if the value has been set.
func (o *SmbStorageImportSpec) GetMountServerOk() (*MountServerSettingsImportSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MountServer, true
}

// SetMountServer sets field value
func (o *SmbStorageImportSpec) SetMountServer(v MountServerSettingsImportSpec) {
	o.MountServer = v
}

func (o SmbStorageImportSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["tag"] = o.Tag
	}
	if true {
		toSerialize["type"] = o.Type
	}
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

type NullableSmbStorageImportSpec struct {
	value *SmbStorageImportSpec
	isSet bool
}

func (v NullableSmbStorageImportSpec) Get() *SmbStorageImportSpec {
	return v.value
}

func (v *NullableSmbStorageImportSpec) Set(val *SmbStorageImportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSmbStorageImportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSmbStorageImportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmbStorageImportSpec(val *SmbStorageImportSpec) *NullableSmbStorageImportSpec {
	return &NullableSmbStorageImportSpec{value: val, isSet: true}
}

func (v NullableSmbStorageImportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmbStorageImportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


