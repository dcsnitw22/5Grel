/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




type EcsAddrConfigInfo struct {

	EcsServerAddr EcsServerAddr `json:"ecsServerAddr,omitempty"`

	SpatialValidityCond SpatialValidityCond `json:"spatialValidityCond,omitempty"`
}

// AssertEcsAddrConfigInfoRequired checks if the required fields are not zero-ed
func AssertEcsAddrConfigInfoRequired(obj EcsAddrConfigInfo) error {
	if err := AssertEcsServerAddrRequired(obj.EcsServerAddr); err != nil {
		return err
	}
	if err := AssertSpatialValidityCondRequired(obj.SpatialValidityCond); err != nil {
		return err
	}
	return nil
}

// AssertEcsAddrConfigInfoConstraints checks if the values respects the defined constraints
func AssertEcsAddrConfigInfoConstraints(obj EcsAddrConfigInfo) error {
	return nil
}