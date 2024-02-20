# \SubscriptionCreationAPI

All URIs are relative to *https://example.com/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Subscribe**](SubscriptionCreationAPI.md#Subscribe) | **Post** /{ueId}/sdm-subscriptions | subscribe to notifications



## Subscribe

> SdmSubscription Subscribe(ctx, ueId).SdmSubscription(sdmSubscription).SharedDataIds(sharedDataIds).Execute()

subscribe to notifications

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
    sdmSubscription := *openapiclient.NewSdmSubscription("NfInstanceId_example", "CallbackReference_example", []string{"MonitoredResourceUris_example"}) // SdmSubscription | 
    sharedDataIds := []string{"Inner_example"} // []string | List of IDs identifying shared Data already available at and subscribed by the NF service consumer  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionCreationAPI.Subscribe(context.Background(), ueId).SdmSubscription(sdmSubscription).SharedDataIds(sharedDataIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionCreationAPI.Subscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Subscribe`: SdmSubscription
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionCreationAPI.Subscribe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identity of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sdmSubscription** | [**SdmSubscription**](SdmSubscription.md) |  | 
 **sharedDataIds** | **[]string** | List of IDs identifying shared Data already available at and subscribed by the NF service consumer  | 

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

