# LcsSubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfiguredLmfId** | Pointer to **string** | LMF identification. | [optional] 
**PruInd** | Pointer to [**PruInd**](PruInd.md) |  | [optional] 
**LpHapType** | Pointer to [**LpHapType**](LpHapType.md) |  | [optional] 
**UserPlanePosIndLmf** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewLcsSubscriptionData

`func NewLcsSubscriptionData() *LcsSubscriptionData`

NewLcsSubscriptionData instantiates a new LcsSubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLcsSubscriptionDataWithDefaults

`func NewLcsSubscriptionDataWithDefaults() *LcsSubscriptionData`

NewLcsSubscriptionDataWithDefaults instantiates a new LcsSubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguredLmfId

`func (o *LcsSubscriptionData) GetConfiguredLmfId() string`

GetConfiguredLmfId returns the ConfiguredLmfId field if non-nil, zero value otherwise.

### GetConfiguredLmfIdOk

`func (o *LcsSubscriptionData) GetConfiguredLmfIdOk() (*string, bool)`

GetConfiguredLmfIdOk returns a tuple with the ConfiguredLmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredLmfId

`func (o *LcsSubscriptionData) SetConfiguredLmfId(v string)`

SetConfiguredLmfId sets ConfiguredLmfId field to given value.

### HasConfiguredLmfId

`func (o *LcsSubscriptionData) HasConfiguredLmfId() bool`

HasConfiguredLmfId returns a boolean if a field has been set.

### GetPruInd

`func (o *LcsSubscriptionData) GetPruInd() PruInd`

GetPruInd returns the PruInd field if non-nil, zero value otherwise.

### GetPruIndOk

`func (o *LcsSubscriptionData) GetPruIndOk() (*PruInd, bool)`

GetPruIndOk returns a tuple with the PruInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruInd

`func (o *LcsSubscriptionData) SetPruInd(v PruInd)`

SetPruInd sets PruInd field to given value.

### HasPruInd

`func (o *LcsSubscriptionData) HasPruInd() bool`

HasPruInd returns a boolean if a field has been set.

### GetLpHapType

`func (o *LcsSubscriptionData) GetLpHapType() LpHapType`

GetLpHapType returns the LpHapType field if non-nil, zero value otherwise.

### GetLpHapTypeOk

`func (o *LcsSubscriptionData) GetLpHapTypeOk() (*LpHapType, bool)`

GetLpHapTypeOk returns a tuple with the LpHapType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpHapType

`func (o *LcsSubscriptionData) SetLpHapType(v LpHapType)`

SetLpHapType sets LpHapType field to given value.

### HasLpHapType

`func (o *LcsSubscriptionData) HasLpHapType() bool`

HasLpHapType returns a boolean if a field has been set.

### GetUserPlanePosIndLmf

`func (o *LcsSubscriptionData) GetUserPlanePosIndLmf() bool`

GetUserPlanePosIndLmf returns the UserPlanePosIndLmf field if non-nil, zero value otherwise.

### GetUserPlanePosIndLmfOk

`func (o *LcsSubscriptionData) GetUserPlanePosIndLmfOk() (*bool, bool)`

GetUserPlanePosIndLmfOk returns a tuple with the UserPlanePosIndLmf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPlanePosIndLmf

`func (o *LcsSubscriptionData) SetUserPlanePosIndLmf(v bool)`

SetUserPlanePosIndLmf sets UserPlanePosIndLmf field to given value.

### HasUserPlanePosIndLmf

`func (o *LcsSubscriptionData) HasUserPlanePosIndLmf() bool`

HasUserPlanePosIndLmf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


