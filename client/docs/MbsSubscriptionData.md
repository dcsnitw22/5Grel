# MbsSubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsAllowed** | Pointer to **bool** |  | [optional] [default to false]
**MbsSessionIdList** | Pointer to [**[]MbsSessionId**](MbsSessionId.md) |  | [optional] 
**UeMbsAssistanceInfo** | Pointer to [**[]MbsSessionId**](MbsSessionId.md) |  | [optional] 

## Methods

### NewMbsSubscriptionData

`func NewMbsSubscriptionData() *MbsSubscriptionData`

NewMbsSubscriptionData instantiates a new MbsSubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsSubscriptionDataWithDefaults

`func NewMbsSubscriptionDataWithDefaults() *MbsSubscriptionData`

NewMbsSubscriptionDataWithDefaults instantiates a new MbsSubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsAllowed

`func (o *MbsSubscriptionData) GetMbsAllowed() bool`

GetMbsAllowed returns the MbsAllowed field if non-nil, zero value otherwise.

### GetMbsAllowedOk

`func (o *MbsSubscriptionData) GetMbsAllowedOk() (*bool, bool)`

GetMbsAllowedOk returns a tuple with the MbsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsAllowed

`func (o *MbsSubscriptionData) SetMbsAllowed(v bool)`

SetMbsAllowed sets MbsAllowed field to given value.

### HasMbsAllowed

`func (o *MbsSubscriptionData) HasMbsAllowed() bool`

HasMbsAllowed returns a boolean if a field has been set.

### GetMbsSessionIdList

`func (o *MbsSubscriptionData) GetMbsSessionIdList() []MbsSessionId`

GetMbsSessionIdList returns the MbsSessionIdList field if non-nil, zero value otherwise.

### GetMbsSessionIdListOk

`func (o *MbsSubscriptionData) GetMbsSessionIdListOk() (*[]MbsSessionId, bool)`

GetMbsSessionIdListOk returns a tuple with the MbsSessionIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSessionIdList

`func (o *MbsSubscriptionData) SetMbsSessionIdList(v []MbsSessionId)`

SetMbsSessionIdList sets MbsSessionIdList field to given value.

### HasMbsSessionIdList

`func (o *MbsSubscriptionData) HasMbsSessionIdList() bool`

HasMbsSessionIdList returns a boolean if a field has been set.

### GetUeMbsAssistanceInfo

`func (o *MbsSubscriptionData) GetUeMbsAssistanceInfo() []MbsSessionId`

GetUeMbsAssistanceInfo returns the UeMbsAssistanceInfo field if non-nil, zero value otherwise.

### GetUeMbsAssistanceInfoOk

`func (o *MbsSubscriptionData) GetUeMbsAssistanceInfoOk() (*[]MbsSessionId, bool)`

GetUeMbsAssistanceInfoOk returns a tuple with the UeMbsAssistanceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMbsAssistanceInfo

`func (o *MbsSubscriptionData) SetUeMbsAssistanceInfo(v []MbsSessionId)`

SetUeMbsAssistanceInfo sets UeMbsAssistanceInfo field to given value.

### HasUeMbsAssistanceInfo

`func (o *MbsSubscriptionData) HasUeMbsAssistanceInfo() bool`

HasUeMbsAssistanceInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


