# TimeSyncSubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfReqAuthorizations** | [**[]AfRequestAuthorization**](AfRequestAuthorization.md) |  | 
**ServiceIds** | [**[]TimeSyncServiceId**](TimeSyncServiceId.md) |  | 

## Methods

### NewTimeSyncSubscriptionData

`func NewTimeSyncSubscriptionData(afReqAuthorizations []AfRequestAuthorization, serviceIds []TimeSyncServiceId, ) *TimeSyncSubscriptionData`

NewTimeSyncSubscriptionData instantiates a new TimeSyncSubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSyncSubscriptionDataWithDefaults

`func NewTimeSyncSubscriptionDataWithDefaults() *TimeSyncSubscriptionData`

NewTimeSyncSubscriptionDataWithDefaults instantiates a new TimeSyncSubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfReqAuthorizations

`func (o *TimeSyncSubscriptionData) GetAfReqAuthorizations() []AfRequestAuthorization`

GetAfReqAuthorizations returns the AfReqAuthorizations field if non-nil, zero value otherwise.

### GetAfReqAuthorizationsOk

`func (o *TimeSyncSubscriptionData) GetAfReqAuthorizationsOk() (*[]AfRequestAuthorization, bool)`

GetAfReqAuthorizationsOk returns a tuple with the AfReqAuthorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfReqAuthorizations

`func (o *TimeSyncSubscriptionData) SetAfReqAuthorizations(v []AfRequestAuthorization)`

SetAfReqAuthorizations sets AfReqAuthorizations field to given value.


### GetServiceIds

`func (o *TimeSyncSubscriptionData) GetServiceIds() []TimeSyncServiceId`

GetServiceIds returns the ServiceIds field if non-nil, zero value otherwise.

### GetServiceIdsOk

`func (o *TimeSyncSubscriptionData) GetServiceIdsOk() (*[]TimeSyncServiceId, bool)`

GetServiceIdsOk returns a tuple with the ServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIds

`func (o *TimeSyncSubscriptionData) SetServiceIds(v []TimeSyncServiceId)`

SetServiceIds sets ServiceIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


