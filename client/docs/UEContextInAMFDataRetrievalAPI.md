# \UEContextInAMFDataRetrievalAPI

All URIs are relative to *https://example.com/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUeCtxInAmfData**](UEContextInAMFDataRetrievalAPI.md#GetUeCtxInAmfData) | **Get** /{supi}/ue-context-in-amf-data | retrieve a UE&#39;s UE Context In AMF Data



## GetUeCtxInAmfData

> UeContextInAmfData GetUeCtxInAmfData(ctx, supi).SupportedFeatures(supportedFeatures).Execute()

retrieve a UE's UE Context In AMF Data

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UEContextInAMFDataRetrievalAPI.GetUeCtxInAmfData(context.Background(), supi).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UEContextInAMFDataRetrievalAPI.GetUeCtxInAmfData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUeCtxInAmfData`: UeContextInAmfData
    fmt.Fprintf(os.Stdout, "Response from `UEContextInAMFDataRetrievalAPI.GetUeCtxInAmfData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUeCtxInAmfDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**UeContextInAmfData**](UeContextInAmfData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

