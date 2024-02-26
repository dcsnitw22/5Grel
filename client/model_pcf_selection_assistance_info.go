/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the PcfSelectionAssistanceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PcfSelectionAssistanceInfo{}

// PcfSelectionAssistanceInfo struct for PcfSelectionAssistanceInfo
type PcfSelectionAssistanceInfo struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn"`
	SingleNssai Snssai `json:"singleNssai"`
}

type _PcfSelectionAssistanceInfo PcfSelectionAssistanceInfo

// NewPcfSelectionAssistanceInfo instantiates a new PcfSelectionAssistanceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcfSelectionAssistanceInfo(dnn string, singleNssai Snssai) *PcfSelectionAssistanceInfo {
	this := PcfSelectionAssistanceInfo{}
	this.Dnn = dnn
	this.SingleNssai = singleNssai
	return &this
}

// NewPcfSelectionAssistanceInfoWithDefaults instantiates a new PcfSelectionAssistanceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcfSelectionAssistanceInfoWithDefaults() *PcfSelectionAssistanceInfo {
	this := PcfSelectionAssistanceInfo{}
	return &this
}

// GetDnn returns the Dnn field value
func (o *PcfSelectionAssistanceInfo) GetDnn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *PcfSelectionAssistanceInfo) GetDnnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *PcfSelectionAssistanceInfo) SetDnn(v string) {
	o.Dnn = v
}

// GetSingleNssai returns the SingleNssai field value
func (o *PcfSelectionAssistanceInfo) GetSingleNssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.SingleNssai
}

// GetSingleNssaiOk returns a tuple with the SingleNssai field value
// and a boolean to check if the value has been set.
func (o *PcfSelectionAssistanceInfo) GetSingleNssaiOk() (*Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SingleNssai, true
}

// SetSingleNssai sets field value
func (o *PcfSelectionAssistanceInfo) SetSingleNssai(v Snssai) {
	o.SingleNssai = v
}

func (o PcfSelectionAssistanceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PcfSelectionAssistanceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dnn"] = o.Dnn
	toSerialize["singleNssai"] = o.SingleNssai
	return toSerialize, nil
}

func (o *PcfSelectionAssistanceInfo) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dnn",
		"singleNssai",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPcfSelectionAssistanceInfo := _PcfSelectionAssistanceInfo{}

	err = json.Unmarshal(bytes, &varPcfSelectionAssistanceInfo)

	if err != nil {
		return err
	}

	*o = PcfSelectionAssistanceInfo(varPcfSelectionAssistanceInfo)

	return err
}

type NullablePcfSelectionAssistanceInfo struct {
	value *PcfSelectionAssistanceInfo
	isSet bool
}

func (v NullablePcfSelectionAssistanceInfo) Get() *PcfSelectionAssistanceInfo {
	return v.value
}

func (v *NullablePcfSelectionAssistanceInfo) Set(val *PcfSelectionAssistanceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePcfSelectionAssistanceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePcfSelectionAssistanceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcfSelectionAssistanceInfo(val *PcfSelectionAssistanceInfo) *NullablePcfSelectionAssistanceInfo {
	return &NullablePcfSelectionAssistanceInfo{value: val, isSet: true}
}

func (v NullablePcfSelectionAssistanceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcfSelectionAssistanceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


