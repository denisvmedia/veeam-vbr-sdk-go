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

// AmazonS3StorageProxyApplianceModel Amazon S3 proxy appliance.
type AmazonS3StorageProxyApplianceModel struct {
	// EC2 instance type for the proxy appliance.
	Ec2InstanceType *string `json:"ec2InstanceType,omitempty"`
	// Amazon VPC where Veeam Backup & Replication launches the target instance.
	Vpc *string `json:"vpc,omitempty"`
	// Subnet for the proxy appliance.
	Subnet *string `json:"subnet,omitempty"`
	// Security group associated with the proxy appliance.
	SecurityGroup *string `json:"securityGroup,omitempty"`
	// TCP port used to route requests between the proxy appliance and backup infrastructure components.
	RedirectorPort *int32 `json:"redirectorPort,omitempty"`
}

// NewAmazonS3StorageProxyApplianceModel instantiates a new AmazonS3StorageProxyApplianceModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonS3StorageProxyApplianceModel() *AmazonS3StorageProxyApplianceModel {
	this := AmazonS3StorageProxyApplianceModel{}
	return &this
}

// NewAmazonS3StorageProxyApplianceModelWithDefaults instantiates a new AmazonS3StorageProxyApplianceModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonS3StorageProxyApplianceModelWithDefaults() *AmazonS3StorageProxyApplianceModel {
	this := AmazonS3StorageProxyApplianceModel{}
	return &this
}

// GetEc2InstanceType returns the Ec2InstanceType field value if set, zero value otherwise.
func (o *AmazonS3StorageProxyApplianceModel) GetEc2InstanceType() string {
	if o == nil || isNil(o.Ec2InstanceType) {
		var ret string
		return ret
	}
	return *o.Ec2InstanceType
}

// GetEc2InstanceTypeOk returns a tuple with the Ec2InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonS3StorageProxyApplianceModel) GetEc2InstanceTypeOk() (*string, bool) {
	if o == nil || isNil(o.Ec2InstanceType) {
    return nil, false
	}
	return o.Ec2InstanceType, true
}

// HasEc2InstanceType returns a boolean if a field has been set.
func (o *AmazonS3StorageProxyApplianceModel) HasEc2InstanceType() bool {
	if o != nil && !isNil(o.Ec2InstanceType) {
		return true
	}

	return false
}

// SetEc2InstanceType gets a reference to the given string and assigns it to the Ec2InstanceType field.
func (o *AmazonS3StorageProxyApplianceModel) SetEc2InstanceType(v string) {
	o.Ec2InstanceType = &v
}

// GetVpc returns the Vpc field value if set, zero value otherwise.
func (o *AmazonS3StorageProxyApplianceModel) GetVpc() string {
	if o == nil || isNil(o.Vpc) {
		var ret string
		return ret
	}
	return *o.Vpc
}

// GetVpcOk returns a tuple with the Vpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonS3StorageProxyApplianceModel) GetVpcOk() (*string, bool) {
	if o == nil || isNil(o.Vpc) {
    return nil, false
	}
	return o.Vpc, true
}

// HasVpc returns a boolean if a field has been set.
func (o *AmazonS3StorageProxyApplianceModel) HasVpc() bool {
	if o != nil && !isNil(o.Vpc) {
		return true
	}

	return false
}

// SetVpc gets a reference to the given string and assigns it to the Vpc field.
func (o *AmazonS3StorageProxyApplianceModel) SetVpc(v string) {
	o.Vpc = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *AmazonS3StorageProxyApplianceModel) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonS3StorageProxyApplianceModel) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *AmazonS3StorageProxyApplianceModel) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *AmazonS3StorageProxyApplianceModel) SetSubnet(v string) {
	o.Subnet = &v
}

// GetSecurityGroup returns the SecurityGroup field value if set, zero value otherwise.
func (o *AmazonS3StorageProxyApplianceModel) GetSecurityGroup() string {
	if o == nil || isNil(o.SecurityGroup) {
		var ret string
		return ret
	}
	return *o.SecurityGroup
}

// GetSecurityGroupOk returns a tuple with the SecurityGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonS3StorageProxyApplianceModel) GetSecurityGroupOk() (*string, bool) {
	if o == nil || isNil(o.SecurityGroup) {
    return nil, false
	}
	return o.SecurityGroup, true
}

// HasSecurityGroup returns a boolean if a field has been set.
func (o *AmazonS3StorageProxyApplianceModel) HasSecurityGroup() bool {
	if o != nil && !isNil(o.SecurityGroup) {
		return true
	}

	return false
}

// SetSecurityGroup gets a reference to the given string and assigns it to the SecurityGroup field.
func (o *AmazonS3StorageProxyApplianceModel) SetSecurityGroup(v string) {
	o.SecurityGroup = &v
}

// GetRedirectorPort returns the RedirectorPort field value if set, zero value otherwise.
func (o *AmazonS3StorageProxyApplianceModel) GetRedirectorPort() int32 {
	if o == nil || isNil(o.RedirectorPort) {
		var ret int32
		return ret
	}
	return *o.RedirectorPort
}

// GetRedirectorPortOk returns a tuple with the RedirectorPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmazonS3StorageProxyApplianceModel) GetRedirectorPortOk() (*int32, bool) {
	if o == nil || isNil(o.RedirectorPort) {
    return nil, false
	}
	return o.RedirectorPort, true
}

// HasRedirectorPort returns a boolean if a field has been set.
func (o *AmazonS3StorageProxyApplianceModel) HasRedirectorPort() bool {
	if o != nil && !isNil(o.RedirectorPort) {
		return true
	}

	return false
}

// SetRedirectorPort gets a reference to the given int32 and assigns it to the RedirectorPort field.
func (o *AmazonS3StorageProxyApplianceModel) SetRedirectorPort(v int32) {
	o.RedirectorPort = &v
}

func (o AmazonS3StorageProxyApplianceModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ec2InstanceType) {
		toSerialize["ec2InstanceType"] = o.Ec2InstanceType
	}
	if !isNil(o.Vpc) {
		toSerialize["vpc"] = o.Vpc
	}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !isNil(o.SecurityGroup) {
		toSerialize["securityGroup"] = o.SecurityGroup
	}
	if !isNil(o.RedirectorPort) {
		toSerialize["redirectorPort"] = o.RedirectorPort
	}
	return json.Marshal(toSerialize)
}

type NullableAmazonS3StorageProxyApplianceModel struct {
	value *AmazonS3StorageProxyApplianceModel
	isSet bool
}

func (v NullableAmazonS3StorageProxyApplianceModel) Get() *AmazonS3StorageProxyApplianceModel {
	return v.value
}

func (v *NullableAmazonS3StorageProxyApplianceModel) Set(val *AmazonS3StorageProxyApplianceModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonS3StorageProxyApplianceModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonS3StorageProxyApplianceModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonS3StorageProxyApplianceModel(val *AmazonS3StorageProxyApplianceModel) *NullableAmazonS3StorageProxyApplianceModel {
	return &NullableAmazonS3StorageProxyApplianceModel{value: val, isSet: true}
}

func (v NullableAmazonS3StorageProxyApplianceModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonS3StorageProxyApplianceModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


