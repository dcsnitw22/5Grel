# \SessionManagementSubscriptionDataRetrievalAPI

All URIs are relative to *https://example.com/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSmData**](SessionManagementSubscriptionDataRetrievalAPI.md#GetSmData) | **Get** /{supi}/sm-data | retrieve a UE&#39;s Session Management Subscription Data



## GetSmData

> SmSubsData GetSmData(ctx, supi).SupportedFeatures(supportedFeatures).SingleNssai(singleNssai).Dnn(dnn).PlmnId(plmnId).DisasterRoamingInd(disasterRoamingInd).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()

retrieve a UE's Session Management Subscription Data

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    supi := "supi_example" // string | Identifier of the UE
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)
    singleNssai := *openapiclient.NewSnssai(int32(123)) // Snssai |  (optional)
    dnn := "dnn_example" // string |  (optional)
    plmnId := *openapiclient.NewPlmnId("Mcc_example", "Mnc_example") // PlmnId |  (optional)
    disasterRoamingInd := true // bool | Indication whether Disaster Roaming service is applied or not (optional) (default to false)
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 9110, 3.2 (optional)
    ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 9110, 3.3 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionManagementSubscriptionDataRetrievalAPI.GetSmData(context.Background(), supi).SupportedFeatures(supportedFeatures).SingleNssai(singleNssai).Dnn(dnn).PlmnId(plmnId).DisasterRoamingInd(disasterRoamingInd).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionManagementSubscriptionDataRetrievalAPI.GetSmData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSmData`: SmSubsData
    fmt.Fprintf(os.Stdout, "Response from `SessionManagementSubscriptionDataRetrievalAPI.GetSmData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 
 **singleNssai** | [**Snssai**](Snssai.md) |  | 
 **dnn** | **string** |  | 
 **plmnId** | [**PlmnId**](PlmnId.md) |  | 
 **disasterRoamingInd** | **bool** | Indication whether Disaster Roaming service is applied or not | [default to false]
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 9110, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 9110, 3.3 | 

### Return type

[**SmSubsData**](SmSubsData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

