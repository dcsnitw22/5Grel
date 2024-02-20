# \GPSIToSUPITranslationOrSUPIToGPSITranslationAPI

All URIs are relative to *https://example.com/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSupiOrGpsi**](GPSIToSUPITranslationOrSUPIToGPSITranslationAPI.md#GetSupiOrGpsi) | **Get** /{ueId}/id-translation-result | retrieve a UE&#39;s SUPI or GPSI



## GetSupiOrGpsi

> IdTranslationResult GetSupiOrGpsi(ctx, ueId).SupportedFeatures(supportedFeatures).AfId(afId).AppPortId(appPortId).AfServiceId(afServiceId).MtcProviderInfo(mtcProviderInfo).RequestedGpsiType(requestedGpsiType).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()

retrieve a UE's SUPI or GPSI

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
    ueId := "ueId_example" // string | Identifier of the UE
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)
    afId := "afId_example" // string | AF identifier (optional)
    appPortId := *openapiclient.NewAppPortId() // AppPortId | Application port identifier (optional)
    afServiceId := "afServiceId_example" // string | AF Service Identifier (optional)
    mtcProviderInfo := "mtcProviderInfo_example" // string | MTC Provider Information (optional)
    requestedGpsiType := *openapiclient.NewGpsiType() // GpsiType | Requested GPSI Type (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 9110, 3.2 (optional)
    ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 9110, 3.3 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GPSIToSUPITranslationOrSUPIToGPSITranslationAPI.GetSupiOrGpsi(context.Background(), ueId).SupportedFeatures(supportedFeatures).AfId(afId).AppPortId(appPortId).AfServiceId(afServiceId).MtcProviderInfo(mtcProviderInfo).RequestedGpsiType(requestedGpsiType).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GPSIToSUPITranslationOrSUPIToGPSITranslationAPI.GetSupiOrGpsi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupiOrGpsi`: IdTranslationResult
    fmt.Fprintf(os.Stdout, "Response from `GPSIToSUPITranslationOrSUPIToGPSITranslationAPI.GetSupiOrGpsi`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupiOrGpsiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 
 **afId** | **string** | AF identifier | 
 **appPortId** | [**AppPortId**](AppPortId.md) | Application port identifier | 
 **afServiceId** | **string** | AF Service Identifier | 
 **mtcProviderInfo** | **string** | MTC Provider Information | 
 **requestedGpsiType** | [**GpsiType**](GpsiType.md) | Requested GPSI Type | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 9110, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 9110, 3.3 | 

### Return type

[**IdTranslationResult**](IdTranslationResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

