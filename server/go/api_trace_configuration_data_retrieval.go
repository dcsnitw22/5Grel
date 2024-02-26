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

// TraceConfigurationDataRetrievalAPIController binds http requests to an api service and writes the service results to the http response
type TraceConfigurationDataRetrievalAPIController struct {
	service      TraceConfigurationDataRetrievalAPIServicer
	errorHandler ErrorHandler
}

// TraceConfigurationDataRetrievalAPIOption for how the controller is set up.
type TraceConfigurationDataRetrievalAPIOption func(*TraceConfigurationDataRetrievalAPIController)

// WithTraceConfigurationDataRetrievalAPIErrorHandler inject ErrorHandler into controller
func WithTraceConfigurationDataRetrievalAPIErrorHandler(h ErrorHandler) TraceConfigurationDataRetrievalAPIOption {
	return func(c *TraceConfigurationDataRetrievalAPIController) {
		c.errorHandler = h
	}
}

// NewTraceConfigurationDataRetrievalAPIController creates a default api controller
func NewTraceConfigurationDataRetrievalAPIController(s TraceConfigurationDataRetrievalAPIServicer, opts ...TraceConfigurationDataRetrievalAPIOption) Router {
	controller := &TraceConfigurationDataRetrievalAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the TraceConfigurationDataRetrievalAPIController
func (c *TraceConfigurationDataRetrievalAPIController) Routes() Routes {
	return Routes{
		"GetTraceConfigData": Route{
			strings.ToUpper("Get"),
			"/nudm-sdm/v2/{supi}/trace-data",
			c.GetTraceConfigData,
		},
	}
}

// GetTraceConfigData - retrieve a UE's Trace Configuration Data
func (c *TraceConfigurationDataRetrievalAPIController) GetTraceConfigData(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.GetTraceConfigData(r.Context(), supiParam, supportedFeaturesParam, plmnIdParam, ifNoneMatchParam, ifModifiedSinceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
