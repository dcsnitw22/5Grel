# \TriggerSORInfoUpdateAPI

All URIs are relative to *https://example.com/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSORInfo**](TriggerSORInfoUpdateAPI.md#UpdateSORInfo) | **Post** /{supi}/am-data/update-sor | Nudm_Sdm custom operation to trigger SOR info update



## UpdateSORInfo

> SorInfo UpdateSORInfo(ctx, supi).SorUpdateInfo(sorUpdateInfo).Execute()

Nudm_Sdm custom operation to trigger SOR info update

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
    sorUpdateInfo := *openapiclient.NewSorUpdateInfo(*openapiclient.NewPlmnId("Mcc_example", "Mnc_example")) // SorUpdateInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TriggerSORInfoUpdateAPI.UpdateSORInfo(context.Background(), supi).SorUpdateInfo(sorUpdateInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggerSORInfoUpdateAPI.UpdateSORInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSORInfo`: SorInfo
    fmt.Fprintf(os.Stdout, "Response from `TriggerSORInfoUpdateAPI.UpdateSORInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSORInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sorUpdateInfo** | [**SorUpdateInfo**](SorUpdateInfo.md) |  | 

### Return type

[**SorInfo**](SorInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

