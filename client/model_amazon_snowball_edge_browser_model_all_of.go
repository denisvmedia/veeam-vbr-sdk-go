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

// AmazonSnowballEdgeBrowserModelAllOf struct for AmazonSnowballEdgeBrowserModelAllOf
type AmazonSnowballEdgeBrowserModelAllOf struct {
	// ID of a server used to connect to the AWS Snowball Edge device.
	HostId *string `json:"hostId,omitempty"`
	// Service point address and port number of the AWS Snowball Edge device.
	ConnectionPoint *string `json:"connectionPoint,omitempty"`
	// Array of regions.
	Regions []AmazonSnowballEdgeRegionBrowserModel `json:"regions,omitempty"`
}

// NewAmazonSnowballEdgeBrowserModelAllOf instantiates a new AmazonSnowballEdgeBrowserModelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonSnowballEdgeBrowserModelAllOf() *AmazonSnowballEdgeBrowserModelAllOf {
	this := AmazonSnowballEdgeBrowserModelAllOf{}
	return &this
}

// NewAmazonSnowballEdgeBrowserModelAllOfWithDefaults instantiates a new AmazonSnowballEdgeBrowserModelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonSnowballEdgeBrowserModelAllOfWithDefaults() *AmazonSnowballEdgeBrowserModelAllOf {
	this := AmazonSnowballEdgeBrowserModelAllOf{}
	return &this
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *AmazonSnowballEdgeBrowserModelAllOf) GetHostId() string {
	if o == nil || isNil(o.HostId) {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeBrowserModelAllOf) GetHostIdOk() (*string, bool) {
	if o == nil || isNil(o.HostId) {
    return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *AmazonSnowballEdgeBrowserModelAllOf) HasHostId() bool {
	if o != nil && !isNil(o.HostId) {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *AmazonSnowballEdgeBrowserModelAllOf) SetHostId(v string) {
	o.HostId = &v
}

// GetConnectionPoint returns the ConnectionPoint field value if set, zero value otherwise.
func (o *AmazonSnowballEdgeBrowserModelAllOf) GetConnectionPoint() string {
	if o == nil || isNil(o.ConnectionPoint) {
		var ret string
		return ret
	}
	return *o.ConnectionPoint
}

// GetConnectionPointOk returns a tuple with the ConnectionPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeBrowserModelAllOf) GetConnectionPointOk() (*string, bool) {
	if o == nil || isNil(o.ConnectionPoint) {
    return nil, false
	}
	return o.ConnectionPoint, true
}

// HasConnectionPoint returns a boolean if a field has been set.
func (o *AmazonSnowballEdgeBrowserModelAllOf) HasConnectionPoint() bool {
	if o != nil && !isNil(o.ConnectionPoint) {
		return true
	}

	return false
}

// SetConnectionPoint gets a reference to the given string and assigns it to the ConnectionPoint field.
func (o *AmazonSnowballEdgeBrowserModelAllOf) SetConnectionPoint(v string) {
	o.ConnectionPoint = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *AmazonSnowballEdgeBrowserModelAllOf) GetRegions() []AmazonSnowballEdgeRegionBrowserModel {
	if o == nil || isNil(o.Regions) {
		var ret []AmazonSnowballEdgeRegionBrowserModel
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonSnowballEdgeBrowserModelAllOf) GetRegionsOk() ([]AmazonSnowballEdgeRegionBrowserModel, bool) {
	if o == nil || isNil(o.Regions) {
    return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *AmazonSnowballEdgeBrowserModelAllOf) HasRegions() bool {
	if o != nil && !isNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []AmazonSnowballEdgeRegionBrowserModel and assigns it to the Regions field.
func (o *AmazonSnowballEdgeBrowserModelAllOf) SetRegions(v []AmazonSnowballEdgeRegionBrowserModel) {
	o.Regions = v
}

func (o AmazonSnowballEdgeBrowserModelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.HostId) {
		toSerialize["hostId"] = o.HostId
	}
	if !isNil(o.ConnectionPoint) {
		toSerialize["connectionPoint"] = o.ConnectionPoint
	}
	if !isNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonSnowballEdgeBrowserModelAllOf struct {
	value *AmazonSnowballEdgeBrowserModelAllOf
	isSet bool
}

func (v NullableAmazonSnowballEdgeBrowserModelAllOf) Get() *AmazonSnowballEdgeBrowserModelAllOf {
	return v.value
}

func (v *NullableAmazonSnowballEdgeBrowserModelAllOf) Set(val *AmazonSnowballEdgeBrowserModelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonSnowballEdgeBrowserModelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonSnowballEdgeBrowserModelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonSnowballEdgeBrowserModelAllOf(val *AmazonSnowballEdgeBrowserModelAllOf) *NullableAmazonSnowballEdgeBrowserModelAllOf {
	return &NullableAmazonSnowballEdgeBrowserModelAllOf{value: val, isSet: true}
}

func (v NullableAmazonSnowballEdgeBrowserModelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonSnowballEdgeBrowserModelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


