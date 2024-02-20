# \MultipleIdentifiersAPI

All URIs are relative to *https://example.com/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMultipleIdentifiers**](MultipleIdentifiersAPI.md#GetMultipleIdentifiers) | **Get** /multiple-identifiers | Mapping of UE Identifiers



## GetMultipleIdentifiers

> map[string]SupiInfo GetMultipleIdentifiers(ctx).GpsiList(gpsiList).SupportedFeatures(supportedFeatures).Execute()

Mapping of UE Identifiers

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
    gpsiList := []string{"Inner_example"} // []string | list of the GPSIs
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultipleIdentifiersAPI.GetMultipleIdentifiers(context.Background()).GpsiList(gpsiList).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultipleIdentifiersAPI.GetMultipleIdentifiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMultipleIdentifiers`: map[string]SupiInfo
    fmt.Fprintf(os.Stdout, "Response from `MultipleIdentifiersAPI.GetMultipleIdentifiers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMultipleIdentifiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gpsiList** | **[]string** | list of the GPSIs | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**map[string]SupiInfo**](SupiInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

