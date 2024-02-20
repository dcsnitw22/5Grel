# \UEContextInSMSFDataRetrievalAPI

All URIs are relative to *https://example.com/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUeCtxInSmsfData**](UEContextInSMSFDataRetrievalAPI.md#GetUeCtxInSmsfData) | **Get** /{supi}/ue-context-in-smsf-data | retrieve a UE&#39;s UE Context In SMSF Data



## GetUeCtxInSmsfData

> UeContextInSmsfData GetUeCtxInSmsfData(ctx, supi).SupportedFeatures(supportedFeatures).Execute()

retrieve a UE's UE Context In SMSF Data

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
    resp, r, err := apiClient.UEContextInSMSFDataRetrievalAPI.GetUeCtxInSmsfData(context.Background(), supi).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UEContextInSMSFDataRetrievalAPI.GetUeCtxInSmsfData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUeCtxInSmsfData`: UeContextInSmsfData
    fmt.Fprintf(os.Stdout, "Response from `UEContextInSMSFDataRetrievalAPI.GetUeCtxInSmsfData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUeCtxInSmsfDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**UeContextInSmsfData**](UeContextInSmsfData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

