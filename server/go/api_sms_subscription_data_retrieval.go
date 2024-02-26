/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// SMSSubscriptionDataRetrievalAPIController binds http requests to an api service and writes the service results to the http response
type SMSSubscriptionDataRetrievalAPIController struct {
	service      SMSSubscriptionDataRetrievalAPIServicer
	errorHandler ErrorHandler
}

// SMSSubscriptionDataRetrievalAPIOption for how the controller is set up.
type SMSSubscriptionDataRetrievalAPIOption func(*SMSSubscriptionDataRetrievalAPIController)

// WithSMSSubscriptionDataRetrievalAPIErrorHandler inject ErrorHandler into controller
func WithSMSSubscriptionDataRetrievalAPIErrorHandler(h ErrorHandler) SMSSubscriptionDataRetrievalAPIOption {
	return func(c *SMSSubscriptionDataRetrievalAPIController) {
		c.errorHandler = h
	}
}

// NewSMSSubscriptionDataRetrievalAPIController creates a default api controller
func NewSMSSubscriptionDataRetrievalAPIController(s SMSSubscriptionDataRetrievalAPIServicer, opts ...SMSSubscriptionDataRetrievalAPIOption) Router {
	controller := &SMSSubscriptionDataRetrievalAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the SMSSubscriptionDataRetrievalAPIController
func (c *SMSSubscriptionDataRetrievalAPIController) Routes() Routes {
	return Routes{
		"GetSmsData": Route{
			strings.ToUpper("Get"),
			"/nudm-sdm/v2/{supi}/sms-data",
			c.GetSmsData,
		},
	}
}

// GetSmsData - retrieve a UE's SMS Subscription Data
func (c *SMSSubscriptionDataRetrievalAPIController) GetSmsData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	supiParam := params["supi"]
	if supiParam == "" {
		c.errorHandler(w, r, &RequiredError{"supi"}, nil)
		return
	}
	var supportedFeaturesParam string
	if query.Has("supported-features") {
		supportedFeaturesParam = query.Get("supported-features")
	}
	var plmnIdParam PlmnId
	if query.Has("plmn-id") {
		// plmnIdParam = query.Get("plmn-id")
		fmt.Println(plmnIdParam)
	}
	ifNoneMatchParam := r.Header.Get("If-None-Match")
	ifModifiedSinceParam := r.Header.Get("If-Modified-Since")
	result, err := c.service.GetSmsData(r.Context(), supiParam, supportedFeaturesParam, plmnIdParam, ifNoneMatchParam, ifModifiedSinceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
