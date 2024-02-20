# \RetrievalOfSharedDataAPI

All URIs are relative to *https://example.com/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSharedData**](RetrievalOfSharedDataAPI.md#GetSharedData) | **Get** /shared-data | retrieve shared data



## GetSharedData

> []SharedData GetSharedData(ctx).SharedDataIds(sharedDataIds).SupportedFeatures(supportedFeatures).SupportedFeatures2(supportedFeatures2).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()

retrieve shared data

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
    sharedDataIds := []string{"Inner_example"} // []string | List of shared data ids
    supportedFeatures := "supportedFeatures_example" // string | Supported Features; this query parameter should not be used (optional)
    supportedFeatures2 := "supportedFeatures_example" // string | Supported Features (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 9110, 3.2 (optional)
    ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 9110, 3.3 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetrievalOfSharedDataAPI.GetSharedData(context.Background()).SharedDataIds(sharedDataIds).SupportedFeatures(supportedFeatures).SupportedFeatures2(supportedFeatures2).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetrievalOfSharedDataAPI.GetSharedData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSharedData`: []SharedData
    fmt.Fprintf(os.Stdout, "Response from `RetrievalOfSharedDataAPI.GetSharedData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSharedDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sharedDataIds** | **[]string** | List of shared data ids | 
 **supportedFeatures** | **string** | Supported Features; this query parameter should not be used | 
 **supportedFeatures2** | **string** | Supported Features | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 9110, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 9110, 3.3 | 

### Return type

[**[]SharedData**](SharedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

