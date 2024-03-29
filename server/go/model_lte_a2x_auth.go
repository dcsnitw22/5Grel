/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




// LteA2xAuth - Contains LTE A2X services authorized information.
type LteA2xAuth struct {

	UavUeAuth UeAuth `json:"uavUeAuth,omitempty"`
}

// AssertLteA2xAuthRequired checks if the required fields are not zero-ed
func AssertLteA2xAuthRequired(obj LteA2xAuth) error {
	if err := AssertUeAuthRequired(obj.UavUeAuth); err != nil {
		return err
	}
	return nil
}

// AssertLteA2xAuthConstraints checks if the values respects the defined constraints
func AssertLteA2xAuthConstraints(obj LteA2xAuth) error {
	return nil
}
