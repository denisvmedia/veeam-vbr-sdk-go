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

// LinuxCredentialsSpec struct for LinuxCredentialsSpec
type LinuxCredentialsSpec struct {
	CredentialsSpec
	// Tag used to identify the credentials record.
	Tag *string `json:"tag,omitempty"`
	// SSH port used to connect to a Linux server.
	SSHPort *int32 `json:"SSHPort,omitempty"`
	// If *true*, the permissions of the account are automatically elevated to the root user.
	AutoElevated *bool `json:"autoElevated,omitempty"`
	// If *true*, the account is automatically added to the sudoers file.
	AddToSudoers *bool `json:"addToSudoers,omitempty"`
	// If *true*, the `su` command is used for Linux distributions where the `sudo` command is not available.
	UseSu *bool `json:"useSu,omitempty"`
	// Private key.
	PrivateKey *string `json:"privateKey,omitempty"`
	// Passphrase that protects the private key.
	Passphrase *string `json:"passphrase,omitempty"`
	// Password for the root account.
	RootPassword *string `json:"rootPassword,omitempty"`
}

// NewLinuxCredentialsSpec instantiates a new LinuxCredentialsSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinuxCredentialsSpec(username string, type_ ECredentialsType) *LinuxCredentialsSpec {
	this := LinuxCredentialsSpec{}
	this.Username = username
	this.Type = type_
	return &this
}

// NewLinuxCredentialsSpecWithDefaults instantiates a new LinuxCredentialsSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinuxCredentialsSpecWithDefaults() *LinuxCredentialsSpec {
	this := LinuxCredentialsSpec{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *LinuxCredentialsSpec) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxCredentialsSpec) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *LinuxCredentialsSpec) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *LinuxCredentialsSpec) SetTag(v string) {
	o.Tag = &v
}

// GetSSHPort returns the SSHPort field value if set, zero value otherwise.
func (o *LinuxCredentialsSpec) GetSSHPort() int32 {
	if o == nil || o.SSHPort == nil {
		var ret int32
		return ret
	}
	return *o.SSHPort
}

// GetSSHPortOk returns a tuple with the SSHPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxCredentialsSpec) GetSSHPortOk() (*int32, bool) {
	if o == nil || o.SSHPort == nil {
		return nil, false
	}
	return o.SSHPort, true
}

// HasSSHPort returns a boolean if a field has been set.
func (o *LinuxCredentialsSpec) HasSSHPort() bool {
	if o != nil && o.SSHPort != nil {
		return true
	}

	return false
}

// SetSSHPort gets a reference to the given int32 and assigns it to the SSHPort field.
func (o *LinuxCredentialsSpec) SetSSHPort(v int32) {
	o.SSHPort = &v
}

// GetAutoElevated returns the AutoElevated field value if set, zero value otherwise.
func (o *LinuxCredentialsSpec) GetAutoElevated() bool {
	if o == nil || o.AutoElevated == nil {
		var ret bool
		return ret
	}
	return *o.AutoElevated
}

// GetAutoElevatedOk returns a tuple with the AutoElevated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxCredentialsSpec) GetAutoElevatedOk() (*bool, bool) {
	if o == nil || o.AutoElevated == nil {
		return nil, false
	}
	return o.AutoElevated, true
}

// HasAutoElevated returns a boolean if a field has been set.
func (o *LinuxCredentialsSpec) HasAutoElevated() bool {
	if o != nil && o.AutoElevated != nil {
		return true
	}

	return false
}

// SetAutoElevated gets a reference to the given bool and assigns it to the AutoElevated field.
func (o *LinuxCredentialsSpec) SetAutoElevated(v bool) {
	o.AutoElevated = &v
}

// GetAddToSudoers returns the AddToSudoers field value if set, zero value otherwise.
func (o *LinuxCredentialsSpec) GetAddToSudoers() bool {
	if o == nil || o.AddToSudoers == nil {
		var ret bool
		return ret
	}
	return *o.AddToSudoers
}

// GetAddToSudoersOk returns a tuple with the AddToSudoers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxCredentialsSpec) GetAddToSudoersOk() (*bool, bool) {
	if o == nil || o.AddToSudoers == nil {
		return nil, false
	}
	return o.AddToSudoers, true
}

// HasAddToSudoers returns a boolean if a field has been set.
func (o *LinuxCredentialsSpec) HasAddToSudoers() bool {
	if o != nil && o.AddToSudoers != nil {
		return true
	}

	return false
}

// SetAddToSudoers gets a reference to the given bool and assigns it to the AddToSudoers field.
func (o *LinuxCredentialsSpec) SetAddToSudoers(v bool) {
	o.AddToSudoers = &v
}

// GetUseSu returns the UseSu field value if set, zero value otherwise.
func (o *LinuxCredentialsSpec) GetUseSu() bool {
	if o == nil || o.UseSu == nil {
		var ret bool
		return ret
	}
	return *o.UseSu
}

// GetUseSuOk returns a tuple with the UseSu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxCredentialsSpec) GetUseSuOk() (*bool, bool) {
	if o == nil || o.UseSu == nil {
		return nil, false
	}
	return o.UseSu, true
}

// HasUseSu returns a boolean if a field has been set.
func (o *LinuxCredentialsSpec) HasUseSu() bool {
	if o != nil && o.UseSu != nil {
		return true
	}

	return false
}

// SetUseSu gets a reference to the given bool and assigns it to the UseSu field.
func (o *LinuxCredentialsSpec) SetUseSu(v bool) {
	o.UseSu = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *LinuxCredentialsSpec) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxCredentialsSpec) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *LinuxCredentialsSpec) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *LinuxCredentialsSpec) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *LinuxCredentialsSpec) GetPassphrase() string {
	if o == nil || o.Passphrase == nil {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxCredentialsSpec) GetPassphraseOk() (*string, bool) {
	if o == nil || o.Passphrase == nil {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *LinuxCredentialsSpec) HasPassphrase() bool {
	if o != nil && o.Passphrase != nil {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *LinuxCredentialsSpec) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetRootPassword returns the RootPassword field value if set, zero value otherwise.
func (o *LinuxCredentialsSpec) GetRootPassword() string {
	if o == nil || o.RootPassword == nil {
		var ret string
		return ret
	}
	return *o.RootPassword
}

// GetRootPasswordOk returns a tuple with the RootPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinuxCredentialsSpec) GetRootPasswordOk() (*string, bool) {
	if o == nil || o.RootPassword == nil {
		return nil, false
	}
	return o.RootPassword, true
}

// HasRootPassword returns a boolean if a field has been set.
func (o *LinuxCredentialsSpec) HasRootPassword() bool {
	if o != nil && o.RootPassword != nil {
		return true
	}

	return false
}

// SetRootPassword gets a reference to the given string and assigns it to the RootPassword field.
func (o *LinuxCredentialsSpec) SetRootPassword(v string) {
	o.RootPassword = &v
}

func (o LinuxCredentialsSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCredentialsSpec, errCredentialsSpec := json.Marshal(o.CredentialsSpec)
	if errCredentialsSpec != nil {
		return []byte{}, errCredentialsSpec
	}
	errCredentialsSpec = json.Unmarshal([]byte(serializedCredentialsSpec), &toSerialize)
	if errCredentialsSpec != nil {
		return []byte{}, errCredentialsSpec
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.SSHPort != nil {
		toSerialize["SSHPort"] = o.SSHPort
	}
	if o.AutoElevated != nil {
		toSerialize["autoElevated"] = o.AutoElevated
	}
	if o.AddToSudoers != nil {
		toSerialize["addToSudoers"] = o.AddToSudoers
	}
	if o.UseSu != nil {
		toSerialize["useSu"] = o.UseSu
	}
	if o.PrivateKey != nil {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if o.Passphrase != nil {
		toSerialize["passphrase"] = o.Passphrase
	}
	if o.RootPassword != nil {
		toSerialize["rootPassword"] = o.RootPassword
	}
	return json.Marshal(toSerialize)
}

type NullableLinuxCredentialsSpec struct {
	value *LinuxCredentialsSpec
	isSet bool
}

func (v NullableLinuxCredentialsSpec) Get() *LinuxCredentialsSpec {
	return v.value
}

func (v *NullableLinuxCredentialsSpec) Set(val *LinuxCredentialsSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableLinuxCredentialsSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableLinuxCredentialsSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinuxCredentialsSpec(val *LinuxCredentialsSpec) *NullableLinuxCredentialsSpec {
	return &NullableLinuxCredentialsSpec{value: val, isSet: true}
}

func (v NullableLinuxCredentialsSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinuxCredentialsSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


