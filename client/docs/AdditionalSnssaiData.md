# AdditionalSnssaiData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequiredAuthnAuthz** | Pointer to **bool** |  | [optional] 
**SubscribedUeSliceMbr** | Pointer to [**NullableSliceMbr**](SliceMbr.md) |  | [optional] 
**SubscribedNsSrgList** | Pointer to **[]string** |  | [optional] 
**NsacMode** | Pointer to [**NsacAdmissionMode**](NsacAdmissionMode.md) |  | [optional] 
**ValidTimePeriod** | Pointer to [**ValidTimePeriod**](ValidTimePeriod.md) |  | [optional] 
**DeregInactTimer** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**OnDemand** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAdditionalSnssaiData

`func NewAdditionalSnssaiData() *AdditionalSnssaiData`

NewAdditionalSnssaiData instantiates a new AdditionalSnssaiData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalSnssaiDataWithDefaults

`func NewAdditionalSnssaiDataWithDefaults() *AdditionalSnssaiData`

NewAdditionalSnssaiDataWithDefaults instantiates a new AdditionalSnssaiData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequiredAuthnAuthz

`func (o *AdditionalSnssaiData) GetRequiredAuthnAuthz() bool`

GetRequiredAuthnAuthz returns the RequiredAuthnAuthz field if non-nil, zero value otherwise.

### GetRequiredAuthnAuthzOk

`func (o *AdditionalSnssaiData) GetRequiredAuthnAuthzOk() (*bool, bool)`

GetRequiredAuthnAuthzOk returns a tuple with the RequiredAuthnAuthz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAuthnAuthz

`func (o *AdditionalSnssaiData) SetRequiredAuthnAuthz(v bool)`

SetRequiredAuthnAuthz sets RequiredAuthnAuthz field to given value.

### HasRequiredAuthnAuthz

`func (o *AdditionalSnssaiData) HasRequiredAuthnAuthz() bool`

HasRequiredAuthnAuthz returns a boolean if a field has been set.

### GetSubscribedUeSliceMbr

`func (o *AdditionalSnssaiData) GetSubscribedUeSliceMbr() SliceMbr`

GetSubscribedUeSliceMbr returns the SubscribedUeSliceMbr field if non-nil, zero value otherwise.

### GetSubscribedUeSliceMbrOk

`func (o *AdditionalSnssaiData) GetSubscribedUeSliceMbrOk() (*SliceMbr, bool)`

GetSubscribedUeSliceMbrOk returns a tuple with the SubscribedUeSliceMbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedUeSliceMbr

`func (o *AdditionalSnssaiData) SetSubscribedUeSliceMbr(v SliceMbr)`

SetSubscribedUeSliceMbr sets SubscribedUeSliceMbr field to given value.

### HasSubscribedUeSliceMbr

`func (o *AdditionalSnssaiData) HasSubscribedUeSliceMbr() bool`

HasSubscribedUeSliceMbr returns a boolean if a field has been set.

### SetSubscribedUeSliceMbrNil

`func (o *AdditionalSnssaiData) SetSubscribedUeSliceMbrNil(b bool)`

 SetSubscribedUeSliceMbrNil sets the value for SubscribedUeSliceMbr to be an explicit nil

### UnsetSubscribedUeSliceMbr
`func (o *AdditionalSnssaiData) UnsetSubscribedUeSliceMbr()`

UnsetSubscribedUeSliceMbr ensures that no value is present for SubscribedUeSliceMbr, not even an explicit nil
### GetSubscribedNsSrgList

`func (o *AdditionalSnssaiData) GetSubscribedNsSrgList() []string`

GetSubscribedNsSrgList returns the SubscribedNsSrgList field if non-nil, zero value otherwise.

### GetSubscribedNsSrgListOk

`func (o *AdditionalSnssaiData) GetSubscribedNsSrgListOk() (*[]string, bool)`

GetSubscribedNsSrgListOk returns a tuple with the SubscribedNsSrgList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedNsSrgList

`func (o *AdditionalSnssaiData) SetSubscribedNsSrgList(v []string)`

SetSubscribedNsSrgList sets SubscribedNsSrgList field to given value.

### HasSubscribedNsSrgList

`func (o *AdditionalSnssaiData) HasSubscribedNsSrgList() bool`

HasSubscribedNsSrgList returns a boolean if a field has been set.

### GetNsacMode

`func (o *AdditionalSnssaiData) GetNsacMode() NsacAdmissionMode`

GetNsacMode returns the NsacMode field if non-nil, zero value otherwise.

### GetNsacModeOk

`func (o *AdditionalSnssaiData) GetNsacModeOk() (*NsacAdmissionMode, bool)`

GetNsacModeOk returns a tuple with the NsacMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsacMode

`func (o *AdditionalSnssaiData) SetNsacMode(v NsacAdmissionMode)`

SetNsacMode sets NsacMode field to given value.

### HasNsacMode

`func (o *AdditionalSnssaiData) HasNsacMode() bool`

HasNsacMode returns a boolean if a field has been set.

### GetValidTimePeriod

`func (o *AdditionalSnssaiData) GetValidTimePeriod() ValidTimePeriod`

GetValidTimePeriod returns the ValidTimePeriod field if non-nil, zero value otherwise.

### GetValidTimePeriodOk

`func (o *AdditionalSnssaiData) GetValidTimePeriodOk() (*ValidTimePeriod, bool)`

GetValidTimePeriodOk returns a tuple with the ValidTimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTimePeriod

`func (o *AdditionalSnssaiData) SetValidTimePeriod(v ValidTimePeriod)`

SetValidTimePeriod sets ValidTimePeriod field to given value.

### HasValidTimePeriod

`func (o *AdditionalSnssaiData) HasValidTimePeriod() bool`

HasValidTimePeriod returns a boolean if a field has been set.

### GetDeregInactTimer

`func (o *AdditionalSnssaiData) GetDeregInactTimer() int32`

GetDeregInactTimer returns the DeregInactTimer field if non-nil, zero value otherwise.

### GetDeregInactTimerOk

`func (o *AdditionalSnssaiData) GetDeregInactTimerOk() (*int32, bool)`

GetDeregInactTimerOk returns a tuple with the DeregInactTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeregInactTimer

`func (o *AdditionalSnssaiData) SetDeregInactTimer(v int32)`

SetDeregInactTimer sets DeregInactTimer field to given value.

### HasDeregInactTimer

`func (o *AdditionalSnssaiData) HasDeregInactTimer() bool`

HasDeregInactTimer returns a boolean if a field has been set.

### GetOnDemand

`func (o *AdditionalSnssaiData) GetOnDemand() bool`

GetOnDemand returns the OnDemand field if non-nil, zero value otherwise.

### GetOnDemandOk

`func (o *AdditionalSnssaiData) GetOnDemandOk() (*bool, bool)`

GetOnDemandOk returns a tuple with the OnDemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemand

`func (o *AdditionalSnssaiData) SetOnDemand(v bool)`

SetOnDemand sets OnDemand field to given value.

### HasOnDemand

`func (o *AdditionalSnssaiData) HasOnDemand() bool`

HasOnDemand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


