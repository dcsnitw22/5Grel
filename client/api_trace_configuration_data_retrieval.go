/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// TraceConfigurationDataRetrievalAPIService TraceConfigurationDataRetrievalAPI service
type TraceConfigurationDataRetrievalAPIService service

type ApiGetTraceConfigDataRequest struct {
	ctx context.Context
	ApiService *TraceConfigurationDataRetrievalAPIService
	supi string
	supportedFeatures *string
	plmnId *PlmnId
	ifNoneMatch *string
	ifModifiedSince *string
}

// Supported Features
func (r ApiGetTraceConfigDataRequest) SupportedFeatures(supportedFeatures string) ApiGetTraceConfigDataRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

// serving PLMN ID
func (r ApiGetTraceConfigDataRequest) PlmnId(plmnId PlmnId) ApiGetTraceConfigDataRequest {
	r.plmnId = &plmnId
	return r
}

// Validator for conditional requests, as described in RFC 9110, 3.2
func (r ApiGetTraceConfigDataRequest) IfNoneMatch(ifNoneMatch string) ApiGetTraceConfigDataRequest {
	r.ifNoneMatch = &ifNoneMatch
	return r
}

// Validator for conditional requests, as described in RFC 9110, 3.3
func (r ApiGetTraceConfigDataRequest) IfModifiedSince(ifModifiedSince string) ApiGetTraceConfigDataRequest {
	r.ifModifiedSince = &ifModifiedSince
	return r
}

func (r ApiGetTraceConfigDataRequest) Execute() (*TraceDataResponse, *http.Response, error) {
	return r.ApiService.GetTraceConfigDataExecute(r)
}

/*
GetTraceConfigData retrieve a UE's Trace Configuration Data

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param supi Identifier of the UE
 @return ApiGetTraceConfigDataRequest
*/
func (a *TraceConfigurationDataRetrievalAPIService) GetTraceConfigData(ctx context.Context, supi string) ApiGetTraceConfigDataRequest {
	return ApiGetTraceConfigDataRequest{
		ApiService: a,
		ctx: ctx,
		supi: supi,
	}
}

// Execute executes the request
//  @return TraceDataResponse
func (a *TraceConfigurationDataRetrievalAPIService) GetTraceConfigDataExecute(r ApiGetTraceConfigDataRequest) (*TraceDataResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TraceDataResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TraceConfigurationDataRetrievalAPIService.GetTraceConfigData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{supi}/trace-data"
	localVarPath = strings.Replace(localVarPath, "{"+"supi"+"}", url.PathEscape(parameterValueToString(r.supi, "supi")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.supportedFeatures != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "supported-features", r.supportedFeatures, "")
	}
	if r.plmnId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "plmn-id", r.plmnId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ifNoneMatch != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-None-Match", r.ifNoneMatch, "")
	}
	if r.ifModifiedSince != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Modified-Since", r.ifModifiedSince, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 502 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}