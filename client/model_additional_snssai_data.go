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

// checks if the AdditionalSnssaiData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalSnssaiData{}

// AdditionalSnssaiData struct for AdditionalSnssaiData
type AdditionalSnssaiData struct {
	RequiredAuthnAuthz *bool `json:"requiredAuthnAuthz,omitempty"`
	SubscribedUeSliceMbr NullableSliceMbr `json:"subscribedUeSliceMbr,omitempty"`
	SubscribedNsSrgList []string `json:"subscribedNsSrgList,omitempty"`
	NsacMode *NsacAdmissionMode `json:"nsacMode,omitempty"`
	ValidTimePeriod *ValidTimePeriod `json:"validTimePeriod,omitempty"`
	// indicating a time in seconds.
	DeregInactTimer *int32 `json:"deregInactTimer,omitempty"`
	OnDemand *bool `json:"onDemand,omitempty"`
}

// NewAdditionalSnssaiData instantiates a new AdditionalSnssaiData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalSnssaiData() *AdditionalSnssaiData {
	this := AdditionalSnssaiData{}
	var onDemand bool = false
	this.OnDemand = &onDemand
	return &this
}

// NewAdditionalSnssaiDataWithDefaults instantiates a new AdditionalSnssaiData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalSnssaiDataWithDefaults() *AdditionalSnssaiData {
	this := AdditionalSnssaiData{}
	var onDemand bool = false
	this.OnDemand = &onDemand
	return &this
}

// GetRequiredAuthnAuthz returns the RequiredAuthnAuthz field value if set, zero value otherwise.
func (o *AdditionalSnssaiData) GetRequiredAuthnAuthz() bool {
	if o == nil || IsNil(o.RequiredAuthnAuthz) {
		var ret bool
		return ret
	}
	return *o.RequiredAuthnAuthz
}

// GetRequiredAuthnAuthzOk returns a tuple with the RequiredAuthnAuthz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalSnssaiData) GetRequiredAuthnAuthzOk() (*bool, bool) {
	if o == nil || IsNil(o.RequiredAuthnAuthz) {
		return nil, false
	}
	return o.RequiredAuthnAuthz, true
}

// HasRequiredAuthnAuthz returns a boolean if a field has been set.
func (o *AdditionalSnssaiData) HasRequiredAuthnAuthz() bool {
	if o != nil && !IsNil(o.RequiredAuthnAuthz) {
		return true
	}

	return false
}

// SetRequiredAuthnAuthz gets a reference to the given bool and assigns it to the RequiredAuthnAuthz field.
func (o *AdditionalSnssaiData) SetRequiredAuthnAuthz(v bool) {
	o.RequiredAuthnAuthz = &v
}

// GetSubscribedUeSliceMbr returns the SubscribedUeSliceMbr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdditionalSnssaiData) GetSubscribedUeSliceMbr() SliceMbr {
	if o == nil || IsNil(o.SubscribedUeSliceMbr.Get()) {
		var ret SliceMbr
		return ret
	}
	return *o.SubscribedUeSliceMbr.Get()
}

// GetSubscribedUeSliceMbrOk returns a tuple with the SubscribedUeSliceMbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdditionalSnssaiData) GetSubscribedUeSliceMbrOk() (*SliceMbr, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscribedUeSliceMbr.Get(), o.SubscribedUeSliceMbr.IsSet()
}

// HasSubscribedUeSliceMbr returns a boolean if a field has been set.
func (o *AdditionalSnssaiData) HasSubscribedUeSliceMbr() bool {
	if o != nil && o.SubscribedUeSliceMbr.IsSet() {
		return true
	}

	return false
}

// SetSubscribedUeSliceMbr gets a reference to the given NullableSliceMbr and assigns it to the SubscribedUeSliceMbr field.
func (o *AdditionalSnssaiData) SetSubscribedUeSliceMbr(v SliceMbr) {
	o.SubscribedUeSliceMbr.Set(&v)
}
// SetSubscribedUeSliceMbrNil sets the value for SubscribedUeSliceMbr to be an explicit nil
func (o *AdditionalSnssaiData) SetSubscribedUeSliceMbrNil() {
	o.SubscribedUeSliceMbr.Set(nil)
}

// UnsetSubscribedUeSliceMbr ensures that no value is present for SubscribedUeSliceMbr, not even an explicit nil
func (o *AdditionalSnssaiData) UnsetSubscribedUeSliceMbr() {
	o.SubscribedUeSliceMbr.Unset()
}

// GetSubscribedNsSrgList returns the SubscribedNsSrgList field value if set, zero value otherwise.
func (o *AdditionalSnssaiData) GetSubscribedNsSrgList() []string {
	if o == nil || IsNil(o.SubscribedNsSrgList) {
		var ret []string
		return ret
	}
	return o.SubscribedNsSrgList
}

// GetSubscribedNsSrgListOk returns a tuple with the SubscribedNsSrgList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalSnssaiData) GetSubscribedNsSrgListOk() ([]string, bool) {
	if o == nil || IsNil(o.SubscribedNsSrgList) {
		return nil, false
	}
	return o.SubscribedNsSrgList, true
}

// HasSubscribedNsSrgList returns a boolean if a field has been set.
func (o *AdditionalSnssaiData) HasSubscribedNsSrgList() bool {
	if o != nil && !IsNil(o.SubscribedNsSrgList) {
		return true
	}

	return false
}

// SetSubscribedNsSrgList gets a reference to the given []string and assigns it to the SubscribedNsSrgList field.
func (o *AdditionalSnssaiData) SetSubscribedNsSrgList(v []string) {
	o.SubscribedNsSrgList = v
}

// GetNsacMode returns the NsacMode field value if set, zero value otherwise.
func (o *AdditionalSnssaiData) GetNsacMode() NsacAdmissionMode {
	if o == nil || IsNil(o.NsacMode) {
		var ret NsacAdmissionMode
		return ret
	}
	return *o.NsacMode
}

// GetNsacModeOk returns a tuple with the NsacMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalSnssaiData) GetNsacModeOk() (*NsacAdmissionMode, bool) {
	if o == nil || IsNil(o.NsacMode) {
		return nil, false
	}
	return o.NsacMode, true
}

// HasNsacMode returns a boolean if a field has been set.
func (o *AdditionalSnssaiData) HasNsacMode() bool {
	if o != nil && !IsNil(o.NsacMode) {
		return true
	}

	return false
}

// SetNsacMode gets a reference to the given NsacAdmissionMode and assigns it to the NsacMode field.
func (o *AdditionalSnssaiData) SetNsacMode(v NsacAdmissionMode) {
	o.NsacMode = &v
}

// GetValidTimePeriod returns the ValidTimePeriod field value if set, zero value otherwise.
func (o *AdditionalSnssaiData) GetValidTimePeriod() ValidTimePeriod {
	if o == nil || IsNil(o.ValidTimePeriod) {
		var ret ValidTimePeriod
		return ret
	}
	return *o.ValidTimePeriod
}

// GetValidTimePeriodOk returns a tuple with the ValidTimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalSnssaiData) GetValidTimePeriodOk() (*ValidTimePeriod, bool) {
	if o == nil || IsNil(o.ValidTimePeriod) {
		return nil, false
	}
	return o.ValidTimePeriod, true
}

// HasValidTimePeriod returns a boolean if a field has been set.
func (o *AdditionalSnssaiData) HasValidTimePeriod() bool {
	if o != nil && !IsNil(o.ValidTimePeriod) {
		return true
	}

	return false
}

// SetValidTimePeriod gets a reference to the given ValidTimePeriod and assigns it to the ValidTimePeriod field.
func (o *AdditionalSnssaiData) SetValidTimePeriod(v ValidTimePeriod) {
	o.ValidTimePeriod = &v
}

// GetDeregInactTimer returns the DeregInactTimer field value if set, zero value otherwise.
func (o *AdditionalSnssaiData) GetDeregInactTimer() int32 {
	if o == nil || IsNil(o.DeregInactTimer) {
		var ret int32
		return ret
	}
	return *o.DeregInactTimer
}

// GetDeregInactTimerOk returns a tuple with the DeregInactTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalSnssaiData) GetDeregInactTimerOk() (*int32, bool) {
	if o == nil || IsNil(o.DeregInactTimer) {
		return nil, false
	}
	return o.DeregInactTimer, true
}

// HasDeregInactTimer returns a boolean if a field has been set.
func (o *AdditionalSnssaiData) HasDeregInactTimer() bool {
	if o != nil && !IsNil(o.DeregInactTimer) {
		return true
	}

	return false
}

// SetDeregInactTimer gets a reference to the given int32 and assigns it to the DeregInactTimer field.
func (o *AdditionalSnssaiData) SetDeregInactTimer(v int32) {
	o.DeregInactTimer = &v
}

// GetOnDemand returns the OnDemand field value if set, zero value otherwise.
func (o *AdditionalSnssaiData) GetOnDemand() bool {
	if o == nil || IsNil(o.OnDemand) {
		var ret bool
		return ret
	}
	return *o.OnDemand
}

// GetOnDemandOk returns a tuple with the OnDemand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalSnssaiData) GetOnDemandOk() (*bool, bool) {
	if o == nil || IsNil(o.OnDemand) {
		return nil, false
	}
	return o.OnDemand, true
}

// HasOnDemand returns a boolean if a field has been set.
func (o *AdditionalSnssaiData) HasOnDemand() bool {
	if o != nil && !IsNil(o.OnDemand) {
		return true
	}

	return false
}

// SetOnDemand gets a reference to the given bool and assigns it to the OnDemand field.
func (o *AdditionalSnssaiData) SetOnDemand(v bool) {
	o.OnDemand = &v
}

func (o AdditionalSnssaiData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalSnssaiData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequiredAuthnAuthz) {
		toSerialize["requiredAuthnAuthz"] = o.RequiredAuthnAuthz
	}
	if o.SubscribedUeSliceMbr.IsSet() {
		toSerialize["subscribedUeSliceMbr"] = o.SubscribedUeSliceMbr.Get()
	}
	if !IsNil(o.SubscribedNsSrgList) {
		toSerialize["subscribedNsSrgList"] = o.SubscribedNsSrgList
	}
	if !IsNil(o.NsacMode) {
		toSerialize["nsacMode"] = o.NsacMode
	}
	if !IsNil(o.ValidTimePeriod) {
		toSerialize["validTimePeriod"] = o.ValidTimePeriod
	}
	if !IsNil(o.DeregInactTimer) {
		toSerialize["deregInactTimer"] = o.DeregInactTimer
	}
	if !IsNil(o.OnDemand) {
		toSerialize["onDemand"] = o.OnDemand
	}
	return toSerialize, nil
}

type NullableAdditionalSnssaiData struct {
	value *AdditionalSnssaiData
	isSet bool
}

func (v NullableAdditionalSnssaiData) Get() *AdditionalSnssaiData {
	return v.value
}

func (v *NullableAdditionalSnssaiData) Set(val *AdditionalSnssaiData) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalSnssaiData) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalSnssaiData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalSnssaiData(val *AdditionalSnssaiData) *NullableAdditionalSnssaiData {
	return &NullableAdditionalSnssaiData{value: val, isSet: true}
}

func (v NullableAdditionalSnssaiData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalSnssaiData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


