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
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// ProvidingAcknowledgementOfSNSSAIsUpdateAPIController binds http requests to an api service and writes the service results to the http response
type ProvidingAcknowledgementOfSNSSAIsUpdateAPIController struct {
	service      ProvidingAcknowledgementOfSNSSAIsUpdateAPIServicer
	errorHandler ErrorHandler
}

// ProvidingAcknowledgementOfSNSSAIsUpdateAPIOption for how the controller is set up.
type ProvidingAcknowledgementOfSNSSAIsUpdateAPIOption func(*ProvidingAcknowledgementOfSNSSAIsUpdateAPIController)

// WithProvidingAcknowledgementOfSNSSAIsUpdateAPIErrorHandler inject ErrorHandler into controller
func WithProvidingAcknowledgementOfSNSSAIsUpdateAPIErrorHandler(h ErrorHandler) ProvidingAcknowledgementOfSNSSAIsUpdateAPIOption {
	return func(c *ProvidingAcknowledgementOfSNSSAIsUpdateAPIController) {
		c.errorHandler = h
	}
}

// NewProvidingAcknowledgementOfSNSSAIsUpdateAPIController creates a default api controller
func NewProvidingAcknowledgementOfSNSSAIsUpdateAPIController(s ProvidingAcknowledgementOfSNSSAIsUpdateAPIServicer, opts ...ProvidingAcknowledgementOfSNSSAIsUpdateAPIOption) Router {
	controller := &ProvidingAcknowledgementOfSNSSAIsUpdateAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ProvidingAcknowledgementOfSNSSAIsUpdateAPIController
func (c *ProvidingAcknowledgementOfSNSSAIsUpdateAPIController) Routes() Routes {
	return Routes{
		"SNSSAIsAck": Route{
			strings.ToUpper("Put"),
			"/nudm-sdm/v2/{supi}/am-data/subscribed-snssais-ack",
			c.SNSSAIsAck,
		},
	}
}

// SNSSAIsAck - Nudm_Sdm Info operation for S-NSSAIs acknowledgement
func (c *ProvidingAcknowledgementOfSNSSAIsUpdateAPIController) SNSSAIsAck(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	supiParam := params["supi"]
	if supiParam == "" {
		c.errorHandler(w, r, &RequiredError{"supi"}, nil)
		return
	}
	acknowledgeInfoParam := AcknowledgeInfo{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&acknowledgeInfoParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertAcknowledgeInfoRequired(acknowledgeInfoParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertAcknowledgeInfoConstraints(acknowledgeInfoParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.SNSSAIsAck(r.Context(), supiParam, acknowledgeInfoParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
