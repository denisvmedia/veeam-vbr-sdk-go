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

// StandardCredentialsSpecAllOf struct for StandardCredentialsSpecAllOf
type StandardCredentialsSpecAllOf struct {
	// Tag used to identify the credentials record.
	Tag *string `json:"tag,omitempty"`
}

// NewStandardCredentialsSpecAllOf instantiates a new StandardCredentialsSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardCredentialsSpecAllOf() *StandardCredentialsSpecAllOf {
	this := StandardCredentialsSpecAllOf{}
	return &this
}

// NewStandardCredentialsSpecAllOfWithDefaults instantiates a new StandardCredentialsSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardCredentialsSpecAllOfWithDefaults() *StandardCredentialsSpecAllOf {
	this := StandardCredentialsSpecAllOf{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *StandardCredentialsSpecAllOf) GetTag() string {
	if o == nil || isNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardCredentialsSpecAllOf) GetTagOk() (*string, bool) {
	if o == nil || isNil(o.Tag) {
    return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *StandardCredentialsSpecAllOf) HasTag() bool {
	if o != nil && !isNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *StandardCredentialsSpecAllOf) SetTag(v string) {
	o.Tag = &v
}

func (o StandardCredentialsSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableStandardCredentialsSpecAllOf struct {
	value *StandardCredentialsSpecAllOf
	isSet bool
}

func (v NullableStandardCredentialsSpecAllOf) Get() *StandardCredentialsSpecAllOf {
	return v.value
}

func (v *NullableStandardCredentialsSpecAllOf) Set(val *StandardCredentialsSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardCredentialsSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardCredentialsSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardCredentialsSpecAllOf(val *StandardCredentialsSpecAllOf) *NullableStandardCredentialsSpecAllOf {
	return &NullableStandardCredentialsSpecAllOf{value: val, isSet: true}
}

func (v NullableStandardCredentialsSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardCredentialsSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


