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

// checks if the Guami type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Guami{}

// Guami Globally Unique AMF Identifier constructed out of PLMN, Network and AMF identity.
type Guami struct {
	PlmnId PlmnIdNid `json:"plmnId"`
	// String identifying the AMF ID composed of AMF Region ID (8 bits), AMF Set ID (10 bits) and AMF  Pointer (6 bits) as specified in clause 2.10.1 of 3GPP TS 23.003. It is encoded as a string of  6 hexadecimal characters (i.e., 24 bits).  
	AmfId string `json:"amfId"`
}

type _Guami Guami

// NewGuami instantiates a new Guami object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuami(plmnId PlmnIdNid, amfId string) *Guami {
	this := Guami{}
	this.PlmnId = plmnId
	this.AmfId = amfId
	return &this
}

// NewGuamiWithDefaults instantiates a new Guami object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuamiWithDefaults() *Guami {
	this := Guami{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *Guami) GetPlmnId() PlmnIdNid {
	if o == nil {
		var ret PlmnIdNid
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *Guami) GetPlmnIdOk() (*PlmnIdNid, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *Guami) SetPlmnId(v PlmnIdNid) {
	o.PlmnId = v
}

// GetAmfId returns the AmfId field value
func (o *Guami) GetAmfId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmfId
}

// GetAmfIdOk returns a tuple with the AmfId field value
// and a boolean to check if the value has been set.
func (o *Guami) GetAmfIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmfId, true
}

// SetAmfId sets field value
func (o *Guami) SetAmfId(v string) {
	o.AmfId = v
}

func (o Guami) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Guami) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plmnId"] = o.PlmnId
	toSerialize["amfId"] = o.AmfId
	return toSerialize, nil
}

func (o *Guami) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"plmnId",
		"amfId",
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

	varGuami := _Guami{}

	err = json.Unmarshal(bytes, &varGuami)

	if err != nil {
		return err
	}

	*o = Guami(varGuami)

	return err
}

type NullableGuami struct {
	value *Guami
	isSet bool
}

func (v NullableGuami) Get() *Guami {
	return v.value
}

func (v *NullableGuami) Set(val *Guami) {
	v.value = val
	v.isSet = true
}

func (v NullableGuami) IsSet() bool {
	return v.isSet
}

func (v *NullableGuami) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuami(val *Guami) *NullableGuami {
	return &NullableGuami{value: val, isSet: true}
}

func (v NullableGuami) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuami) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


