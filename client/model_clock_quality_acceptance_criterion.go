/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ClockQualityAcceptanceCriterion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClockQualityAcceptanceCriterion{}

// ClockQualityAcceptanceCriterion Contains a Clock Quality Acceptance Criterion.
type ClockQualityAcceptanceCriterion struct {
	SynchronizationState *SynchronizationState `json:"synchronizationState,omitempty"`
	ClockQuality *ClockQuality `json:"clockQuality,omitempty"`
	ParentTimeSource *TimeSource `json:"parentTimeSource,omitempty"`
}

// NewClockQualityAcceptanceCriterion instantiates a new ClockQualityAcceptanceCriterion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClockQualityAcceptanceCriterion() *ClockQualityAcceptanceCriterion {
	this := ClockQualityAcceptanceCriterion{}
	return &this
}

// NewClockQualityAcceptanceCriterionWithDefaults instantiates a new ClockQualityAcceptanceCriterion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClockQualityAcceptanceCriterionWithDefaults() *ClockQualityAcceptanceCriterion {
	this := ClockQualityAcceptanceCriterion{}
	return &this
}

// GetSynchronizationState returns the SynchronizationState field value if set, zero value otherwise.
func (o *ClockQualityAcceptanceCriterion) GetSynchronizationState() SynchronizationState {
	if o == nil || IsNil(o.SynchronizationState) {
		var ret SynchronizationState
		return ret
	}
	return *o.SynchronizationState
}

// GetSynchronizationStateOk returns a tuple with the SynchronizationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClockQualityAcceptanceCriterion) GetSynchronizationStateOk() (*SynchronizationState, bool) {
	if o == nil || IsNil(o.SynchronizationState) {
		return nil, false
	}
	return o.SynchronizationState, true
}

// HasSynchronizationState returns a boolean if a field has been set.
func (o *ClockQualityAcceptanceCriterion) HasSynchronizationState() bool {
	if o != nil && !IsNil(o.SynchronizationState) {
		return true
	}

	return false
}

// SetSynchronizationState gets a reference to the given SynchronizationState and assigns it to the SynchronizationState field.
func (o *ClockQualityAcceptanceCriterion) SetSynchronizationState(v SynchronizationState) {
	o.SynchronizationState = &v
}

// GetClockQuality returns the ClockQuality field value if set, zero value otherwise.
func (o *ClockQualityAcceptanceCriterion) GetClockQuality() ClockQuality {
	if o == nil || IsNil(o.ClockQuality) {
		var ret ClockQuality
		return ret
	}
	return *o.ClockQuality
}

// GetClockQualityOk returns a tuple with the ClockQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClockQualityAcceptanceCriterion) GetClockQualityOk() (*ClockQuality, bool) {
	if o == nil || IsNil(o.ClockQuality) {
		return nil, false
	}
	return o.ClockQuality, true
}

// HasClockQuality returns a boolean if a field has been set.
func (o *ClockQualityAcceptanceCriterion) HasClockQuality() bool {
	if o != nil && !IsNil(o.ClockQuality) {
		return true
	}

	return false
}

// SetClockQuality gets a reference to the given ClockQuality and assigns it to the ClockQuality field.
func (o *ClockQualityAcceptanceCriterion) SetClockQuality(v ClockQuality) {
	o.ClockQuality = &v
}

// GetParentTimeSource returns the ParentTimeSource field value if set, zero value otherwise.
func (o *ClockQualityAcceptanceCriterion) GetParentTimeSource() TimeSource {
	if o == nil || IsNil(o.ParentTimeSource) {
		var ret TimeSource
		return ret
	}
	return *o.ParentTimeSource
}

// GetParentTimeSourceOk returns a tuple with the ParentTimeSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClockQualityAcceptanceCriterion) GetParentTimeSourceOk() (*TimeSource, bool) {
	if o == nil || IsNil(o.ParentTimeSource) {
		return nil, false
	}
	return o.ParentTimeSource, true
}

// HasParentTimeSource returns a boolean if a field has been set.
func (o *ClockQualityAcceptanceCriterion) HasParentTimeSource() bool {
	if o != nil && !IsNil(o.ParentTimeSource) {
		return true
	}

	return false
}

// SetParentTimeSource gets a reference to the given TimeSource and assigns it to the ParentTimeSource field.
func (o *ClockQualityAcceptanceCriterion) SetParentTimeSource(v TimeSource) {
	o.ParentTimeSource = &v
}

func (o ClockQualityAcceptanceCriterion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClockQualityAcceptanceCriterion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SynchronizationState) {
		toSerialize["synchronizationState"] = o.SynchronizationState
	}
	if !IsNil(o.ClockQuality) {
		toSerialize["clockQuality"] = o.ClockQuality
	}
	if !IsNil(o.ParentTimeSource) {
		toSerialize["parentTimeSource"] = o.ParentTimeSource
	}
	return toSerialize, nil
}

type NullableClockQualityAcceptanceCriterion struct {
	value *ClockQualityAcceptanceCriterion
	isSet bool
}

func (v NullableClockQualityAcceptanceCriterion) Get() *ClockQualityAcceptanceCriterion {
	return v.value
}

func (v *NullableClockQualityAcceptanceCriterion) Set(val *ClockQualityAcceptanceCriterion) {
	v.value = val
	v.isSet = true
}

func (v NullableClockQualityAcceptanceCriterion) IsSet() bool {
	return v.isSet
}

func (v *NullableClockQualityAcceptanceCriterion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClockQualityAcceptanceCriterion(val *ClockQualityAcceptanceCriterion) *NullableClockQualityAcceptanceCriterion {
	return &NullableClockQualityAcceptanceCriterion{value: val, isSet: true}
}

func (v NullableClockQualityAcceptanceCriterion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClockQualityAcceptanceCriterion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


