# \SubscriptionDeletionForSharedDataAPI

All URIs are relative to *https://example.com/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UnsubscribeForSharedData**](SubscriptionDeletionForSharedDataAPI.md#UnsubscribeForSharedData) | **Delete** /shared-data-subscriptions/{subscriptionId} | unsubscribe from notifications for shared data



## UnsubscribeForSharedData

> UnsubscribeForSharedData(ctx, subscriptionId).Execute()

unsubscribe from notifications for shared data

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
    subscriptionId := "subscriptionId_example" // string | Id of the Shared data Subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubscriptionDeletionForSharedDataAPI.UnsubscribeForSharedData(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionDeletionForSharedDataAPI.UnsubscribeForSharedData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Id of the Shared data Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeForSharedDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

