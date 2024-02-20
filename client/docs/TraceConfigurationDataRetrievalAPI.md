# \TraceConfigurationDataRetrievalAPI

All URIs are relative to *https://example.com/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTraceConfigData**](TraceConfigurationDataRetrievalAPI.md#GetTraceConfigData) | **Get** /{supi}/trace-data | retrieve a UE&#39;s Trace Configuration Data



## GetTraceConfigData

> TraceDataResponse GetTraceConfigData(ctx, supi).SupportedFeatures(supportedFeatures).PlmnId(plmnId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()

retrieve a UE's Trace Configuration Data

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
    plmnId := *openapiclient.NewPlmnId("Mcc_example", "Mnc_example") // PlmnId | serving PLMN ID (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 9110, 3.2 (optional)
    ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 9110, 3.3 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TraceConfigurationDataRetrievalAPI.GetTraceConfigData(context.Background(), supi).SupportedFeatures(supportedFeatures).PlmnId(plmnId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TraceConfigurationDataRetrievalAPI.GetTraceConfigData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTraceConfigData`: TraceDataResponse
    fmt.Fprintf(os.Stdout, "Response from `TraceConfigurationDataRetrievalAPI.GetTraceConfigData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceConfigDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 
 **plmnId** | [**PlmnId**](PlmnId.md) | serving PLMN ID | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 9110, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 9110, 3.3 | 

### Return type

[**TraceDataResponse**](TraceDataResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

