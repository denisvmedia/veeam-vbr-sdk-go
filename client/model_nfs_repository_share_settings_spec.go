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

// NfsRepositoryShareSettingsSpec struct for NfsRepositoryShareSettingsSpec
type NfsRepositoryShareSettingsSpec struct {
	// Path to the shared folder that that is used as a backup repository.
	SharePath string `json:"sharePath"`
	GatewayServer *RepositoryShareGatewayImportSpec `json:"gatewayServer,omitempty"`
}

// NewNfsRepositoryShareSettingsSpec instantiates a new NfsRepositoryShareSettingsSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfsRepositoryShareSettingsSpec(sharePath string) *NfsRepositoryShareSettingsSpec {
	this := NfsRepositoryShareSettingsSpec{}
	this.SharePath = sharePath
	return &this
}

// NewNfsRepositoryShareSettingsSpecWithDefaults instantiates a new NfsRepositoryShareSettingsSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfsRepositoryShareSettingsSpecWithDefaults() *NfsRepositoryShareSettingsSpec {
	this := NfsRepositoryShareSettingsSpec{}
	return &this
}

// GetSharePath returns the SharePath field value
func (o *NfsRepositoryShareSettingsSpec) GetSharePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SharePath
}

// GetSharePathOk returns a tuple with the SharePath field value
// and a boolean to check if the value has been set.
func (o *NfsRepositoryShareSettingsSpec) GetSharePathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SharePath, true
}

// SetSharePath sets field value
func (o *NfsRepositoryShareSettingsSpec) SetSharePath(v string) {
	o.SharePath = v
}

// GetGatewayServer returns the GatewayServer field value if set, zero value otherwise.
func (o *NfsRepositoryShareSettingsSpec) GetGatewayServer() RepositoryShareGatewayImportSpec {
	if o == nil || o.GatewayServer == nil {
		var ret RepositoryShareGatewayImportSpec
		return ret
	}
	return *o.GatewayServer
}

// GetGatewayServerOk returns a tuple with the GatewayServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfsRepositoryShareSettingsSpec) GetGatewayServerOk() (*RepositoryShareGatewayImportSpec, bool) {
	if o == nil || o.GatewayServer == nil {
		return nil, false
	}
	return o.GatewayServer, true
}

// HasGatewayServer returns a boolean if a field has been set.
func (o *NfsRepositoryShareSettingsSpec) HasGatewayServer() bool {
	if o != nil && o.GatewayServer != nil {
		return true
	}

	return false
}

// SetGatewayServer gets a reference to the given RepositoryShareGatewayImportSpec and assigns it to the GatewayServer field.
func (o *NfsRepositoryShareSettingsSpec) SetGatewayServer(v RepositoryShareGatewayImportSpec) {
	o.GatewayServer = &v
}

func (o NfsRepositoryShareSettingsSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sharePath"] = o.SharePath
	}
	if o.GatewayServer != nil {
		toSerialize["gatewayServer"] = o.GatewayServer
	}
	return json.Marshal(toSerialize)
}

type NullableNfsRepositoryShareSettingsSpec struct {
	value *NfsRepositoryShareSettingsSpec
	isSet bool
}

func (v NullableNfsRepositoryShareSettingsSpec) Get() *NfsRepositoryShareSettingsSpec {
	return v.value
}

func (v *NullableNfsRepositoryShareSettingsSpec) Set(val *NfsRepositoryShareSettingsSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableNfsRepositoryShareSettingsSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableNfsRepositoryShareSettingsSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfsRepositoryShareSettingsSpec(val *NfsRepositoryShareSettingsSpec) *NullableNfsRepositoryShareSettingsSpec {
	return &NullableNfsRepositoryShareSettingsSpec{value: val, isSet: true}
}

func (v NullableNfsRepositoryShareSettingsSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfsRepositoryShareSettingsSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


