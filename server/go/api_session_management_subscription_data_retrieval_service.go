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
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis/v8"
)

// SessionManagementSubscriptionDataRetrievalAPIService is a service that implements the logic for the SessionManagementSubscriptionDataRetrievalAPIServicer
// This service should implement the business logic for every endpoint for the SessionManagementSubscriptionDataRetrievalAPI API.
// Include any external packages or services that will be required by this service.
type SessionManagementSubscriptionDataRetrievalAPIService struct {
	rdb *redis.Client
	ctx context.Context
}

// NewSessionManagementSubscriptionDataRetrievalAPIService creates a default api service
func NewSessionManagementSubscriptionDataRetrievalAPIService(rdb *redis.Client, ctx context.Context) SessionManagementSubscriptionDataRetrievalAPIServicer {
	return &SessionManagementSubscriptionDataRetrievalAPIService{
		rdb: rdb,
		ctx: ctx,
	}
}

// GetSmData - retrieve a UE&#39;s Session Management Subscription Data
func (s *SessionManagementSubscriptionDataRetrievalAPIService) GetSmData(ctx context.Context, supi string, supportedFeatures string, singleNssai Snssai, dnn string, plmnId PlmnId, disasterRoamingInd bool, ifNoneMatch string, ifModifiedSince string) (ImplResponse, error) {

	strData, err := s.rdb.Get(ctx, supi).Result()

	if err != nil {
		// If there's an error during the retrieval of data from Redis, return a 500 Internal Server Error response
		return Response(http.StatusInternalServerError, nil), err
	}

	// if there is no error then unmarshal the data
	var subscriptionData *SessionManagementSubscriptionData
	if err := json.Unmarshal([]byte(strData), &subscriptionData); err != nil {
		return Response(http.StatusInternalServerError, nil), err
	}

	// Check if the subscription data exists
	if subscriptionData == nil {
		// If no valid subscription data is found, return a 404 Not Found response
		problemDetails := ProblemDetails{Detail: "No valid subscription data found for the UE"}
		return Response(http.StatusNotFound, problemDetails), nil
	}

	// Check if the requested session management subscription is available
	if !isRequestedSessionAvailable(subscriptionData, singleNssai, dnn, supportedFeatures, plmnId) {
		// If the requested session management subscription is not available, return a 404 Not Found response
		problemDetails := ProblemDetails{Detail: "Requested session management subscription is not available"}
		return Response(http.StatusNotFound, problemDetails), nil
	}
	// If all checks pass, return the subscription data with a 200 OK response
	return Response(http.StatusOK, subscriptionData), nil

	// TODO - update GetSmData with the required logic for this service method.
	// Add api_session_management_subscription_data_retrieval_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, SmSubsData{}) or use other options such as http.Ok ...
	// return Response(200, SmSubsData{}), nil

	// TODO: Uncomment the next line to return response Response(400, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(400, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(401, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(401, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(403, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(403, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(404, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(404, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(406, {}) or use other options such as http.Ok ...
	// return Response(406, nil),nil

	// TODO: Uncomment the next line to return response Response(429, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(429, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(500, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(500, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(502, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(502, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(503, ProblemDetails{}) or use other options such as http.Ok ...
	// return Response(503, ProblemDetails{}), nil

	// TODO: Uncomment the next line to return response Response(0, {}) or use other options such as http.Ok ...
	// return Response(0, nil),nil

	// return Response(http.StatusNotImplemented, nil), errors.New("GetSmData method not implemented")
}

// Function to check if the requested session management subscription is available
func isRequestedSessionAvailable(subscriptionData *SessionManagementSubscriptionData, singleNssai Snssai, dnn string, supportedFeatures string, plmnId PlmnId) bool {
	// Compare singleNssai
	// if subscriptionData.SingleNssai.Sst != singleNssai.Sst || subscriptionData.SingleNssai.Sd != singleNssai.Sd {
	// 	return false
	// }

	// // Compare plmnId
	// if subscriptionData.SupportedFeatures != supportedFeatures {
	// 	return false
	// }

	return true
}
