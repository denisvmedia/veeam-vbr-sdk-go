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

// JobStartSpec struct for JobStartSpec
type JobStartSpec struct {
	PerformActiveFull bool `json:"performActiveFull"`
	StartChainedJobs *bool `json:"startChainedJobs,omitempty"`
}

// NewJobStartSpec instantiates a new JobStartSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobStartSpec(performActiveFull bool) *JobStartSpec {
	this := JobStartSpec{}
	this.PerformActiveFull = performActiveFull
	var startChainedJobs bool = false
	this.StartChainedJobs = &startChainedJobs
	return &this
}

// NewJobStartSpecWithDefaults instantiates a new JobStartSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobStartSpecWithDefaults() *JobStartSpec {
	this := JobStartSpec{}
	var performActiveFull bool = false
	this.PerformActiveFull = performActiveFull
	var startChainedJobs bool = false
	this.StartChainedJobs = &startChainedJobs
	return &this
}

// GetPerformActiveFull returns the PerformActiveFull field value
func (o *JobStartSpec) GetPerformActiveFull() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PerformActiveFull
}

// GetPerformActiveFullOk returns a tuple with the PerformActiveFull field value
// and a boolean to check if the value has been set.
func (o *JobStartSpec) GetPerformActiveFullOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PerformActiveFull, true
}

// SetPerformActiveFull sets field value
func (o *JobStartSpec) SetPerformActiveFull(v bool) {
	o.PerformActiveFull = v
}

// GetStartChainedJobs returns the StartChainedJobs field value if set, zero value otherwise.
func (o *JobStartSpec) GetStartChainedJobs() bool {
	if o == nil || o.StartChainedJobs == nil {
		var ret bool
		return ret
	}
	return *o.StartChainedJobs
}

// GetStartChainedJobsOk returns a tuple with the StartChainedJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStartSpec) GetStartChainedJobsOk() (*bool, bool) {
	if o == nil || o.StartChainedJobs == nil {
		return nil, false
	}
	return o.StartChainedJobs, true
}

// HasStartChainedJobs returns a boolean if a field has been set.
func (o *JobStartSpec) HasStartChainedJobs() bool {
	if o != nil && o.StartChainedJobs != nil {
		return true
	}

	return false
}

// SetStartChainedJobs gets a reference to the given bool and assigns it to the StartChainedJobs field.
func (o *JobStartSpec) SetStartChainedJobs(v bool) {
	o.StartChainedJobs = &v
}

func (o JobStartSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["performActiveFull"] = o.PerformActiveFull
	}
	if o.StartChainedJobs != nil {
		toSerialize["startChainedJobs"] = o.StartChainedJobs
	}
	return json.Marshal(toSerialize)
}

type NullableJobStartSpec struct {
	value *JobStartSpec
	isSet bool
}

func (v NullableJobStartSpec) Get() *JobStartSpec {
	return v.value
}

func (v *NullableJobStartSpec) Set(val *JobStartSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableJobStartSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableJobStartSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobStartSpec(val *JobStartSpec) *NullableJobStartSpec {
	return &NullableJobStartSpec{value: val, isSet: true}
}

func (v NullableJobStartSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobStartSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


