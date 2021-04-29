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

// EmailCustomNotificationType struct for EmailCustomNotificationType
type EmailCustomNotificationType struct {
	// Notification subject. Use the following variables in the subject: *%Time%* (completion time), *%JobName%*, *%JobResult%*, *%ObjectCount%* (number of VMs in the job) and *%Issues%* (number of VMs in the job that have finished with the Warning or Failed status). 
	Subject *string `json:"subject,omitempty"`
	// If *true*, email notifications are sent when the job completes successfully.
	NotifyOnSuccess *bool `json:"notifyOnSuccess,omitempty"`
	// If *true*, email notifications are sent when the job completes with a warning.
	NotifyOnWarning *bool `json:"notifyOnWarning,omitempty"`
	// If *true*, email notifications are sent when the job fails.
	NotifyOnError *bool `json:"notifyOnError,omitempty"`
	// If *true*, email notifications are sent about the final job status only (not per every job retry).
	SuppressNotificationUntilLastRetry *bool `json:"SuppressNotificationUntilLastRetry,omitempty"`
}

// NewEmailCustomNotificationType instantiates a new EmailCustomNotificationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailCustomNotificationType() *EmailCustomNotificationType {
	this := EmailCustomNotificationType{}
	return &this
}

// NewEmailCustomNotificationTypeWithDefaults instantiates a new EmailCustomNotificationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailCustomNotificationTypeWithDefaults() *EmailCustomNotificationType {
	this := EmailCustomNotificationType{}
	return &this
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *EmailCustomNotificationType) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailCustomNotificationType) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *EmailCustomNotificationType) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *EmailCustomNotificationType) SetSubject(v string) {
	o.Subject = &v
}

// GetNotifyOnSuccess returns the NotifyOnSuccess field value if set, zero value otherwise.
func (o *EmailCustomNotificationType) GetNotifyOnSuccess() bool {
	if o == nil || o.NotifyOnSuccess == nil {
		var ret bool
		return ret
	}
	return *o.NotifyOnSuccess
}

// GetNotifyOnSuccessOk returns a tuple with the NotifyOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailCustomNotificationType) GetNotifyOnSuccessOk() (*bool, bool) {
	if o == nil || o.NotifyOnSuccess == nil {
		return nil, false
	}
	return o.NotifyOnSuccess, true
}

// HasNotifyOnSuccess returns a boolean if a field has been set.
func (o *EmailCustomNotificationType) HasNotifyOnSuccess() bool {
	if o != nil && o.NotifyOnSuccess != nil {
		return true
	}

	return false
}

// SetNotifyOnSuccess gets a reference to the given bool and assigns it to the NotifyOnSuccess field.
func (o *EmailCustomNotificationType) SetNotifyOnSuccess(v bool) {
	o.NotifyOnSuccess = &v
}

// GetNotifyOnWarning returns the NotifyOnWarning field value if set, zero value otherwise.
func (o *EmailCustomNotificationType) GetNotifyOnWarning() bool {
	if o == nil || o.NotifyOnWarning == nil {
		var ret bool
		return ret
	}
	return *o.NotifyOnWarning
}

// GetNotifyOnWarningOk returns a tuple with the NotifyOnWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailCustomNotificationType) GetNotifyOnWarningOk() (*bool, bool) {
	if o == nil || o.NotifyOnWarning == nil {
		return nil, false
	}
	return o.NotifyOnWarning, true
}

// HasNotifyOnWarning returns a boolean if a field has been set.
func (o *EmailCustomNotificationType) HasNotifyOnWarning() bool {
	if o != nil && o.NotifyOnWarning != nil {
		return true
	}

	return false
}

// SetNotifyOnWarning gets a reference to the given bool and assigns it to the NotifyOnWarning field.
func (o *EmailCustomNotificationType) SetNotifyOnWarning(v bool) {
	o.NotifyOnWarning = &v
}

// GetNotifyOnError returns the NotifyOnError field value if set, zero value otherwise.
func (o *EmailCustomNotificationType) GetNotifyOnError() bool {
	if o == nil || o.NotifyOnError == nil {
		var ret bool
		return ret
	}
	return *o.NotifyOnError
}

// GetNotifyOnErrorOk returns a tuple with the NotifyOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailCustomNotificationType) GetNotifyOnErrorOk() (*bool, bool) {
	if o == nil || o.NotifyOnError == nil {
		return nil, false
	}
	return o.NotifyOnError, true
}

// HasNotifyOnError returns a boolean if a field has been set.
func (o *EmailCustomNotificationType) HasNotifyOnError() bool {
	if o != nil && o.NotifyOnError != nil {
		return true
	}

	return false
}

// SetNotifyOnError gets a reference to the given bool and assigns it to the NotifyOnError field.
func (o *EmailCustomNotificationType) SetNotifyOnError(v bool) {
	o.NotifyOnError = &v
}

// GetSuppressNotificationUntilLastRetry returns the SuppressNotificationUntilLastRetry field value if set, zero value otherwise.
func (o *EmailCustomNotificationType) GetSuppressNotificationUntilLastRetry() bool {
	if o == nil || o.SuppressNotificationUntilLastRetry == nil {
		var ret bool
		return ret
	}
	return *o.SuppressNotificationUntilLastRetry
}

// GetSuppressNotificationUntilLastRetryOk returns a tuple with the SuppressNotificationUntilLastRetry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailCustomNotificationType) GetSuppressNotificationUntilLastRetryOk() (*bool, bool) {
	if o == nil || o.SuppressNotificationUntilLastRetry == nil {
		return nil, false
	}
	return o.SuppressNotificationUntilLastRetry, true
}

// HasSuppressNotificationUntilLastRetry returns a boolean if a field has been set.
func (o *EmailCustomNotificationType) HasSuppressNotificationUntilLastRetry() bool {
	if o != nil && o.SuppressNotificationUntilLastRetry != nil {
		return true
	}

	return false
}

// SetSuppressNotificationUntilLastRetry gets a reference to the given bool and assigns it to the SuppressNotificationUntilLastRetry field.
func (o *EmailCustomNotificationType) SetSuppressNotificationUntilLastRetry(v bool) {
	o.SuppressNotificationUntilLastRetry = &v
}

func (o EmailCustomNotificationType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.NotifyOnSuccess != nil {
		toSerialize["notifyOnSuccess"] = o.NotifyOnSuccess
	}
	if o.NotifyOnWarning != nil {
		toSerialize["notifyOnWarning"] = o.NotifyOnWarning
	}
	if o.NotifyOnError != nil {
		toSerialize["notifyOnError"] = o.NotifyOnError
	}
	if o.SuppressNotificationUntilLastRetry != nil {
		toSerialize["SuppressNotificationUntilLastRetry"] = o.SuppressNotificationUntilLastRetry
	}
	return json.Marshal(toSerialize)
}

type NullableEmailCustomNotificationType struct {
	value *EmailCustomNotificationType
	isSet bool
}

func (v NullableEmailCustomNotificationType) Get() *EmailCustomNotificationType {
	return v.value
}

func (v *NullableEmailCustomNotificationType) Set(val *EmailCustomNotificationType) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailCustomNotificationType) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailCustomNotificationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailCustomNotificationType(val *EmailCustomNotificationType) *NullableEmailCustomNotificationType {
	return &NullableEmailCustomNotificationType{value: val, isSet: true}
}

func (v NullableEmailCustomNotificationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailCustomNotificationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


