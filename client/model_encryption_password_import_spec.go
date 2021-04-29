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

// EncryptionPasswordImportSpec struct for EncryptionPasswordImportSpec
type EncryptionPasswordImportSpec struct {
	// Password.
	Password string `json:"password"`
	// Hint for the encryption password.
	Hint string `json:"hint"`
	// Tag for the encryption password.
	Tag *string `json:"tag,omitempty"`
}

// NewEncryptionPasswordImportSpec instantiates a new EncryptionPasswordImportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptionPasswordImportSpec(password string, hint string) *EncryptionPasswordImportSpec {
	this := EncryptionPasswordImportSpec{}
	this.Password = password
	this.Hint = hint
	return &this
}

// NewEncryptionPasswordImportSpecWithDefaults instantiates a new EncryptionPasswordImportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionPasswordImportSpecWithDefaults() *EncryptionPasswordImportSpec {
	this := EncryptionPasswordImportSpec{}
	return &this
}

// GetPassword returns the Password field value
func (o *EncryptionPasswordImportSpec) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *EncryptionPasswordImportSpec) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *EncryptionPasswordImportSpec) SetPassword(v string) {
	o.Password = v
}

// GetHint returns the Hint field value
func (o *EncryptionPasswordImportSpec) GetHint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hint
}

// GetHintOk returns a tuple with the Hint field value
// and a boolean to check if the value has been set.
func (o *EncryptionPasswordImportSpec) GetHintOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hint, true
}

// SetHint sets field value
func (o *EncryptionPasswordImportSpec) SetHint(v string) {
	o.Hint = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *EncryptionPasswordImportSpec) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionPasswordImportSpec) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *EncryptionPasswordImportSpec) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *EncryptionPasswordImportSpec) SetTag(v string) {
	o.Tag = &v
}

func (o EncryptionPasswordImportSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["hint"] = o.Hint
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableEncryptionPasswordImportSpec struct {
	value *EncryptionPasswordImportSpec
	isSet bool
}

func (v NullableEncryptionPasswordImportSpec) Get() *EncryptionPasswordImportSpec {
	return v.value
}

func (v *NullableEncryptionPasswordImportSpec) Set(val *EncryptionPasswordImportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptionPasswordImportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptionPasswordImportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptionPasswordImportSpec(val *EncryptionPasswordImportSpec) *NullableEncryptionPasswordImportSpec {
	return &NullableEncryptionPasswordImportSpec{value: val, isSet: true}
}

func (v NullableEncryptionPasswordImportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptionPasswordImportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


