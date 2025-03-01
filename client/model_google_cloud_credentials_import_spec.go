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

// GoogleCloudCredentialsImportSpec struct for GoogleCloudCredentialsImportSpec
type GoogleCloudCredentialsImportSpec struct {
	Type ECloudCredentialsType `json:"type"`
	// Description of the cloud credentials record.
	Description *string `json:"description,omitempty"`
	// Tag used to identify the cloud credentials record.
	Tag string `json:"tag"`
	// Access ID of a Google HMAC key.
	AccessKey string `json:"accessKey"`
	// Secret linked to the access ID.
	SecretKey string `json:"secretKey"`
}

// NewGoogleCloudCredentialsImportSpec instantiates a new GoogleCloudCredentialsImportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleCloudCredentialsImportSpec(type_ ECloudCredentialsType, tag string, accessKey string, secretKey string) *GoogleCloudCredentialsImportSpec {
	this := GoogleCloudCredentialsImportSpec{}
	this.Type = type_
	this.Tag = tag
	this.AccessKey = accessKey
	this.SecretKey = secretKey
	return &this
}

// NewGoogleCloudCredentialsImportSpecWithDefaults instantiates a new GoogleCloudCredentialsImportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleCloudCredentialsImportSpecWithDefaults() *GoogleCloudCredentialsImportSpec {
	this := GoogleCloudCredentialsImportSpec{}
	return &this
}

// GetType returns the Type field value
func (o *GoogleCloudCredentialsImportSpec) GetType() ECloudCredentialsType {
	if o == nil {
		var ret ECloudCredentialsType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GoogleCloudCredentialsImportSpec) GetTypeOk() (*ECloudCredentialsType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GoogleCloudCredentialsImportSpec) SetType(v ECloudCredentialsType) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GoogleCloudCredentialsImportSpec) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleCloudCredentialsImportSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GoogleCloudCredentialsImportSpec) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GoogleCloudCredentialsImportSpec) SetDescription(v string) {
	o.Description = &v
}

// GetTag returns the Tag field value
func (o *GoogleCloudCredentialsImportSpec) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *GoogleCloudCredentialsImportSpec) GetTagOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *GoogleCloudCredentialsImportSpec) SetTag(v string) {
	o.Tag = v
}

// GetAccessKey returns the AccessKey field value
func (o *GoogleCloudCredentialsImportSpec) GetAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value
// and a boolean to check if the value has been set.
func (o *GoogleCloudCredentialsImportSpec) GetAccessKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AccessKey, true
}

// SetAccessKey sets field value
func (o *GoogleCloudCredentialsImportSpec) SetAccessKey(v string) {
	o.AccessKey = v
}

// GetSecretKey returns the SecretKey field value
func (o *GoogleCloudCredentialsImportSpec) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
func (o *GoogleCloudCredentialsImportSpec) GetSecretKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SecretKey, true
}

// SetSecretKey sets field value
func (o *GoogleCloudCredentialsImportSpec) SetSecretKey(v string) {
	o.SecretKey = v
}

func (o GoogleCloudCredentialsImportSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["tag"] = o.Tag
	}
	if true {
		toSerialize["accessKey"] = o.AccessKey
	}
	if true {
		toSerialize["secretKey"] = o.SecretKey
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleCloudCredentialsImportSpec struct {
	value *GoogleCloudCredentialsImportSpec
	isSet bool
}

func (v NullableGoogleCloudCredentialsImportSpec) Get() *GoogleCloudCredentialsImportSpec {
	return v.value
}

func (v *NullableGoogleCloudCredentialsImportSpec) Set(val *GoogleCloudCredentialsImportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleCloudCredentialsImportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleCloudCredentialsImportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleCloudCredentialsImportSpec(val *GoogleCloudCredentialsImportSpec) *NullableGoogleCloudCredentialsImportSpec {
	return &NullableGoogleCloudCredentialsImportSpec{value: val, isSet: true}
}

func (v NullableGoogleCloudCredentialsImportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleCloudCredentialsImportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


