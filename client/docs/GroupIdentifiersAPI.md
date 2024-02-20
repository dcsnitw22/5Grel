# \GroupIdentifiersAPI

All URIs are relative to *https://example.com/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroupIdentifiers**](GroupIdentifiersAPI.md#GetGroupIdentifiers) | **Get** /group-data/group-identifiers | Mapping of Group Identifiers



## GetGroupIdentifiers

> GroupIdentifiers GetGroupIdentifiers(ctx).ExtGroupId(extGroupId).IntGroupId(intGroupId).UeIdInd(ueIdInd).SupportedFeatures(supportedFeatures).AfId(afId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()

Mapping of Group Identifiers

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
    extGroupId := "extGroupId_example" // string | External Group Identifier (optional)
    intGroupId := "intGroupId_example" // string | Internal Group Identifier (optional)
    ueIdInd := true // bool | Indication whether UE identifiers are required or not (optional) (default to false)
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)
    afId := "afId_example" // string | AF identifier (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 9110, 3.2 (optional)
    ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 9110, 3.3 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupIdentifiersAPI.GetGroupIdentifiers(context.Background()).ExtGroupId(extGroupId).IntGroupId(intGroupId).UeIdInd(ueIdInd).SupportedFeatures(supportedFeatures).AfId(afId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupIdentifiersAPI.GetGroupIdentifiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupIdentifiers`: GroupIdentifiers
    fmt.Fprintf(os.Stdout, "Response from `GroupIdentifiersAPI.GetGroupIdentifiers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupIdentifiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extGroupId** | **string** | External Group Identifier | 
 **intGroupId** | **string** | Internal Group Identifier | 
 **ueIdInd** | **bool** | Indication whether UE identifiers are required or not | [default to false]
 **supportedFeatures** | **string** | Supported Features | 
 **afId** | **string** | AF identifier | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 9110, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 9110, 3.3 | 

### Return type

[**GroupIdentifiers**](GroupIdentifiers.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

