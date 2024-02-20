# \SubscriptionModificationAPI

All URIs are relative to *https://example.com/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Modify**](SubscriptionModificationAPI.md#Modify) | **Patch** /{ueId}/sdm-subscriptions/{subscriptionId} | modify the subscription
[**ModifySharedDataSubs**](SubscriptionModificationAPI.md#ModifySharedDataSubs) | **Patch** /shared-data-subscriptions/{subscriptionId} | modify the subscription



## Modify

> Modify200Response Modify(ctx, ueId, subscriptionId).SdmSubsModification(sdmSubsModification).SupportedFeatures(supportedFeatures).Execute()

modify the subscription

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
    ueId := "ueId_example" // string | Identity of the user
    subscriptionId := "subscriptionId_example" // string | Id of the SDM Subscription
    sdmSubsModification := *openapiclient.NewSdmSubsModification() // SdmSubsModification | 
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionModificationAPI.Modify(context.Background(), ueId, subscriptionId).SdmSubsModification(sdmSubsModification).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionModificationAPI.Modify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Modify`: Modify200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionModificationAPI.Modify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identity of the user | 
**subscriptionId** | **string** | Id of the SDM Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sdmSubsModification** | [**SdmSubsModification**](SdmSubsModification.md) |  | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**Modify200Response**](Modify200Response.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifySharedDataSubs

> Modify200Response ModifySharedDataSubs(ctx, subscriptionId).SdmSubsModification(sdmSubsModification).SupportedFeatures(supportedFeatures).Execute()

modify the subscription

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
    subscriptionId := "subscriptionId_example" // string | Id of the SDM Subscription
    sdmSubsModification := *openapiclient.NewSdmSubsModification() // SdmSubsModification | 
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionModificationAPI.ModifySharedDataSubs(context.Background(), subscriptionId).SdmSubsModification(sdmSubsModification).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionModificationAPI.ModifySharedDataSubs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifySharedDataSubs`: Modify200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionModificationAPI.ModifySharedDataSubs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Id of the SDM Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifySharedDataSubsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sdmSubsModification** | [**SdmSubsModification**](SdmSubsModification.md) |  | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**Modify200Response**](Modify200Response.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

