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

// HostConnectionSpec struct for HostConnectionSpec
type HostConnectionSpec struct {
	// Full DNS name or IP address of the server.
	ServerName string `json:"serverName"`
	// ID of a credentials record used to connect to the server.
	CredentialsId string `json:"credentialsId"`
	Type EManagedServerType `json:"type"`
	// Port used to communicate with the server.
	Port *int32 `json:"port,omitempty"`
}

// NewHostConnectionSpec instantiates a new HostConnectionSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostConnectionSpec(serverName string, credentialsId string, type_ EManagedServerType) *HostConnectionSpec {
	this := HostConnectionSpec{}
	this.ServerName = serverName
	this.CredentialsId = credentialsId
	this.Type = type_
	return &this
}

// NewHostConnectionSpecWithDefaults instantiates a new HostConnectionSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostConnectionSpecWithDefaults() *HostConnectionSpec {
	this := HostConnectionSpec{}
	return &this
}

// GetServerName returns the ServerName field value
func (o *HostConnectionSpec) GetServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value
// and a boolean to check if the value has been set.
func (o *HostConnectionSpec) GetServerNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServerName, true
}

// SetServerName sets field value
func (o *HostConnectionSpec) SetServerName(v string) {
	o.ServerName = v
}

// GetCredentialsId returns the CredentialsId field value
func (o *HostConnectionSpec) GetCredentialsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CredentialsId
}

// GetCredentialsIdOk returns a tuple with the CredentialsId field value
// and a boolean to check if the value has been set.
func (o *HostConnectionSpec) GetCredentialsIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CredentialsId, true
}

// SetCredentialsId sets field value
func (o *HostConnectionSpec) SetCredentialsId(v string) {
	o.CredentialsId = v
}

// GetType returns the Type field value
func (o *HostConnectionSpec) GetType() EManagedServerType {
	if o == nil {
		var ret EManagedServerType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HostConnectionSpec) GetTypeOk() (*EManagedServerType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *HostConnectionSpec) SetType(v EManagedServerType) {
	o.Type = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *HostConnectionSpec) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostConnectionSpec) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *HostConnectionSpec) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *HostConnectionSpec) SetPort(v int32) {
	o.Port = &v
}

func (o HostConnectionSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serverName"] = o.ServerName
	}
	if true {
		toSerialize["credentialsId"] = o.CredentialsId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableHostConnectionSpec struct {
	value *HostConnectionSpec
	isSet bool
}

func (v NullableHostConnectionSpec) Get() *HostConnectionSpec {
	return v.value
}

func (v *NullableHostConnectionSpec) Set(val *HostConnectionSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableHostConnectionSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableHostConnectionSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostConnectionSpec(val *HostConnectionSpec) *NullableHostConnectionSpec {
	return &NullableHostConnectionSpec{value: val, isSet: true}
}

func (v NullableHostConnectionSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostConnectionSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


