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

// AvailableRanVisibleQoeMetric The enumeration AvailableRanVisibleQoeMetric indicates different available RAN-visible QoE metrics to the gNB. It shall comply with the provisions defined in TS 29.571, table 5.6.3.22-1. 
type AvailableRanVisibleQoeMetric struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AvailableRanVisibleQoeMetric) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AvailableRanVisibleQoeMetric)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AvailableRanVisibleQoeMetric) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAvailableRanVisibleQoeMetric struct {
	value *AvailableRanVisibleQoeMetric
	isSet bool
}

func (v NullableAvailableRanVisibleQoeMetric) Get() *AvailableRanVisibleQoeMetric {
	return v.value
}

func (v *NullableAvailableRanVisibleQoeMetric) Set(val *AvailableRanVisibleQoeMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableRanVisibleQoeMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableRanVisibleQoeMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableRanVisibleQoeMetric(val *AvailableRanVisibleQoeMetric) *NullableAvailableRanVisibleQoeMetric {
	return &NullableAvailableRanVisibleQoeMetric{value: val, isSet: true}
}

func (v NullableAvailableRanVisibleQoeMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableRanVisibleQoeMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


