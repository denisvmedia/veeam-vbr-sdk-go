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

// AzureBlobStorageImportSpec struct for AzureBlobStorageImportSpec
type AzureBlobStorageImportSpec struct {
	// Name of the object storage repository.
	Name string `json:"name"`
	// Description of the object storage repository.
	Description string `json:"description"`
	// Tag that identifies the object storage repository.
	Tag string `json:"tag"`
	Type ERepositoryType `json:"type"`
	// If *true*, the maximum number of concurrent tasks is limited.
	EnableTaskLimit *bool `json:"enableTaskLimit,omitempty"`
	// Maximum number of concurrent tasks.
	MaxTaskCount *int32 `json:"maxTaskCount,omitempty"`
	Account AzureBlobStorageAccountImportModel `json:"account"`
	Container AzureBlobStorageContainerModel `json:"container"`
	ProxyAppliance *AzureStorageProxyModel `json:"proxyAppliance,omitempty"`
	MountServer MountServerSettingsImportSpec `json:"mountServer"`
}

// NewAzureBlobStorageImportSpec instantiates a new AzureBlobStorageImportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureBlobStorageImportSpec(name string, description string, tag string, type_ ERepositoryType, account AzureBlobStorageAccountImportModel, container AzureBlobStorageContainerModel, mountServer MountServerSettingsImportSpec) *AzureBlobStorageImportSpec {
	this := AzureBlobStorageImportSpec{}
	this.Name = name
	this.Description = description
	this.Tag = tag
	this.Type = type_
	this.Account = account
	this.Container = container
	this.MountServer = mountServer
	return &this
}

// NewAzureBlobStorageImportSpecWithDefaults instantiates a new AzureBlobStorageImportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureBlobStorageImportSpecWithDefaults() *AzureBlobStorageImportSpec {
	this := AzureBlobStorageImportSpec{}
	return &this
}

// GetName returns the Name field value
func (o *AzureBlobStorageImportSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageImportSpec) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AzureBlobStorageImportSpec) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *AzureBlobStorageImportSpec) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageImportSpec) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AzureBlobStorageImportSpec) SetDescription(v string) {
	o.Description = v
}

// GetTag returns the Tag field value
func (o *AzureBlobStorageImportSpec) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageImportSpec) GetTagOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *AzureBlobStorageImportSpec) SetTag(v string) {
	o.Tag = v
}

// GetType returns the Type field value
func (o *AzureBlobStorageImportSpec) GetType() ERepositoryType {
	if o == nil {
		var ret ERepositoryType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageImportSpec) GetTypeOk() (*ERepositoryType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AzureBlobStorageImportSpec) SetType(v ERepositoryType) {
	o.Type = v
}

// GetEnableTaskLimit returns the EnableTaskLimit field value if set, zero value otherwise.
func (o *AzureBlobStorageImportSpec) GetEnableTaskLimit() bool {
	if o == nil || isNil(o.EnableTaskLimit) {
		var ret bool
		return ret
	}
	return *o.EnableTaskLimit
}

// GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageImportSpec) GetEnableTaskLimitOk() (*bool, bool) {
	if o == nil || isNil(o.EnableTaskLimit) {
    return nil, false
	}
	return o.EnableTaskLimit, true
}

// HasEnableTaskLimit returns a boolean if a field has been set.
func (o *AzureBlobStorageImportSpec) HasEnableTaskLimit() bool {
	if o != nil && !isNil(o.EnableTaskLimit) {
		return true
	}

	return false
}

// SetEnableTaskLimit gets a reference to the given bool and assigns it to the EnableTaskLimit field.
func (o *AzureBlobStorageImportSpec) SetEnableTaskLimit(v bool) {
	o.EnableTaskLimit = &v
}

// GetMaxTaskCount returns the MaxTaskCount field value if set, zero value otherwise.
func (o *AzureBlobStorageImportSpec) GetMaxTaskCount() int32 {
	if o == nil || isNil(o.MaxTaskCount) {
		var ret int32
		return ret
	}
	return *o.MaxTaskCount
}

// GetMaxTaskCountOk returns a tuple with the MaxTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageImportSpec) GetMaxTaskCountOk() (*int32, bool) {
	if o == nil || isNil(o.MaxTaskCount) {
    return nil, false
	}
	return o.MaxTaskCount, true
}

// HasMaxTaskCount returns a boolean if a field has been set.
func (o *AzureBlobStorageImportSpec) HasMaxTaskCount() bool {
	if o != nil && !isNil(o.MaxTaskCount) {
		return true
	}

	return false
}

// SetMaxTaskCount gets a reference to the given int32 and assigns it to the MaxTaskCount field.
func (o *AzureBlobStorageImportSpec) SetMaxTaskCount(v int32) {
	o.MaxTaskCount = &v
}

// GetAccount returns the Account field value
func (o *AzureBlobStorageImportSpec) GetAccount() AzureBlobStorageAccountImportModel {
	if o == nil {
		var ret AzureBlobStorageAccountImportModel
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageImportSpec) GetAccountOk() (*AzureBlobStorageAccountImportModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *AzureBlobStorageImportSpec) SetAccount(v AzureBlobStorageAccountImportModel) {
	o.Account = v
}

// GetContainer returns the Container field value
func (o *AzureBlobStorageImportSpec) GetContainer() AzureBlobStorageContainerModel {
	if o == nil {
		var ret AzureBlobStorageContainerModel
		return ret
	}

	return o.Container
}

// GetContainerOk returns a tuple with the Container field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageImportSpec) GetContainerOk() (*AzureBlobStorageContainerModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Container, true
}

// SetContainer sets field value
func (o *AzureBlobStorageImportSpec) SetContainer(v AzureBlobStorageContainerModel) {
	o.Container = v
}

// GetProxyAppliance returns the ProxyAppliance field value if set, zero value otherwise.
func (o *AzureBlobStorageImportSpec) GetProxyAppliance() AzureStorageProxyModel {
	if o == nil || isNil(o.ProxyAppliance) {
		var ret AzureStorageProxyModel
		return ret
	}
	return *o.ProxyAppliance
}

// GetProxyApplianceOk returns a tuple with the ProxyAppliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageImportSpec) GetProxyApplianceOk() (*AzureStorageProxyModel, bool) {
	if o == nil || isNil(o.ProxyAppliance) {
    return nil, false
	}
	return o.ProxyAppliance, true
}

// HasProxyAppliance returns a boolean if a field has been set.
func (o *AzureBlobStorageImportSpec) HasProxyAppliance() bool {
	if o != nil && !isNil(o.ProxyAppliance) {
		return true
	}

	return false
}

// SetProxyAppliance gets a reference to the given AzureStorageProxyModel and assigns it to the ProxyAppliance field.
func (o *AzureBlobStorageImportSpec) SetProxyAppliance(v AzureStorageProxyModel) {
	o.ProxyAppliance = &v
}

// GetMountServer returns the MountServer field value
func (o *AzureBlobStorageImportSpec) GetMountServer() MountServerSettingsImportSpec {
	if o == nil {
		var ret MountServerSettingsImportSpec
		return ret
	}

	return o.MountServer
}

// GetMountServerOk returns a tuple with the MountServer field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageImportSpec) GetMountServerOk() (*MountServerSettingsImportSpec, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MountServer, true
}

// SetMountServer sets field value
func (o *AzureBlobStorageImportSpec) SetMountServer(v MountServerSettingsImportSpec) {
	o.MountServer = v
}

func (o AzureBlobStorageImportSpec) MarshalJSON() ([]byte, error) {
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
		toSerialize["container"] = o.Container
	}
	if !isNil(o.ProxyAppliance) {
		toSerialize["proxyAppliance"] = o.ProxyAppliance
	}
	if true {
		toSerialize["mountServer"] = o.MountServer
	}
	return json.Marshal(toSerialize)
}

type NullableAzureBlobStorageImportSpec struct {
	value *AzureBlobStorageImportSpec
	isSet bool
}

func (v NullableAzureBlobStorageImportSpec) Get() *AzureBlobStorageImportSpec {
	return v.value
}

func (v *NullableAzureBlobStorageImportSpec) Set(val *AzureBlobStorageImportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureBlobStorageImportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureBlobStorageImportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureBlobStorageImportSpec(val *AzureBlobStorageImportSpec) *NullableAzureBlobStorageImportSpec {
	return &NullableAzureBlobStorageImportSpec{value: val, isSet: true}
}

func (v NullableAzureBlobStorageImportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureBlobStorageImportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


