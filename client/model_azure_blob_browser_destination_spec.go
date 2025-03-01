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

// AzureBlobBrowserDestinationSpec struct for AzureBlobBrowserDestinationSpec
type AzureBlobBrowserDestinationSpec struct {
	CloudBrowserNewFolderSpec
	// ID of a server you want to use to connect to the object storage. You can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used.
	HostId *string `json:"hostId,omitempty"`
	RegionType EAzureRegionType `json:"regionType"`
	// Name of the container where you want to store your backup data.
	ContainerName string `json:"containerName"`
	FolderType *ECloudBrowserFolderType `json:"folderType,omitempty"`
}

// NewAzureBlobBrowserDestinationSpec instantiates a new AzureBlobBrowserDestinationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureBlobBrowserDestinationSpec(regionType EAzureRegionType, containerName string, credentialsId string, serviceType ECloudServiceType, newFolderName string) *AzureBlobBrowserDestinationSpec {
	this := AzureBlobBrowserDestinationSpec{}
	this.CredentialsId = credentialsId
	this.ServiceType = serviceType
	this.NewFolderName = newFolderName
	this.RegionType = regionType
	this.ContainerName = containerName
	return &this
}

// NewAzureBlobBrowserDestinationSpecWithDefaults instantiates a new AzureBlobBrowserDestinationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureBlobBrowserDestinationSpecWithDefaults() *AzureBlobBrowserDestinationSpec {
	this := AzureBlobBrowserDestinationSpec{}
	return &this
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *AzureBlobBrowserDestinationSpec) GetHostId() string {
	if o == nil || isNil(o.HostId) {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobBrowserDestinationSpec) GetHostIdOk() (*string, bool) {
	if o == nil || isNil(o.HostId) {
    return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *AzureBlobBrowserDestinationSpec) HasHostId() bool {
	if o != nil && !isNil(o.HostId) {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *AzureBlobBrowserDestinationSpec) SetHostId(v string) {
	o.HostId = &v
}

// GetRegionType returns the RegionType field value
func (o *AzureBlobBrowserDestinationSpec) GetRegionType() EAzureRegionType {
	if o == nil {
		var ret EAzureRegionType
		return ret
	}

	return o.RegionType
}

// GetRegionTypeOk returns a tuple with the RegionType field value
// and a boolean to check if the value has been set.
func (o *AzureBlobBrowserDestinationSpec) GetRegionTypeOk() (*EAzureRegionType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RegionType, true
}

// SetRegionType sets field value
func (o *AzureBlobBrowserDestinationSpec) SetRegionType(v EAzureRegionType) {
	o.RegionType = v
}

// GetContainerName returns the ContainerName field value
func (o *AzureBlobBrowserDestinationSpec) GetContainerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value
// and a boolean to check if the value has been set.
func (o *AzureBlobBrowserDestinationSpec) GetContainerNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ContainerName, true
}

// SetContainerName sets field value
func (o *AzureBlobBrowserDestinationSpec) SetContainerName(v string) {
	o.ContainerName = v
}

// GetFolderType returns the FolderType field value if set, zero value otherwise.
func (o *AzureBlobBrowserDestinationSpec) GetFolderType() ECloudBrowserFolderType {
	if o == nil || isNil(o.FolderType) {
		var ret ECloudBrowserFolderType
		return ret
	}
	return *o.FolderType
}

// GetFolderTypeOk returns a tuple with the FolderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobBrowserDestinationSpec) GetFolderTypeOk() (*ECloudBrowserFolderType, bool) {
	if o == nil || isNil(o.FolderType) {
    return nil, false
	}
	return o.FolderType, true
}

// HasFolderType returns a boolean if a field has been set.
func (o *AzureBlobBrowserDestinationSpec) HasFolderType() bool {
	if o != nil && !isNil(o.FolderType) {
		return true
	}

	return false
}

// SetFolderType gets a reference to the given ECloudBrowserFolderType and assigns it to the FolderType field.
func (o *AzureBlobBrowserDestinationSpec) SetFolderType(v ECloudBrowserFolderType) {
	o.FolderType = &v
}

func (o AzureBlobBrowserDestinationSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBrowserNewFolderSpec, errCloudBrowserNewFolderSpec := json.Marshal(o.CloudBrowserNewFolderSpec)
	if errCloudBrowserNewFolderSpec != nil {
		return []byte{}, errCloudBrowserNewFolderSpec
	}
	errCloudBrowserNewFolderSpec = json.Unmarshal([]byte(serializedCloudBrowserNewFolderSpec), &toSerialize)
	if errCloudBrowserNewFolderSpec != nil {
		return []byte{}, errCloudBrowserNewFolderSpec
	}
	if !isNil(o.HostId) {
		toSerialize["hostId"] = o.HostId
	}
	if true {
		toSerialize["regionType"] = o.RegionType
	}
	if true {
		toSerialize["containerName"] = o.ContainerName
	}
	if !isNil(o.FolderType) {
		toSerialize["folderType"] = o.FolderType
	}
	return json.Marshal(toSerialize)
}

type NullableAzureBlobBrowserDestinationSpec struct {
	value *AzureBlobBrowserDestinationSpec
	isSet bool
}

func (v NullableAzureBlobBrowserDestinationSpec) Get() *AzureBlobBrowserDestinationSpec {
	return v.value
}

func (v *NullableAzureBlobBrowserDestinationSpec) Set(val *AzureBlobBrowserDestinationSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureBlobBrowserDestinationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureBlobBrowserDestinationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureBlobBrowserDestinationSpec(val *AzureBlobBrowserDestinationSpec) *NullableAzureBlobBrowserDestinationSpec {
	return &NullableAzureBlobBrowserDestinationSpec{value: val, isSet: true}
}

func (v NullableAzureBlobBrowserDestinationSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureBlobBrowserDestinationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


