# \ProvidingAcknowledgementOfUEParametersUpdateAPI

All URIs are relative to *https://example.com/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpuAck**](ProvidingAcknowledgementOfUEParametersUpdateAPI.md#UpuAck) | **Put** /{supi}/am-data/upu-ack | Nudm_Sdm Info for UPU service operation



## UpuAck

> UpuAck(ctx, supi).AcknowledgeInfo(acknowledgeInfo).Execute()

Nudm_Sdm Info for UPU service operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    supi := "supi_example" // string | Identifier of the UE
    acknowledgeInfo := *openapiclient.NewAcknowledgeInfo(time.Now()) // AcknowledgeInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProvidingAcknowledgementOfUEParametersUpdateAPI.UpuAck(context.Background(), supi).AcknowledgeInfo(acknowledgeInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidingAcknowledgementOfUEParametersUpdateAPI.UpuAck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpuAckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acknowledgeInfo** | [**AcknowledgeInfo**](AcknowledgeInfo.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

