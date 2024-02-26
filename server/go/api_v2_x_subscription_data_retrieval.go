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
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// V2XSubscriptionDataRetrievalAPIController binds http requests to an api service and writes the service results to the http response
type V2XSubscriptionDataRetrievalAPIController struct {
	service      V2XSubscriptionDataRetrievalAPIServicer
	errorHandler ErrorHandler
}

// V2XSubscriptionDataRetrievalAPIOption for how the controller is set up.
type V2XSubscriptionDataRetrievalAPIOption func(*V2XSubscriptionDataRetrievalAPIController)

// WithV2XSubscriptionDataRetrievalAPIErrorHandler inject ErrorHandler into controller
func WithV2XSubscriptionDataRetrievalAPIErrorHandler(h ErrorHandler) V2XSubscriptionDataRetrievalAPIOption {
	return func(c *V2XSubscriptionDataRetrievalAPIController) {
		c.errorHandler = h
	}
}

// NewV2XSubscriptionDataRetrievalAPIController creates a default api controller
func NewV2XSubscriptionDataRetrievalAPIController(s V2XSubscriptionDataRetrievalAPIServicer, opts ...V2XSubscriptionDataRetrievalAPIOption) Router {
	controller := &V2XSubscriptionDataRetrievalAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the V2XSubscriptionDataRetrievalAPIController
func (c *V2XSubscriptionDataRetrievalAPIController) Routes() Routes {
	return Routes{
		"GetV2xData": Route{
			strings.ToUpper("Get"),
			"/nudm-sdm/v2/{supi}/v2x-data",
			c.GetV2xData,
		},
	}
}

// GetV2xData - retrieve a UE's V2X Subscription Data
func (c *V2XSubscriptionDataRetrievalAPIController) GetV2xData(w http.ResponseWriter, r *http.Request) {
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
	ifNoneMatchParam := r.Header.Get("If-None-Match")
	ifModifiedSinceParam := r.Header.Get("If-Modified-Since")
	result, err := c.service.GetV2xData(r.Context(), supiParam, supportedFeaturesParam, ifNoneMatchParam, ifModifiedSinceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}