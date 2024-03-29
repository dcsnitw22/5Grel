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

// NsacAdmissionMode Indicates the NSAC admission mode applied in roaming case. 
type NsacAdmissionMode struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NsacAdmissionMode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(NsacAdmissionMode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NsacAdmissionMode) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNsacAdmissionMode struct {
	value *NsacAdmissionMode
	isSet bool
}

func (v NullableNsacAdmissionMode) Get() *NsacAdmissionMode {
	return v.value
}

func (v *NullableNsacAdmissionMode) Set(val *NsacAdmissionMode) {
	v.value = val
	v.isSet = true
}

func (v NullableNsacAdmissionMode) IsSet() bool {
	return v.isSet
}

func (v *NullableNsacAdmissionMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsacAdmissionMode(val *NsacAdmissionMode) *NullableNsacAdmissionMode {
	return &NullableNsacAdmissionMode{value: val, isSet: true}
}

func (v NullableNsacAdmissionMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsacAdmissionMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


