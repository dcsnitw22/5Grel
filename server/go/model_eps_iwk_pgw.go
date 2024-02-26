/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




// EpsIwkPgw - This datatype signifies the PGW FQDN, SMF instance ID and the PLMNId location.
type EpsIwkPgw struct {

	// Fully Qualified Domain Name
	PgwFqdn string `json:"pgwFqdn"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	SmfInstanceId string `json:"smfInstanceId"`

	PlmnId PlmnId `json:"plmnId,omitempty"`
}

// AssertEpsIwkPgwRequired checks if the required fields are not zero-ed
func AssertEpsIwkPgwRequired(obj EpsIwkPgw) error {
	elements := map[string]interface{}{
		"pgwFqdn": obj.PgwFqdn,
		"smfInstanceId": obj.SmfInstanceId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertPlmnIdRequired(obj.PlmnId); err != nil {
		return err
	}
	return nil
}

// AssertEpsIwkPgwConstraints checks if the values respects the defined constraints
func AssertEpsIwkPgwConstraints(obj EpsIwkPgw) error {
	return nil
}