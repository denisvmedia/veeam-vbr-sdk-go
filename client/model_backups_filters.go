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
	"time"
)

// BackupsFilters struct for BackupsFilters
type BackupsFilters struct {
	Skip *int32 `json:"skip,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	OrderColumn *EBackupsFiltersOrderColumn `json:"orderColumn,omitempty"`
	OrderAsc *bool `json:"orderAsc,omitempty"`
	NameFilter *string `json:"nameFilter,omitempty"`
	CreatedAfterFilter *time.Time `json:"createdAfterFilter,omitempty"`
	CreatedBeforeFilter *time.Time `json:"createdBeforeFilter,omitempty"`
	PlatformIdFilter *string `json:"platformIdFilter,omitempty"`
	JobIdFilter *string `json:"jobIdFilter,omitempty"`
	PolicyTagFilter *string `json:"policyTagFilter,omitempty"`
}

// NewBackupsFilters instantiates a new BackupsFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupsFilters() *BackupsFilters {
	this := BackupsFilters{}
	return &this
}

// NewBackupsFiltersWithDefaults instantiates a new BackupsFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupsFiltersWithDefaults() *BackupsFilters {
	this := BackupsFilters{}
	return &this
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *BackupsFilters) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupsFilters) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *BackupsFilters) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *BackupsFilters) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *BackupsFilters) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupsFilters) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *BackupsFilters) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *BackupsFilters) SetLimit(v int32) {
	o.Limit = &v
}

// GetOrderColumn returns the OrderColumn field value if set, zero value otherwise.
func (o *BackupsFilters) GetOrderColumn() EBackupsFiltersOrderColumn {
	if o == nil || o.OrderColumn == nil {
		var ret EBackupsFiltersOrderColumn
		return ret
	}
	return *o.OrderColumn
}

// GetOrderColumnOk returns a tuple with the OrderColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupsFilters) GetOrderColumnOk() (*EBackupsFiltersOrderColumn, bool) {
	if o == nil || o.OrderColumn == nil {
		return nil, false
	}
	return o.OrderColumn, true
}

// HasOrderColumn returns a boolean if a field has been set.
func (o *BackupsFilters) HasOrderColumn() bool {
	if o != nil && o.OrderColumn != nil {
		return true
	}

	return false
}

// SetOrderColumn gets a reference to the given EBackupsFiltersOrderColumn and assigns it to the OrderColumn field.
func (o *BackupsFilters) SetOrderColumn(v EBackupsFiltersOrderColumn) {
	o.OrderColumn = &v
}

// GetOrderAsc returns the OrderAsc field value if set, zero value otherwise.
func (o *BackupsFilters) GetOrderAsc() bool {
	if o == nil || o.OrderAsc == nil {
		var ret bool
		return ret
	}
	return *o.OrderAsc
}

// GetOrderAscOk returns a tuple with the OrderAsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupsFilters) GetOrderAscOk() (*bool, bool) {
	if o == nil || o.OrderAsc == nil {
		return nil, false
	}
	return o.OrderAsc, true
}

// HasOrderAsc returns a boolean if a field has been set.
func (o *BackupsFilters) HasOrderAsc() bool {
	if o != nil && o.OrderAsc != nil {
		return true
	}

	return false
}

// SetOrderAsc gets a reference to the given bool and assigns it to the OrderAsc field.
func (o *BackupsFilters) SetOrderAsc(v bool) {
	o.OrderAsc = &v
}

// GetNameFilter returns the NameFilter field value if set, zero value otherwise.
func (o *BackupsFilters) GetNameFilter() string {
	if o == nil || o.NameFilter == nil {
		var ret string
		return ret
	}
	return *o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupsFilters) GetNameFilterOk() (*string, bool) {
	if o == nil || o.NameFilter == nil {
		return nil, false
	}
	return o.NameFilter, true
}

// HasNameFilter returns a boolean if a field has been set.
func (o *BackupsFilters) HasNameFilter() bool {
	if o != nil && o.NameFilter != nil {
		return true
	}

	return false
}

// SetNameFilter gets a reference to the given string and assigns it to the NameFilter field.
func (o *BackupsFilters) SetNameFilter(v string) {
	o.NameFilter = &v
}

// GetCreatedAfterFilter returns the CreatedAfterFilter field value if set, zero value otherwise.
func (o *BackupsFilters) GetCreatedAfterFilter() time.Time {
	if o == nil || o.CreatedAfterFilter == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAfterFilter
}

// GetCreatedAfterFilterOk returns a tuple with the CreatedAfterFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupsFilters) GetCreatedAfterFilterOk() (*time.Time, bool) {
	if o == nil || o.CreatedAfterFilter == nil {
		return nil, false
	}
	return o.CreatedAfterFilter, true
}

// HasCreatedAfterFilter returns a boolean if a field has been set.
func (o *BackupsFilters) HasCreatedAfterFilter() bool {
	if o != nil && o.CreatedAfterFilter != nil {
		return true
	}

	return false
}

// SetCreatedAfterFilter gets a reference to the given time.Time and assigns it to the CreatedAfterFilter field.
func (o *BackupsFilters) SetCreatedAfterFilter(v time.Time) {
	o.CreatedAfterFilter = &v
}

// GetCreatedBeforeFilter returns the CreatedBeforeFilter field value if set, zero value otherwise.
func (o *BackupsFilters) GetCreatedBeforeFilter() time.Time {
	if o == nil || o.CreatedBeforeFilter == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedBeforeFilter
}

// GetCreatedBeforeFilterOk returns a tuple with the CreatedBeforeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupsFilters) GetCreatedBeforeFilterOk() (*time.Time, bool) {
	if o == nil || o.CreatedBeforeFilter == nil {
		return nil, false
	}
	return o.CreatedBeforeFilter, true
}

// HasCreatedBeforeFilter returns a boolean if a field has been set.
func (o *BackupsFilters) HasCreatedBeforeFilter() bool {
	if o != nil && o.CreatedBeforeFilter != nil {
		return true
	}

	return false
}

// SetCreatedBeforeFilter gets a reference to the given time.Time and assigns it to the CreatedBeforeFilter field.
func (o *BackupsFilters) SetCreatedBeforeFilter(v time.Time) {
	o.CreatedBeforeFilter = &v
}

// GetPlatformIdFilter returns the PlatformIdFilter field value if set, zero value otherwise.
func (o *BackupsFilters) GetPlatformIdFilter() string {
	if o == nil || o.PlatformIdFilter == nil {
		var ret string
		return ret
	}
	return *o.PlatformIdFilter
}

// GetPlatformIdFilterOk returns a tuple with the PlatformIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupsFilters) GetPlatformIdFilterOk() (*string, bool) {
	if o == nil || o.PlatformIdFilter == nil {
		return nil, false
	}
	return o.PlatformIdFilter, true
}

// HasPlatformIdFilter returns a boolean if a field has been set.
func (o *BackupsFilters) HasPlatformIdFilter() bool {
	if o != nil && o.PlatformIdFilter != nil {
		return true
	}

	return false
}

// SetPlatformIdFilter gets a reference to the given string and assigns it to the PlatformIdFilter field.
func (o *BackupsFilters) SetPlatformIdFilter(v string) {
	o.PlatformIdFilter = &v
}

// GetJobIdFilter returns the JobIdFilter field value if set, zero value otherwise.
func (o *BackupsFilters) GetJobIdFilter() string {
	if o == nil || o.JobIdFilter == nil {
		var ret string
		return ret
	}
	return *o.JobIdFilter
}

// GetJobIdFilterOk returns a tuple with the JobIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupsFilters) GetJobIdFilterOk() (*string, bool) {
	if o == nil || o.JobIdFilter == nil {
		return nil, false
	}
	return o.JobIdFilter, true
}

// HasJobIdFilter returns a boolean if a field has been set.
func (o *BackupsFilters) HasJobIdFilter() bool {
	if o != nil && o.JobIdFilter != nil {
		return true
	}

	return false
}

// SetJobIdFilter gets a reference to the given string and assigns it to the JobIdFilter field.
func (o *BackupsFilters) SetJobIdFilter(v string) {
	o.JobIdFilter = &v
}

// GetPolicyTagFilter returns the PolicyTagFilter field value if set, zero value otherwise.
func (o *BackupsFilters) GetPolicyTagFilter() string {
	if o == nil || o.PolicyTagFilter == nil {
		var ret string
		return ret
	}
	return *o.PolicyTagFilter
}

// GetPolicyTagFilterOk returns a tuple with the PolicyTagFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupsFilters) GetPolicyTagFilterOk() (*string, bool) {
	if o == nil || o.PolicyTagFilter == nil {
		return nil, false
	}
	return o.PolicyTagFilter, true
}

// HasPolicyTagFilter returns a boolean if a field has been set.
func (o *BackupsFilters) HasPolicyTagFilter() bool {
	if o != nil && o.PolicyTagFilter != nil {
		return true
	}

	return false
}

// SetPolicyTagFilter gets a reference to the given string and assigns it to the PolicyTagFilter field.
func (o *BackupsFilters) SetPolicyTagFilter(v string) {
	o.PolicyTagFilter = &v
}

func (o BackupsFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Skip != nil {
		toSerialize["skip"] = o.Skip
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.OrderColumn != nil {
		toSerialize["orderColumn"] = o.OrderColumn
	}
	if o.OrderAsc != nil {
		toSerialize["orderAsc"] = o.OrderAsc
	}
	if o.NameFilter != nil {
		toSerialize["nameFilter"] = o.NameFilter
	}
	if o.CreatedAfterFilter != nil {
		toSerialize["createdAfterFilter"] = o.CreatedAfterFilter
	}
	if o.CreatedBeforeFilter != nil {
		toSerialize["createdBeforeFilter"] = o.CreatedBeforeFilter
	}
	if o.PlatformIdFilter != nil {
		toSerialize["platformIdFilter"] = o.PlatformIdFilter
	}
	if o.JobIdFilter != nil {
		toSerialize["jobIdFilter"] = o.JobIdFilter
	}
	if o.PolicyTagFilter != nil {
		toSerialize["policyTagFilter"] = o.PolicyTagFilter
	}
	return json.Marshal(toSerialize)
}

type NullableBackupsFilters struct {
	value *BackupsFilters
	isSet bool
}

func (v NullableBackupsFilters) Get() *BackupsFilters {
	return v.value
}

func (v *NullableBackupsFilters) Set(val *BackupsFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupsFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupsFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupsFilters(val *BackupsFilters) *NullableBackupsFilters {
	return &NullableBackupsFilters{value: val, isSet: true}
}

func (v NullableBackupsFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupsFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


