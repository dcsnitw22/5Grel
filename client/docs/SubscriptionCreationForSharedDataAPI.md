# \SubscriptionCreationForSharedDataAPI

All URIs are relative to *https://example.com/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscribeToSharedData**](SubscriptionCreationForSharedDataAPI.md#SubscribeToSharedData) | **Post** /shared-data-subscriptions | subscribe to notifications for shared data



## SubscribeToSharedData

> SdmSubscription SubscribeToSharedData(ctx).SdmSubscription(sdmSubscription).Execute()

subscribe to notifications for shared data

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
    sdmSubscription := *openapiclient.NewSdmSubscription("NfInstanceId_example", "CallbackReference_example", []string{"MonitoredResourceUris_example"}) // SdmSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionCreationForSharedDataAPI.SubscribeToSharedData(context.Background()).SdmSubscription(sdmSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionCreationForSharedDataAPI.SubscribeToSharedData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscribeToSharedData`: SdmSubscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionCreationForSharedDataAPI.SubscribeToSharedData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeToSharedDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sdmSubscription** | [**SdmSubscription**](SdmSubscription.md) |  | 

### Return type

[**SdmSubscription**](SdmSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

