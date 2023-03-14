/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br>Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Some request and response bodies refer to reusable schema objects that are defined in the `schemas` section of the REST API specification. Schemas, in turn, may inherit a part of their properties by referring to other schemas, and be polymorphic by referring to multiple alternate schemas.

API version: 1.1-rev0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// EncryptionPasswordModel struct for EncryptionPasswordModel
type EncryptionPasswordModel struct {
	// ID of the encryption password.
	Id string `json:"id"`
	// Hint for the encryption password.
	Hint string `json:"hint"`
	// Date and time when the password was last modified.
	ModificationTime *time.Time `json:"modificationTime,omitempty"`
	// Tag for the encryption password.
	Tag *string `json:"tag,omitempty"`
}

// NewEncryptionPasswordModel instantiates a new EncryptionPasswordModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptionPasswordModel(id string, hint string) *EncryptionPasswordModel {
	this := EncryptionPasswordModel{}
	this.Id = id
	this.Hint = hint
	return &this
}

// NewEncryptionPasswordModelWithDefaults instantiates a new EncryptionPasswordModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionPasswordModelWithDefaults() *EncryptionPasswordModel {
	this := EncryptionPasswordModel{}
	return &this
}

// GetId returns the Id field value
func (o *EncryptionPasswordModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EncryptionPasswordModel) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EncryptionPasswordModel) SetId(v string) {
	o.Id = v
}

// GetHint returns the Hint field value
func (o *EncryptionPasswordModel) GetHint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hint
}

// GetHintOk returns a tuple with the Hint field value
// and a boolean to check if the value has been set.
func (o *EncryptionPasswordModel) GetHintOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Hint, true
}

// SetHint sets field value
func (o *EncryptionPasswordModel) SetHint(v string) {
	o.Hint = v
}

// GetModificationTime returns the ModificationTime field value if set, zero value otherwise.
func (o *EncryptionPasswordModel) GetModificationTime() time.Time {
	if o == nil || isNil(o.ModificationTime) {
		var ret time.Time
		return ret
	}
	return *o.ModificationTime
}

// GetModificationTimeOk returns a tuple with the ModificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionPasswordModel) GetModificationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.ModificationTime) {
    return nil, false
	}
	return o.ModificationTime, true
}

// HasModificationTime returns a boolean if a field has been set.
func (o *EncryptionPasswordModel) HasModificationTime() bool {
	if o != nil && !isNil(o.ModificationTime) {
		return true
	}

	return false
}

// SetModificationTime gets a reference to the given time.Time and assigns it to the ModificationTime field.
func (o *EncryptionPasswordModel) SetModificationTime(v time.Time) {
	o.ModificationTime = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *EncryptionPasswordModel) GetTag() string {
	if o == nil || isNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionPasswordModel) GetTagOk() (*string, bool) {
	if o == nil || isNil(o.Tag) {
    return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *EncryptionPasswordModel) HasTag() bool {
	if o != nil && !isNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *EncryptionPasswordModel) SetTag(v string) {
	o.Tag = &v
}

func (o EncryptionPasswordModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["hint"] = o.Hint
	}
	if !isNil(o.ModificationTime) {
		toSerialize["modificationTime"] = o.ModificationTime
	}
	if !isNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableEncryptionPasswordModel struct {
	value *EncryptionPasswordModel
	isSet bool
}

func (v NullableEncryptionPasswordModel) Get() *EncryptionPasswordModel {
	return v.value
}

func (v *NullableEncryptionPasswordModel) Set(val *EncryptionPasswordModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptionPasswordModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptionPasswordModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptionPasswordModel(val *EncryptionPasswordModel) *NullableEncryptionPasswordModel {
	return &NullableEncryptionPasswordModel{value: val, isSet: true}
}

func (v NullableEncryptionPasswordModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptionPasswordModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


