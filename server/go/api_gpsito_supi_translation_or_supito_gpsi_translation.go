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

// GPSIToSUPITranslationOrSUPIToGPSITranslationAPIController binds http requests to an api service and writes the service results to the http response
type GPSIToSUPITranslationOrSUPIToGPSITranslationAPIController struct {
	service      GPSIToSUPITranslationOrSUPIToGPSITranslationAPIServicer
	errorHandler ErrorHandler
}

// GPSIToSUPITranslationOrSUPIToGPSITranslationAPIOption for how the controller is set up.
type GPSIToSUPITranslationOrSUPIToGPSITranslationAPIOption func(*GPSIToSUPITranslationOrSUPIToGPSITranslationAPIController)

// WithGPSIToSUPITranslationOrSUPIToGPSITranslationAPIErrorHandler inject ErrorHandler into controller
func WithGPSIToSUPITranslationOrSUPIToGPSITranslationAPIErrorHandler(h ErrorHandler) GPSIToSUPITranslationOrSUPIToGPSITranslationAPIOption {
	return func(c *GPSIToSUPITranslationOrSUPIToGPSITranslationAPIController) {
		c.errorHandler = h
	}
}

// NewGPSIToSUPITranslationOrSUPIToGPSITranslationAPIController creates a default api controller
func NewGPSIToSUPITranslationOrSUPIToGPSITranslationAPIController(s GPSIToSUPITranslationOrSUPIToGPSITranslationAPIServicer, opts ...GPSIToSUPITranslationOrSUPIToGPSITranslationAPIOption) Router {
	controller := &GPSIToSUPITranslationOrSUPIToGPSITranslationAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the GPSIToSUPITranslationOrSUPIToGPSITranslationAPIController
func (c *GPSIToSUPITranslationOrSUPIToGPSITranslationAPIController) Routes() Routes {
	return Routes{
		"GetSupiOrGpsi": Route{
			strings.ToUpper("Get"),
			"/nudm-sdm/v2/{ueId}/id-translation-result",
			c.GetSupiOrGpsi,
		},
	}
}

// GetSupiOrGpsi - retrieve a UE's SUPI or GPSI
func (c *GPSIToSUPITranslationOrSUPIToGPSITranslationAPIController) GetSupiOrGpsi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	ueIdParam := params["ueId"]
	if ueIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"ueId"}, nil)
		return
	}
	var supportedFeaturesParam string
	if query.Has("supported-features") {
		supportedFeaturesParam = query.Get("supported-features")
	}
	var afIdParam string
	if query.Has("af-id") {
		afIdParam = query.Get("af-id")
	}
	var appPortIdParam AppPortId
	if query.Has("app-port-id") {
		// appPortIdParam = query.Get("app-port-id")
		fmt.Println(appPortIdParam)
	}
	var afServiceIdParam string
	if query.Has("af-service-id") {
		afServiceIdParam = query.Get("af-service-id")
	}
	var mtcProviderInfoParam string
	if query.Has("mtc-provider-info") {
		mtcProviderInfoParam = query.Get("mtc-provider-info")
	}
	var requestedGpsiTypeParam GpsiType
	if query.Has("requested-gpsi-type") {
		// requestedGpsiTypeParam = query.Get("requested-gpsi-type")
		fmt.Println(requestedGpsiTypeParam)
	}
	ifNoneMatchParam := r.Header.Get("If-None-Match")
	ifModifiedSinceParam := r.Header.Get("If-Modified-Since")
	result, err := c.service.GetSupiOrGpsi(r.Context(), ueIdParam, supportedFeaturesParam, afIdParam, appPortIdParam, afServiceIdParam, mtcProviderInfoParam, requestedGpsiTypeParam, ifNoneMatchParam, ifModifiedSinceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
