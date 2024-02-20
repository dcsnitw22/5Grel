# \RetrievalOfMultipleDataSetsAPI

All URIs are relative to *https://example.com/nudm-sdm/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDataSets**](RetrievalOfMultipleDataSetsAPI.md#GetDataSets) | **Get** /{supi} | retrieve multiple data sets



## GetDataSets

> SubscriptionDataSets GetDataSets(ctx, supi).DatasetNames(datasetNames).PlmnId(plmnId).AdjacentPlmns(adjacentPlmns).SingleNssai(singleNssai).Dnn(dnn).UcPurpose(ucPurpose).DisasterRoamingInd(disasterRoamingInd).SupportedFeatures(supportedFeatures).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()

retrieve multiple data sets

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
    datasetNames := []openapiclient.DataSetName{*openapiclient.NewDataSetName()} // []DataSetName | List of dataset names
    plmnId := *openapiclient.NewPlmnIdNid("Mcc_example", "Mnc_example") // PlmnIdNid | serving PLMN ID (optional)
    adjacentPlmns := []openapiclient.PlmnId{*openapiclient.NewPlmnId("Mcc_example", "Mnc_example")} // []PlmnId | List of PLMNs adjacent to the UE's serving PLMN (optional)
    singleNssai := *openapiclient.NewSnssai(int32(123)) // Snssai |  (optional)
    dnn := "dnn_example" // string |  (optional)
    ucPurpose := *openapiclient.NewUcPurpose() // UcPurpose | User consent purpose (optional)
    disasterRoamingInd := true // bool | Indication whether Disaster Roaming service is applied or not (optional) (default to false)
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 9110, 3.2 (optional)
    ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 9110, 3.3 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetrievalOfMultipleDataSetsAPI.GetDataSets(context.Background(), supi).DatasetNames(datasetNames).PlmnId(plmnId).AdjacentPlmns(adjacentPlmns).SingleNssai(singleNssai).Dnn(dnn).UcPurpose(ucPurpose).DisasterRoamingInd(disasterRoamingInd).SupportedFeatures(supportedFeatures).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetrievalOfMultipleDataSetsAPI.GetDataSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataSets`: SubscriptionDataSets
    fmt.Fprintf(os.Stdout, "Response from `RetrievalOfMultipleDataSetsAPI.GetDataSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasetNames** | [**[]DataSetName**](DataSetName.md) | List of dataset names | 
 **plmnId** | [**PlmnIdNid**](PlmnIdNid.md) | serving PLMN ID | 
 **adjacentPlmns** | [**[]PlmnId**](PlmnId.md) | List of PLMNs adjacent to the UE&#39;s serving PLMN | 
 **singleNssai** | [**Snssai**](Snssai.md) |  | 
 **dnn** | **string** |  | 
 **ucPurpose** | [**UcPurpose**](UcPurpose.md) | User consent purpose | 
 **disasterRoamingInd** | **bool** | Indication whether Disaster Roaming service is applied or not | [default to false]
 **supportedFeatures** | **string** | Supported Features | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 9110, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 9110, 3.3 | 

### Return type

[**SubscriptionDataSets**](SubscriptionDataSets.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

