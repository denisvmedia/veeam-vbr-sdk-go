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

// InstantViVMRecoveryMount VM mount.
type InstantViVMRecoveryMount struct {
	// Mount ID.
	Id string `json:"id"`
	// ID of the restore session. Use the ID to track the progress. For details, see [Get Session](#tag/Sessions/operation/GetSession).
	SessionId string `json:"sessionId"`
	State EInstantRecoveryMountState `json:"state"`
	Spec InstantViVMRecoverySpec `json:"spec"`
	// Name of the recovered VM.
	VmName string `json:"vmName"`
	// Error message.
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// NewInstantViVMRecoveryMount instantiates a new InstantViVMRecoveryMount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstantViVMRecoveryMount(id string, sessionId string, state EInstantRecoveryMountState, spec InstantViVMRecoverySpec, vmName string) *InstantViVMRecoveryMount {
	this := InstantViVMRecoveryMount{}
	this.Id = id
	this.SessionId = sessionId
	this.State = state
	this.Spec = spec
	this.VmName = vmName
	return &this
}

// NewInstantViVMRecoveryMountWithDefaults instantiates a new InstantViVMRecoveryMount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstantViVMRecoveryMountWithDefaults() *InstantViVMRecoveryMount {
	this := InstantViVMRecoveryMount{}
	return &this
}

// GetId returns the Id field value
func (o *InstantViVMRecoveryMount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InstantViVMRecoveryMount) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InstantViVMRecoveryMount) SetId(v string) {
	o.Id = v
}

// GetSessionId returns the SessionId field value
func (o *InstantViVMRecoveryMount) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *InstantViVMRecoveryMount) GetSessionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *InstantViVMRecoveryMount) SetSessionId(v string) {
	o.SessionId = v
}

// GetState returns the State field value
func (o *InstantViVMRecoveryMount) GetState() EInstantRecoveryMountState {
	if o == nil {
		var ret EInstantRecoveryMountState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *InstantViVMRecoveryMount) GetStateOk() (*EInstantRecoveryMountState, bool) {
	if o == nil {
    return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *InstantViVMRecoveryMount) SetState(v EInstantRecoveryMountState) {
	o.State = v
}

// GetSpec returns the Spec field value
func (o *InstantViVMRecoveryMount) GetSpec() InstantViVMRecoverySpec {
	if o == nil {
		var ret InstantViVMRecoverySpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *InstantViVMRecoveryMount) GetSpecOk() (*InstantViVMRecoverySpec, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *InstantViVMRecoveryMount) SetSpec(v InstantViVMRecoverySpec) {
	o.Spec = v
}

// GetVmName returns the VmName field value
func (o *InstantViVMRecoveryMount) GetVmName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VmName
}

// GetVmNameOk returns a tuple with the VmName field value
// and a boolean to check if the value has been set.
func (o *InstantViVMRecoveryMount) GetVmNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VmName, true
}

// SetVmName sets field value
func (o *InstantViVMRecoveryMount) SetVmName(v string) {
	o.VmName = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *InstantViVMRecoveryMount) GetErrorMessage() string {
	if o == nil || isNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstantViVMRecoveryMount) GetErrorMessageOk() (*string, bool) {
	if o == nil || isNil(o.ErrorMessage) {
    return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *InstantViVMRecoveryMount) HasErrorMessage() bool {
	if o != nil && !isNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *InstantViVMRecoveryMount) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o InstantViVMRecoveryMount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["sessionId"] = o.SessionId
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["spec"] = o.Spec
	}
	if true {
		toSerialize["vmName"] = o.VmName
	}
	if !isNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	return json.Marshal(toSerialize)
}

type NullableInstantViVMRecoveryMount struct {
	value *InstantViVMRecoveryMount
	isSet bool
}

func (v NullableInstantViVMRecoveryMount) Get() *InstantViVMRecoveryMount {
	return v.value
}

func (v *NullableInstantViVMRecoveryMount) Set(val *InstantViVMRecoveryMount) {
	v.value = val
	v.isSet = true
}

func (v NullableInstantViVMRecoveryMount) IsSet() bool {
	return v.isSet
}

func (v *NullableInstantViVMRecoveryMount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstantViVMRecoveryMount(val *InstantViVMRecoveryMount) *NullableInstantViVMRecoveryMount {
	return &NullableInstantViVMRecoveryMount{value: val, isSet: true}
}

func (v NullableInstantViVMRecoveryMount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstantViVMRecoveryMount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


