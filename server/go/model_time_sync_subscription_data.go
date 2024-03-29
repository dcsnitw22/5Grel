/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




// TimeSyncSubscriptionData - UE Time Synchronization Subscription Data
type TimeSyncSubscriptionData struct {

	AfReqAuthorizations []AfRequestAuthorization `json:"afReqAuthorizations"`

	ServiceIds []TimeSyncServiceId `json:"serviceIds"`
}

// AssertTimeSyncSubscriptionDataRequired checks if the required fields are not zero-ed
func AssertTimeSyncSubscriptionDataRequired(obj TimeSyncSubscriptionData) error {
	elements := map[string]interface{}{
		"afReqAuthorizations": obj.AfReqAuthorizations,
		"serviceIds": obj.ServiceIds,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.AfReqAuthorizations {
		if err := AssertAfRequestAuthorizationRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ServiceIds {
		if err := AssertTimeSyncServiceIdRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertTimeSyncSubscriptionDataConstraints checks if the values respects the defined constraints
func AssertTimeSyncSubscriptionDataConstraints(obj TimeSyncSubscriptionData) error {
	return nil
}
