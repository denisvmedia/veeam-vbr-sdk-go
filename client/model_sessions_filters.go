/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br>Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Some request and response bodies refer to reusable schema objects that are defined in the `schemas` section of the REST API specification. Schemas, in turn, may inherit a part of their properties by referring to other schemas, and be polymorphic by referring to multiple alternate schemas.

API version: 1.1-rev0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// SessionsFilters struct for SessionsFilters
type SessionsFilters struct {
	Skip *int32 `json:"skip,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	OrderColumn *ESessionsFiltersOrderColumn `json:"orderColumn,omitempty"`
	OrderAsc *bool `json:"orderAsc,omitempty"`
	NameFilter *string `json:"nameFilter,omitempty"`
	CreatedAfterFilter *time.Time `json:"createdAfterFilter,omitempty"`
	CreatedBeforeFilter *time.Time `json:"createdBeforeFilter,omitempty"`
	EndedAfterFilter *time.Time `json:"endedAfterFilter,omitempty"`
	EndedBeforeFilter *time.Time `json:"endedBeforeFilter,omitempty"`
	TypeFilter *ESessionType `json:"typeFilter,omitempty"`
	StateFilter *ESessionState `json:"stateFilter,omitempty"`
	ResultFilter *ESessionResult `json:"resultFilter,omitempty"`
	JobIdFilter *string `json:"jobIdFilter,omitempty"`
}

// NewSessionsFilters instantiates a new SessionsFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionsFilters() *SessionsFilters {
	this := SessionsFilters{}
	return &this
}

// NewSessionsFiltersWithDefaults instantiates a new SessionsFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionsFiltersWithDefaults() *SessionsFilters {
	this := SessionsFilters{}
	return &this
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *SessionsFilters) GetSkip() int32 {
	if o == nil || isNil(o.Skip) {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionsFilters) GetSkipOk() (*int32, bool) {
	if o == nil || isNil(o.Skip) {
    return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *SessionsFilters) HasSkip() bool {
	if o != nil && !isNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *SessionsFilters) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *SessionsFilters) GetLimit() int32 {
	if o == nil || isNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionsFilters) GetLimitOk() (*int32, bool) {
	if o == nil || isNil(o.Limit) {
    return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *SessionsFilters) HasLimit() bool {
	if o != nil && !isNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *SessionsFilters) SetLimit(v int32) {
	o.Limit = &v
}

// GetOrderColumn returns the OrderColumn field value if set, zero value otherwise.
func (o *SessionsFilters) GetOrderColumn() ESessionsFiltersOrderColumn {
	if o == nil || isNil(o.OrderColumn) {
		var ret ESessionsFiltersOrderColumn
		return ret
	}
	return *o.OrderColumn
}

// GetOrderColumnOk returns a tuple with the OrderColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionsFilters) GetOrderColumnOk() (*ESessionsFiltersOrderColumn, bool) {
	if o == nil || isNil(o.OrderColumn) {
    return nil, false
	}
	return o.OrderColumn, true
}

// HasOrderColumn returns a boolean if a field has been set.
func (o *SessionsFilters) HasOrderColumn() bool {
	if o != nil && !isNil(o.OrderColumn) {
		return true
	}

	return false
}

// SetOrderColumn gets a reference to the given ESessionsFiltersOrderColumn and assigns it to the OrderColumn field.
func (o *SessionsFilters) SetOrderColumn(v ESessionsFiltersOrderColumn) {
	o.OrderColumn = &v
}

// GetOrderAsc returns the OrderAsc field value if set, zero value otherwise.
func (o *SessionsFilters) GetOrderAsc() bool {
	if o == nil || isNil(o.OrderAsc) {
		var ret bool
		return ret
	}
	return *o.OrderAsc
}

// GetOrderAscOk returns a tuple with the OrderAsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionsFilters) GetOrderAscOk() (*bool, bool) {
	if o == nil || isNil(o.OrderAsc) {
    return nil, false
	}
	return o.OrderAsc, true
}

// HasOrderAsc returns a boolean if a field has been set.
func (o *SessionsFilters) HasOrderAsc() bool {
	if o != nil && !isNil(o.OrderAsc) {
		return true
	}

	return false
}

// SetOrderAsc gets a reference to the given bool and assigns it to the OrderAsc field.
func (o *SessionsFilters) SetOrderAsc(v bool) {
	o.OrderAsc = &v
}

// GetNameFilter returns the NameFilter field value if set, zero value otherwise.
func (o *SessionsFilters) GetNameFilter() string {
	if o == nil || isNil(o.NameFilter) {
		var ret string
		return ret
	}
	return *o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionsFilters) GetNameFilterOk() (*string, bool) {
	if o == nil || isNil(o.NameFilter) {
    return nil, false
	}
	return o.NameFilter, true
}

// HasNameFilter returns a boolean if a field has been set.
func (o *SessionsFilters) HasNameFilter() bool {
	if o != nil && !isNil(o.NameFilter) {
		return true
	}

	return false
}

// SetNameFilter gets a reference to the given string and assigns it to the NameFilter field.
func (o *SessionsFilters) SetNameFilter(v string) {
	o.NameFilter = &v
}

// GetCreatedAfterFilter returns the CreatedAfterFilter field value if set, zero value otherwise.
func (o *SessionsFilters) GetCreatedAfterFilter() time.Time {
	if o == nil || isNil(o.CreatedAfterFilter) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAfterFilter
}

// GetCreatedAfterFilterOk returns a tuple with the CreatedAfterFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionsFilters) GetCreatedAfterFilterOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAfterFilter) {
    return nil, false
	}
	return o.CreatedAfterFilter, true
}

// HasCreatedAfterFilter returns a boolean if a field has been set.
func (o *SessionsFilters) HasCreatedAfterFilter() bool {
	if o != nil && !isNil(o.CreatedAfterFilter) {
		return true
	}

	return false
}

// SetCreatedAfterFilter gets a reference to the given time.Time and assigns it to the CreatedAfterFilter field.
func (o *SessionsFilters) SetCreatedAfterFilter(v time.Time) {
	o.CreatedAfterFilter = &v
}

// GetCreatedBeforeFilter returns the CreatedBeforeFilter field value if set, zero value otherwise.
func (o *SessionsFilters) GetCreatedBeforeFilter() time.Time {
	if o == nil || isNil(o.CreatedBeforeFilter) {
		var ret time.Time
		return ret
	}
	return *o.CreatedBeforeFilter
}

// GetCreatedBeforeFilterOk returns a tuple with the CreatedBeforeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionsFilters) GetCreatedBeforeFilterOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedBeforeFilter) {
    return nil, false
	}
	return o.CreatedBeforeFilter, true
}

// HasCreatedBeforeFilter returns a boolean if a field has been set.
func (o *SessionsFilters) HasCreatedBeforeFilter() bool {
	if o != nil && !isNil(o.CreatedBeforeFilter) {
		return true
	}

	return false
}

// SetCreatedBeforeFilter gets a reference to the given time.Time and assigns it to the CreatedBeforeFilter field.
func (o *SessionsFilters) SetCreatedBeforeFilter(v time.Time) {
	o.CreatedBeforeFilter = &v
}

// GetEndedAfterFilter returns the EndedAfterFilter field value if set, zero value otherwise.
func (o *SessionsFilters) GetEndedAfterFilter() time.Time {
	if o == nil || isNil(o.EndedAfterFilter) {
		var ret time.Time
		return ret
	}
	return *o.EndedAfterFilter
}

// GetEndedAfterFilterOk returns a tuple with the EndedAfterFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionsFilters) GetEndedAfterFilterOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndedAfterFilter) {
    return nil, false
	}
	return o.EndedAfterFilter, true
}

// HasEndedAfterFilter returns a boolean if a field has been set.
func (o *SessionsFilters) HasEndedAfterFilter() bool {
	if o != nil && !isNil(o.EndedAfterFilter) {
		return true
	}

	return false
}

// SetEndedAfterFilter gets a reference to the given time.Time and assigns it to the EndedAfterFilter field.
func (o *SessionsFilters) SetEndedAfterFilter(v time.Time) {
	o.EndedAfterFilter = &v
}

// GetEndedBeforeFilter returns the EndedBeforeFilter field value if set, zero value otherwise.
func (o *SessionsFilters) GetEndedBeforeFilter() time.Time {
	if o == nil || isNil(o.EndedBeforeFilter) {
		var ret time.Time
		return ret
	}
	return *o.EndedBeforeFilter
}

// GetEndedBeforeFilterOk returns a tuple with the EndedBeforeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionsFilters) GetEndedBeforeFilterOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndedBeforeFilter) {
    return nil, false
	}
	return o.EndedBeforeFilter, true
}

// HasEndedBeforeFilter returns a boolean if a field has been set.
func (o *SessionsFilters) HasEndedBeforeFilter() bool {
	if o != nil && !isNil(o.EndedBeforeFilter) {
		return true
	}

	return false
}

// SetEndedBeforeFilter gets a reference to the given time.Time and assigns it to the EndedBeforeFilter field.
func (o *SessionsFilters) SetEndedBeforeFilter(v time.Time) {
	o.EndedBeforeFilter = &v
}

// GetTypeFilter returns the TypeFilter field value if set, zero value otherwise.
func (o *SessionsFilters) GetTypeFilter() ESessionType {
	if o == nil || isNil(o.TypeFilter) {
		var ret ESessionType
		return ret
	}
	return *o.TypeFilter
}

// GetTypeFilterOk returns a tuple with the TypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionsFilters) GetTypeFilterOk() (*ESessionType, bool) {
	if o == nil || isNil(o.TypeFilter) {
    return nil, false
	}
	return o.TypeFilter, true
}

// HasTypeFilter returns a boolean if a field has been set.
func (o *SessionsFilters) HasTypeFilter() bool {
	if o != nil && !isNil(o.TypeFilter) {
		return true
	}

	return false
}

// SetTypeFilter gets a reference to the given ESessionType and assigns it to the TypeFilter field.
func (o *SessionsFilters) SetTypeFilter(v ESessionType) {
	o.TypeFilter = &v
}

// GetStateFilter returns the StateFilter field value if set, zero value otherwise.
func (o *SessionsFilters) GetStateFilter() ESessionState {
	if o == nil || isNil(o.StateFilter) {
		var ret ESessionState
		return ret
	}
	return *o.StateFilter
}

// GetStateFilterOk returns a tuple with the StateFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionsFilters) GetStateFilterOk() (*ESessionState, bool) {
	if o == nil || isNil(o.StateFilter) {
    return nil, false
	}
	return o.StateFilter, true
}

// HasStateFilter returns a boolean if a field has been set.
func (o *SessionsFilters) HasStateFilter() bool {
	if o != nil && !isNil(o.StateFilter) {
		return true
	}

	return false
}

// SetStateFilter gets a reference to the given ESessionState and assigns it to the StateFilter field.
func (o *SessionsFilters) SetStateFilter(v ESessionState) {
	o.StateFilter = &v
}

// GetResultFilter returns the ResultFilter field value if set, zero value otherwise.
func (o *SessionsFilters) GetResultFilter() ESessionResult {
	if o == nil || isNil(o.ResultFilter) {
		var ret ESessionResult
		return ret
	}
	return *o.ResultFilter
}

// GetResultFilterOk returns a tuple with the ResultFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionsFilters) GetResultFilterOk() (*ESessionResult, bool) {
	if o == nil || isNil(o.ResultFilter) {
    return nil, false
	}
	return o.ResultFilter, true
}

// HasResultFilter returns a boolean if a field has been set.
func (o *SessionsFilters) HasResultFilter() bool {
	if o != nil && !isNil(o.ResultFilter) {
		return true
	}

	return false
}

// SetResultFilter gets a reference to the given ESessionResult and assigns it to the ResultFilter field.
func (o *SessionsFilters) SetResultFilter(v ESessionResult) {
	o.ResultFilter = &v
}

// GetJobIdFilter returns the JobIdFilter field value if set, zero value otherwise.
func (o *SessionsFilters) GetJobIdFilter() string {
	if o == nil || isNil(o.JobIdFilter) {
		var ret string
		return ret
	}
	return *o.JobIdFilter
}

// GetJobIdFilterOk returns a tuple with the JobIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionsFilters) GetJobIdFilterOk() (*string, bool) {
	if o == nil || isNil(o.JobIdFilter) {
    return nil, false
	}
	return o.JobIdFilter, true
}

// HasJobIdFilter returns a boolean if a field has been set.
func (o *SessionsFilters) HasJobIdFilter() bool {
	if o != nil && !isNil(o.JobIdFilter) {
		return true
	}

	return false
}

// SetJobIdFilter gets a reference to the given string and assigns it to the JobIdFilter field.
func (o *SessionsFilters) SetJobIdFilter(v string) {
	o.JobIdFilter = &v
}

func (o SessionsFilters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Skip) {
		toSerialize["skip"] = o.Skip
	}
	if !isNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !isNil(o.OrderColumn) {
		toSerialize["orderColumn"] = o.OrderColumn
	}
	if !isNil(o.OrderAsc) {
		toSerialize["orderAsc"] = o.OrderAsc
	}
	if !isNil(o.NameFilter) {
		toSerialize["nameFilter"] = o.NameFilter
	}
	if !isNil(o.CreatedAfterFilter) {
		toSerialize["createdAfterFilter"] = o.CreatedAfterFilter
	}
	if !isNil(o.CreatedBeforeFilter) {
		toSerialize["createdBeforeFilter"] = o.CreatedBeforeFilter
	}
	if !isNil(o.EndedAfterFilter) {
		toSerialize["endedAfterFilter"] = o.EndedAfterFilter
	}
	if !isNil(o.EndedBeforeFilter) {
		toSerialize["endedBeforeFilter"] = o.EndedBeforeFilter
	}
	if !isNil(o.TypeFilter) {
		toSerialize["typeFilter"] = o.TypeFilter
	}
	if !isNil(o.StateFilter) {
		toSerialize["stateFilter"] = o.StateFilter
	}
	if !isNil(o.ResultFilter) {
		toSerialize["resultFilter"] = o.ResultFilter
	}
	if !isNil(o.JobIdFilter) {
		toSerialize["jobIdFilter"] = o.JobIdFilter
	}
	return json.Marshal(toSerialize)
}

type NullableSessionsFilters struct {
	value *SessionsFilters
	isSet bool
}

func (v NullableSessionsFilters) Get() *SessionsFilters {
	return v.value
}

func (v *NullableSessionsFilters) Set(val *SessionsFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionsFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionsFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionsFilters(val *SessionsFilters) *NullableSessionsFilters {
	return &NullableSessionsFilters{value: val, isSet: true}
}

func (v NullableSessionsFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionsFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


