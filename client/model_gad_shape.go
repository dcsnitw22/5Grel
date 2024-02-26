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

// checks if the GADShape type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GADShape{}

// GADShape Common base type for GAD shapes.
type GADShape struct {
	Shape SupportedGADShapes `json:"shape"`
}

type _GADShape GADShape

// NewGADShape instantiates a new GADShape object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGADShape(shape SupportedGADShapes) *GADShape {
	this := GADShape{}
	this.Shape = shape
	return &this
}

// NewGADShapeWithDefaults instantiates a new GADShape object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGADShapeWithDefaults() *GADShape {
	this := GADShape{}
	return &this
}

// GetShape returns the Shape field value
func (o *GADShape) GetShape() SupportedGADShapes {
	if o == nil {
		var ret SupportedGADShapes
		return ret
	}

	return o.Shape
}

// GetShapeOk returns a tuple with the Shape field value
// and a boolean to check if the value has been set.
func (o *GADShape) GetShapeOk() (*SupportedGADShapes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Shape, true
}

// SetShape sets field value
func (o *GADShape) SetShape(v SupportedGADShapes) {
	o.Shape = v
}

func (o GADShape) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GADShape) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shape"] = o.Shape
	return toSerialize, nil
}

func (o *GADShape) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"shape",
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

	varGADShape := _GADShape{}

	err = json.Unmarshal(bytes, &varGADShape)

	if err != nil {
		return err
	}

	*o = GADShape(varGADShape)

	return err
}

type NullableGADShape struct {
	value *GADShape
	isSet bool
}

func (v NullableGADShape) Get() *GADShape {
	return v.value
}

func (v *NullableGADShape) Set(val *GADShape) {
	v.value = val
	v.isSet = true
}

func (v NullableGADShape) IsSet() bool {
	return v.isSet
}

func (v *NullableGADShape) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGADShape(val *GADShape) *NullableGADShape {
	return &NullableGADShape{value: val, isSet: true}
}

func (v NullableGADShape) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGADShape) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


